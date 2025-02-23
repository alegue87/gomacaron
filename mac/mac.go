package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"mac/hotswap/inter/hutils"

	"github.com/edwingeng/hotswap"
	"github.com/edwingeng/hotswap/demo/hello/g"
	"github.com/google/uuid"

	"mac/auth"
	"net/http"

	"mac/models"

	//hello "example/hotswap/initial"

	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"gorm.io/driver/postgres"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Println("Connecting to db....")
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Rome"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to db!")

	// Migrate the schema
	fmt.Println("Migrate model User...")
	db.AutoMigrate(&models.User{})

	fmt.Println("Migrate model jobs...")
	db.AutoMigrate(&models.Job{})

	m := macaron.Classic()
	m.Map(db)
	fmt.Println("Registering routes...")
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/User", handleGetUser)
	m.Post("/User", binding.BindIgnErr(models.UserForm{}), handlePostUser)
	m.Post("/Login", binding.BindIgnErr(models.UserForm{}), auth.HandlePostLogin)
	m.Get("/Test", handleTest)
	m.Get("/Test1", handleTest1)
	m.Get("/Test2", handleTest2)

	go http.ListenAndServe(":8080", m)
	go m.Run()
	fmt.Println("Ready macaron")
	go jobManager(db)

	main1(db)

}

var s gocron.Scheduler = nil

func jobManager(db *gorm.DB) {
	fmt.Println("Job Manager")

	if s == nil {
		// create a scheduler
		var err error
		s, err = gocron.NewScheduler()
		defer func() { _ = s.Shutdown() }()

		if err != nil {
			fmt.Println("Error during creation of the scheduler")
		}
	}

	// recupera job per l'avvio nello scheduler
	var jobs []models.Job
	db.Find(&jobs)

	if len(jobs) > 0 {
		fmt.Println("jobs total :", len(jobs))
	}

	for _, job := range jobs {

		fmt.Println("JOB: ", job.Uuid, " start ", job.Run)

		/*

			gocron.CronJob, "* 0 0 ? * * *", true),


			start := time.Date(2025, 1, 17, 19, 45, 0, 0, time.UTC)
		*/

		//start := time.Date(2025, 01, 20, 17, 56, 01, 01, time.Local)
		//fmt.Println("localtime: ", time.Now())
		//fmt.Println("Starting at: ", start)
		curMs := int(time.Now().UnixMilli() / 1000)
		//time.Sleep(time.Second)
		nextMs := int(time.Now().UnixMilli() / 1000)

		if job.Run {
			// add a job to the scheduler
			fmt.Println("running")
			fmt.Println("Ms ", nextMs-curMs)

			j, e := s.NewJob(
				gocron.DurationJob(
					time.Millisecond*100,
				),
				gocron.NewTask(serviceRunner, db, job),
				//gocron.WithStartAt(
				//	gocron.WithStartDateTime(start),
				//),
				gocron.WithTags(string(job.ID)),
			)

			if j == nil {
				fmt.Println("Error locating job", e.Error)
			}

			// each job has a unique id
			fmt.Println(j.ID().String())
			fmt.Print(j)
			job.Uuid = j.ID().String()
			job.Run = true
			db.Save(&job)

		} else {
			s.RemoveJob(uuid.MustParse(job.Uuid))
		}
		s.Start()
		fmt.Println("\nScheduler start")
		select {}

		// when you're done, shut it down
		//s.Shutdown()

	}
	// start the scheduler
}

func handleGetUser(ctx *macaron.Context, db *gorm.DB) {
	res := (ctx).Query("ciao")

	fmt.Println(res)

	var users []models.User
	e := db.Find(&users)

	if e != nil {
		fmt.Println(e.Error)
	}

	for u, _ := range users {
		fmt.Println(u)
	}

}

func handlePostUser(ctx *macaron.Context, db *gorm.DB, user models.UserForm) {

	newUser := models.User{Username: user.Username, Token: "", Password: user.Password}

	result := db.Create(&newUser) // pass pointer of data to Create

	if result == nil {
		fmt.Printf("ok")
	} else {
		fmt.Print(result.Error)
	}
}

func main1(db *gorm.DB) {
	var pluginDir string
	var pidFile string
	flag.StringVar(&pluginDir, "pluginDir", "", "the directory holding your plugins")
	flag.StringVar(&pidFile, "pidFile", "", "pid file path")
	flag.Parse()

	absDir, err := filepath.Abs(pluginDir)
	if err != nil {
		panic(err)
	}
	if err := hutils.FindDirectory(absDir, ""); err != nil {
		panic(err)
	}
	if pidFile == "" {
		panic("no --pidFile")
	}

	pid := fmt.Sprint(os.Getpid())
	if err := ioutil.WriteFile(pidFile, []byte(pid), 0644); err != nil {
		panic(err)
	}

	g.PluginManagerSwapper = hotswap.NewPluginManagerSwapper(absDir,
		hotswap.WithLogger(g.Logger),
		hotswap.WithFreeDelay(time.Second*15),
	)
	swapper := g.PluginManagerSwapper
	details, err := swapper.LoadPlugins(nil)
	if err != nil {
		panic(err)
	} else if len(details) == 0 {
		panic("no plugin is found in " + absDir)
	} else {
		g.Logger.Infof("<hotswap> %d plugin(s) loaded. details: [%s]",
			len(details), details)
	}

	/*
		go func() {
			heartbeat()
			for range time.Tick(time.Second * 3) {
				heartbeat()
			}
		}()
	*/

	// Wait for signals
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

loop:
	for {
		select {
		case sig := <-chSignal:
			g.Logger.Infof("signal received: %v", sig)
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				break loop
			case syscall.SIGUSR1:
				g.Logger.Info("<hotswap> reloading...")
				details, err := swapper.Reload(nil)
				if err != nil {
					panic(err)
				} else if len(details) == 0 {
					g.Logger.Infof("no plugin is found in " + absDir)
				} else {
					g.Logger.Infof("<hotswap> %d plugin(s) loaded. details: [%s]",
						len(details), details)
				}
				//heartbeat()
			}
		}
	}

	signal.Reset(syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	g.Logger.Info("THE END")
}

func heartbeat(ctx *macaron.Context, db *gorm.DB) {
	repeat := rand.Intn(3) + 1
	g.PluginManagerSwapper.Current().FindPlugin("world").InvokeFunc("hum", repeat, ctx, db)
}
func heartbeat1(ctx *macaron.Context, db *gorm.DB) {
	repeat := rand.Intn(3) + 1
	g.PluginManagerSwapper.Current().FindPlugin("world").InvokeFunc("hum1", repeat, ctx, db)
}

func heartbeat2(ctx *macaron.Context, db *gorm.DB) {
	repeat := rand.Intn(3) + 1
	g.PluginManagerSwapper.Current().FindPlugin("world1").InvokeFunc("hum", repeat, ctx, db)
}

func serviceRunner(db *gorm.DB, job models.Job) {
	var ctx = macaron.Context{}
	p := &ctx
	//repeat := rand.Intn(3) + 1
	fmt.Println("Starting job ", job.Plugin, job.Service)
	g.PluginManagerSwapper.Current().FindPlugin(job.Plugin).InvokeFunc(job.Service, 5, p, db)
}

func handleTest(ctx *macaron.Context, db *gorm.DB) {
	heartbeat(ctx, db)
}
func handleTest1(ctx *macaron.Context, db *gorm.DB) {
	heartbeat1(ctx, db)
}

func handleTest2(ctx *macaron.Context, db *gorm.DB) {
	heartbeat2(ctx, db)
}

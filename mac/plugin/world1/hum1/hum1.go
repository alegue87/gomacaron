package hum1

import (
	"github.com/edwingeng/hotswap/demo/hello/g"

	"fmt"

	"mac/models"

	"gorm.io/gorm"
)

func Hum2(pluginName string, compileTimeString string, db *gorm.DB, job models.Job) {
	//str := strings.TrimSpace(strings.Repeat("job hum ", repeat))
	g.Logger.Infof("<%s.%s> %s. reloadCounter: %v",
		pluginName, compileTimeString, job.Uuid, g.PluginManagerSwapper.ReloadCounter())

	fmt.Print("!!?")
}

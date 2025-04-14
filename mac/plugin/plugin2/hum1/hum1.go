package hum1

import (
	"mac/models"

	"github.com/edwingeng/hotswap/demo/hello/g"

	"fmt"

	"gorm.io/gorm"
)

func Hum1(pluginName string, compileTimeString string, db *gorm.DB, job models.Job) {
	//str := strings.TrimSpace(strings.Repeat("hum1 ", repeat))
	g.Logger.Infof("<%s.%s> %s. reloadCounter: %v",
		pluginName, compileTimeString, job.Plugin, g.PluginManagerSwapper.ReloadCounter())

	fmt.Print("!!?...", db)
}

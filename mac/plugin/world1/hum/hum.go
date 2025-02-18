package hum

import (
	"strings"

	"github.com/edwingeng/hotswap/demo/hello/g"

	"fmt"

	"gopkg.in/macaron.v1"
	"gorm.io/gorm"
)

func Hum(pluginName string, compileTimeString string, repeat int, ctx *macaron.Context, db *gorm.DB) {
	str := strings.TrimSpace(strings.Repeat("hum ", repeat))
	g.Logger.Infof("<%s.%s> %s. reloadCounter: %v",
		pluginName, compileTimeString, str, g.PluginManagerSwapper.ReloadCounter())

	res := (ctx).Query("ciao")

	fmt.Println(res)
	fmt.Print("!!?")
}

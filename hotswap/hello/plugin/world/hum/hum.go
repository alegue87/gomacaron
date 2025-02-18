package hum

import (
	"strings"

	"github.com/edwingeng/hotswap/demo/hello/g"
)

func Hum(pluginName string, compileTimeString string, repeat int, value int) {
	str := strings.TrimSpace(strings.Repeat("humciao ", repeat))
	g.Logger.Infof("<%s.%s> %s. reloadCounter: %v",
		pluginName, compileTimeString, str, g.PluginManagerSwapper.ReloadCounter())
	g.Logger.Infof("value %v", value)
}

package plugin1

import (
	"mac/mac/plugin/plugin1/hum"
	"mac/mac/plugin/plugin1/hum1"

	"github.com/edwingeng/hotswap/demo/hello/g"
	"github.com/edwingeng/hotswap/vault"

	"gopkg.in/macaron.v1"
	"gorm.io/gorm"

	"mac/models"
)

const (
	pluginName = "plugin1"
)

var (
	CompileTimeString string
)

func OnLoad(data interface{}) error {
	g.Logger.Infof("<%s.%s> OnLoad", pluginName, CompileTimeString)
	return nil
}

func OnInit(sharedVault *vault.Vault) error {
	g.Logger.Infof("<%s.%s> OnInit", pluginName, CompileTimeString)
	return nil
}

func OnFree() {
	g.Logger.Infof("<%s.%s> OnFree", pluginName, CompileTimeString)
}

func Export() interface{} {
	g.Logger.Infof("<%s.%s> Export", pluginName, CompileTimeString)
	return nil
}

func Import() interface{} {
	g.Logger.Infof("<%s.%s> Import", pluginName, CompileTimeString)
	return nil
}

func InvokeFunc(name string, params ...interface{}) (interface{}, error) {
	switch name {
	case "hum":
		repeat := params[0].(int)
		ctx := params[1].(*macaron.Context)
		db := params[2].(*gorm.DB)
		hum.Hum(pluginName, CompileTimeString, repeat, ctx, db)

	case "hum1":
		db := params[0].(*gorm.DB)
		job := params[1].(models.Job)
		hum1.Hum2(pluginName, CompileTimeString, db, job)
	}

	return nil, nil
}

func Reloadable() bool {
	g.Logger.Infof("<%s.%s> Reloadable", pluginName, CompileTimeString)
	return true
}

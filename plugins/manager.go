package plugins

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

func GetPluginManager() Manager {
	if reflect.DeepEqual(pluginManager, Manager{}) {
		log.Info("START:LOAD_PLUGINS")
		data, err := os.ReadFile("plugins.yaml")
		if err != nil {
			log.Fatal(err.Error())
		}
		err = yaml.Unmarshal(data, &pluginManager)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Info("FINISHED:LOAD_PLUGINS", pluginManager.Plugins)
	}
	return pluginManager
}

func (m Manager) IsPluginEnabled(name string) bool {
	for _, plugin := range m.Plugins {
		if plugin.Name == name && plugin.Enabled {
			return true
		}
	}
	return false
}

package configuration

import (
	"github.com/dlintw/goconf"
	"sync"
)

type Config struct {
	configFile *goconf.ConfigFile
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		configFile, err := goconf.ReadConfigFile("config.config")
		println(configFile.GetSections())
		if err != nil {
			panic(err)
		}

		instance = &Config{configFile}
	})
	return instance
}

func (instance *Config) GetFstComponentFileName() string {
	return instance.readValue("fst-component-file")
}

func (instance *Config) GetSndComponentFileName() string {
	return instance.readValue("snd-component-file")
}

func (instance *Config) readValue(valueName string) string {
	value, err := instance.configFile.GetString("default", valueName)
	if err != nil {
		panic(err)
	}
	return value
}

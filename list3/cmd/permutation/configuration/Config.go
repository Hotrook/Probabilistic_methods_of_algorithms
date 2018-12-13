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

		if err != nil {
			panic(err)
		}

		instance = &Config{configFile}
	})
	return instance
}

func (instance *Config) GetDefaultN() int {
	return instance.readIntValue("n")
}

func (instance *Config) GetDefaultStart() int {
	return instance.readIntValue("start")
}

func (instance *Config) GetDefaultStep() int {
	return instance.readIntValue("step")
}

func (instance *Config) GetDefaultProbes() int {
	return instance.readIntValue("probes")
}

func (instance *Config) readStringValue(propertyName string) string {
	value, err := instance.configFile.GetString("default", propertyName)
	if err != nil {
		panic(err)
	}
	return value
}

func (instance *Config) readIntValue(propertyName string) int {
	value, err := instance.configFile.GetInt("default", propertyName)
	if err != nil {
		panic(err)
	}
	return value
}

func (instance *Config) GetFilename() string {
	return instance.readStringValue("filename")
}

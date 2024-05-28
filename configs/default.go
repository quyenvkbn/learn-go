package configs

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigJson map[string]interface{}

type LoadConfig struct {
}

var configData map[string]interface{}

func LoadConfigJson(path string) ConfigJson {

	if len(configData) > 0 {
		return configData
	} else {
		viper.AddConfigPath(path)
		viper.SetConfigType("json")
		viper.SetConfigName("configs")
		viper.AutomaticEnv()
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Could not load environment variables", err)
		}
		err = viper.Unmarshal(&configData)
		return configData
	}
}

func GetConfigByKey(key string) map[string]interface{} {
	return configData[key].(map[string]interface{})
}

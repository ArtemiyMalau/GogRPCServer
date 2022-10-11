package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type config struct {
	GRPC struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"grpc"`
}

var configInstance *config
var once sync.Once

func GetConfig() *config {
	once.Do(func() {
		configInstance = &config{}
		if err := cleanenv.ReadConfig("config.yml", configInstance); err != nil {
			help, _ := cleanenv.GetDescription(configInstance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return configInstance
}

package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type config struct {
	Environment string `yaml:"environment"`
	Listen      struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
	} `yaml:"listen"`
	MongoDB struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Database    string `yaml:"database"`
		AuthDB      string `yaml:"auth_db"`
		Collections struct {
			Product  string `yaml:"product"`
			Document string `yaml:"document"`
		} `yaml:"collections"`
	} `yaml:"mongodb"`
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

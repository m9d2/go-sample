package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

var Conf = Config{}

func InitConfig() {
	file, _ := os.ReadFile("./config.yml")

	err := yaml.Unmarshal(file, &Conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

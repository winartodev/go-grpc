package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	TodoList struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"todolist"`

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Driver   string `yaml:"driver"`
	} `yaml:"database"`
}

func (c *Config) GetConfig() *Config {

	yamlFile, err := ioutil.ReadFile("file/dev.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

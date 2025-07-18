package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// default config
var conf = &Config{
	App: &App{
		Host: "127.0.0.1",
		Port: 8080,
	},
	MySQL: &MySQL{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "123456",
		Database: "test",
		Debug:    true,
	},
}

// get config
func GetConfig() *Config {
	if conf == nil {
		fmt.Println("Failed to get config")
		os.Exit(1)
	}
	return conf
}

// load config
func LoadConfigFromYaml(configFilePath string) error {
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Failed to read config file", err)
		return err
	}
	yaml.Unmarshal(content, conf)
	return nil
}

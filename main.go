package main

import (
	"fmt"
	"os"

	"github.com/XieChengG/bookapi/api"
	"github.com/XieChengG/bookapi/config"
	"github.com/XieChengG/bookapi/model"
	"github.com/gin-gonic/gin"
)

func main() {
	// load config
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		configFilePath = "application.yaml"
	}
	if err := config.LoadConfigFromYaml(configFilePath); err != nil {
		fmt.Println("Failed to load config")
		os.Exit(1)
	}
	conf := config.GetConfig()

	// initial db
	db := conf.MySQL.DB()
	db.AutoMigrate(&model.Book{})

	// create a gin server
	server := gin.Default()

	api.NewBookApiHander().Registry(server)

	// run server
	if err := server.Run(conf.App.Address()); err != nil {
		panic(err)
	}
}

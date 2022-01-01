package main

import (
	"flag"
	"log"

	"social/internal/app"
	"social/internal/config"
	"social/internal/domain/user"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", config.DefaultConfigPath, "Path to configuration file")
	flag.Parse()
}

func main() {
	cfg, err := config.New(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	application := app.New(cfg)
	err = application.Init()
	if err != nil {
		log.Fatalln(err)
	}

	db := application.DB()
	err = db.AutoMigrate(user.User{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("success")
}

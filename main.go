package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kwong21/graphql-go-cardkeeper/models"
	"github.com/kwong21/graphql-go-cardkeeper/server"
	"github.com/kwong21/graphql-go-cardkeeper/service"
)

func main() {
	config := parseConfig()

	service := service.New(config)
	server.Init(service)
}

func parseConfig() models.Config {
	var conf models.Config

	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		log.Fatalf("Unable to read configuration file: %s", err)
	}

	return conf
}

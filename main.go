package main

import (
	"log"

	"github.com/denilbhatt0814/email-scheduler/config"
	"github.com/denilbhatt0814/email-scheduler/internal/api"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalln("Config file is not loaded properly", err)
	}
	api.StartServer(cfg)
}

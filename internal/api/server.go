package api

import (
	"log"

	"github.com/denilbhatt0814/email-scheduler/config"
	"github.com/denilbhatt0814/email-scheduler/internal/api/rest"
	"github.com/denilbhatt0814/email-scheduler/internal/api/rest/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Database connection error: ", err)
	}

	log.Println("Database connected")
	// run migrations // TODO:
	err = db.AutoMigrate()
	if err != nil {
		log.Fatalf("Error on running migration: %v", err.Error())
	}
	log.Println("migration was succefull")

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Config: config,
	}
	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupScheduleHandler(rh)
}

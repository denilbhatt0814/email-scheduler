package api

import (
	"log"

	"github.com/denilbhatt0814/email-scheduler/config"
	"github.com/denilbhatt0814/email-scheduler/internal/api/rest"
	"github.com/denilbhatt0814/email-scheduler/internal/api/rest/handlers"
	"github.com/denilbhatt0814/email-scheduler/internal/domain"
	"github.com/denilbhatt0814/email-scheduler/internal/service"
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
	// run migrations
	err = db.AutoMigrate(&domain.ScheduledEmail{})
	if err != nil {
		log.Fatalf("Error on running migration: %v", err.Error())
	}
	log.Println("migration was succefull")

	// initalizing cron
	cron := service.NewCronService()

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Cron:   cron,
		Config: config,
	}
	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupEmailEmailScheduleHandler(rh)
}

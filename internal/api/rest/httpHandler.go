package rest

import (
	"github.com/denilbhatt0814/email-scheduler/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App    *fiber.App
	DB     *gorm.DB
	Config config.AppConfig
}

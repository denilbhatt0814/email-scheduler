package handlers

import (
	"github.com/denilbhatt0814/email-scheduler/internal/api/rest"
	"github.com/denilbhatt0814/email-scheduler/internal/repository"
	"github.com/denilbhatt0814/email-scheduler/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ScheduleHandler struct {
	svc service.ScheduleService
}

func SetupScheduleHandler(rh *rest.RestHandler) {
	app := rh.App

	svc := service.ScheduleService{
		Repo:   repository.NewScheduleRepository(rh.DB),
		Config: rh.Config,
	}

	handler := ScheduleHandler{
		svc: svc,
	}

	app.Post("/schedule-email", handler.ScheduleEmail)
	app.Get("/scheduled-emails", handler.GetScheduledEmails)
	app.Get("/scheduled-emails/:id", handler.GetScheduledEmail)
	app.Delete("/scheduled-emails/:id", handler.DeleteScheduledEmail)
}

func (h *ScheduleHandler) ScheduleEmail(ctx *fiber.Ctx) error {
	return nil
}
func (h *ScheduleHandler) GetScheduledEmails(ctx *fiber.Ctx) error {
	return nil
}
func (h *ScheduleHandler) GetScheduledEmail(ctx *fiber.Ctx) error {
	return nil
}
func (h *ScheduleHandler) DeleteScheduledEmail(ctx *fiber.Ctx) error {
	return nil
}

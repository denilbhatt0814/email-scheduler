package handlers

import (
	"net/http"
	"strconv"

	"github.com/denilbhatt0814/email-scheduler/internal/api/rest"
	"github.com/denilbhatt0814/email-scheduler/internal/dto"
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

	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"msg": "healthy",
		})
	})
	app.Post("/schedule-email", handler.ScheduleEmail)
	app.Get("/scheduled-emails", handler.GetScheduledEmails)
	app.Get("/scheduled-emails/:id", handler.GetScheduledEmail)
	app.Delete("/scheduled-emails/:id", handler.DeleteScheduledEmail)
}

func (h *ScheduleHandler) ScheduleEmail(ctx *fiber.Ctx) error {
	scheduledEmail := dto.CreateScheduledEmail{}
	err := ctx.BodyParser(&scheduledEmail)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	err = h.svc.ScheduleEmail(&scheduledEmail)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error scheduling email",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "email scheduled successfully",
	})
}

func (h *ScheduleHandler) GetScheduledEmails(ctx *fiber.Ctx) error {

	emails, err := h.svc.GetScheduledEmails()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error retrieving scheduled emails",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "scheduled emails retreived successfully",
		"data":    emails,
	})
}
func (h *ScheduleHandler) GetScheduledEmail(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	email, err := h.svc.GetScheduledEmail(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "error retrieving scheduled email",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "scheduled emails retreived successfully",
		"data":    email,
	})

}
func (h *ScheduleHandler) DeleteScheduledEmail(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	err := h.svc.DeleteScheduledEmail(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error deleting scheduled email",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "scheduled email deleted successfully",
	})
}

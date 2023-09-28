package handler

import (
	"project-prim-aor/struc"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) HomeHandler(c *fiber.Ctx) error {
	return c.JSON(struc.ResponseJson{
		Status: "OK",
		Data: struc.Student{
			ID:        1,
			Firstname: "Prim",
			Lastname:  "Aor",
			JoinDate:  time.Now(),
		},
	})
}

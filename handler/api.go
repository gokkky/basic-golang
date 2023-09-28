package handler

import (
	"project-prim-aor/struc"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) InsertHandler(c *fiber.Ctx) error {

	h.DB.Insert(struc.Student{
		ID:        1,
		Firstname: "test",
		Lastname:  "test",
		JoinDate:  time.Now(),
	})

	return c.JSON(struc.ResponseJson{
		Status: "OK",
	})
}

func (h *Handler) GetHandler(c *fiber.Ctx) error {

	stu := h.DB.Get()

	return c.JSON(struc.ResponseJson{
		Status: "OK",
		Data:   stu,
	})
}

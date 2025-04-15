package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/POSPharm/internal/usecase"
)

type Role struct {
	role usecase.Role
}

func NewRole(role usecase.Role) *Role {
	return &Role{role: role}
}

func (h *Role) List(c *fiber.Ctx) error {
	var limit = c.QueryInt("limit", 10)
	var offset = c.QueryInt("offset", 0)
	roles, err := h.role.List(limit, offset)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": roles})
}

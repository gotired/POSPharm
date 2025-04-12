package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/POSPharm/internal/usecase"
)

type Branch struct {
	branch usecase.Branch
}

func NewBranch(branch usecase.Branch) *Branch {
	return &Branch{branch: branch}
}

func (h *Branch) List(c *fiber.Ctx) error {
	var limit = c.QueryInt("limit", 10)
	var offset = c.QueryInt("offset", 0)
	branches, err := h.branch.List(limit, offset)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": branches})
}

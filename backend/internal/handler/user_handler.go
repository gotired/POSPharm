package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gotired/POSPharm/internal/domain"
	"github.com/gotired/POSPharm/internal/usecase"
)

type User struct {
	userUsecase usecase.User
}

func NewUser(userUsecase usecase.User) *User {
	return &User{userUsecase: userUsecase}
}

func (h *User) List(c *fiber.Ctx) error {
	users, err := h.userUsecase.GetUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": users})
}

func (h *User) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := h.userUsecase.GetUser(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func (h *User) Create(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.userUsecase.RegisterUser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(user)
}

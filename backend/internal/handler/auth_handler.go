package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/POSPharm/internal/domain"
	"github.com/gotired/POSPharm/internal/usecase"
	"gorm.io/gorm"
)

type Auth struct {
	user usecase.User
	auth usecase.Auth
}

func NewAuth(user usecase.User, auth usecase.Auth) *Auth {
	return &Auth{user: user, auth: auth}
}

func (h *Auth) Login(c *fiber.Ctx) error {
	var loginDetail struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginDetail); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	userDetail, err := h.user.GetUserLogin(loginDetail.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(401).JSON(fiber.Map{"error": "mismatch username or password"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if userDetail == nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	if !h.auth.ComparePassword(userDetail.PasswordHash, loginDetail.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "password is not correct"})
	}

	accessToken, err := h.auth.GenerateToken(userDetail.UserID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	encryptAccess, err := h.auth.EncryptToken(accessToken)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	refreshToken, err := h.auth.GenerateRefreshToken(userDetail.UserID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	encryptRefresh, err := h.auth.EncryptToken(refreshToken)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"access": encryptAccess, "refresh": encryptRefresh})
}

func (h *Auth) Register(c *fiber.Ctx) error {
	userDetail := domain.UserCreate{
		UserName:        c.FormValue("username"),
		FirstName:       c.FormValue("first_name"),
		LastName:        c.FormValue("last_name"),
		Email:           c.FormValue("email"),
		Tel:             c.FormValue("tel"),
		Password:        c.FormValue("password"),
		ConfirmPassword: c.FormValue("confirm_password"),
	}
	if err := h.user.ValidateRegister(userDetail); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := h.auth.HashPassword(userDetail.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.user.RegisterUser(userDetail, hashedPassword); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return h.Login(c)
}

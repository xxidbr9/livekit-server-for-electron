package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xxidbr9/livekit-server-for-electron/internal/config"
	"github.com/xxidbr9/livekit-server-for-electron/internal/livekit"
	"github.com/xxidbr9/livekit-server-for-electron/internal/service"
)

var tokenService *service.TokenService

// Initialize the token service
func init() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	tokenGenerator := livekit.NewTokenGenerator(cfg)
	tokenService = service.NewTokenService(tokenGenerator)
}

// GenerateToken generates a token for a participant to join a room
func GenerateToken(c *fiber.Ctx) error {
	var req livekit.TokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.RoomName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "room name is required",
		})
	}

	if req.Identity == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "identity is required",
		})
	}

	token, err := tokenService.GenerateToken(req.RoomName, req.Identity, req.CanPublish, req.CanSubscribe)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "token created",
		"data": fiber.Map{
			"token": token,
		},
	})
}

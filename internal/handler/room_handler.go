package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xxidbr9/livekit-server-for-electron/internal/config"
	"github.com/xxidbr9/livekit-server-for-electron/internal/livekit"
	"github.com/xxidbr9/livekit-server-for-electron/internal/service"
)

var roomService *service.RoomService

func init() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	tokenGenerator := livekit.NewTokenGenerator(cfg)
	roomGenerator := livekit.NewRoomGenerator(cfg, tokenGenerator)
	roomService = service.NewRoomService(roomGenerator)
}

func CreateRoom(c *fiber.Ctx) error {
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
	roomName := req.RoomName

	if req.Identity == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "identity name is required",
		})
	}
	identity := req.Identity

	token, err := roomService.GenerateRoom(roomName, identity)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "token with room created",
		"data": fiber.Map{
			"room": fiber.Map{
				"room_name": roomName,
			},
			"identity": identity,
			"token":    token,
		},
	})
}

// TODO: rest
// func GetRooms(c *fiber.Ctx) error {

// }

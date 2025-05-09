package router

import (
	"github.com/go-zoox/fetch"
	"github.com/gofiber/fiber/v2"
	"github.com/xxidbr9/livekit-server-for-electron/internal/config"
	"github.com/xxidbr9/livekit-server-for-electron/internal/handler"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(app *fiber.App) {
	// API group
	api := app.Group("/api")

	// Room routes
	roomRoutes := api.Group("/rooms")
	roomRoutes.Post("/", handler.CreateRoom)

	// roomRoutes.Get("/", handler.GetRooms)
	// roomRoutes.Get("/:id", handler.GetRoom)
	// roomRoutes.Delete("/:id", handler.DeleteRoom) // for force to close all participants

	// Token routes
	tokenRoutes := api.Group("/token")
	tokenRoutes.Post("/", handler.GenerateToken)

	// WebSocket route for real-time updates
	// app.Get("/ws", handler.HandleWebSocket)

	// Webhook route
	// app.Post("/webhook", handler.HandleWebhook)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		cfg, err := config.Load()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		response, err := fetch.Get(cfg.LiveKitURL)
		if err != nil {
			panic(err)
		}
		if response.StatusCode() != 200 {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}

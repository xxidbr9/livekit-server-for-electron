package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xxidbr9/livekit-server-for-electron/internal/handler"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(app *fiber.App) {
	// API group
	api := app.Group("/api")

	// Room routes
	// roomRoutes := api.Group("/rooms")
	// roomRoutes.Get("/", handler.GetRooms)
	// roomRoutes.Get("/:id", handler.GetRoom)
	// roomRoutes.Post("/", handler.CreateRoom)
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
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}

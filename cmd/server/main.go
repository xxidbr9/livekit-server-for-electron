package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/xxidbr9/livekit-server-for-electron/internal/config"
	"github.com/xxidbr9/livekit-server-for-electron/internal/router"
)

// Run starts the Fiber server
func Run() error {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "LiveKit Fiber App",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Setup routes
	router.SetupRoutes(app)

	// Start server
	serverAddr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Server starting on %s", serverAddr)
	return app.Listen(serverAddr)
}

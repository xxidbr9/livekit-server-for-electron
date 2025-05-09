package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	Port             int
	LiveKitURL       string
	LiveKitAPIKey    string
	LiveKitAPISecret string

	// Webhook key
	WebhookKey string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	port := 8080 // Default port
	if portStr := os.Getenv("PORT"); portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid PORT value: %w", err)
		}
	}

	livekitURL := os.Getenv("LIVEKIT_URL")
	if livekitURL == "" {
		livekitURL = "http://localhost:7880" // Default LiveKit URL
	}

	apiKey := os.Getenv("LIVEKIT_API_KEY")
	if apiKey == "" {
		apiKey = "devkey" // Default API key for development
	}

	apiSecret := os.Getenv("LIVEKIT_API_SECRET")
	if apiSecret == "" {
		apiSecret = "secret" // Default API secret for development
	}

	webhookKey := os.Getenv("WEBHOOK_KEY")
	if webhookKey == "" {
		webhookKey = "webhook_secret" // Default API secret for development
	}

	return &Config{
		Port:             port,
		LiveKitURL:       livekitURL,
		LiveKitAPIKey:    apiKey,
		LiveKitAPISecret: apiSecret,
		WebhookKey:       webhookKey,
	}, nil
}

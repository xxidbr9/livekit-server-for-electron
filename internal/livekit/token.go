package livekit

import (
	"time"

	"github.com/livekit/protocol/auth"
	"github.com/xxidbr9/livekit-server-for-electron/internal/config"
)

// TokenGenerator handles token generation for LiveKit
type TokenGenerator struct {
	config *config.Config
}

// NewTokenGenerator creates a new token generator
func NewTokenGenerator(cfg *config.Config) *TokenGenerator {
	return &TokenGenerator{
		config: cfg,
	}
}

// GenerateToken creates a new token for a participant to join a room
func (g *TokenGenerator) GenerateToken(roomName, identity string, canPublish, canSubscribe *bool) (string, error) {
	at := auth.NewAccessToken(g.config.LiveKitAPIKey, g.config.LiveKitAPISecret)
	grant := &auth.VideoGrant{
		RoomJoin:     true,
		Room:         roomName,
		CanPublish:   canPublish,
		CanSubscribe: canSubscribe,
	}

	// at.AddGrant(grant).
	// 	SetIdentity(identity).
	// 	SetValidFor(24 * time.Hour) // Token valid for 24 hours

	at.SetVideoGrant(grant).
		SetIdentity(identity).
		SetValidFor(24 * time.Hour)

	return at.ToJWT()
}

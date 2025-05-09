package service

import (
	"github.com/xxidbr9/livekit-server-for-electron/internal/livekit"
)

// TokenService handles token-related operations
type TokenService struct {
	tokenGenerator *livekit.TokenGenerator
}

// NewTokenService creates a new token service
func NewTokenService(tokenGenerator *livekit.TokenGenerator) *TokenService {
	return &TokenService{
		tokenGenerator: tokenGenerator,
	}
}

// GenerateToken creates a new token for a participant to join a room
func (s *TokenService) GenerateToken(roomName, identity string, canPublish, canSubscribe *bool) (string, error) {
	return s.tokenGenerator.GenerateToken(roomName, identity, canPublish, canSubscribe)
}

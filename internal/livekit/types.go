package livekit

import (
	"errors"
)

// Common errors
var (
	ErrRoomNotFound = errors.New("room not found")
)

// Room represents a LiveKit room
type Room struct {
	Name       string `json:"name"`
	NumClients int    `json:"numClients"`
	CreatedAt  int64  `json:"createdAt"`
}

// TokenRequest represents a request for a LiveKit token
type TokenRequest struct {
	RoomName     string `json:"roomName"`
	Identity     string `json:"identity"`
	CanPublish   *bool  `json:"canPublish"`
	CanSubscribe *bool  `json:"canSubscribe"`
}

// TokenResponse represents a response containing a LiveKit token
type TokenResponse struct {
	Token string `json:"token"`
}

package livekit

import (
	"errors"
)

// Common errors
var (
	ErrRoomNotFound = errors.New("room not found")
)

type RoomResponse struct {
	RoomName string `json:"room_name"`
	Token    string `token:"token"`
	Identity string `json:"identity"`
}

// TokenRequest represents a request for a LiveKit token
type TokenRequest struct {
	RoomName     string `json:"room_name"`
	Identity     string `json:"identity"`
	CanPublish   *bool  `json:"can_publish"`
	CanSubscribe *bool  `json:"can_subscribe"`
}

package livekit

import (
	"context"

	"github.com/livekit/protocol/livekit"
	livekitSdk "github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/xxidbr9/livekit-server-for-electron/internal/config"
)

/*TODO:
- [X] add room (create or join room if exist)
- [X] get all participant
- [X] remove room
- [X] list room
- [X] remove participant
*/

type TokenLiveKitInterface interface {
	GenerateToken(roomName, identity string, canPublish, canSubscribe *bool) (string, error)
}

type RoomGenerator struct {
	config         *config.Config
	TokenGenerator TokenLiveKitInterface
	roomClient     *lksdk.RoomServiceClient
	rooms          map[string]*livekitSdk.Room
}

func NewRoomGenerator(cfg *config.Config, tokenGenerator TokenLiveKitInterface) *RoomGenerator {
	roomClient := lksdk.NewRoomServiceClient(cfg.LiveKitURL, cfg.LiveKitAPIKey, cfg.LiveKitAPISecret)
	return &RoomGenerator{
		config:         cfg,
		TokenGenerator: tokenGenerator,
		roomClient:     roomClient,
	}
}

// Generate token for the room
func (r *RoomGenerator) GenerateRoomToken(roomName, identity string) (string, error) {
	canPublish := true
	canSubscribe := true
	token, err := r.TokenGenerator.GenerateToken(roomName, identity, &canPublish, &canSubscribe)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Add room or join room if exist
func (r *RoomGenerator) GenerateRoom(roomName, identity string) (string, error) {
	// Check if room exist
	if r.rooms[roomName] != nil {
		token, err := r.GenerateRoomToken(roomName, identity)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	room, err := r.roomClient.CreateRoom(context.Background(), &livekitSdk.CreateRoomRequest{
		Name: roomName,
		// MaxParticipants: 20,
		EmptyTimeout: 10 * 60, // 10 minutes
	})

	if err != nil {
		return "", err
	}

	// store room in r.rooms
	r.rooms[roomName] = room

	// generate token for the room
	token, err := r.GenerateRoomToken(roomName, identity)
	if err != nil {
		return "", err
	}

	return token, nil
}

// List room
func (r *RoomGenerator) ListRoom() ([]*livekitSdk.Room, error) {
	rooms, err := r.roomClient.ListRooms(context.Background(), &livekitSdk.ListRoomsRequest{})
	if err != nil {
		return nil, err
	}
	return rooms.Rooms, nil
}

// Get all participant
func (r *RoomGenerator) GetAllParticipant(roomName string) ([]*livekitSdk.ParticipantInfo, error) {
	participants, err := r.roomClient.ListParticipants(context.Background(), &livekit.ListParticipantsRequest{
		Room: roomName,
	})
	if err != nil {
		return nil, err
	}
	return participants.Participants, nil
}

// Remove room
func (r *RoomGenerator) RemoveRoom(roomName string) error {
	_, err := r.roomClient.DeleteRoom(context.Background(), &livekitSdk.DeleteRoomRequest{
		Room: roomName,
	})
	if err != nil {
		return err
	}
	return nil
}

// Remove participant
func (r *RoomGenerator) RemoveParticipant(roomName, identity string) error {
	_, err := r.roomClient.RemoveParticipant(context.Background(), &livekitSdk.RoomParticipantIdentity{
		Room:     roomName,
		Identity: identity,
	})
	if err != nil {
		return err
	}
	return nil
}

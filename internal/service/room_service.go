package service

import (
	livekitSdk "github.com/livekit/protocol/livekit"
	"github.com/xxidbr9/livekit-server-for-electron/internal/livekit"
)

type RoomService struct {
	roomGenerator *livekit.RoomGenerator
}

func NewRoomService(roomGenerator *livekit.RoomGenerator) *RoomService {
	return &RoomService{
		roomGenerator,
	}
}

func (s *RoomService) GenerateRoom(roomName, identity string) (string, error) {
	return s.roomGenerator.GenerateRoom(roomName, identity)
}

func (s *RoomService) GetRoomList() ([]*livekitSdk.Room, error) {
	return s.roomGenerator.ListRoom()
}

func (s *RoomService) RemoveRoom(roomName string) error {
	return s.roomGenerator.RemoveRoom(roomName)
}

func (s *RoomService) RemoveParticipant(roomName, identity string) error {
	return s.roomGenerator.RemoveParticipant(roomName, identity)
}

func (s *RoomService) GetAllParticipant(roomName string) ([]*livekitSdk.ParticipantInfo, error) {
	return s.roomGenerator.GetAllParticipant(roomName)
}

package convertors

import (
	"github.com/mksmstpck/telegram-miniapp/server/internal/dto"
	"github.com/mksmstpck/telegram-miniapp/server/internal/models"
)

func FriendDTOToDB(friend dto.Friend) models.Friend {
	return models.Friend(friend)
}

func FriendDBToDTO(friend models.Friend) dto.Friend {
	return dto.Friend(friend)
}

package convertors

import (
	"github.com/mksmstpck/telegram-miniapp/server/internal/dto"
	"github.com/mksmstpck/telegram-miniapp/server/internal/models"
)

func UserDTOToDB(user dto.User) models.User {
	return models.User(user)
}

func UserDBToDTO(user models.User) dto.User {
	return dto.User(user)
}

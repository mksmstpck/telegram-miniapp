package convertors

import (
	"github.com/mksmstpck/telegram-miniapp/server/internal/dto"
	"github.com/mksmstpck/telegram-miniapp/server/internal/models"
)

func TokenDBToDTO(token models.Token) dto.Token {
	return dto.Token(token)
}

func TokenDTOToDB(token dto.Token) models.Token {
	return models.Token(token)
}

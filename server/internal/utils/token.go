package utils

import (
	"github.com/mksmstpck/telegram-miniapp/server/internal/dto"
	"github.com/pborman/uuid"
)

func CreateToken(telegramID string) dto.Token {
	var token dto.Token

	token.Token = uuid.New()
	token.TelegramID = telegramID

	return token
}

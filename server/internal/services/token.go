package services

import (
	"github.com/mksmstpck/telegram-miniapp/server/internal/convertors"
	"github.com/mksmstpck/telegram-miniapp/server/internal/dto"
	"github.com/mksmstpck/telegram-miniapp/server/internal/utils"
)

func (s *Token) Create(telegramID string) (*dto.Token, error) {
	token := utils.CreateToken(telegramID)

	err := s.tDB.Create(convertors.TokenDTOToDB(token))
	if err != nil {
		return nil, err
	}

	return &token, err
}

func (s *Token) GetByToken(tokenStr string) (*dto.Token, error) {
	tokenDB, err := s.tDB.GetByToken(tokenStr)
	if err != nil {
		return nil, err
	}

	token := convertors.TokenDBToDTO(*tokenDB)

	return &token, nil
}

func (s *Token) GetByID(ID string) (*dto.Token, error) {
	tokenDB, err := s.tDB.GetByID(ID)
	if err != nil {
		return nil, err
	}

	token := convertors.TokenDBToDTO(*tokenDB)

	return &token, nil
}

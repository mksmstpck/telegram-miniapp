package services

import (
	"github.com/mksmstpck/telegram-miniapp/server/internal/db"
	"github.com/mksmstpck/telegram-miniapp/server/internal/dto"
)

type Token struct {
	tDB db.Tokens
}

func NewToken(db db.DB) Tokens {
	return &Token{
		tDB: db.Token,
	}
}

type Tokens interface {
	Create(string) (*dto.Token, error)
	GetByToken(tokenStr string) (*dto.Token, error)
	GetByID(id string) (*dto.Token, error)
}

type Services struct {
	Token Tokens
}

func NewServices(db db.DB) *Services {
	return &Services{
		Token: NewToken(db),
	}
}

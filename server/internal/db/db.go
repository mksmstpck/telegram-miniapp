package db

import (
	"github.com/jackc/pgx/v4"
	"github.com/mksmstpck/telegram-miniapp/server/internal/models"
)

type Token struct {
	conn *pgx.Conn
}

func NewToken(conn *pgx.Conn) Tokens {
	return &Token{
		conn: conn,
	}
}

type Tokens interface {
	Create(token models.Token) error
	GetByToken(tokenStr string) (*models.Token, error)
	GetByID(id string) (*models.Token, error)
}

type DB struct {
	Token Tokens
}

func NewDB(conn *pgx.Conn) *DB {
	return &DB{
		Token: NewToken(conn),
	}
}

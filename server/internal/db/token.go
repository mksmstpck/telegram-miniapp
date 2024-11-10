package db

import (
	"context"

	"github.com/mksmstpck/telegram-miniapp/server/internal/models"
	"github.com/sirupsen/logrus"
)

func (d *Token) Create(token models.Token) error {
	_, err := d.conn.Exec(context.Background(), "INSERT INTO token(id, token, telegramID) VALUES($1, $2, $3)", token.Token, token.TelegramID)
	if err != nil {
		logrus.Info(err)
		return err
	}

	return nil
}

func (d *Token) GetByToken(tokenStr string) (*models.Token, error) {
	var token models.Token
	err := d.conn.QueryRow(context.Background(), "SELECT * FROM tokens WHERE token=$1", tokenStr).Scan(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (d *Token) GetByID(id string) (*models.Token, error) {
	var token models.Token
	err := d.conn.QueryRow(context.Background(), "SELECT * FROM tokens WHERE id=$1", id).Scan(&token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

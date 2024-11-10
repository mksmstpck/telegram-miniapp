package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/telegram-miniapp/server/internal/config"
	"github.com/mksmstpck/telegram-miniapp/server/internal/services"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	config *config.Config
	s      *services.Services
}

func NewHandlers(config *config.Config, s *services.Services) *Handlers {
	return &Handlers{
		config: config,
		s:      s,
	}
}

func (h *Handlers) HandleAll() {
	r := gin.Default()

	r.POST("/token", h.verifyUserHandler)

	if err := r.Run(h.config.GinUrl); err != nil {
		logrus.Fatal(err)
	}
}

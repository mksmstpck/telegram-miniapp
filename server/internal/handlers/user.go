package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apimodels "github.com/mksmstpck/telegram-miniapp/server/internal/handlers/apiModels"
	"github.com/mksmstpck/telegram-miniapp/server/internal/utils"
)

func (h *Handlers) verifyUserHandler(c *gin.Context) {
	// Parse JSON from Vue front-end
	var user apimodels.UserValidate
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		return
	}

	botToken := "YOUR_TELEGRAM_BOT_TOKEN"
	if utils.VerifyTelegramAuth(user.User, user.Hash, botToken) {
		token, err := h.s.Token.Create(user.User.TelegramID)
		if err != nil {
			c.AbortWithStatusJSON(500, apimodels.Message{Message: err.Error()})
		}
		c.JSON(201, apimodels.Message{Message: token.Token})

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, nil)

	}
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mksmstpck/telegram-miniapp/server/internal/utils"
)

func verifyUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON from Vue front-end
	var req struct {
		User map[string]string `json:"user"`
		Hash string            `json:"hash"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	botToken := "YOUR_TELEGRAM_BOT_TOKEN"
	if utils.VerifyTelegramAuth(req.User, req.Hash, botToken) {
		// User is verified, proceed to create a session or JWT
		// Return success response
		w.WriteHeader(http.StatusOK)
	} else {
		// Verification failed
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

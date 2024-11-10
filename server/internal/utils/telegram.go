package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"

	apimodels "github.com/mksmstpck/telegram-miniapp/server/internal/handlers/apiModels"
)

func VerifyTelegramAuth(userData apimodels.UserTelegram, hash string, botToken string) bool {
	// Create a sorted string of userData as required by Telegram
	dataStrings := []string{
		"id=" + userData.TelegramID,
		"username=" + userData.Username,
		"first_name=" + userData.FirstName,
		"last_name=" + userData.LastName,
		"photo_url=" + userData.PhotoLink,
		"auth_date=" + userData.AuthDate,
	}
	sort.Strings(dataStrings)
	dataCheckString := strings.Join(dataStrings, "\n")

	// Create a hash using HMAC SHA-256 with bot token
	secretKey := sha256.Sum256([]byte(botToken))
	mac := hmac.New(sha256.New, secretKey[:])
	mac.Write([]byte(dataCheckString))
	expectedHash := hex.EncodeToString(mac.Sum(nil))

	// Compare expected hash to the hash provided by Telegram
	return expectedHash == hash
}

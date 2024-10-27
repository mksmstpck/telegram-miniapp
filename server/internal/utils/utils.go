package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

func VerifyTelegramAuth(userData map[string]string, hash string, botToken string) bool {
	// Create a sorted string of userData as required by Telegram
	var dataStrings []string
	for k, v := range userData {
		dataStrings = append(dataStrings, k+"="+v)
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

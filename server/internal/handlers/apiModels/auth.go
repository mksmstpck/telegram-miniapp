package apimodels

type UserValidate struct {
	User UserTelegram `json:"user"`
	Hash string       `json:"hash"`
}

type UserTelegram struct {
	TelegramID string `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	PhotoLink  string `json:"photo_url"`
	AuthDate   string `json:"auth_date"`
}

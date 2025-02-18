package models

// Profile – модель профиля пользователя
type Profile struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	FullName   string `json:"full_name"`
	Phone      string `json:"phone"`
	Company    string `json:"company"`
	Position   string `json:"position"`
	Requisites string `json:"requisites"`
	PhotoURL   string `json:"photo_url"`
}

package models

type Address struct {
	ID          int    `json:"id"`
	ProfileID   int    `json:"profile_id"`
	AddressType string `json:"address_type"` // "physical", "legal", ...
	AddressText string `json:"address_text"`
}

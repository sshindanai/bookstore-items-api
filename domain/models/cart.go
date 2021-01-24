package models

type Cart struct {
	Items      []BuyingItem `json:"items"`
	TotalPrice float32      `json:"total_price"`
}

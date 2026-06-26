package models

import "time"

type StockTransaction struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Username        string    `json:"username,omitempty"`
	ProductID       int       `json:"product_id"`
	ProductName     string    `json:"product_name,omitempty"`
	TransactionType string    `json:"transaction_type"`
	Quantity        int       `json:"quantity"`
	CreatedAt       time.Time `json:"created_at"`
}

package models

import "time"

type Transaction struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Amount    float64   `json:"amount"`
	OrderDate time.Time `json:"order_date"`
	Status    string    `json:"status"`
}

// Add any additional fields or methods as needed for Transaction model

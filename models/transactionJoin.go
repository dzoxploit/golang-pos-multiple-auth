package models

import "time"

type TransactionJoin struct {
	ID          uint      `json:"id"`
	ProductID   uint      `json:"-"`
	UserID      uint      `json:"-"`
	Product     string    `json:"product_name"`
	Username    string    `json:"username"`
	Quantity    int       `json:"quantity"`
	Amount      float64   `json:"amount"`
	OrderDate   time.Time `json:"order_date"`
	Status      string    `json:"status"`
}
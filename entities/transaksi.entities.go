package entities

import "time"

type Transaction struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Type        string    `json:"type"` // "income" atau "expense"
	Amount      float64   `json:"amount"`
	Buyer       *string   `json:"buyer"`
	Phone       *string   `json:"phone"`
	Address     *string   `json:"address"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

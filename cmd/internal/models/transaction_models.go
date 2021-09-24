package models

import "time"

type TransactionType string

const (
	TransferTransaction TransactionType = "transfer"
	RefillTransaction   TransactionType = "refill"
)

type Transaction struct {
	ID        string          `gorm:"column:id;type:varchar(255);primary_key" json:"id"`
	Value     int             `gorm:"column:value;type:bigint(20) NOT NULL;" json:"value"`
	Type      TransactionType `gorm:"column:type;type:varchar(100);not null" json:"type"`
	CreatedAt time.Time       `gorm:"column:created_at;type:datetime NOT NULL;default:CURRENT_TIMESTAMP" json:"created_at"`
}

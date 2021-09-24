package models

import "time"

// FOR JWT AUTH
//type AccessToken struct {
//	AccessToken string `json:"access_token"`
//	ExpiresAt   int64  `json:"expires_at"`
//}

type Token struct {
	Token     string    `gorm:"column:token;type:varchar(255);primary_key" json:"id"`
	UserID    int       `gorm:"column:user_id;type:bigint(20);not null" json:"value"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime NOT NULL;default:CURRENT_TIMESTAMP" json:"created_at"`
}

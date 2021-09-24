package models

import "time"

type User struct {
	ID        int       `gorm:"column:id;type:bigint(20) unsigned NOT NULL;primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"column:email;type:varchar(255);not null;unique" json:"email"`
	Hash      string    `gorm:"column:hash;type:varchar(255);not null;" json:"-"` // might be hash
	CreatedAt time.Time `gorm:"column:created_at;type:datetime NOT NULL;default:CURRENT_TIMESTAMP" json:"created_at"`
}

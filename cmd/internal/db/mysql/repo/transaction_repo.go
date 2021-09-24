package repo

import (
	"UserWallet2021_09_24/cmd/internal/models"
	"github.com/jinzhu/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	db.CreateTable(&models.Transaction{})
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		AutoMigrate(&models.Transaction{})

	return TransactionRepo{db: db}
}

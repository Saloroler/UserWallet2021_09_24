package repo

import (
	"UserWallet2021_09_24/cmd/internal/models"
	"UserWallet2021_09_24/cmd/pkg/logger"
	"github.com/jinzhu/gorm"
)

type TokenRepo struct {
	db *gorm.DB
}

func NewTokenRepo(db *gorm.DB) TokenRepo {
	db.CreateTable(&models.Token{})
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		AutoMigrate(&models.Token{})

	return TokenRepo{db: db}
}

func (r *TokenRepo) Create(token models.Token) error {
	err := r.db.Table("token").Create(&token).Error
	if err != nil {
		logger.LogError("Failed to create token err: ", err)
		return err
	}

	return err
}

func (r *TokenRepo) FindByUserID(userID int) (models.Token, error) {
	var token models.Token
	err := r.db.Table("token").Where("user_id = ?", userID).First(&token).Error
	if err != nil {
		logger.LogError("Failed to get user err: ", err)
		return models.Token{}, err
	}

	return token, nil
}

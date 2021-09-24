package repo

import (
	"UserWallet2021_09_24/cmd/internal/models"
	"UserWallet2021_09_24/cmd/pkg/logger"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	db.CreateTable(&models.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		AutoMigrate(&models.User{})

	return UserRepo{db: db}
}

func (r *UserRepo) Find(id string) (models.User, error) {
	var user = models.User{}
	err := r.db.Table("user").First(&user, id).Error
	if err != nil {
		logger.LogError("Failed to get user err: ", err)
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepo) Create(user models.User) (models.User, error) {
	result := r.db.Table("user").Create(&user) // ID will be inserted
	if result.Error != nil {
		logger.LogError("Failed to add user err: ", result.Error) // always print errors where they happened
		return models.User{}, result.Error
	}

	return user, nil
}

package repository

import (
	"gorm.io/gorm"
	"upgrade/internal/models"
)

type UserModel struct {
	Db *gorm.DB
}

func (m *UserModel) Create(user models.User) error {
	result := m.Db.Create(&user)
	return result.Error
}

func (m *UserModel) FindOne(telegramId int64) (*models.User, error) {
	existUser := models.User{}

	result := m.Db.First(&existUser, models.User{TelegramId: telegramId})

	if result.Error != nil {
		return nil, result.Error
	}

	return &existUser, nil
}

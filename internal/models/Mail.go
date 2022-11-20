package models

import "gorm.io/gorm"

type Mail struct {
	gorm.Model
    ID              uid       	`json:"id"`
    Letter          string      `json:"letter"`
    ReceivedDate    string      `json:"date"`
    User            User
}

type MailModel struct {
	Db *gorm.DB
}

func (m *MailModel) Create(mail Mail) error {

	result := m.Db.Create(&mail)

	return result.Error
}

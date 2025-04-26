package database

import (
	"log"

	"github.com/posiposi/project/backend/internal/orm/model"
	"gorm.io/gorm"
)

type DBManagerInterface interface {
	Seed() error
}

type DBManager struct {
	db *gorm.DB
}

func NewDBManager(d *gorm.DB) DBManagerInterface {
	return &DBManager{db: d}
}

func (m *DBManager) Seed() error {
	var count int64

	if err := m.db.Model(&model.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		log.Println("User Data already exists!")
		return nil
	}

	if err := m.seedUsers(); err != nil {
		return err
	}

	return nil
}

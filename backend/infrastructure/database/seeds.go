package database

import (
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/posiposi/project/backend/internal/orm/model"
)

func (m *DBManager) seedUsers() error {
	log.Println("User Seeder Start!")
	for i := 0; i < 3; i++ {
		id := uuid.NewString()
		name := "test_user" + strconv.Itoa(i)
		user := model.User{
			Id:        id,
			Name:      name,
			Email:     name + "@example.com",
			Password:  "password",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := m.db.FirstOrCreate(&user).Error; err != nil {
			log.Fatalf("Seeder error: %v", err)
			return err
		}
	}

	log.Println("User Seeder Done!")
	return nil
}

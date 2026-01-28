package config

import (
	"fmt"
	"log"
	"tasker_go/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=localhost user=tasker_user password=12345678 dbname=tasker port=5432 sslmode=disable",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	err = db.AutoMigrate(&models.Task{}, &models.User{}, &models.Judge{})
	if err != nil {
		log.Fatal("Migration error")
	}

	log.Println("PostgreSQL подключён")
	return db
}

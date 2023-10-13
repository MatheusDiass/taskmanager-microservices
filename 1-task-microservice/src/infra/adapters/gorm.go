package adapters

import (
	"fmt"
	"log"
	"task-microservice/src/main/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter() GormAdapter {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
		config.DatabaseServer,
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseName,
		config.DatabasePort)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return GormAdapter{db}
}

func (adapter GormAdapter) GetDb() *gorm.DB {
	return adapter.db
}

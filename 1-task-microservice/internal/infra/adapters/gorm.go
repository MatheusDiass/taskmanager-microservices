package adapters

import (
	"fmt"
	"log"

	"github.com/matheusdiass/taskmanager-microservices/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter() GormAdapter {
	configs.SetupEnvironmentVariables()

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
		configs.DatabaseServer,
		configs.DatabaseUser,
		configs.DatabasePassword,
		configs.DatabaseName,
		configs.DatabasePort)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return GormAdapter{db}
}

func (adapter GormAdapter) GetDb() *gorm.DB {
	return adapter.db
}

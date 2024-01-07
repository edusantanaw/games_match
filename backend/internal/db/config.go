package db

import (
	"fmt"
	"os"

	"github.com/edusantanaw/games_match.git/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

func GetConnection() *gorm.DB {
	if database == nil {
		OpenConnection()
	}
	return database
}

func OpenConnection() {
	var connectionString = os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		panic("Erro ao conectar com o banco de dados!")
	}
	db.AutoMigrate(&entities.Game{})
	database = db
}

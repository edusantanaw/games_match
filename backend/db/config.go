package db

import (
	"github.com/edusantanaw/games_match.git/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

func getConnection() *gorm.DB {
	if database == nil {
		openConnection()
	}
	return database
}

func openConnection() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar com o banco de dados!")
	}
	db.AutoMigrate(&entities.Game{})
	database = db
}

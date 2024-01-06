package db

import (
	"fmt"

	"github.com/edusantanaw/games_match.git/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

const host = "localhost:5432"
const databaseName = "game"
const username = "postgres"
const password = "eduardo123"

const connectionString = "postgres://" + username + ":" + password + "@" + host + "/" + databaseName

func GetConnection() *gorm.DB {
	if database == nil {
		OpenConnection()
	}
	return database
}

func OpenConnection() {
	fmt.Println(connectionString)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		panic("Erro ao conectar com o banco de dados!")
	}
	db.AutoMigrate(&entities.Game{})
	database = db
}

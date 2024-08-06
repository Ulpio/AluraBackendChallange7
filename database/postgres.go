package database

import (
	"github.com/Ulpio/Alura_challange_BackEnd_7/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar no DB")
	}
	db.AutoMigrate(&models.Depoimentos{})
	db.AutoMigrate(&models.Destinos{})
	DB = db
}

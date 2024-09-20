package database

import (
	"fmt"

	"github.com/bmendescrypex/CRUD_g/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=db user=root password=BM2106 dbname=crud_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro ao conectar ao db:", err)
	}

	DB.AutoMigrate(&models.User{})

	if err == nil {
		fmt.Println("Conectado ao db com sucesso")
	}
}

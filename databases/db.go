package databases

import (
	"log"

	"github.com/sirakytk/api-with-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() {
	dsn := "host=localhost user=postgres password=postgres404 dbname=go_jwt_auth port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("koneksi database gagal:")
	}

	DB = db
	db.AutoMigrate(&models.User{})
}

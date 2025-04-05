package gorm

import (
	"log"
	"os"

	models "github.com/Natan-Lucena/products-api/internal/infraestructure/persistence"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar arquivo .env: %v", err)
    }
	key := "DB_URL"
	dsn := os.Getenv(key)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	err = db.AutoMigrate(&models.ProductModel{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
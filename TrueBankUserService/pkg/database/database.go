package database

import (
	"TrueBankUserService/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func InitDb() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	log.Println("Database connected:", Db != nil)

	dns := os.Getenv("DB_DNS")

	var err error
	Db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := Db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}

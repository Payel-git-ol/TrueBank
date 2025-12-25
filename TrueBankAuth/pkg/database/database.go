package database

import (
	"TrueBankAuth/pkg/models"
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

	dsn := os.Getenv("DB_DNS")
	if dsn == "" {
		log.Fatalf("DB_DNS not found in environment")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	Db = db

	if err := Db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("migration failed:", err)
	}

	connectDB, err := Db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := connectDB.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}
}

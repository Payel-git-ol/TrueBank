package database

import (
	"TrueBankTransactionService/pkg/models/dbModels"
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

	if os.Getenv("RUN_MIGRATION") == "true" {
		if err := Db.AutoMigrate(&dbModels.ListTransaction{}, &dbModels.HistoryTransaction{}, &dbModels.RemittanceHistory{}); err != nil {
			log.Fatalf("migration failed: %v", err)
		}
	}
}

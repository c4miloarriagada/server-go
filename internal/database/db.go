package db

import (
	"log"
	"os"
	"serverpackage/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserFile struct {
	Users []models.User
}

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dsn := "host=127.0.0.1 user=" + dbUser + " password=" + dbPassword + " dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	var dropModels = []interface{}{&models.User{}, &models.Account{}, &models.Transactions{}, &models.Type{}}

	db.Migrator().DropTable(dropModels...)

	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		return
	}

	log.Print("*******DB READY*******")

	DB = db
}

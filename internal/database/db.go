package db

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	UserID    int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Image     string    `json:"image"`
}

type Account struct {
	AccountID int     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int     `json:"userid" gorm:"foreignKey:UserRefer"`
	Balance   float64 `json:"balance"`
}

type Transactions struct {
	TrasnctionID int       `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID    int       `json:"accountid" gorm:"foreignKey:AccountRefer"`
	TimeStamp    time.Time `json:"timestamp"`
	Descriptions string    `json:"descriptions"`
	Total        float64   `json:"total"`
	Type         int       `json:"typeid" gorm:"foreignKey:"`
}

type Type struct {
	TypeID int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Type   string `json:"type" `
}

type UserFile struct {
	Users []User
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
	var models = []interface{}{&User{}, &Account{}, &Transactions{}, &Type{}}

	db.Migrator().DropTable(models...)
	if err != nil {
		return
	}

	err = db.AutoMigrate(&User{})

	if err != nil {
		return
	}

	log.Print("*******DB READY*******")

	DB = db
}

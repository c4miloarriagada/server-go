package models

import "time"

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

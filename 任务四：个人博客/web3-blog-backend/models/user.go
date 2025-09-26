package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id;"`
	UserName  string    `gorm:"not null;unique" json:"user_name"`
	Password  string    `json:"-" gorm:"not null;column:password;type:longtext"`
	Emall     string    `gorm:"not null;unique"`
	CreatedAt time.Time `json:"created-at"`
	Posts     []Post    `gorm:"foreignKey:UserID"`
	Comments  []Comment `gorm:"foreignKey:UserID"`
}

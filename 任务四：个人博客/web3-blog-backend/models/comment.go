package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"not null" json:"content"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID"`
	CreatedAt time.Time `json:"created-at"`
}

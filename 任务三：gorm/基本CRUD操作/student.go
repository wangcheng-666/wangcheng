package models

import "gorm.io/gorm"

type Student struct {
	Id      uint `gorm:"primary_key"`
	Name    string
	Age     int
	Grade   string
	DeletAt gorm.DeletedAt `gorm:"index"`
}

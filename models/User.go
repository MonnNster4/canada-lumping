package models

import "gorm.io/gorm"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
	Name     string
	Type     string
	Email     string
	Deleted  gorm.DeletedAt
}

package models

import "gorm.io/gorm"

type Department struct {
	ID       			uint `gorm:"primaryKey"`
	User 	 			User
	UserID   			string
	Name				string
	Code				string
	Deleted  			gorm.DeletedAt
}

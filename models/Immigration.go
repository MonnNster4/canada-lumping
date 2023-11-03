package models

import "gorm.io/gorm"

type Immigration struct {
	ID       			uint `gorm:"primaryKey"`
	Country 			string
	Name     			string
	Email    			string
	Message		   	 	string
	Deleted  			gorm.DeletedAt
}

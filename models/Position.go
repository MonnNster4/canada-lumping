package models

import "gorm.io/gorm"

type Position struct {
	ID       			uint `gorm:"primaryKey"`
	User 	 			User
	UserID   			string
	Title     			string
	Salary     			string
	Location    	 	string
	Description     	string
	Responsibility      string
	Requirements		string
	LastUpdate 			string
	Country				string
	Count 				string
	Type 				string
	Status				string	
	Deleted  			gorm.DeletedAt
}

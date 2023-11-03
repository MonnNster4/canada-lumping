package models

import "gorm.io/gorm"

type Supervisor struct {
	ID       			uint `gorm:"primaryKey"`
	User 	 			User
	UserID   			string
	Department 	 		Department
	DepartmentID   		string
	Name				string
	Deleted  			gorm.DeletedAt
}

package models

import "gorm.io/gorm"

type Applicant struct {
	ID       			uint `gorm:"primaryKey"`
	Position 	 		Position
	PositionID   		string
	Name     			string
	Email    			string
	Link		   	 	string
	CV			     	string
	Letter		        string
	Status				string
	Deleted  			gorm.DeletedAt
}

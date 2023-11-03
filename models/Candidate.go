package models

import "gorm.io/gorm"

type Candidate struct {
	ID       			uint `gorm:"primaryKey"`
	Applicant 	 		Applicant
	ApplicantID   		string
	User				User
	UserID				string	
	Status				string
	LastUpdate 			string
	Deleted  			gorm.DeletedAt
}

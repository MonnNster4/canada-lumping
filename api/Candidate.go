package api

import (
	"net/http"
	"time"
	"strconv"

	"github.com/dafalo/Mobile-Tracker-App/models"
)


func AddCandidate(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	position := models.Applicant{}
	id:= r.FormValue("id")
	user, _ := r.Cookie("id")
	now := time.Now()
	date := now.Format("2006-01-02")
	datetime := now.Format("15:04:05")


	db.Where("id", id).Find(&position)

		position.Status = "Transfered"
		db.Save(&position)

		candidate := models.Candidate{}

		candidate.ApplicantID = id
		candidate.UserID = user.Value
		candidate.Status = "Pending"
		candidate.LastUpdate = date + " " + datetime
		db.Save(&candidate)
}

func GetCandidate(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Candidate{}
	db.Preload("User").Preload("Applicant").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DeleteCandidate(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Applicant{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


func EditCandidate(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	position := models.Candidate{}
	status:= r.FormValue("status")
	
	db.Where("id", id).Find(&position)
		position.Status = status
		db.Save(&position)
}








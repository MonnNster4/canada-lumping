package api

import (
	"net/http"
	"time"
	"strconv"

	"github.com/dafalo/Mobile-Tracker-App/models"
)


func AddPosition(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	position := models.Position{}
	location:= r.FormValue("location")
	title:= r.FormValue("title")
	salary:= r.FormValue("salary")
	desc:= r.FormValue("desc")
	resp := r.FormValue("resp")
	number:= r.FormValue("number")
	ptype := r.FormValue("type")
	status := r.FormValue("status")
	recquire := r.FormValue("recquire")
	country := r.FormValue("country")
	user, _ := r.Cookie("id")
	now := time.Now()
	date := now.Format("2006-01-02")
	datetime := now.Format("15:04:05")

		position.Location = location
		position.Title = title
		position.Salary = salary
		position.Description = desc
		position.Responsibility = resp
		position.Count = number
		position.Type = ptype
		position.Status = status
		position.Requirements = recquire
		position.Country = country
		position.UserID = user.Value
		position.LastUpdate = date + " " + datetime
		db.Save(&position)
}

func GetPosition(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Position{}
	db.Preload("User").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditPosition(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	position := models.Position{}
	location:= r.FormValue("location")
	title:= r.FormValue("title")
	salary:= r.FormValue("salary")
	desc:= r.FormValue("desc")
	resp := r.FormValue("resp")
	number:= r.FormValue("number")
	ptype := r.FormValue("type")
	status := r.FormValue("status")
	recquire := r.FormValue("recquire")
	country := r.FormValue("country")
	user, _ := r.Cookie("id")
	now := time.Now()
	date := now.Format("2006-01-02")
	datetime := now.Format("15:04:05")

	db.Where("id", id).Find(&position)

		position.Location = location
		position.Title = title
		position.Salary = salary
		position.Description = desc
		position.Responsibility = resp
		position.Count = number
		position.Type = ptype
		position.Status = status
		position.Requirements = recquire
		position.Country = country
		position.UserID = user.Value
		position.LastUpdate = date + " " + datetime
		db.Save(&position)
	
	

}

func DeletePosition(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Position{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}








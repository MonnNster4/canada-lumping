package api

import (
	"net/http"
	"time"
	"strconv"
	"math/rand"
	"fmt"

	"github.com/dafalo/Mobile-Tracker-App/models"
)


func AddDepartment(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	position := models.Department{}
	name:= r.FormValue("name")
	user, _ := r.Cookie("id")
	rand.Seed(time.Now().UnixNano())

		position.Name = name
		position.Code = fmt.Sprint(rand.Intn(99999))
		position.UserID = user.Value
	
		db.Save(&position)
}

func GetDepartment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Department{}
	db.Preload("User").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditDepartment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	position := models.Department{}
	name:= r.FormValue("name")
	
	db.Where("id", id).Find(&position)
		position.Name = name
		db.Save(&position)
}

func DeleteDepartment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Department{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}








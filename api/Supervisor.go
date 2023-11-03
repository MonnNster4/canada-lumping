package api

import (
	"net/http"
	"strconv"
	"net/smtp"
	"fmt"
	
	"github.com/dafalo/Mobile-Tracker-App/models"
)


func AddSupervisor(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	position := models.Supervisor{}
	name:= r.FormValue("name")
	id:= r.FormValue("id")
	user, _ := r.Cookie("id")
	
		position.DepartmentID = id
		position.Name = name
		position.UserID = user.Value
	
		db.Save(&position)
}

func GetSupervisor(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Supervisor{}
	db.Preload("User").Preload("Department").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditSupervisor(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	position := models.Supervisor{}
	name := r.FormValue("name")
	depart := r.FormValue("depart")
	user, _ := r.Cookie("id")
	db.Where("id", id).Find(&position)
	position.UserID = user.Value
		position.Name = name
		position.DepartmentID = depart
		db.Save(&position)
}

func DeleteSupervisor(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Supervisor{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


func ImmigrationEmail(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	position := models.Immigration{}
	name:= r.FormValue("name")
	email:= r.FormValue("email")
	message:= r.FormValue("message")
	country:= r.FormValue("country")
	part := r.FormValue("part")
	
		position.Name = name
		position.Email = email
		position.Message = message
		position.Country = country

		db.Save(&position)

		if part == "Canada"{
			from := "groupnb23@gmail.com"
			pass := "dwxlbnyzqaxzjtbx"
			to := "info@groupnb.ca"

			msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Immigration \n\n" +
			"Good day,\n\n Email Details \n\n" +
			"Country:" + country + "\n\n" +
			"Name:" + name + "\n\n" +
			"Email:" + email + "\n\n" +
			"Message:" + message + "\n\n" +
			"Thank You"
			

		data := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))

		if data != nil {
			fmt.Printf("smtp error: %s", data)
			return
		}
		fmt.Println("Successfully sended to " + to)

		}else if part == "Usa"{
		from := "groupnb23@gmail.com"
		pass := "dwxlbnyzqaxzjtbx"
		to := "info@groupnb-usa.com"

		msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Immigration \n\n" +
		"Good day,\n\n Email Details \n\n" +
		"Country:" + country + "\n\n" +
		"Name:" + name + "\n\n" +
		"Email:" + email + "\n\n" +
		"Message:" + message + "\n\n" +
		 "Thank You"
		

	data := smtp.SendMail("smtp.gmail.com:587",
	smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if data != nil {
		fmt.Printf("smtp error: %s", data)
		return
	}
	fmt.Println("Successfully sended to " + to)
		}else if part == "Immigration"{
			from := "groupnb23@gmail.com"
			pass := "dwxlbnyzqaxzjtbx"
			to := "immigration@groupnb.ca"

			msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Immigration \n\n" +
			"Good day,\n\n Email Details \n\n" +
			"Country:" + country + "\n\n" +
			"Name:" + name + "\n\n" +
			"Email:" + email + "\n\n" +
			"Message:" + message + "\n\n" +
			 "Thank You"
			
	
		data := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))
	
		if data != nil {
			fmt.Printf("smtp error: %s", data)
			return
		}
		fmt.Println("Successfully sended to " + to)
		}
		


}


func GetEmail(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Immigration{}
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}








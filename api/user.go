package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/dafalo/Mobile-Tracker-App/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := models.User{}
	name := r.FormValue("name")
	password := r.FormValue("password")
	username := r.FormValue("username")
	email := r.FormValue("email")
	position := r.FormValue("position")

	contact.Name = name
	contact.Password = hashPassword(password)
	contact.Username = username
	contact.Email = email
	contact.Type = position

	db.Save(&contact)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func ChangePassword(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	product := models.User{}
	id, _ := strconv.Atoi(r.FormValue("id"))
	password := r.FormValue("password")


	db.Where("id", id).Find(&product)
	product.Password = hashPassword(password)
	db.Save(&product)

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := []models.User{}
	db.Find(&contact)

	data := map[string]interface{}{
		"status": "ok",
		"item":   contact,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetApplicant(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := []models.Applicant{}
	db.Preload("Position").Find(&contact)

	data := map[string]interface{}{
		"status": "ok",
		"item":   contact,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func Login(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := models.User{}
	report := []models.User{}

	db.Where("username = ?", username).Find(&user)
	db.Where("username = ?", username).Find(&report)

	println(username, user.ID)

	if CheckPasswordHash(password, user.Password) {
		result := "1"

		newSession := uuid.NewString()

		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "session",
			Value: newSession,
		})

		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "id",
			Value: fmt.Sprint(user.ID),
		})
		data := map[string]interface{}{
			"status":  "ok",
			"results": result,
			"reports": report,
		}
		ReturnJSON(w, r, data)
	} else {
		result := "0"
		data := map[string]interface{}{
			"status":  "error",
			"results": result,
		}
		ReturnJSON(w, r, data)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()


}

func EditUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	product := models.User{}
	name := r.FormValue("name")
	username := r.FormValue("username")
	email := r.FormValue("email")
	position := r.FormValue("position")

	db.Where("id", id).Find(&product)
	product.Name = name
	product.Username = username
	product.Email = email
	product.Type = position
	db.Save(&product)
	

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.User{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func ResetPassword(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	product := models.User{}
	id, _ := strconv.Atoi(r.FormValue("id"))


	db.Where("id", id).Find(&product)
	product.Password = hashPassword("password123")
	db.Save(&product)

}



func GetActiveSession(r *http.Request) string {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	return ""
}

func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func GormDB() *gorm.DB {
	dsn := "root:GroupNB2023@tcp(127.0.0.1:3306)/groupnb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}

func ReturnJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		response := map[string]interface{}{
			"status":    "error",
			"error_msg": fmt.Sprintf("unable to encode JSON. %s", err),
		}
		b, _ = json.MarshalIndent(response, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	w.Write(b)
}

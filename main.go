package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dafalo/Mobile-Tracker-App/api"
	"github.com/dafalo/Mobile-Tracker-App/models"
	"github.com/dafalo/Mobile-Tracker-App/views"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	BindIP = "0.0.0.0"
	Port   = ":2027"
)

func main() {
	u, _ := url.Parse("http://" + BindIP + Port)
	fmt.Printf("Server Started: %v\n", u)

	CreateDB("groupnb")
	MigrateDB()
	CreateDefaultUser()
	Handlers()

	http.ListenAndServe(Port, nil)
}

func Handlers() {
	fmt.Println("Handlers")
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/api/", api.APIHandler)
	http.HandleFunc("/portal", views.LoginHandler)
	http.HandleFunc("/", views.MainHandler)
	http.HandleFunc("/team", views.TeamHandler)
	http.HandleFunc("/team_fr", views.TeamFRHandler)
	http.HandleFunc("/Department", views.DepartmentHandler)
	http.HandleFunc("/Supervisor", views.SupervisorHandler)
	http.HandleFunc("/canada_immigration", views.Im_CanadaHandler)
	http.HandleFunc("/canada_immigration_fr", views.Im_CanadaFRHandler)
	http.HandleFunc("/canada_lumping", views.L_CanadaHandler)
	http.HandleFunc("/canada_lumping_fr", views.L_CanadaFRHandler)
	http.HandleFunc("/usa_immigration", views.Im_UsaHandler)
	http.HandleFunc("/usa_lumping", views.L_UsaHandler)
	http.HandleFunc("/dashboard", views.IndexHandler)
	http.HandleFunc("/canada/apply", views.ClientHandler)
	http.HandleFunc("/usa/apply", views.USAHandler)
	http.HandleFunc("/job", views.JobHandler)
	http.HandleFunc("/usa/job", views.USHandler)
	http.HandleFunc("/position", views.PositionHandler)
	http.HandleFunc("/applicant", views.ApplicantHandler)
	http.HandleFunc("/candidate", views.CandidateHandler)
	http.HandleFunc("/profile", views.ProfileHandler)
	http.HandleFunc("/register", views.UserHandler)
	http.HandleFunc("/logout", views.LogOutHandler)

}

func CreateDB(name string) *sql.DB {
	fmt.Println("Database Created")
	db, err := sql.Open("mysql", "root:GroupNB2023@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func CreateDefaultUser() {

	db := GormDB()

	user := []models.User{}
	db.Find(&user)

	defaultUser := []models.User{
		{
			Username: "admin",
			Password: hashPassword("admin"),
			Name:     "Software Developer",
			Type:     "Administrator",
		},
		{
			Username: "user",
			Password: hashPassword("user"),
			Name:     "Software Developer",
			Type:     "User",
		},
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Username == users.Username {
				isExisting = true
				break
			}
		}

		if !isExisting {
			fmt.Println("Create Default User")
			db.Save(&defaultUser[i])
		}
	}

}

func MigrateDB() {
	fmt.Println("Database Migrated")
	user := models.User{}
	position := models.Position{}
	applicant := models.Applicant{}
	department := models.Department{}
	supervisor := models.Supervisor{}
	candidate := models.Candidate{}
	immigration := models.Immigration{}

	db := GormDB()
	db.AutoMigrate(&user, &position,&applicant,&department,&supervisor,&candidate,&immigration)
}

func GormDB() *gorm.DB {
	dsn := "root:GroupNB2023@tcp(127.0.0.1:3306)/groupnb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

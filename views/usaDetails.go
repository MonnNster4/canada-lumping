package views

import (
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"text/template"
	
	

	"github.com/dafalo/Mobile-Tracker-App/models"
	
)


func USHandler(w http.ResponseWriter, r *http.Request) {
	
		tmpl := template.Must(template.ParseFiles("./templates/us_details.html"))
		data := map[string]interface{}{}
		data["Title"] = "Register | POS_SYSTEM"

		if r.Method == "POST" {
			db := GormDB()
			id := r.FormValue("id")
			title := r.FormValue("title")
			name := r.FormValue("name")
			email := r.FormValue("email")
			link := r.FormValue("link")
			letter := r.FormValue("letter")
			item := models.Applicant{}

			mydir, _ := os.Getwd()
		b := mydir + "/static/resume/"
		err := r.ParseMultipartForm(200000) 
		if err != nil {
			fmt.Fprintln(w, err)

		}
		formdata := r.MultipartForm 
		files := formdata.File["productImages"] 
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				fmt.Fprintln(w, err)

			}

			out, err := os.Create(b + files[i].Filename)

			defer out.Close()

			if err != nil {
				fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")

			}

			_, err = io.Copy(out, file)

			if err != nil {
				fmt.Fprintln(w, err)

			}

			item.PositionID = id
			item.Name = name
			item.Email = email
			item.Link = link
			item.Letter = letter
			item.Status = "Available"
			item.CV ="/static/resume/" + files[i].Filename

			db.Save(&item)

			from := "groupnb23@gmail.com"
			pass := "dwxlbnyzqaxzjtbx"
			to := "nbrecruit@groupnb.ca"
	
		msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject:" + title + "\n\n" +
			"Good day,\n\n Applicant Details \n\n" +
			"Name:" + name + "\n\n" +
			"Resume:" + " " + "http://97.74.83.143:2027"+"/static/resume/" + files[i].Filename + "\n\n" +
			 "Thank You"
			

		data := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))
	
		if data != nil {
			fmt.Printf("smtp error: %s", data)
			return
		}
		fmt.Println("Successfully sended to " + to)
	
			fmt.Println("Files uploaded successfully : ")
			fmt.Println(files[i].Filename + "\n")
		}
		http.Redirect(w, r, "/usa/job", http.StatusSeeOther)
		
	}
		tmpl.Execute(w, data)
}

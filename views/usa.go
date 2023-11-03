package views

import (
	"net/http"
	"text/template"
)

func USAHandler(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/usa.html"))
		data := map[string]interface{}{}
		data["Title"] = "Register | POS_SYSTEM"
		tmpl.Execute(w, data)
}

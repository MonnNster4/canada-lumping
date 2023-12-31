package views

import (
	"net/http"
	"text/template"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	activeSession := GetActiveSession(r)

	if activeSession == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		tmpl := template.Must(template.ParseFiles("./templates/Profile.html"))
		data := map[string]interface{}{}
		data["Title"] = "Register | POS_SYSTEM"
		tmpl.Execute(w, data)

	}

}

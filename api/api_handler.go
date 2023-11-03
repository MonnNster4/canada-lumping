package api

import (
	"net/http"
	"strings"
)

// APIHandler !
func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	
	if strings.HasPrefix(r.URL.Path, "user") {
		AddUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_user") {
		GetUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_user") {
		EditUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_user") {
		DeleteUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "reset_user") {
		ResetPassword(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "login") {
		Login(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "position") {
		AddPosition(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_position") {
		GetPosition(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_position") {
		EditPosition(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_position") {
		DeletePosition(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "change_pass") {
		ChangePassword(w, r)
		return
	}


	if strings.HasPrefix(r.URL.Path, "get_applicant") {
		GetApplicant(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "department") {
		AddDepartment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_department") {
		GetDepartment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_department") {
		EditDepartment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_department") {
		DeleteDepartment(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "supervisor") {
		AddSupervisor(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_supervisor") {
		GetSupervisor(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_supervisor") {
		EditSupervisor(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_supervisor") {
		DeleteSupervisor(w, r)
		return
	}


	if strings.HasPrefix(r.URL.Path, "candidate") {
		AddCandidate(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_candidate") {
		GetCandidate(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_candidate") {
		DeleteCandidate(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_candidate") {
		EditCandidate(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "immigration") {
		ImmigrationEmail(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_email") {
		GetEmail(w, r)
		return
	}
}

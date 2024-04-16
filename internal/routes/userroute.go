package routes

import (
	"html/template"
	"log"
	"net/http"
	"school23/internal/services"
	"school23/internal/utils"
	"strconv"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	users := services.GetUsers()
	tmpl, err := template.ParseFiles("assets/html/admin/user/users.html")
	if err != nil {
		log.Println(err)
		http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, users)
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("assets/html/admin/user/user_create.html")
		if err != nil {
			log.Println(err)
			http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}
	username := r.FormValue("username")
	role := r.FormValue("role")
	password := r.FormValue("password")
	err := services.UserCreate(username, password, role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/operations/user", http.StatusSeeOther)
}

func userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, utils.ServerError, http.StatusInternalServerError)
			return
		}

		user, err := services.GetUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("assets/html/admin/user/user_update.html")
		if err != nil {
			log.Println(err)
			http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, user)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println(err)
		http.Error(w, utils.ServerError, http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	role := r.FormValue("role")
	status := r.FormValue("status") == "active"
	err = services.UserUpdate(id, username, role, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/operations/user", http.StatusSeeOther)

}

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println(err)
		http.Error(w, utils.ServerError, http.StatusInternalServerError)
		return
	}
	err = services.DeleteUserById(id)
	if err != nil {
		http.Error(w, utils.ServerError, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/operations/user", http.StatusSeeOther)
}

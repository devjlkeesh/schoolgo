package routes

import (
	"html/template"
	"log"
	"net/http"
	"school23/internal/services"
	"school23/internal/utils"
	"strconv"
)

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("assets/html/admin/teacher/teacher.html")
	if err != nil {
		log.Println(err)
		http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
		return
	}
	teachers := services.GetTeachers()
	tmpl.Execute(w, teachers)
}

func teacherCreateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		fullname := r.FormValue("fullname")
		birthdate := r.FormValue("birthdate")
		phone := r.FormValue("phone")
		subject := r.FormValue("subject")
		category := r.FormValue("category")
		file, header, _ := r.FormFile("img")
		isImgPublic := r.FormValue("isImgPublic") == "yes"

		err := services.TeacherCreate(fullname, birthdate, phone, subject, category, file, header, isImgPublic)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/operations/teacher", http.StatusSeeOther)
	} else {
		tmpl, err := template.ParseFiles("assets/html/admin/teacher/teacher_create.html")
		if err != nil {
			log.Println(err)
			http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func teacherUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, utils.ServerError, http.StatusInternalServerError)
			return
		}
		fullname := r.FormValue("fullname")
		birthdate := r.FormValue("birthdate")
		phone := r.FormValue("phone")
		subject := r.FormValue("subject")
		category := r.FormValue("category")
		file, header, _ := r.FormFile("img")
		isImgPublic := r.FormValue("isImgPublic") == "yes"

		err = services.TeacherUpdate(id, fullname, birthdate, phone, subject, category, file, header, isImgPublic)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/operations/teacher", http.StatusSeeOther)
	} else {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, utils.ServerError, http.StatusInternalServerError)
			return
		}

		teacher, err := services.GetTeacher(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("assets/html/admin/teacher/teacher_update.html")
		if err != nil {
			log.Println(err)
			http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, teacher)
	}
}

func teacherDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println(err)
		http.Error(w, utils.ServerError, http.StatusInternalServerError)
		return
	}

	err = services.DeleteTeacherById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/operations/teacher", http.StatusSeeOther)
}

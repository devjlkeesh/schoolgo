package routes

import (
	"html/template"
	"log"
	"net/http"
	"school23/internal/models"
	"school23/internal/services"
	"school23/internal/utils"
	"strconv"
)

func articleHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("assets/html/admin/news/news.html")
	if err != nil {
		log.Println(err)
		http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
		return
	}
	page := utils.ParseTOIntSafe(r.URL.Query().Get("page"))
	articlePage := services.GetAllNews(page, false)
	articlePage.Path = "article"
	articlePage.TitleName = "Maqola"
	tmpl.Execute(w, articlePage)
}

func articleCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		overview := r.FormValue("overview")
		body := r.FormValue("body")
		err := services.NewsCreate(title, overview, body, false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/operations/article", http.StatusSeeOther)
	} else {
		tmpl, err := template.ParseFiles("assets/html/admin/news/news_create.html")
		if err != nil {
			log.Println(err)
			http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
			return
		}
		data := models.NewsCreate{TitleName: "Maqola", Path: "article"}
		tmpl.Execute(w, data)
	}
}

func articleUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, utils.ServerError, http.StatusInternalServerError)
			return
		}

		title := r.FormValue("title")
		overview := r.FormValue("overview")
		body := r.FormValue("body")
		accepted := r.FormValue("publish") == "yes"
		err = services.NewsUpdate(id, title, overview, body, false, accepted)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/operations/article", http.StatusSeeOther)
	} else {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			http.Error(w, utils.ServerError, http.StatusInternalServerError)
			return
		}

		news, err := services.GetNews(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("assets/html/admin/news/news_update.html")
		if err != nil {
			log.Println(err)
			http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
			return
		}

		data := models.NewsUpdate{
			TitleName: "Maqola",
			Path:      "article",
			Title:     news.Title,
			Overview:  news.Overview,
			Body:      news.Body,
		}
		tmpl.Execute(w, data)
	}
}

func articleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println(err)
		http.Error(w, utils.ServerError, http.StatusInternalServerError)
		return
	}
	err = services.DeleteNewsById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/operations/article", http.StatusSeeOther)
}

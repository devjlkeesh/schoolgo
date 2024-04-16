package routes

import (
	"html/template"
	"log"
	"net/http"
	"school23/internal/models"
	"school23/internal/services"
	"school23/internal/utils"
)

func publicIndexHandler(w http.ResponseWriter, r *http.Request) {
	if !(r.Method == "GET") {
		http.Error(w, utils.MethodNotAllowed, http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("assets/html/public/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, utils.PageNotFound, http.StatusNotFound)
		return
	}

	NewsPage := services.GetAllAcceptedNews(0, true, true)
	ArticlePage := services.GetAllAcceptedNews(0, false, true)
	TeacherPage := services.GetTeachersPage(0)
	data := models.MainPage{
		News:               NewsPage.NewsArray,
		NewsPagination:     NewsPage.Pagination,
		Article:            ArticlePage.NewsArray,
		ArticlePagination:  ArticlePage.Pagination,
		Teachers:           TeacherPage.Teachers,
		TeachersPagination: TeacherPage.Pagination,
	}
	tmpl.Execute(w, data)
}

func publicNewsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("assets/html/public/news.html")
	if err != nil {
		log.Println(err)
		http.Error(w, utils.PageNotFound, http.StatusNotFound)
		return
	}

	NewsPage := services.GetAllAcceptedNews(0, true, true)
	tmpl.Execute(w, NewsPage)
}

func publicArticleHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("assets/html/public/article.html")
	if err != nil {
		log.Println(err)
		http.Error(w, utils.PageNotFound, http.StatusNotFound)
		return
	}

	ArticlePage := services.GetAllAcceptedNews(0, false, true)
	tmpl.Execute(w, ArticlePage)
}

func publicTeacherHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("assets/html/public/teacher.html")
	if err != nil {
		log.Println(err)
		http.Error(w, utils.PageNotFound, http.StatusInternalServerError)
		return
	}
	page := utils.ParseTOIntSafe(r.URL.Query().Get("page"))
	TeacherPage := services.GetTeachersPage(page)
	tmpl.Execute(w, TeacherPage)
}

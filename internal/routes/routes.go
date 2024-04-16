package routes

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"school23/internal/xzy"
)

type application struct {
	auth struct {
		username string
		password string
	}
}

func NewRouter() http.Handler {

	app := new(application)
	app.auth.username = xzy.A1
	app.auth.password = xzy.A2

	mux := http.NewServeMux()
	mux.HandleFunc("/home", publicIndexHandler)
	mux.HandleFunc("/news", publicNewsHandler)
	mux.HandleFunc("/article", publicArticleHandler)
	mux.HandleFunc("/teacher", publicTeacherHandler)

	mux.HandleFunc("/operations/user", app.basicAuth(userHandler))
	mux.HandleFunc("/operations/user/create", app.basicAuth(userCreateHandler))
	mux.HandleFunc("/operations/user/update/{id}", app.basicAuth(userUpdateHandler))
	mux.HandleFunc("/operations/user/delete/{id}", app.basicAuth(userDeleteHandler))

	mux.HandleFunc("/operations/teacher", app.basicAuth(teachersHandler))
	mux.HandleFunc("/operations/teacher/create", app.basicAuth(teacherCreateHandler))
	mux.HandleFunc("/operations/teacher/update/{id}", app.basicAuth(teacherUpdateHandler))
	mux.HandleFunc("/operations/teacher/delete/{id}", app.basicAuth(teacherDeleteHandler))

	mux.HandleFunc("/operations/news", app.basicAuth(newsHandler))
	mux.HandleFunc("/operations/news/", app.basicAuth(newsHandler))
	mux.HandleFunc("/operations/news/create", app.basicAuth(newsCreateHandler))
	mux.HandleFunc("/operations/news/update/{id}", app.basicAuth(newsUpdateHandler))
	mux.HandleFunc("/operations/news/delete/{id}", app.basicAuth(newsDeleteHandler))

	mux.HandleFunc("/operations/article", app.basicAuth(articleHandler))
	mux.HandleFunc("/operations/article/create", app.basicAuth(articleCreateHandler))
	mux.HandleFunc("/operations/article/update/{id}", app.basicAuth(articleUpdateHandler))
	mux.HandleFunc("/operations/article/delete/{id}", app.basicAuth(articleDeleteHandler))

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	return mux
}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

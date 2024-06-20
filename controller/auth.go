package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func Login(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			username := r.FormValue("username")
			password := r.FormValue("password")

			var dbUsername, dbPassword string
			err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&dbUsername, &dbPassword)
			if err != nil || password != dbPassword {
				http.Redirect(w, r, "/admin/login?error=invalid_credentials", http.StatusSeeOther)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:  "authenticated",
				Value: "true",
				Path:  "/",
			})

			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		fp := filepath.Join("views", "admin", "login.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Logout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "authenticated",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

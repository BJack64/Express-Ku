package masinisController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateMasinis(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			id_masinis := r.Form["id_masinis"][0]
			nama_masinis := r.Form["nama_masinis"][0]
			email_masinis := r.Form["email_masinis"][0]
			gender_masinis := r.Form["gender_masinis"][0]
			exp_masinis := r.Form["exp_masinis"][0]
			_, err := db.Exec("INSERT INTO masinis (id_masinis, nama_masinis, email_masinis, gender_masinis, exp_masinis) VALUES (?, ?, ?, ?, ?)", id_masinis, nama_masinis, email_masinis, gender_masinis, exp_masinis)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
			http.Redirect(w, r, "/admin/masinis", http.StatusMovedPermanently)
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "admin/masinis/create.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
		}
	}
}

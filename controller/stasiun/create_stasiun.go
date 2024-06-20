package stasiunController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateStasiun(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			id_stasiun := r.Form["id_stasiun"][0]
			nama_stasiun := r.Form["nama_stasiun"][0]
			kota_stasiun := r.Form["kota_stasiun"][0]
			_, err := db.Exec("INSERT INTO stasiun (id_stasiun, nama_stasiun, kota_stasiun) VALUES (?, ?, ?)", id_stasiun, nama_stasiun, kota_stasiun)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
			http.Redirect(w, r, "/admin/stasiun", http.StatusMovedPermanently)
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "admin/stasiun/create.html")
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

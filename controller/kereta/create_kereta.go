package keretaController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateKereta(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			id_kereta := r.Form["id_kereta"][0]
			nama_kereta := r.Form["nama_kereta"][0]
			kelas := r.Form["kelas"][0]
			_, err := db.Exec("INSERT INTO kereta (id_kereta, nama_kereta, kelas) VALUES (?, ?, ?)", id_kereta, nama_kereta, kelas)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
			http.Redirect(w, r, "/admin/kereta", http.StatusMovedPermanently)
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "admin/kereta/create.html")
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

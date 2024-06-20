package keretaController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func EditKereta(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_kereta := r.URL.Query().Get("id_kereta")
		if r.Method == "POST" {
			r.ParseForm()

			nama_kereta := r.Form["nama_kereta"][0]
			kelas := r.Form["kelas"][0]
			id_kereta = r.Form["id_kereta"][0]
			_, err := db.Exec("UPDATE kereta SET nama_kereta=?, kelas=? WHERE id_kereta=?", nama_kereta, kelas, id_kereta)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			http.Redirect(w, r, "/admin/kereta", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			var kereta Kereta
			row := db.QueryRow("SELECT nama_kereta, kelas FROM kereta WHERE id_kereta = ?", id_kereta)
			err := row.Scan(
				&kereta.NamaKereta,
				&kereta.Kelas,
			)
			kereta.IdKereta = id_kereta

			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			fp := filepath.Join("views", "admin/kereta/edit.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			data := make(map[string]any)
			data["kereta"] = kereta

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
		}
	}
}

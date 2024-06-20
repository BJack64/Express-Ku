package masinisController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func EditMasinis(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_masinis := r.URL.Query().Get("id_masinis")
		if r.Method == "POST" {
			r.ParseForm()

			nama_masinis := r.Form["nama_masinis"][0]
			email_masinis := r.Form["email_masinis"][0]
			gender_masinis := r.Form["gender_masinis"][0]
			exp_masinis := r.Form["exp_masinis"][0]
			id_masinis = r.Form["id_masinis"][0]
			_, err := db.Exec("UPDATE masinis SET nama_masinis=?, email_masinis=?, gender_masinis=?, exp_masinis=? WHERE id_masinis=?", nama_masinis, email_masinis, gender_masinis, exp_masinis, id_masinis)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			http.Redirect(w, r, "/admin/masinis", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			var masinis Masinis
			row := db.QueryRow("SELECT nama_masinis, email_masinis, gender_masinis, exp_masinis FROM masinis WHERE id_masinis = ?", id_masinis)
			err := row.Scan(
				&masinis.NamaMasinis,
				&masinis.EmailMasinis,
				&masinis.GenderMasinis,
				&masinis.ExpMasinis,
			)
			masinis.IdMasinis = id_masinis

			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			fp := filepath.Join("views", "admin/masinis/edit.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			data := make(map[string]any)
			data["masinis"] = masinis

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
		}
	}
}

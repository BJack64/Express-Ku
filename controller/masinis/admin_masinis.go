package masinisController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Masinis struct {
	IdMasinis     string
	NamaMasinis   string
	EmailMasinis  string
	GenderMasinis string
	ExpMasinis    string
}

func AdminMasinis(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id_masinis, nama_masinis, email_masinis, gender_masinis, exp_masinis FROM masinis")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var masiniss []Masinis
		for rows.Next() {
			var masinis Masinis
			err := rows.Scan(
				&masinis.IdMasinis,
				&masinis.NamaMasinis,
				&masinis.EmailMasinis,
				&masinis.GenderMasinis,
				&masinis.ExpMasinis,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			masiniss = append(masiniss, masinis)
		}

		fp := filepath.Join("views", "admin", "masinis", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]interface{})
		data["masiniss"] = masiniss

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

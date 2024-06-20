package keretaController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Kereta struct {
	IdKereta   string
	NamaKereta string
	Kelas      string
}

func AdminKereta(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id_kereta, nama_kereta, kelas FROM kereta")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var keretas []Kereta
		for rows.Next() {
			var kereta Kereta
			err := rows.Scan(
				&kereta.IdKereta,
				&kereta.NamaKereta,
				&kereta.Kelas,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			keretas = append(keretas, kereta)
		}

		fp := filepath.Join("views", "admin", "kereta", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]interface{})
		data["keretas"] = keretas

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

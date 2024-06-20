package stasiunController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Stasiun struct {
	IdStasiun   string
	NamaStasiun string
	KotaStasiun string
}

func AdminStasiun(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id_stasiun, nama_stasiun, kota_stasiun FROM stasiun")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var stasiuns []Stasiun
		for rows.Next() {
			var stasiun Stasiun
			err := rows.Scan(
				&stasiun.IdStasiun,
				&stasiun.NamaStasiun,
				&stasiun.KotaStasiun,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			stasiuns = append(stasiuns, stasiun)
		}

		fp := filepath.Join("views", "admin", "stasiun", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]interface{})
		data["stasiuns"] = stasiuns

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

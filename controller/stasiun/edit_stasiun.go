package stasiunController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func EditStasiun(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_stasiun := r.URL.Query().Get("id_stasiun")
		if r.Method == "POST" {
			r.ParseForm()

			nama_stasiun := r.Form["nama_stasiun"][0]
			kota_stasiun := r.Form["kota_stasiun"][0]
			id_stasiun = r.Form["id_stasiun"][0]
			_, err := db.Exec("UPDATE stasiun SET nama_stasiun=?, kota_stasiun=? WHERE id_stasiun=?", nama_stasiun, kota_stasiun, id_stasiun)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			http.Redirect(w, r, "/admin/stasiun", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			var stasiun Stasiun
			row := db.QueryRow("SELECT nama_stasiun, kota_stasiun FROM stasiun WHERE id_stasiun = ?", id_stasiun)
			err := row.Scan(
				&stasiun.NamaStasiun,
				&stasiun.KotaStasiun,
			)
			stasiun.IdStasiun = id_stasiun

			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			fp := filepath.Join("views", "admin/stasiun/edit.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			data := make(map[string]any)
			data["stasiun"] = stasiun

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
		}
	}
}

package jadwalController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateJadwal(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			id_jadwal := r.Form["id_jadwal"][0]
			id_kereta := r.Form["id_kereta"][0]
			id_masinis := r.Form["id_masinis"][0]
			id_stasiun_asal := r.Form["id_stasiun_asal"][0]
			id_stasiun_tujuan := r.Form["id_stasiun_tujuan"][0]
			tgl_berangkat := r.Form["tgl_berangkat"][0]
			waktu_berangkat := r.Form["waktu_berangkat"][0]
			tgl_tiba := r.Form["tgl_tiba"][0]
			waktu_tiba := r.Form["waktu_tiba"][0]
			_, err := db.Exec("INSERT INTO jadwal (id_jadwal, id_kereta, id_masinis, id_stasiun_asal, id_stasiun_tujuan, tgl_berangkat, waktu_berangkat, tgl_tiba, waktu_tiba) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", id_jadwal, id_kereta, id_masinis, id_stasiun_asal, id_stasiun_tujuan, tgl_berangkat, waktu_berangkat, tgl_tiba, waktu_tiba)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
			http.Redirect(w, r, "/admin/jadwal", http.StatusMovedPermanently)
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "admin/jadwal/create.html")
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

package buktiController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateBukti(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			id_tiket := r.Form["id_tiket"][0]
			tgl_pembelian := r.Form["tgl_pembelian"][0]
			id_jadwal := r.Form["id_jadwal"][0]
			nama_penumpang := r.Form["nama_penumpang"][0]
			no_telp := r.Form["no_telp"][0]
			_, err := db.Exec("INSERT INTO bukti (id_tiket, tgl_pembelian, id_jadwal, nama_penumpang, no_telp) VALUES (?, ?, ?, ?, ?)", id_tiket, tgl_pembelian, id_jadwal, nama_penumpang, no_telp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
			http.Redirect(w, r, "/admin/bukti", http.StatusMovedPermanently)
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "admin/bukti/create.html")
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

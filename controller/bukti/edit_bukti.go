package buktiController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func EditBukti(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_tiket := r.URL.Query().Get("id_tiket")
		if r.Method == "POST" {
			r.ParseForm()

			tgl_pembelian := r.Form["tgl_pembelian"][0]
			id_jadwal := r.Form["id_jadwal"][0]
			nama_penumpang := r.Form["nama_penumpang"][0]
			no_telp := r.Form["no_telp"][0]
			id_tiket = r.Form["id_tiket"][0] // Ambil id_tiket dari form data

			_, err := db.Exec("UPDATE bukti SET tgl_pembelian=?, id_jadwal=?, nama_penumpang=?, no_telp=? WHERE id_tiket=?", tgl_pembelian, id_jadwal, nama_penumpang, no_telp, id_tiket)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/admin/bukti", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			var bukti Bukti
			row := db.QueryRow("SELECT tgl_pembelian, id_jadwal, nama_penumpang, no_telp FROM bukti WHERE id_tiket=?", id_tiket)
			err := row.Scan(
				&bukti.TglPembelian,
				&bukti.IdJadwal,
				&bukti.NamaPenumpang,
				&bukti.NoTelp,
			)
			bukti.IdTiket = id_tiket

			if err != nil {
				if err == sql.ErrNoRows {
					http.NotFound(w, r)
				} else {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
				}
				return
			}

			fp := filepath.Join("views", "admin/bukti/edit.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]interface{})
			data["bukti"] = bukti

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

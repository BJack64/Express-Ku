package buktiController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Bukti struct {
	IdTiket       string
	TglPembelian  string
	IdJadwal      string
	NamaPenumpang string
	NoTelp        string
}

func AdminBukti(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id_tiket, tgl_pembelian, id_jadwal, nama_penumpang, no_telp FROM bukti")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var buktis []Bukti
		for rows.Next() {
			var bukti Bukti
			err := rows.Scan(
				&bukti.IdTiket,
				&bukti.TglPembelian,
				&bukti.IdJadwal,
				&bukti.NamaPenumpang,
				&bukti.NoTelp,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			buktis = append(buktis, bukti)
		}

		fp := filepath.Join("views", "admin", "bukti", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]interface{})
		data["buktis"] = buktis

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

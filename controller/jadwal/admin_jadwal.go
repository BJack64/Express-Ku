package jadwalController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Jadwal struct {
	IdJadwal        string
	IdKereta        string
	IdMasinis       string
	IdStasiunAsal   string
	IdStasiunTujuan string
	TglBerangkat    string
	WaktuBerangkat  string
	TglTiba         string
	WaktuTiba       string
}

func AdminJadwal(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id_jadwal, id_kereta, id_masinis, id_stasiun_asal, id_stasiun_tujuan, tgl_berangkat, waktu_berangkat, tgl_tiba, waktu_tiba FROM jadwal")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var jadwals []Jadwal
		for rows.Next() {
			var jadwal Jadwal
			err := rows.Scan(
				&jadwal.IdJadwal,
				&jadwal.IdKereta,
				&jadwal.IdMasinis,
				&jadwal.IdStasiunAsal,
				&jadwal.IdStasiunTujuan,
				&jadwal.TglBerangkat,
				&jadwal.WaktuBerangkat,
				&jadwal.TglTiba,
				&jadwal.WaktuTiba,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			jadwals = append(jadwals, jadwal)
		}

		fp := filepath.Join("views", "admin", "jadwal", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]interface{})
		data["jadwals"] = jadwals

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

package jadwalController

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func EditJadwal(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_jadwal := r.URL.Query().Get("id_jadwal")
		if r.Method == "POST" {
			r.ParseForm()

			id_kereta := r.Form["id_kereta"][0]
			id_masinis := r.Form["id_masinis"][0]
			id_stasiun_asal := r.Form["id_stasiun_asal"][0]
			id_stasiun_tujuan := r.Form["id_stasiun_tujuan"][0]
			tgl_berangkat := r.Form["tgl_berangkat"][0]
			waktu_berangkat := r.Form["waktu_berangkat"][0]
			tgl_tiba := r.Form["tgl_tiba"][0]
			waktu_tiba := r.Form["waktu_tiba"][0]
			id_jadwal = r.Form["id_jadwal"][0]
			_, err := db.Exec("UPDATE jadwal SET id_kereta=?, id_masinis=?, id_stasiun_asal=?, id_stasiun_tujuan=?, tgl_berangkat=?, waktu_berangkat=?, tgl_tiba=?, waktu_tiba=? WHERE id_jadwal=?", id_kereta, id_masinis, id_stasiun_asal, id_stasiun_tujuan, tgl_berangkat, waktu_berangkat, tgl_tiba, waktu_tiba, id_jadwal)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			http.Redirect(w, r, "/admin/jadwal", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			var jadwal Jadwal
			row := db.QueryRow("SELECT id_kereta, id_masinis, id_stasiun_asal, id_stasiun_tujuan, tgl_berangkat, waktu_berangkat, tgl_tiba, waktu_tiba FROM jadwal WHERE id_jadwal=?", id_jadwal)
			err := row.Scan(
				&jadwal.IdKereta,
				&jadwal.IdMasinis,
				&jadwal.IdStasiunAsal,
				&jadwal.IdStasiunTujuan,
				&jadwal.TglBerangkat,
				&jadwal.WaktuBerangkat,
				&jadwal.TglTiba,
				&jadwal.WaktuTiba,
			)
			jadwal.IdJadwal = id_jadwal

			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			fp := filepath.Join("views", "admin/jadwal/edit.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			data := make(map[string]any)
			data["jadwal"] = jadwal

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
		}
	}
}

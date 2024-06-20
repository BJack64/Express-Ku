package jadwalController

import (
	"database/sql"
	"net/http"
)

func DeleteJadwal(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_jadwal := r.URL.Query().Get("id_jadwal")

		_, err := db.Exec("DELETE FROM jadwal WHERE id_jadwal = ?", id_jadwal)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader((http.StatusInternalServerError))
			return
		}
		http.Redirect(w, r, "/admin/jadwal", http.StatusMovedPermanently)
	}
}

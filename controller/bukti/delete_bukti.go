package buktiController

import (
	"database/sql"
	"net/http"
)

func DeleteBukti(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_tiket := r.URL.Query().Get("id_tiket")

		_, err := db.Exec("DELETE FROM bukti WHERE id_tiket = ?", id_tiket)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader((http.StatusInternalServerError))
			return
		}
		http.Redirect(w, r, "/admin/bukti", http.StatusMovedPermanently)
	}
}

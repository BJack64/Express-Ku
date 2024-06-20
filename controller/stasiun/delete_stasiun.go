package stasiunController

import (
	"database/sql"
	"net/http"
)

func DeleteStasiun(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_stasiun := r.URL.Query().Get("id_stasiun")

		_, err := db.Exec("DELETE FROM stasiun WHERE id_stasiun = ?", id_stasiun)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader((http.StatusInternalServerError))
			return
		}
		http.Redirect(w, r, "/admin/stasiun", http.StatusMovedPermanently)
	}
}

package masinisController

import (
	"database/sql"
	"net/http"
)

func DeleteMasinis(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_masinis := r.URL.Query().Get("id_masinis")

		_, err := db.Exec("DELETE FROM masinis WHERE id_masinis = ?", id_masinis)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader((http.StatusInternalServerError))
			return
		}
		http.Redirect(w, r, "/admin/masinis", http.StatusMovedPermanently)
	}
}

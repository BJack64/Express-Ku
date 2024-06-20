package keretaController

import (
	"database/sql"
	"net/http"
)

func DeleteKereta(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id_kereta := r.URL.Query().Get("id_kereta")

		_, err := db.Exec("DELETE FROM kereta WHERE id_kereta = ?", id_kereta)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader((http.StatusInternalServerError))
			return
		}
		http.Redirect(w, r, "/admin/kereta", http.StatusMovedPermanently)
	}
}

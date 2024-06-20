package routes

import (
	"database/sql"
	"express-ku/controller"
	buktiController "express-ku/controller/bukti"
	jadwalController "express-ku/controller/jadwal"
	keretaController "express-ku/controller/kereta"
	masinisController "express-ku/controller/masinis"
	stasiunController "express-ku/controller/stasiun"
	"express-ku/middleware"
	"net/http"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.BookingController())
	server.HandleFunc("/finalisasi", controller.FinalisasiController(db))
	server.HandleFunc("/bukti", controller.BuktiController(db))
	server.HandleFunc("/browsing", controller.BrowsingController(db))
	server.HandleFunc("/rekomendasi", controller.RekomendasiController())

	server.HandleFunc("/admin/login", controller.Login(db))
	server.Handle("/admin/logout", http.HandlerFunc(controller.Logout()))
	server.Handle("/admin", middleware.AuthRequired(http.HandlerFunc(controller.AdminController())))

	server.Handle("/admin/stasiun", middleware.AuthRequired(http.HandlerFunc(stasiunController.AdminStasiun(db))))
	server.Handle("/admin/stasiun/create", middleware.AuthRequired(http.HandlerFunc(stasiunController.CreateStasiun(db))))
	server.Handle("/admin/stasiun/edit", middleware.AuthRequired(http.HandlerFunc(stasiunController.EditStasiun(db))))
	server.Handle("/admin/stasiun/delete", middleware.AuthRequired(http.HandlerFunc(stasiunController.DeleteStasiun(db))))

	server.Handle("/admin/kereta", middleware.AuthRequired(http.HandlerFunc(keretaController.AdminKereta(db))))
	server.Handle("/admin/kereta/create", middleware.AuthRequired(http.HandlerFunc(keretaController.CreateKereta(db))))
	server.Handle("/admin/kereta/edit", middleware.AuthRequired(http.HandlerFunc(keretaController.EditKereta(db))))
	server.Handle("/admin/kereta/delete", middleware.AuthRequired(http.HandlerFunc(keretaController.DeleteKereta(db))))

	server.Handle("/admin/masinis", middleware.AuthRequired(http.HandlerFunc(masinisController.AdminMasinis(db))))
	server.Handle("/admin/masinis/create", middleware.AuthRequired(http.HandlerFunc(masinisController.CreateMasinis(db))))
	server.Handle("/admin/masinis/edit", middleware.AuthRequired(http.HandlerFunc(masinisController.EditMasinis(db))))
	server.Handle("/admin/masinis/delete", middleware.AuthRequired(http.HandlerFunc(masinisController.DeleteMasinis(db))))

	server.Handle("/admin/jadwal", middleware.AuthRequired(http.HandlerFunc(jadwalController.AdminJadwal(db))))
	server.Handle("/admin/jadwal/create", middleware.AuthRequired(http.HandlerFunc(jadwalController.CreateJadwal(db))))
	server.Handle("/admin/jadwal/edit", middleware.AuthRequired(http.HandlerFunc(jadwalController.EditJadwal(db))))
	server.Handle("/admin/jadwal/delete", middleware.AuthRequired(http.HandlerFunc(jadwalController.DeleteJadwal(db))))

	server.Handle("/admin/bukti", middleware.AuthRequired(http.HandlerFunc(buktiController.AdminBukti(db))))
	server.Handle("/admin/bukti/create", middleware.AuthRequired(http.HandlerFunc(buktiController.CreateBukti(db))))
	server.Handle("/admin/bukti/edit", middleware.AuthRequired(http.HandlerFunc(buktiController.EditBukti(db))))
	server.Handle("/admin/bukti/delete", middleware.AuthRequired(http.HandlerFunc(buktiController.DeleteBukti(db))))
}

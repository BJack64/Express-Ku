package controller

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Search struct {
	IdStasiunAsal   string
	StasiunAsal     string
	IdStasiunTujuan string
	StasiunTujuan   string
	IdJadwal        string
	IdKereta        string
	TglBerangkat    string
	WaktuBerangkat  string
	TglTiba         string
	WaktuTiba       string
	NamaKereta      string
	Kelas           string
}

type Bukti struct {
	IdTiket       int
	IdJadwal      string
	TglPembelian  string
	NamaPenumpang string
	NoTelp        string
}

type Jadwal struct {
	IdJadwal         string
	IdKereta         string
	IdMasinis        string
	NamaMasinis      string
	IdStasiunAsal    string
	StasiunAsal      string
	IdStasiunTujuan  string
	StasiunTujuan    string
	KotaAsal         string
	KotaTujuan       string
	NamaKereta       string
	Kelas            string
	TanggalBerangkat string
	WaktuBerangkat   string
	TanggalTiba      string
	WaktuTiba        string
}

type Rekomendasi struct {
	Kota string
	Recc genai.Part
}

var Id_Jadwal string

func AdminController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fp := filepath.Join("views", "admin", "index.html")
		tmpl, err := template.ParseFiles(fp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func BookingController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fp := filepath.Join("views", "public", "booking.html")
		tmpl, err := template.ParseFiles(fp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func FinalisasiController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var searches []Search
			kota_asal := r.URL.Query().Get("kotaAsal")
			kota_tujuan := r.URL.Query().Get("kotaTujuan")
			tgl_berangkat := r.URL.Query().Get("tglBerangkat")
			var IdStasiunAsal string
			var IdStasiunTujuan string
			var IdKereta string

			//stasiun_asal
			rows1, err := db.Query("SELECT id_stasiun, nama_stasiun FROM stasiun WHERE kota_stasiun=?", kota_asal)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer rows1.Close()

			var search Search
			for rows1.Next() {
				err := rows1.Scan(
					&search.IdStasiunAsal,
					&search.StasiunAsal,
				)
				IdStasiunAsal = search.IdStasiunAsal

				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			//stasiun_tujuan
			rows2, err := db.Query("SELECT id_stasiun, nama_stasiun FROM stasiun WHERE kota_stasiun=?", kota_tujuan)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer rows2.Close()

			for rows2.Next() {
				err := rows2.Scan(
					&search.IdStasiunTujuan,
					&search.StasiunTujuan,
				)
				IdStasiunTujuan = search.IdStasiunTujuan

				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			//jadwal
			rows3, err := db.Query("SELECT id_jadwal, id_kereta, waktu_berangkat, tgl_tiba, waktu_tiba FROM jadwal WHERE id_stasiun_asal=? && id_stasiun_tujuan=? && tgl_berangkat=?", IdStasiunAsal, IdStasiunTujuan, tgl_berangkat)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer rows3.Close()

			for rows3.Next() {
				err := rows3.Scan(
					&search.IdJadwal,
					&search.IdKereta,
					&search.WaktuBerangkat,
					&search.TglTiba,
					&search.WaktuTiba,
				)
				IdKereta = search.IdKereta
				Id_Jadwal = search.IdJadwal
				search.TglBerangkat = tgl_berangkat

				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			//kereta
			rows4, err := db.Query("SELECT nama_kereta, kelas FROM kereta WHERE id_kereta=?", IdKereta)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer rows4.Close()

			for rows4.Next() {
				err := rows4.Scan(
					&search.NamaKereta,
					&search.Kelas,
				)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				searches = append(searches, search)
			}

			data := make(map[string]interface{})
			data["searches"] = searches

			fp := filepath.Join("views", "public", "finalisasi.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

func BuktiController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			r.ParseForm()
			var buktis []Bukti
			tgl_pembelian := r.URL.Query().Get("tglPembelian")
			nama_penumpang := r.URL.Query().Get("namaPenumpang")
			no_telp := r.URL.Query().Get("noTelp")
			id_jadwal := r.URL.Query().Get("idJadwal")

			_, err := db.Exec("INSERT INTO bukti (tgl_pembelian, id_jadwal, nama_penumpang, no_telp) VALUES (?, ?, ?, ?)", tgl_pembelian, id_jadwal, nama_penumpang, no_telp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			var bukti Bukti
			row := db.QueryRow("SELECT id_tiket, tgl_pembelian, id_jadwal, nama_penumpang, no_telp FROM bukti WHERE nama_penumpang = ? && id_jadwal = ?", nama_penumpang, Id_Jadwal)
			err = row.Scan(
				&bukti.IdTiket,
				&bukti.TglPembelian,
				&bukti.IdJadwal,
				&bukti.NamaPenumpang,
				&bukti.NoTelp,
			)
			buktis = append(buktis, bukti)

			if err != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			fp := filepath.Join("views", "public", "bukti.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			data := make(map[string]any)
			data["buktis"] = buktis

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
		}
	}
}

func BrowsingController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var jadwals []Jadwal

			//stasiun_asal
			rows1, err := db.Query("SELECT id_jadwal, id_kereta, id_masinis, id_stasiun_asal, id_stasiun_tujuan, tgl_berangkat, waktu_berangkat, tgl_tiba, waktu_tiba FROM jadwal")
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer rows1.Close()

			//jadwal
			for rows1.Next() {
				var jadwal Jadwal
				err := rows1.Scan(
					&jadwal.IdJadwal,
					&jadwal.IdKereta,
					&jadwal.IdMasinis,
					&jadwal.IdStasiunAsal,
					&jadwal.IdStasiunTujuan,
					&jadwal.TanggalBerangkat,
					&jadwal.WaktuBerangkat,
					&jadwal.TanggalTiba,
					&jadwal.WaktuTiba,
				)

				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				//stasiun_asal
				rows2, err := db.Query("SELECT nama_stasiun, kota_stasiun FROM stasiun WHERE id_stasiun=?", jadwal.IdStasiunAsal)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				defer rows2.Close()

				for rows2.Next() {
					err := rows2.Scan(
						&jadwal.StasiunAsal,
						&jadwal.KotaAsal,
					)

					if err != nil {
						w.Write([]byte(err.Error()))
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
				}

				//stasiun_tujuan
				rows3, err := db.Query("SELECT nama_stasiun, kota_stasiun FROM stasiun WHERE id_stasiun=?", jadwal.IdStasiunTujuan)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				defer rows3.Close()

				for rows3.Next() {
					err := rows3.Scan(
						&jadwal.StasiunTujuan,
						&jadwal.KotaTujuan,
					)

					if err != nil {
						w.Write([]byte(err.Error()))
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
				}

				//kereta
				rows4, err := db.Query("SELECT nama_kereta, kelas FROM kereta WHERE id_kereta=?", jadwal.IdKereta)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				defer rows4.Close()

				for rows4.Next() {
					err := rows4.Scan(
						&jadwal.NamaKereta,
						&jadwal.Kelas,
					)
					if err != nil {
						w.Write([]byte(err.Error()))
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
					jadwals = append(jadwals, jadwal)
				}
			}

			data := make(map[string]interface{})
			data["jadwals"] = jadwals

			fp := filepath.Join("views", "public", "browsing.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

func RekomendasiController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			r.ParseForm()
			var rekomendasi Rekomendasi
			kota := r.URL.Query().Get("kota")
			rekomendasi.Kota = kota

			ctx := context.Background()

			apiKey := ("AIzaSyArXtxCk-GWKWk9xGyqBtWY5MXYLx4WnIk")
			if apiKey == "" {
				http.Error(w, "GENAI_API_KEY environment variable is not set", http.StatusInternalServerError)
				return
			}

			client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer client.Close()

			model := client.GenerativeModel("gemini-1.5-flash")

			if kota == "" {
				http.Error(w, "Kota belum dimasukkan", http.StatusBadRequest)
				return
			}

			resp, err := model.GenerateContent(ctx, genai.Text("Berikan rekomendasi destinasi wisata dalam 200 kata atau kurang di kota "+kota))
			if err != nil {
				log.Fatal(err)
			}

			if resp != nil && len(resp.Candidates) > 0 {
				recc := (resp.Candidates[0].Content.Parts[0])
				rekomendasi.Recc = recc
			}

			fp := filepath.Join("views", "public", "rekomendasi.html")
			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}

			data := make(map[string]any)
			data["rekomendasi"] = rekomendasi

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader((http.StatusInternalServerError))
				return
			}
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Saham struct {
	Kode            string
	NamaPerusahaan  string
	Harga           float64
	VolumeTransaksi int
}

type Transaksi struct {
	KodeSaham string
	IsBeli    bool
	Jumlah    int
	Harga     float64
}

type Portofolio struct {
	Transaksi []Transaksi
	Saldo     float64
}

type Aplikasi struct {
	Saham      []Saham
	Portofolio Portofolio
	reader     *bufio.Reader
}

func BaruAplikasi() *Aplikasi {
	saham := []Saham{
		{"BBCA", "Bank Central Asia Tbk.", 8800.00, 1500000},
		{"TLKM", "Telekomunikasi Indonesia Tbk.", 3500.00, 2000000},
		{"BMRI", "Bank Mandiri (Persero) Tbk.", 7300.00, 1200000},
		{"ASII", "Astra International Tbk.", 6200.00, 900000},
		{"BBRI", "Bank Rakyat Indonesia (Persero) Tbk.", 4100.00, 1700000},
		{"UNVR", "Unilever Indonesia Tbk.", 42000.00, 300000},
		{"ICBP", "Indofood CBP Sukses Makmur Tbk.", 10200.00, 400000},
		{"HMSP", "HM Sampoerna Tbk.", 12500.00, 600000},
		{"EXCL", "XL Axiata Tbk.", 2800.00, 1000000},
		{"PGAS", "Perusahaan Gas Negara Tbk.", 2700.00, 1100000},
		{"GGRM", "Gudang Garam Tbk.", 61000.00, 250000},
		{"MEDC", "Medco Energi Internasional Tbk.", 1200.00, 1800000},
	}

	app := &Aplikasi{
		Saham: saham,
		Portofolio: Portofolio{
			Saldo:     100000.00,
			Transaksi: []Transaksi{},
		},
		reader: bufio.NewReader(os.Stdin),
	}

	sort.Slice(app.Saham, func(i, j int) bool {
		return app.Saham[i].Kode < app.Saham[j].Kode
	})
	return app
}

func (app *Aplikasi) Jalankan() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("=== Selamat datang di Aplikasi Simulasi Pasar Saham Virtual ===")
	for {
		app.simulasiPerubahanHarga()

		fmt.Println("\nSaldo Anda:", formatRupiah(app.Portofolio.Saldo))
		fmt.Println("Menu:")
		fmt.Println("1. Tambah transaksi pembelian/penjualan saham")
		fmt.Println("2. Ubah transaksi")
		fmt.Println("3. Hapus transaksi")
		fmt.Println("4. Lihat portofolio & nilai terkini")
		fmt.Println("5. Cari saham")
		fmt.Println("6. Urutkan saham")
		fmt.Println("7. Statistik keuntungan/kerugian")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu (0-7): ")
		input, _ := app.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			app.tambahTransaksi()
		case "2":
			app.ubahTransaksi()
		case "3":
			app.hapusTransaksi()
		case "4":
			app.tampilkanPortofolio()
		case "5":
			app.menuCariSaham()
		case "6":
			app.menuUrutSaham()
		case "7":
			app.tampilkanStatistikUntungRugi()
		case "0":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func formatRupiah(jumlah float64) string {
	return fmt.Sprintf("Rp %0.2f", jumlah)
}

func (app *Aplikasi) simulasiPerubahanHarga() {
	for i := range app.Saham {
		perubahan := (rand.Float64() * 0.1) - 0.05
		hargaBaru := app.Saham[i].Harga * (1 + perubahan)
		if hargaBaru < 1.0 {
			hargaBaru = 1.0
		}
		app.Saham[i].Harga = hargaBaru
	}
}

func (app *Aplikasi) tambahTransaksi() {
	fmt.Println("--- Tambah Transaksi ---")
	fmt.Print("Masukkan kode saham: ")
	kode := app.bacaInput()

	idx := app.cariSahamByKodeBinary(kode)
	if idx == -1 {
		fmt.Println("Saham tidak ditemukan.")
		return
	}
	saham := app.Saham[idx]

	fmt.Print("Jenis transaksi (beli/jual): ")
	jenis := strings.ToLower(app.bacaInput())
	if jenis != "beli" && jenis != "jual" {
		fmt.Println("Jenis transaksi tidak valid.")
		return
	}

	fmt.Print("Jumlah saham yang ingin ditransaksikan: ")
	jmlStr := app.bacaInput()
	jml, err := strconv.Atoi(jmlStr)
	if err != nil || jml <= 0 {
		fmt.Println("Jumlah tidak valid.")
		return
	}

	if jenis == "beli" {
		totalBiaya := saham.Harga * float64(jml)
		if totalBiaya > app.Portofolio.Saldo {
			fmt.Println("Saldo tidak cukup untuk transaksi ini.")
			return
		}
		app.Portofolio.Saldo -= totalBiaya
	} else {
		dimiliki := app.hitungSahamDimiliki(saham.Kode)
		if jml > dimiliki {
			fmt.Printf("Anda tidak memiliki cukup saham %s untuk dijual (dimiliki: %d).\n", saham.Kode, dimiliki)
			return
		}
		app.Portofolio.Saldo += saham.Harga * float64(jml)
	}

	tx := Transaksi{
		KodeSaham: saham.Kode,
		IsBeli:    jenis == "beli",
		Jumlah:    jml,
		Harga:     saham.Harga,
	}
	app.Portofolio.Transaksi = append(app.Portofolio.Transaksi, tx)
	fmt.Println("Transaksi berhasil ditambahkan.")
}

func (app *Aplikasi) ubahTransaksi() {
	if len(app.Portofolio.Transaksi) == 0 {
		fmt.Println("Tidak ada transaksi untuk diubah.")
		return
	}
	fmt.Println("--- Daftar Transaksi ---")
	for i, tx := range app.Portofolio.Transaksi {
		jenis := "Jual"
		if tx.IsBeli {
			jenis = "Beli"
		}
		fmt.Printf("%d) %s %d saham %s @%s\n", i+1, jenis, tx.Jumlah, tx.KodeSaham, formatRupiah(tx.Harga))
	}
	fmt.Print("Pilih nomor transaksi yang ingin diubah: ")
	input := app.bacaInput()
	idx, err := strconv.Atoi(input)
	if err != nil || idx < 1 || idx > len(app.Portofolio.Transaksi) {
		fmt.Println("Nomor transaksi tidak valid.")
		return
	}
	idx--

	tx := &app.Portofolio.Transaksi[idx]

	fmt.Printf("Transaksi terpilih: %s %d saham %s @%s\n",
		func() string {
			if tx.IsBeli {
				return "Beli"
			}
			return "Jual"
		}(), tx.Jumlah, tx.KodeSaham, formatRupiah(tx.Harga))

	fmt.Print("Masukkan jumlah baru (atau enter untuk tidak mengubah): ")
	baruJumlahStr := app.bacaInput()
	if baruJumlahStr != "" {
		baruJumlah, err := strconv.Atoi(baruJumlahStr)
		if err != nil || baruJumlah <= 0 {
			fmt.Println("Jumlah tidak valid. Perubahan dibatalkan.")
			return
		}
		idxSaham := app.cariSahamByKodeBinary(tx.KodeSaham)
		if idxSaham == -1 {
			fmt.Println("Saham tidak ditemukan. Perubahan dibatalkan.")
			return
		}
		hargaSaham := app.Saham[idxSaham].Harga

		if tx.IsBeli {
			app.Portofolio.Saldo += tx.Harga * float64(tx.Jumlah)
			totalBiaya := hargaSaham * float64(baruJumlah)
			if totalBiaya > app.Portofolio.Saldo {
				fmt.Println("Saldo tidak cukup untuk mengubah transaksi ini.")
				app.Portofolio.Saldo -= tx.Harga * float64(tx.Jumlah)
				return
			}
			app.Portofolio.Saldo -= totalBiaya
		} else {
			app.Portofolio.Saldo -= tx.Harga * float64(tx.Jumlah)
			dimiliki := app.hitungSahamDimiliki(tx.KodeSaham) + tx.Jumlah
			if baruJumlah > dimiliki {
				fmt.Printf("Tidak cukup saham untuk dijual. Dimiliki: %d\n", dimiliki)
				app.Portofolio.Saldo += tx.Harga * float64(tx.Jumlah)
				return
			}
			app.Portofolio.Saldo += hargaSaham * float64(baruJumlah)
		}
		tx.Jumlah = baruJumlah
		tx.Harga = hargaSaham
	}

	fmt.Print("Ubah jenis transaksi? (beli/jual, enter untuk tidak mengubah): ")
	baruJenis := strings.ToLower(app.bacaInput())
	if baruJenis != "" && baruJenis != "beli" && baruJenis != "jual" {
		fmt.Println("Jenis transaksi tidak valid. Perubahan dibatalkan.")
		return
	}
	if baruJenis != "" && ((baruJenis == "beli") != tx.IsBeli) {
		idxSaham := app.cariSahamByKodeBinary(tx.KodeSaham)
		if idxSaham == -1 {
			fmt.Println("Saham tidak ditemukan. Perubahan dibatalkan.")
			return
		}
		hargaSaham := app.Saham[idxSaham].Harga

		if tx.IsBeli {
			app.Portofolio.Saldo += tx.Harga * float64(tx.Jumlah)
		} else {
			app.Portofolio.Saldo -= tx.Harga * float64(tx.Jumlah)
		}

		if baruJenis == "beli" {
			totalBiaya := hargaSaham * float64(tx.Jumlah)
			if totalBiaya > app.Portofolio.Saldo {
				fmt.Println("Saldo tidak cukup untuk mengubah jenis transaksi.")
				if tx.IsBeli {
					app.Portofolio.Saldo -= tx.Harga * float64(tx.Jumlah)
				} else {
					app.Portofolio.Saldo += tx.Harga * float64(tx.Jumlah)
				}
				return
			}
			app.Portofolio.Saldo -= totalBiaya
			tx.IsBeli = true
			tx.Harga = hargaSaham
		} else {
			dimiliki := app.hitungSahamDimiliki(tx.KodeSaham)
			if tx.Jumlah > dimiliki {
				fmt.Printf("Tidak cukup saham untuk dijual. Dimiliki: %d\n", dimiliki)
				if tx.IsBeli {
					app.Portofolio.Saldo -= tx.Harga * float64(tx.Jumlah)
				} else {
					app.Portofolio.Saldo += tx.Harga * float64(tx.Jumlah)
				}
				return
			}
			app.Portofolio.Saldo += hargaSaham * float64(tx.Jumlah)
			tx.IsBeli = false
			tx.Harga = hargaSaham
		}
	}
	fmt.Println("Transaksi berhasil diperbarui.")
}

func (app *Aplikasi) hapusTransaksi() {
	if len(app.Portofolio.Transaksi) == 0 {
		fmt.Println("Tidak ada transaksi untuk dihapus.")
		return
	}
	fmt.Println("--- Daftar Transaksi ---")
	for i, tx := range app.Portofolio.Transaksi {
		jenis := "Jual"
		if tx.IsBeli {
			jenis = "Beli"
		}
		fmt.Printf("%d) %s %d saham %s @%s\n", i+1, jenis, tx.Jumlah, tx.KodeSaham, formatRupiah(tx.Harga))
	}
	fmt.Print("Pilih nomor transaksi yang ingin dihapus: ")
	input := app.bacaInput()
	idx, err := strconv.Atoi(input)
	if err != nil || idx < 1 || idx > len(app.Portofolio.Transaksi) {
		fmt.Println("Nomor transaksi tidak valid.")
		return
	}
	idx--

	tx := app.Portofolio.Transaksi[idx]

	if tx.IsBeli {
		app.Portofolio.Saldo += tx.Harga * float64(tx.Jumlah)
	} else {
		app.Portofolio.Saldo -= tx.Harga * float64(tx.Jumlah)
	}

	app.Portofolio.Transaksi = append(app.Portofolio.Transaksi[:idx], app.Portofolio.Transaksi[idx+1:]...)
	fmt.Println("Transaksi berhasil dihapus.")
}

func (app *Aplikasi) tampilkanPortofolio() {
	fmt.Println("--- Portofolio Anda ---")
	if len(app.Portofolio.Transaksi) == 0 {
		fmt.Println("Anda belum melakukan transaksi apapun.")
		return
	}

	// Hitung jumlah saham yang dipunya
	pegang := make(map[string]int)
	for _, tx := range app.Portofolio.Transaksi {
		if tx.IsBeli {
			pegang[tx.KodeSaham] += tx.Jumlah
		} else {
			pegang[tx.KodeSaham] -= tx.Jumlah
		}
	}

	// Hapus saham yang tidak dijual
	for kode, jumlah := range pegang {
		if jumlah <= 0 {
			delete(pegang, kode)
		}
	}

	totalNilai := 0.0
	fmt.Printf("%-8s %-25s %-10s %-15s\n", "Kode", "Nama Perusahaan", "Jumlah", "Nilai Saat Ini")
	for kode, jumlah := range pegang {
		idx := app.cariSahamByKodeBinary(kode)
		if idx == -1 {
			continue
		}
		s := app.Saham[idx]
		nilai := s.Harga * float64(jumlah)
		totalNilai += nilai
		fmt.Printf("%-8s %-25s %-10d %s\n", s.Kode, s.NamaPerusahaan, jumlah, formatRupiah(nilai))
	}

	fmt.Println("Total nilai portofolio saham:", formatRupiah(totalNilai))
	fmt.Println("Saldo tersedia:", formatRupiah(app.Portofolio.Saldo))
	fmt.Println("Total nilai portofolio + saldo:", formatRupiah(totalNilai+app.Portofolio.Saldo))
}

func (app *Aplikasi) hitungSahamDimiliki(kodeSaham string) int {
	total := 0
	for _, tx := range app.Portofolio.Transaksi {
		if tx.KodeSaham == kodeSaham {
			if tx.IsBeli {
				total += tx.Jumlah
			} else {
				total -= tx.Jumlah
			}
		}
	}
	if total < 0 {
		total = 0
	}
	return total
}

func (app *Aplikasi) menuCariSaham() {
	fmt.Println("--- Cari Saham ---")
	fmt.Println("1. Pencarian Sequential")
	fmt.Println("2. Pencarian Binary")
	fmt.Print("Pilih metode pencarian: ")
	input := app.bacaInput()

	fmt.Print("Masukkan kode atau nama perusahaan: ")
	keyword := strings.ToLower(app.bacaInput())

	switch input {
	case "1":
		results := app.cariSahamSequential(keyword)
		if len(results) == 0 {
			fmt.Println("Saham tidak ditemukan.")
			return
		}
		app.cetakSaham(results)
	case "2":
		index := app.cariSahamByKodeBinary(strings.ToUpper(keyword))
		if index == -1 {
			fmt.Println("Saham tidak ditemukan.")
			return
		}
		app.cetakSaham([]Saham{app.Saham[index]})
	default:
		fmt.Println("Metode pencarian tidak valid.")
	}
}

func (app *Aplikasi) cariSahamSequential(kataKunci string) []Saham {
	var hasil []Saham
	for _, s := range app.Saham {
		if strings.Contains(strings.ToLower(s.Kode), kataKunci) ||
			strings.Contains(strings.ToLower(s.NamaPerusahaan), kataKunci) {
			hasil = append(hasil, s)
		}
	}
	return hasil
}

func (app *Aplikasi) cariSahamByKodeBinary(kode string) int {
	awal := 0
	akhir := len(app.Saham) - 1
	for awal <= akhir {
		tengah := (awal + akhir) / 2
		if app.Saham[tengah].Kode == kode {
			return tengah
		} else if app.Saham[tengah].Kode < kode {
			awal = tengah + 1
		} else {
			akhir = tengah - 1
		}
	}
	return -1
}

func (app *Aplikasi) cetakSaham(saham []Saham) {
	fmt.Printf("%-8s %-30s %-15s %-15s\n", "Kode", "Nama Perusahaan", "Harga", "Volume Transaksi")
	for _, s := range saham {
		fmt.Printf("%-8s %-30s %-15s %-15d\n", s.Kode, s.NamaPerusahaan, formatRupiah(s.Harga), s.VolumeTransaksi)
	}
}

func (app *Aplikasi) menuUrutSaham() {
	fmt.Println("--- Urutkan Saham ---")
	fmt.Println("1. Selection Sort berdasarkan Harga Tertinggi")
	fmt.Println("2. Insertion Sort berdasarkan Harga Tertinggi")
	fmt.Println("3. Selection Sort berdasarkan Volume Transaksi")
	fmt.Println("4. Insertion Sort berdasarkan Volume Transaksi")
	fmt.Print("Pilih metode pengurutan: ")
	input := app.bacaInput()

	switch input {
	case "1":
		app.selectionSortHargaDescending()
		fmt.Println("Saham diurutkan berdasarkan harga tertinggi menggunakan Selection Sort:")
		app.cetakSaham(app.Saham)
	case "2":
		app.insertionSortHargaDescending()
		fmt.Println("Saham diurutkan berdasarkan harga tertinggi menggunakan Insertion Sort:")
		app.cetakSaham(app.Saham)
	case "3":
		app.selectionSortVolumeDescending()
		fmt.Println("Saham diurutkan berdasarkan volume transaksi menggunakan Selection Sort:")
		app.cetakSaham(app.Saham)
	case "4":
		app.insertionSortVolumeDescending()
		fmt.Println("Saham diurutkan berdasarkan volume transaksi menggunakan Insertion Sort:")
		app.cetakSaham(app.Saham)
	default:
		fmt.Println("Metode pengurutan tidak valid.")
	}

	if input == "3" || input == "4" {
		sort.Slice(app.Saham, func(i, j int) bool {
			return app.Saham[i].Kode < app.Saham[j].Kode
		})
	}
}

func (app *Aplikasi) selectionSortHargaDescending() {
	n := len(app.Saham)
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if app.Saham[j].Harga > app.Saham[idxMax].Harga {
				idxMax = j
			}
		}
		app.Saham[i], app.Saham[idxMax] = app.Saham[idxMax], app.Saham[i]
	}
}

func (app *Aplikasi) insertionSortHargaDescending() {
	n := len(app.Saham)
	for i := 1; i < n; i++ {
		key := app.Saham[i]
		j := i - 1
		for j >= 0 && app.Saham[j].Harga < key.Harga {
			app.Saham[j+1] = app.Saham[j]
			j--
		}
		app.Saham[j+1] = key
	}
}

func (app *Aplikasi) selectionSortVolumeDescending() {
	n := len(app.Saham)
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if app.Saham[j].VolumeTransaksi > app.Saham[idxMax].VolumeTransaksi {
				idxMax = j
			}
		}
		app.Saham[i], app.Saham[idxMax] = app.Saham[idxMax], app.Saham[i]
	}
}

func (app *Aplikasi) insertionSortVolumeDescending() {
	n := len(app.Saham)
	for i := 1; i < n; i++ {
		key := app.Saham[i]
		j := i - 1
		for j >= 0 && app.Saham[j].VolumeTransaksi < key.VolumeTransaksi {
			app.Saham[j+1] = app.Saham[j]
			j--
		}
		app.Saham[j+1] = key
	}
}

func (app *Aplikasi) tampilkanStatistikUntungRugi() {
	fmt.Println("--- Statistik Keuntungan dan Kerugian ---")
	if len(app.Portofolio.Transaksi) == 0 {
		fmt.Println("Tidak ada transaksi untuk dihitung keuntungan/kerugian.")
		return
	}

	type StatistikSaham struct {
		TotalBeliJumlah     int
		TotalBeliBiaya      float64
		TotalJualJumlah     int
		TotalJualPendapatan float64
	}

	stats := make(map[string]*StatistikSaham)
	for _, tx := range app.Portofolio.Transaksi {
		if _, ok := stats[tx.KodeSaham]; !ok {
			stats[tx.KodeSaham] = &StatistikSaham{}
		}
		s := stats[tx.KodeSaham]
		if tx.IsBeli {
			s.TotalBeliJumlah += tx.Jumlah
			s.TotalBeliBiaya += tx.Harga * float64(tx.Jumlah)
		} else {
			s.TotalJualJumlah += tx.Jumlah
			s.TotalJualPendapatan += tx.Harga * float64(tx.Jumlah)
		}
	}

	totalUntung := 0.0
	fmt.Printf("%-8s %-25s %-15s %-15s %-15s\n", "Kode", "Nama Perusahaan", "Harga Beli Rata2", "Harga Jual Rata2", "Untung/Rugi")
	for kode, s := range stats {
		rataBeli := 0.0
		rataJual := 0.0
		untung := 0.0
		if s.TotalBeliJumlah > 0 {
			rataBeli = s.TotalBeliBiaya / float64(s.TotalBeliJumlah)
		}
		if s.TotalJualJumlah > 0 {
			rataJual = s.TotalJualPendapatan / float64(s.TotalJualJumlah)
			untung = (rataJual - rataBeli) * float64(s.TotalJualJumlah)
		}
		idx := app.cariSahamByKodeBinary(kode)
		if idx == -1 {
			continue
		}
		fmt.Printf("%-8s %-25s %s %s %s\n",
			kode,
			app.Saham[idx].NamaPerusahaan,
			formatRupiah(rataBeli),
			formatRupiah(rataJual),
			formatRupiah(untung),
		)
		totalUntung += untung
	}
	fmt.Println("Total untung/rugi dari penjualan:", formatRupiah(totalUntung))
}

func (app *Aplikasi) bacaInput() string {
	val, _ := app.reader.ReadString('\n')
	return strings.TrimSpace(val)
}

func main() {
	app := BaruAplikasi()
	app.Jalankan()
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Saham struct {
	Kode            string
	NamaPerusahaan  string
	Harga           float64
	VolumeTransaksi int
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

	
	app.insertionSortByKode()
	return app
}



func (app *Aplikasi) simulasiPerubahanHarga() {
	rand.Seed(time.Now().UnixNano())
	for i := range app.Saham {
		perubahan := (rand.Float64() * 0.1) - 0.05
		hargaBaru := app.Saham[i].Harga * (1 + perubahan)
		if hargaBaru < 1.0 {
			hargaBaru = 1.0
		}
		app.Saham[i].Harga = hargaBaru
	}
}

func (app *Aplikasi) cetakSaham(saham []Saham) {
	fmt.Printf("%-8s %-30s %-15s %-15s\n", "Kode", "Nama Perusahaan", "Harga", "Volume Transaksi")
	for _, s := range saham {
		fmt.Printf("%-8s %-30s %-15s %-15d\n", s.Kode, s.NamaPerusahaan, formatRupiah(s.Harga), s.VolumeTransaksi)
	}
}

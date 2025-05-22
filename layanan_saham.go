package main

import (
	"fmt"
	"strings"
)

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

func (app *Aplikasi) menuUrutSaham() {
	fmt.Println("\n=== Urutkan Saham ===")
	fmt.Println("1. Selection Sort (Harga Tertinggi)")
	fmt.Println("2. Insertion Sort (Harga Tertinggi)")
	fmt.Println("3. Selection Sort (Volume Tertinggi)")
	fmt.Println("4. Insertion Sort (Volume Tertinggi)")
	fmt.Print("Pilih metode (1-4): ")
	input := app.bacaInput()

	// Simpan data asli untuk reset
	backupSaham := make([]Saham, len(app.Saham))
	copy(backupSaham, app.Saham)

	switch input {
	case "1":
		app.selectionSortHargaDescending()
		fmt.Println("\nHasil Selection Sort (Harga Tertinggi):")
		app.cetakSaham(app.Saham)
	case "2":
		app.insertionSortHargaDescending()
		fmt.Println("\nHasil Insertion Sort (Harga Tertinggi):")
		app.cetakSaham(app.Saham)
	case "3":
		app.selectionSortVolumeDescending()
		fmt.Println("\nHasil Selection Sort (Volume Tertinggi):")
		app.cetakSaham(app.Saham)
	case "4":
		app.insertionSortVolumeDescending()
		fmt.Println("\nHasil Insertion Sort (Volume Tertinggi):")
		app.cetakSaham(app.Saham)
	default:
		fmt.Println("Input tidak valid!")
		return
	}

	// Kembalikan ke urutan Kode jika sorting Volume
	if input == "3" || input == "4" {
		app.insertionSortByKode()
	}
}

// ===== Algoritma Sorting Manual =====

// Selection Sort (Harga Tertinggi → Terendah)
func (app *Aplikasi) selectionSortHargaDescending() {
	n := len(app.Saham)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if app.Saham[j].Harga > app.Saham[maxIdx].Harga {
				maxIdx = j
			}
		}
		app.Saham[i], app.Saham[maxIdx] = app.Saham[maxIdx], app.Saham[i]
	}
}

// Insertion Sort (Harga Tertinggi → Terendah)
func (app *Aplikasi) insertionSortHargaDescending() {
	for i := 1; i < len(app.Saham); i++ {
		key := app.Saham[i]
		j := i - 1
		for j >= 0 && app.Saham[j].Harga < key.Harga {
			app.Saham[j+1] = app.Saham[j]
			j--
		}
		app.Saham[j+1] = key
	}
}

// Selection Sort (Volume Tertinggi → Terendah)
func (app *Aplikasi) selectionSortVolumeDescending() {
	n := len(app.Saham)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if app.Saham[j].VolumeTransaksi > app.Saham[maxIdx].VolumeTransaksi {
				maxIdx = j
			}
		}
		app.Saham[i], app.Saham[maxIdx] = app.Saham[maxIdx], app.Saham[i]
	}
}

// Insertion Sort (Volume Tertinggi → Terendah)
func (app *Aplikasi) insertionSortVolumeDescending() {
	for i := 1; i < len(app.Saham); i++ {
		key := app.Saham[i]
		j := i - 1
		for j >= 0 && app.Saham[j].VolumeTransaksi < key.VolumeTransaksi {
			app.Saham[j+1] = app.Saham[j]
			j--
		}
		app.Saham[j+1] = key
	}
}

// Insertion Sort untuk mengembalikan urutan Kode (A → Z)
func (app *Aplikasi) insertionSortByKode() {
	for i := 1; i < len(app.Saham); i++ {
		key := app.Saham[i]
		j := i - 1
		for j >= 0 && app.Saham[j].Kode > key.Kode {
			app.Saham[j+1] = app.Saham[j]
			j--
		}
		app.Saham[j+1] = key
	}
}
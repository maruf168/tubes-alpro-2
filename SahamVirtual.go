package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Saham struct {
    Kode       string
    Nama       string
    Harga      int
    Perubahan  int
}

type Transaksi struct {
    KodeSaham string
    Jumlah    int
    Total     int
    Jenis     string // beli atau jual
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

var portofolio = make(map[string]int)
var riwayatTransaksi []Transaksi

func main() {
    rand.Seed(time.Now().UnixNano())
    for {
        fmt.Println("\n=== APLIKASI SIMULASI PASAR SAHAM VIRTUAL ===")
        fmt.Println("1. Lihat Daftar Saham")
        fmt.Println("2. Beli Saham")
        fmt.Println("3. Jual Saham")
        fmt.Println("4. Lihat Portofolio")
        fmt.Println("5. Cari Saham")
        fmt.Println("6. Keluar")
        fmt.Print("Pilih menu: ")

        var pilihan int
        fmt.Scan(&pilihan)

        switch pilihan {
        case 1:
            tampilkanDaftarSaham()
        case 2:
            beliSaham()
        case 3:
            jualSaham()
        case 4:
            tampilkanPortofolio()
        case 5:
            cariSaham()
        case 6:
            fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
            return
        default:
            fmt.Println("Pilihan tidak valid.")
        }
    }
}

func tampilkanDaftarSaham() {
    fmt.Println("\nKode\tNama\t\tHarga\tPerubahan")
    for i := range daftarSaham {
        perubahan := rand.Intn(200) - 100
        daftarSaham[i].Harga += perubahan
        daftarSaham[i].Perubahan = perubahan
        fmt.Printf("%s\t%s\t%d\t%d\n", daftarSaham[i].Kode, daftarSaham[i].Nama, daftarSaham[i].Harga, perubahan)
    }
}

func beliSaham() {
    var kode string
    var jumlah int
    fmt.Print("Masukkan kode saham yang ingin dibeli: ")
    fmt.Scan(&kode)
    fmt.Print("Masukkan jumlah saham yang ingin dibeli: ")
    fmt.Scan(&jumlah)

    for _, saham := range daftarSaham {
        if saham.Kode == kode {
            total := jumlah * saham.Harga
            portofolio[kode] += jumlah
            riwayatTransaksi = append(riwayatTransaksi, Transaksi{kode, jumlah, total, "beli"})
            fmt.Printf("Berhasil membeli %d saham %s seharga %d\n", jumlah, kode, total)
            return
        }
    }
    fmt.Println("Saham tidak ditemukan.")
}

func jualSaham() {
    var kode string
    var jumlah int
    fmt.Print("Masukkan kode saham yang ingin dijual: ")
    fmt.Scan(&kode)
    fmt.Print("Masukkan jumlah saham yang ingin dijual: ")
    fmt.Scan(&jumlah)

    if portofolio[kode] >= jumlah {
        for _, saham := range daftarSaham {
            if saham.Kode == kode {
                total := jumlah * saham.Harga
                portofolio[kode] -= jumlah
                riwayatTransaksi = append(riwayatTransaksi, Transaksi{kode, jumlah, total, "jual"})
                fmt.Printf("Berhasil menjual %d saham %s seharga %d\n", jumlah, kode, total)
                return
            }
        }
    } else {
        fmt.Println("Jumlah saham tidak mencukupi.")
    }
}

func tampilkanPortofolio() {
    fmt.Println("\n--- Portofolio Saham ---")
    for kode, jumlah := range portofolio {
        fmt.Printf("Kode: %s | Jumlah: %d\n", kode, jumlah)
    }
}

func cariSaham() {
    var kode string
    fmt.Print("Masukkan kode saham yang ingin dicari: ")
    fmt.Scan(&kode)

    found := false
    for _, saham := range daftarSaham {
        if saham.Kode == kode {
            fmt.Printf("Ditemukan: %s - %s | Harga: %d\n", saham.Kode, saham.Nama, saham.Harga)
            found = true
            break
        }
    }

    if !found {
        fmt.Println("Saham tidak ditemukan.")
    }
}


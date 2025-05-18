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

var daftarSaham = []Saham{
    {"BBCA", "Bank BCA", 8500, 0},
    {"TLKM", "Telkom Indonesia", 4200, 0},
    {"BBRI", "Bank BRI", 5100, 0},
    {"UNVR", "Unilever", 4100, 0},
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


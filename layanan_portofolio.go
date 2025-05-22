package main

import (
    "fmt"
    "strings"
)

func (app *Aplikasi) tampilkanPortofolio() {
    fmt.Println("--- Portofolio Anda ---")
    if len(app.Portofolio.Transaksi) == 0 {
        fmt.Println("Anda belum melakukan transaksi apapun.")
        return
    }

   
    pegang := make(map[string]int)
    for _, tx := range app.Portofolio.Transaksi {
        if tx.IsBeli {
            pegang[tx.KodeSaham] += tx.Jumlah
        } else {
            pegang[tx.KodeSaham] -= tx.Jumlah
        }
    }

	
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

func (app *Aplikasi) Jalankan() {
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
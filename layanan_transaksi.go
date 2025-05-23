package main

import (
    "fmt"
    "strconv"
	 "strings"
)

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
        fmt.Printf("%d) %s %d saham %s %s\n", i+1, jenis, tx.Jumlah, tx.KodeSaham, formatRupiah(tx.Harga))
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

    fmt.Printf("Transaksi terpilih: %s %d saham %s %s\n",
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
        fmt.Printf("%d) %s %d saham %s %s\n", i+1, jenis, tx.Jumlah, tx.KodeSaham, formatRupiah(tx.Harga))
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

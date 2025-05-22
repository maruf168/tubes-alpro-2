package main

import  "bufio"
   


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
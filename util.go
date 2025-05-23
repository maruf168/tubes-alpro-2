package main

import (
    "fmt"
    "strings"  
)

func formatRupiah(jumlah float64) string {
    return fmt.Sprintf("Rp %0.2f", jumlah)
}

func (app *Aplikasi) bacaInput() string {
    val, _ := app.reader.ReadString('\n')
    return strings.TrimSpace(val)  
}

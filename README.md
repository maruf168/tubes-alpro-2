# Tubes-Alpro-2

## aplikasi simulasi pasar saham virtual 

### Group 5
1. Ma'ruf Sarifudin (103112400128)  
2. RAFAEL ADITYA NUGROHO (103112400110)  

---

### Deskripsi Program

Aplikasi yang dikembangkan menggunakan bahasa pemrograman Go (Golang) ini adalah sebuah simulasi pasar saham virtual berbasis teks yang bertujuan untuk memberikan pengalaman belajar investasi saham secara sederhana dan interaktif bagi pengguna, khususnya mahasiswa dan pemula.
Aplikasi ini menggunakan struktur data array dan slice untuk menyimpan data saham dan mendukung fitur dasar pencarian dan pengurutan untuk memudahkan pengelolaan data saham.
Fitur utama aplikasi meliputi pendaftaran data saham (kode saham, nama perusahaan, harga, jumlah saham), pengelolaan data saham seperti mengubah dan menghapus data, serta pencarian saham menggunakan metode Sequential Search.
Pengurutan data dilakukan menggunakan algoritma dasar seperti Selection Sort atau Insertion Sort pada salah satu kriteria, misalnya berdasarkan harga atau nama perusahaan.
Aplikasi menggunakan menu interaktif berbasis teks yang sederhana dan mudah dipahami, dengan validasi input dasar untuk menjaga kelancaran penggunaan. Fitur beli dan jual saham, laporan portofolio, dan pencarian lanjutan seperti Binary Search masih dapat dikembangkan pada versi selanjutnya.

### Fitur

1.	Tambah Data Saham: Pengguna dapat memasukkan informasi kode saham, nama perusahaan, harga saham, dan jumlah saham yang tersedia.
2.	Ubah Data Saham: Memungkinkan pengguna memperbarui informasi saham yang sudah terdaftar.
3.	Hapus Data Saham: Menghapus data saham tertentu dari daftar.
4.	Pencarian Data: 
      •	Sequential Search untuk mencari saham berdasarkan nama perusahaan.
        (Biasanya belum ada Binary Search karena butuh data yang terurut dan fungsi sorting yang lengkap). 
5.	Pengurutan Data:
      •	Mungkin baru ada satu metode sorting saja, misal Selection Sort untuk harga atau Insertion Sort untuk nama, tapi belum keduanya.
6.	Menu Interaktif: Menyediakan navigasi menu dasar untuk mengakses fitur aplikasi.
7.	Keluar: Menghentikan program

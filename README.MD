# Employee Management API

API ini adalah aplikasi sederhana untuk mengelola data karyawan (employee). Dibangun menggunakan bahasa pemrograman Go dengan menggunakan framework Gin dan database PostgreSQL.

## Fitur

1. Menambahkan data karyawan baru.
2. Mengambil data semua karyawan.
3. Mengambil data karyawan berdasarkan ID.
4. Mengupdate data karyawan.
5. Menghapus data karyawan.

## Setup

1. Pastikan Anda telah menginstall Golang di komputer Anda.
2. Clone repository ini ke dalam direktori kerja Anda.
3. Buatlah file `.env` dan atur konfigurasi koneksi database PostgreSQL Anda.
4. Jalankan perintah berikut untuk menginstall semua dependensi: `go mod download`
5. Jalankan aplikasi dengan perintah: `go run main.go`

   Aplikasi akan berjalan di `http://localhost:8080`.

## Migrasi Tabel Employee

Untuk melakukan migrasi tabel Employee, ikuti langkah-langkah berikut:

1. Buka terminal atau command prompt.

2. Pindah ke direktori proyek Anda menggunakan perintah `cd` (Change Directory). Misalnya, jika proyek Anda berada di direktori `/path/to/project`, jalankan perintah berikut:

```
cd /path/to/project
```

3. Pastikan Anda sudah memiliki file `migrate.go` yang berisi kode migrasi. Jika belum, buatlah file tersebut dan letakkan kode migrasi di dalamnya.

4. Jalankan perintah berikut untuk menjalankan proses migrasi:

```shell
go run migrate/migrate.go
```

Pastikan Anda memberikan jalur yang benar ke file `migrate.go` berdasarkan struktur direktori proyek Anda.

5. Perintah `go run` akan menjalankan file `migrate.go` dan memulai proses migrasi tabel. Anda akan melihat pesan-pesan yang menunjukkan status migrasi, seperti "Migrating table Employee... Done!" atau pesan kesalahan jika terjadi masalah.

6. Setelah proses migrasi selesai, tabel-tabel yang didefinisikan dalam kode migrasi akan terbuat di database. Anda dapat menggunakan tabel-tabel tersebut dalam aplikasi Anda.

Pastikan Anda sudah mengganti `/path/to/project` dengan jalur direktori proyek Anda yang sesungguhnya.
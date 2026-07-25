Studi Kasus: Data Mata Kuliah

Sebuah kampus ingin membuat aplikasi sederhana untuk menampilkan informasi mata kuliah berdasarkan kode mata kuliah.

Ketentuan:
1. Buat sebuah struct MataKuliah yang memiliki atribut:
Kode
Nama
SKS

2. Pada Repository, simpan data mata kuliah berikut menggunakan slice:
Kode   Nama Mata Kuliah   SKS
IF201  Struktur Data      3
IF202  Basis Data         3
IF203  Pemrograman Web    3
IF204  Jaringan Komputer  3
IF205  Sistem Operasi     3

3. Pada Service, buat fungsi untuk mencari data mata kuliah berdasarkan kode yang diterima dari Handler.

4. Pada Handler, ambil parameter kode dari URL, kemudian panggil Service untuk mendapatkan data mata kuliah.

5. Jika data mata kuliah ditemukan, tampilkan informasi mata kuliah tersebut.
Jika data mata kuliah tidak ditemukan, tampilkan:Data mata kuliah tidak ditemukan.

untuk service gunakan perulangan untuk cari datanya
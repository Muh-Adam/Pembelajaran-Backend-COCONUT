1. Buatlah sebuah program Go dengan ketentuan berikut:

2. Buat fugsi bernama cekUmur(umur int).
Jika nilai umur kurang dari 17, tampilkan panic dengan pesan "Umur belum mencukupi".

3. Gunakan recover() agar program tidak berhenti ketika terjadi panic, lalu tampilkan pesan :  "Terjadi kesalahan : umur belum mencukupi".

4. Gunakan defer untuk selalu menampilkan pesan : "Pemeriksaan selesai"

5. Pada fungsi main(), panggil fungsi cekUmur() dan setelah itu tampilkan program selesai
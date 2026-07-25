Buat fungsi validasiPassword(pass string).
1. Jika len(pass) < 6, panic dengan pesan "Password terlalu pendek".
2. Gunakan recover() agar program tidak berhenti, tampilkan: "Validasi gagal : password terlalu pendek"
3. Gunakan defer untuk menampilkan: "Log validasi tersimpan"
4. Pada main(), pangil validasiPassword("abc") lalu tampilkan "silahkan coba lagi"

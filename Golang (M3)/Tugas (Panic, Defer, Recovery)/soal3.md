Buat fungsi tarikSaldo(saldo, jumlah int).
1. Jika jumlah > saldo, panic dengan pesan "saldo tidak mencukupi"
2. Gunakan recover() agar program tetap jalan, tampilkan : "Transaksi ditolak : saldo tidak mencukupi".
3. Gunakan defer untuk selalu menampilkan : "Rekening di cek"

4. Pada main(), panggil tarikSaldo(50000, 100000) lalu tampilkan "Terima kasih telah meenggunakan layanan kami".


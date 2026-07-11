-- ADVANCED SQL, DATA PROCCESSING, DAN DATABASE SECURITY

CREATE DATABASE Akademik;

USE Akademik;

CREATE TABLE Mahasiswa (
    NIM varchar(10) PRIMARY KEY,
    mahasiswa varchar(100),
    jurusan varchar(50)
);

INSERT INTO Mahasiswa (NIM, mahasiswa, jurusan) VALUES ('D1234567', 'Noah', 'Teknik Informatika');
INSERT INTO Mahasiswa (NIM, mahasiswa, jurusan) VALUES ('D3949182', 'Sandy', 'Teknik Mesin');
INSERT INTO Mahasiswa (NIM, mahasiswa, jurusan) VALUES ('D7530929', 'Satria', 'Teknik Sipil');

-- ## ALTER TABLE (Mengubah Struktur Table)

-- ADD COLUMN
ALTER TABLE Mahasiswa
ADD angkatan INT;

-- MENGUBAH TIPE DATA

ALTER TABLE Mahasiswa
MODIFY jurusan VARCHAR(100);

-- MENGUBAH NAMA KOLOM

ALTER TABLE Mahasiswa
CHANGE mahasiswa nama_mahasiswa varchar(100);

-- MENGHAPUS KOLOM

ALTER TABLE Mahasiswa
DROP COLUMN angkatan;

ALTER TABLE Mahasiswa 
ADD email varchar(100) unique;

UPDATE Mahasiswa SET email = 'noah@gmail.com' WHERE NIM = 'D1234567';
UPDATE Mahasiswa SET email = 'sandy@gmail.com' WHERE NIM = 'D3949182';
UPDATE Mahasiswa SET email = 'satria@gmail.com' WHERE NIM = 'D7530929';


-- ##  UPDATE DENGAN PERHITUNGAN (Menggunakan Operasi Aritmetika)

-- PENJUMLAHAN
ALTER TABLE Mahasiswa
ADD SKS int;

UPDATE Mahasiswa
SET SKS = 20;

UPDATE Mahasiswa
SET SKS = SKS + 6;

-- PNGURNGAN
ALTER TABLE Mahasiswa
ADD SKS int;

UPDATE Mahasiswa
SET SKS = 20;

UPDATE Mahasiswa
SET SKS = SKS - 6;

-- PERKALIAN
ALTER TABLE Mahasiswa
ADD SKS int;

UPDATE Mahasiswa
SET SKS = 20;

UPDATE Mahasiswa
SET SKS = SKS * 6;

-- PEMBAGIAN
ALTER TABLE Mahasiswa
ADD SKS int;

UPDATE Mahasiswa
SET SKS = 20;

UPDATE Mahasiswa
SET SKS = SKS / 6;



-- ADD Kolom Semester
ALTER TABLE Mahasiswa
ADD Semester int;

UPDATE Mahasiswa
SET semester = 3;

UPDATE Mahasiswa
SET semester = semester + semester + semester WHERE NIM = 'D1234567';

UPDATE Mahasiswa
SET semester = semester - 2 WHERE NIM = 'D7530929';


-- ## CREATE TABLE AS SELECT (Membuat Tabel Baru Berdasarkan Hasil Query) -- Table Sementara Berdasarkan Data Yang Sudah Ada

-- SALIN SEMUA DATA
CREATE TABLE Backup_Mahasiswa AS
SELECT * FROM Mahasiswa;

-- SALIN KOLOM TERTENTU
CREATE TABLE NIM_Mahasiswa AS
SELECT NIM FROM Mahasiswa;

CREATE TABLE Daftar_Mahasiswa AS
SELECT NIM, nama_mahasiswa FROM Mahasiswa;

-- SALIN BERDASARKAN KONDISI
CREATE TABLE Informatika_Mahasiswa AS
SELECT NIM, nama_mahasiswa FROM Mahasiswa WHERE jurusan = 'Teknik Informatika';

-- SALIN TANPA DATA
CREATE TABLE struktur AS
SELECT * FROM Mahasiswa WHERE 1=0;

INSERT INTO struktur (NIM) VALUES ('1');


-- ## INDEXING (Struktur Untuk Menemukan Data Lebih Cepat)

-- MEMBUAT INDEX
CREATE INDEX idx_jurusan ON mahasiswa(jurusan);

-- MELIHAT INDEX
SHOW INDEX FROM mahasiswa;

-- MENGHAPUS INDEX
DROP INDEX idx_jurusan ON Mahasiswa;
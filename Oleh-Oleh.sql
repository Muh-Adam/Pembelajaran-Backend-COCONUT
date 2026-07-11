CREATE database inventaris_db;

CREATE TABLE barang (
    id int PRIMARY KEY AUTO_INCREMENT,
    nama_barang varchar(100),
    jumlah int,
    harga int,
    supplier varchar(100)
);

CREATE TABLE kategori (
    id int PRIMARY KEY AUTO_INCREMENT,
    nama_kategori varchar(100),
    deskripsi text
);

CREATE TABLE supplier (
    id int PRIMARY KEY AUTO_INCREMENT,
    nama_supplier varchar(100),
    alamat varchar(100),
    no_telepon varchar(100)
);

CREATE TABLE log_aktivitas (
    id int PRIMARY KEY AUTO_INCREMENT,
    nama_aktivitas varchar(100),
    waktu datetime
);

-- ## INSERT
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Baju Adat', 100, 500000, 'Supplier A');
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Kipas Tangan', 200, 25000, 'Supplier B');
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Topi Khas', 150, 75000, 'Supplier C');
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Patung Kayu', 50, 150000, 'Supplier D');
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Tas Anyaman', 300, 125000, 'Supplier E');
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Sepatu Kulit', 75, 450000, 'Supplier F');
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Keramik Hias', 125, 350000, 'Supplier G');
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Kalung Etnik', 250, 175000, 'Supplier H');
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Sarung Tenun', 80, 275000, 'Supplier I');
INSERT INTO barang (nama_barang, jumlah, harga, supplier) VALUES ('Tasbih Kayu', 150, 85000, 'Supplier J');

INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Pakaian', 'Baju dan aksesoris pakaian');
INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Kerajinan Tangan', 'Barang hasil kerajinan tangan');
INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Sovenir', 'Barang kenang-kenangan');
INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Makanan', 'Makanan khas daerah');
INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Minuman', 'Minuman khas daerah');
INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Aksesoris', 'Aksesoris pelengkap');
INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Peralatan Rumah Tangga', 'Barang untuk rumah tangga');
INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Alat Musik', 'Alat musik tradisional');
INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Mainan', 'Mainan tradisional');
INSERT INTO kategori (nama_kategori, deskripsi) VALUES ('Peralatan Elektronik', 'Barang elektronik');

INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier A', 'Jl. Sudirman No. 1', '08123456789');
INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier B', 'Jl. Thamrin No. 2', '08123456788');
INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier C', 'Jl. Gatot Subroto No. 3', '08123456787');
INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier D', 'Jl. Sudirman No. 4', '08123456786');
INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier E', 'Jl. Thamrin No. 5', '08123456785');
INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier F', 'Jl. Gatot Subroto No. 6', '08123456784');
INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier G', 'Jl. Sudirman No. 7', '08123456783');
INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier H', 'Jl. Thamrin No. 8', '08123456782');
INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier I', 'Jl. Gatot Subroto No. 9', '08123456781');
INSERT INTO supplier (nama_supplier, alamat, no_telepon) VALUES ('Supplier J', 'Jl. Sudirman No. 10', '08123456780');

INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('Admin menambahkan data', NOW());
INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('Admin mengupdate data', NOW());
INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('Admin menghapus data', NOW());
INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('Staff menambahkan data', NOW());
INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('Staff mengupdate data', NOW());
INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('Staff menghapus data', NOW());
INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('Auditor menambahkan data', NOW());
INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('Auditor mengupdate data', NOW());
INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('Auditor menghapus data', NOW());
INSERT INTO log_aktivitas (nama_aktivitas, waktu) VALUES ('User menambahkan data', NOW());

CREATE USER 'admin'@'localhost'

GRANT ALL PRIVILEGES ON inventaris_db.* TO 'admin'@'localhost';

CREATE USER 'staff'@'localhost'

GRANT SELECT, INSERT ON inventaris_db.* TO 'staff'@'localhost';

CREATE USER 'auditor'@'localhost'

GRANT SELECT ON inventaris_db.* TO 'auditor'@'localhost';
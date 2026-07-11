-- ## MEMBUAT USER
CREATE USER 'dosen'@'localhost'
IDENTIFIED BY 'dosen123';

SELECT user, host FROM mysql.user;

DROP USER 'dosen'@'localhost';

CREATE USER 'mahasiswa'@'localhost' IDENTIFIED BY 'mahasiswa1';

-- ## GRANT PRIVILEGE (MEMBATASI DAN MEMBERIKAN HAK AKSES)
GRANT SELECT ON Akademik.Mahasiswa TO 'dosen'@'localhost';

GRANT UPDATE, DELETE ON Akademik.Mahasiswa TO 'dosen'@'localhost';

SHOW GRANTS FOR 'dosen'@'localhost';

GRANT SELECT ON Akademik.Mahasiswa TO 'dosen'@'localhost', 'mahasiswa'@'localhost'; 

GRANT ALL PRIVILEGES ON Akademik.* TO 'dosen'@'localhost', 'mahasiswa'@'localhost'; 

-- ## REVOKE (CABUT HAK AKSES)
REVOKE DELETE ON Akademik.Mahasiswa FROM 'dosen'@'localhost';

-- ## UBAH USERNAME/PASSWORD (Butuh Akses Dari Root)
RENAME USER 'dosen'@'localhost' TO 'orang'@'localhost';

ALTER USER 'dosen'@'localhost' IDENTIFIED BY '123456';
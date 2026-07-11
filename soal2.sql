CREATE USER 'operator'@'localhost' IDENTIFIED BY 'op123';

GRANT SELECT, INSERT ON Akademik.Mahasiswa TO 'operator'@'localhost';

SELECT user, host FROM mysql.user;

SHOW GRANTS FOR 'operator'@'localhost';

REVOKE INSERT ON Akademik.Mahasiswa FROM 'operator'@'localhost';

SHOW GRANTS FOR 'operator'@'localhost';
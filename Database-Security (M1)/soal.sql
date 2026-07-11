ALTER TABLE Mahasiswa ADD IPK  DECIMAL(3,2);

UPDATE Mahasiswa SET IPK = 3.00;

UPDATE Mahasiswa SET IPK = IPK + 0.25;

UPDATE Mahasiswa SET IPK = IPK + 0.25 WHERE jurusan = 'Teknik Informatika';

ALTER TABLE Mahasiswa MODIFY jurusan VARCHAR(250);

CREATE TABLE mahasiswa_berprestasi AS SELECT NIM, nama_mahasiswa, jurusan, IPK FROM Mahasiswa WHERE IPK > 3.30;

CREATE INDEX idx_jurusan ON Mahasiswa(jurusan);

SHOW INDEX FROM Mahasiswa;
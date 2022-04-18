CREATE TABLE `Penyakit`
(
    `NamaPenyakit` VARCHAR(100),
    `SequenceDNA` VARCHAR(1000) NOT NULL,
    PRIMARY KEY (`NamaPenyakit`)
);

CREATE Table `Hasil`
(
    `Tanggal` DATE NOT NULL,
    `NamaPengguna` VARCHAR(100) NOT NULL,
    `NamaPenyakit` VARCHAR(100) NOT NULL,
    `Persentase` TINYINT(3) NOT NULL,
    `Hasil` TINYINT(1) NOT NULL,
    FOREIGN KEY (`NamaPenyakit`) REFERENCES `Penyakit` (`NamaPenyakit`)
);

INSERT INTO `Penyakit` (NamaPenyakit, SequenceDNA)
    VALUES
        ('HIV', 'ATTCGTAACTAGTAAGTTA');

INSERT INTO `Hasil` (Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil)
    VALUES
        ('2022-04-01', 'Mhs IF', 'HIV', 75, False);
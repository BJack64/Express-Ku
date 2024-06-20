-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 20, 2024 at 01:48 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `express-ku`
--

-- --------------------------------------------------------

--
-- Table structure for table `bukti`
--

CREATE TABLE `bukti` (
  `id_tiket` int(255) NOT NULL,
  `tgl_pembelian` varchar(255) NOT NULL,
  `id_jadwal` varchar(255) NOT NULL,
  `nama_penumpang` varchar(255) NOT NULL,
  `no_telp` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `bukti`
--

INSERT INTO `bukti` (`id_tiket`, `tgl_pembelian`, `id_jadwal`, `nama_penumpang`, `no_telp`) VALUES
(1, '2024-06-17', 'JA001', 'Giovanni Daniel Setiadi', '12345678910'),
(2, '2024-06-17', 'JA005', 'Giovanni Daniel Setiadi', '12345678');

-- --------------------------------------------------------

--
-- Table structure for table `jadwal`
--

CREATE TABLE `jadwal` (
  `id_jadwal` varchar(255) NOT NULL,
  `id_kereta` varchar(255) NOT NULL,
  `id_masinis` varchar(255) NOT NULL,
  `id_stasiun_asal` varchar(255) NOT NULL,
  `id_stasiun_tujuan` varchar(255) NOT NULL,
  `tgl_berangkat` varchar(255) NOT NULL,
  `waktu_berangkat` varchar(255) NOT NULL,
  `tgl_tiba` varchar(255) NOT NULL,
  `waktu_tiba` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `jadwal`
--

INSERT INTO `jadwal` (`id_jadwal`, `id_kereta`, `id_masinis`, `id_stasiun_asal`, `id_stasiun_tujuan`, `tgl_berangkat`, `waktu_berangkat`, `tgl_tiba`, `waktu_tiba`) VALUES
('JA001', 'KAI002', 'MA002', 'ST002', 'ST001', '2024-06-18', '10:00 WIB', '2024-06-19', '00:00 WIB'),
('JA002', 'KAI001', 'MA001', 'ST001', 'ST003', '2024-06-19', '10:00 WIB', '2024-06-19', '12:00 WIB'),
('JA003', 'KAI003', 'MA003', 'ST002', 'ST003', '2024-06-18', '12:00 WIB', '2024-06-18', '15:00 WIB'),
('JA004', 'KAI004', 'MA004', 'ST003', 'ST004', '2024-06-19', '22:00 WIB', '2024-06-20', '02:00 WIB'),
('JA005', 'KAI005', 'MA005', 'ST004', 'ST005', '2024-06-19', '12:00 WIB', '2024-06-19', '22:00 WIB'),
('JA006', 'KAI006', 'MA006', 'ST005', 'ST006', '2024-06-20', '10:00 WIB', '2024-06-20', '22:00 WIB'),
('JA007', 'KAI007', 'MA007', 'ST006', 'ST007', '2024-06-20', '23:00 WIB', '2024-06-21', '05:00 WIB'),
('JA008', 'KAI008', 'MA008', 'ST007', 'ST008', '2024-06-18', '07:00 WIB', '2024-06-18', '12:00 WIB'),
('JA009', 'KAI009', 'MA009', 'ST008', 'ST009', '2024-06-20', '02:00 WIB', '2024-06-20', '07:00 WIB'),
('JA010', 'KAI010', 'MA010', 'ST009', 'ST008', '2024-06-21', '13:00 WIB', '2024-06-21', '16:00 WIB');

-- --------------------------------------------------------

--
-- Table structure for table `kereta`
--

CREATE TABLE `kereta` (
  `id_kereta` varchar(255) NOT NULL,
  `nama_kereta` varchar(255) NOT NULL,
  `kelas` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `kereta`
--

INSERT INTO `kereta` (`id_kereta`, `nama_kereta`, `kelas`) VALUES
('KAI001', 'Kereta Api Pandawa', 'Ekonomi'),
('KAI002', 'Kereta Api Punakawan', 'Bisnis'),
('KAI003', 'Kereta Api Kertajaya', 'Ekonomi'),
('KAI004', 'Kereta Api Malabar', 'Eksekutif'),
('KAI005', 'Kereta Api Panembangan', 'Bisnis'),
('KAI006', 'Kereta Api Airlangga', 'Ekonomi'),
('KAI007', 'Kereta Api Brawijaya', 'Eksekutif'),
('KAI008', 'Kereta Api Jayabaya', 'Ekonomi'),
('KAI009', 'Kereta Api Kertanegara', 'Ekonomi'),
('KAI010', 'Kereta Api Siliwangi', 'Ekonomi');

-- --------------------------------------------------------

--
-- Table structure for table `masinis`
--

CREATE TABLE `masinis` (
  `id_masinis` varchar(255) NOT NULL,
  `nama_masinis` varchar(255) NOT NULL,
  `email_masinis` varchar(255) NOT NULL,
  `gender_masinis` varchar(255) NOT NULL,
  `exp_masinis` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `masinis`
--

INSERT INTO `masinis` (`id_masinis`, `nama_masinis`, `email_masinis`, `gender_masinis`, `exp_masinis`) VALUES
('MA001', 'Duda Permana', 'duda_permana@yahoo.com', 'Pria', '7 Tahun'),
('MA002', 'Rina Ambyar', 'rina_ambyar@gmail.com', 'Wanita', '15 Tahun'),
('MA003', 'Krida Permana', 'krida_permana@gmail.com', 'Wanita', '12 Tahun'),
('MA004', 'Ridho Gusti', 'ridho_gusti@gmail.com', 'Pria', '3 Tahun'),
('MA005', 'Sena M.Yantung', 'senam_yantung@yahoo.com', 'Pria', '25 Tahun'),
('MA006', 'Siti Astuti', 'siti_astuti@gmail.com', 'Wanita', '20 Tahun'),
('MA007', 'Rini Paini', 'rini_paini@gmail.com', 'Wanita', '12 Tahun'),
('MA008', 'Projo Admojoyo', 'projo_admojoyo@gmail.com', 'Pria', '15 Tahun'),
('MA009', 'Tito Asmoro', 'tito_asmoro@yahoo.com', 'Pria', '5 Tahun'),
('MA010', 'Puji Admuji', 'puji_admuji@gmail.com', 'Pria', '8 Tahun');

-- --------------------------------------------------------

--
-- Table structure for table `stasiun`
--

CREATE TABLE `stasiun` (
  `id_stasiun` varchar(255) NOT NULL,
  `nama_stasiun` varchar(255) NOT NULL,
  `kota_stasiun` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `stasiun`
--

INSERT INTO `stasiun` (`id_stasiun`, `nama_stasiun`, `kota_stasiun`) VALUES
('ST001', 'Stasiun Kendalsari', 'Banyuwangi'),
('ST002', 'Stasiun Prambanan', 'Malang'),
('ST003', 'Stasiun Bogorami', 'Surabaya'),
('ST004', 'Stasiun Ambarawa', 'Mojokerto'),
('ST005', 'Stasiun Blimbing', 'Batu'),
('ST006', 'Stasiun Tembangan', 'Lawang'),
('ST007', 'Stasiun Banjarmasin', 'Lumajang'),
('ST008', 'Stasiun Kidungkandang', 'Probolinggo'),
('ST009', 'Stasiun Padjajaran', 'Tulungagung'),
('ST010', 'Stasiun Blambangan', 'Kediri');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`) VALUES
(1, 'Admin', 'admin');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `bukti`
--
ALTER TABLE `bukti`
  ADD PRIMARY KEY (`id_tiket`),
  ADD KEY `bukti_id_jadwal_foreign` (`id_jadwal`);

--
-- Indexes for table `jadwal`
--
ALTER TABLE `jadwal`
  ADD PRIMARY KEY (`id_jadwal`),
  ADD KEY `jadwal_id_kereta_foreign` (`id_kereta`),
  ADD KEY `jadwal_id_masinis_foreign` (`id_masinis`),
  ADD KEY `jadwal_id_stasiun_asal_foreign` (`id_stasiun_asal`),
  ADD KEY `jadwal_id_stasiun_tujuan_foreign` (`id_stasiun_tujuan`);

--
-- Indexes for table `kereta`
--
ALTER TABLE `kereta`
  ADD PRIMARY KEY (`id_kereta`);

--
-- Indexes for table `masinis`
--
ALTER TABLE `masinis`
  ADD PRIMARY KEY (`id_masinis`);

--
-- Indexes for table `stasiun`
--
ALTER TABLE `stasiun`
  ADD PRIMARY KEY (`id_stasiun`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `bukti`
--
ALTER TABLE `bukti`
  MODIFY `id_tiket` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `bukti`
--
ALTER TABLE `bukti`
  ADD CONSTRAINT `bukti_id_jadwal_foreign` FOREIGN KEY (`id_jadwal`) REFERENCES `jadwal` (`id_jadwal`) ON DELETE CASCADE;

--
-- Constraints for table `jadwal`
--
ALTER TABLE `jadwal`
  ADD CONSTRAINT `jadwal_id_kereta_foreign` FOREIGN KEY (`id_kereta`) REFERENCES `kereta` (`id_kereta`) ON DELETE CASCADE,
  ADD CONSTRAINT `jadwal_id_masinis_foreign` FOREIGN KEY (`id_masinis`) REFERENCES `masinis` (`id_masinis`) ON DELETE CASCADE,
  ADD CONSTRAINT `jadwal_id_stasiun_asal_foreign` FOREIGN KEY (`id_stasiun_asal`) REFERENCES `stasiun` (`id_stasiun`) ON DELETE CASCADE,
  ADD CONSTRAINT `jadwal_id_stasiun_tujuan_foreign` FOREIGN KEY (`id_stasiun_tujuan`) REFERENCES `stasiun` (`id_stasiun`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

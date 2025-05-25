-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 25, 2025 at 06:01 PM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `core_chat`
--

-- --------------------------------------------------------

--
-- Table structure for table `person`
--

CREATE TABLE `person` (
  `identifier` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('admin','user') NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(255) NOT NULL,
  `email_validated` tinyint(1) NOT NULL DEFAULT 0,
  `date_of_birth` date NOT NULL,
  `description` text DEFAULT NULL,
  `picture_path` varchar(255) NOT NULL DEFAULT 'upload\\image\\profile\\default.png',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `person`
--

INSERT INTO `person` (`identifier`, `password`, `role`, `name`, `email`, `email_validated`, `date_of_birth`, `description`, `picture_path`, `created_at`) VALUES
('darren1', '$2a$10$Y7.kVj3aLFJUlA9VdFmW8e.xrNS1nQcF29yRQU5zMGSUE22Q3Uyq6', 'user', 'Darren1', 'darren1@email.com', 0, '2000-01-01', '', 'upload\\image\\profile\\darren1.jpg', '2025-05-25 14:56:06'),
('test', '$2a$10$PMmdIf03p/LzeSET6tQZUOpDBNDW1RGQ2Uhjvmllt/rJqalo7TPNa', 'user', 'test', 'test@email.com', 1, '2000-12-01', NULL, 'user', '2025-05-18 14:30:46');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `person`
--
ALTER TABLE `person`
  ADD PRIMARY KEY (`identifier`),
  ADD UNIQUE KEY `email` (`email`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

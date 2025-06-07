-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 07, 2025 at 12:59 PM
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
-- Table structure for table `chat_image`
--

CREATE TABLE `chat_image` (
  `id` int(11) NOT NULL,
  `sender` varchar(50) NOT NULL,
  `receiver` varchar(50) NOT NULL,
  `image_path` text NOT NULL,
  `inserted_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `chat_image`
--

INSERT INTO `chat_image` (`id`, `sender`, `receiver`, `image_path`, `inserted_at`) VALUES
(1, 'test', 'darren1', 'upload\\image\\chat\\14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-06 02:29:19');

-- --------------------------------------------------------

--
-- Table structure for table `chat_message`
--

CREATE TABLE `chat_message` (
  `id` bigint(20) NOT NULL,
  `sender` varchar(50) NOT NULL,
  `receiver` varchar(50) NOT NULL,
  `type` varchar(50) NOT NULL,
  `title` varchar(255) DEFAULT '',
  `body` text DEFAULT '',
  `payload` text DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `read_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `chat_message`
--

INSERT INTO `chat_message` (`id`, `sender`, `receiver`, `type`, `title`, `body`, `payload`, `created_at`, `read_at`) VALUES
(1, 'test', 'darren1', 'chat_message', '', '', '', '2025-06-07 10:29:03', NULL);

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
-- Indexes for table `chat_image`
--
ALTER TABLE `chat_image`
  ADD PRIMARY KEY (`id`),
  ADD KEY `sender` (`sender`),
  ADD KEY `receiver` (`receiver`);

--
-- Indexes for table `chat_message`
--
ALTER TABLE `chat_message`
  ADD PRIMARY KEY (`id`),
  ADD KEY `sender` (`sender`),
  ADD KEY `receiver` (`receiver`);

--
-- Indexes for table `person`
--
ALTER TABLE `person`
  ADD PRIMARY KEY (`identifier`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `chat_image`
--
ALTER TABLE `chat_image`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `chat_message`
--
ALTER TABLE `chat_message`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `chat_image`
--
ALTER TABLE `chat_image`
  ADD CONSTRAINT `chat_image_ibfk_1` FOREIGN KEY (`sender`) REFERENCES `person` (`identifier`),
  ADD CONSTRAINT `chat_image_ibfk_2` FOREIGN KEY (`receiver`) REFERENCES `person` (`identifier`);

--
-- Constraints for table `chat_message`
--
ALTER TABLE `chat_message`
  ADD CONSTRAINT `chat_message_ibfk_1` FOREIGN KEY (`sender`) REFERENCES `person` (`identifier`),
  ADD CONSTRAINT `chat_message_ibfk_2` FOREIGN KEY (`receiver`) REFERENCES `person` (`identifier`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

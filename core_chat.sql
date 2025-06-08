-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 08, 2025 at 03:36 PM
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
(1, 'test', 'darren1', 'chat_message', '', '', '', '2025-06-07 10:29:03', '2025-06-08 11:04:48'),
(2, 'test', 'darren1', 'image', '', 'test image', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-07 11:08:11', '2025-06-08 10:59:02'),
(3, 'test', 'darren1', 'image', '', 'test image 2', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 09:09:53', '2025-06-08 11:15:54'),
(4, 'test', 'darren1', 'chat_message', '', 'test message', '', '2025-06-08 09:15:35', '2025-06-08 11:16:23'),
(5, 'test', 'darren1', 'image', '', 'test dua koneksi', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:22:07', '2025-06-08 11:17:21'),
(6, 'test', 'darren1', 'image', '', 'test dua koneksi 2 after update', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:26:04', NULL),
(7, 'test', 'darren1', 'image', '', 'test dua koneksi 3 after update', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:26:26', NULL),
(8, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room after update', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:28:47', NULL),
(9, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room 2 after update', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:30:12', NULL),
(10, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room, 3 after update', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:31:02', NULL),
(11, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room, 4 after update', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:33:52', NULL),
(12, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room, 5 after update', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:34:16', NULL),
(13, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room, 6 after update logic', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:50:26', NULL),
(14, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room, 7 after update logic', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:50:50', '2025-06-08 11:02:37'),
(15, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room, 8 after update logic, harusnya bener ini', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 10:51:39', '2025-06-08 10:51:39'),
(16, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room, 8 after update logic, harusnya bener ini', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 11:18:48', '2025-06-08 11:18:48'),
(17, 'test', 'darren1', 'image', '', 'test dua koneksi dengan join room, 8 after update logic, harusnya bener ini', 'localhost:8888/chat/image?image_name=14375fbb-2295-4a3a-ad20-c39b5e3f6b74.jpg', '2025-06-08 11:19:37', NULL);

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
  `email_validated` tinyint(1) NOT NULL DEFAULT 1,
  `date_of_birth` date NOT NULL,
  `description` text DEFAULT NULL,
  `picture_path` varchar(255) NOT NULL DEFAULT 'upload\\image\\profile\\default.png',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `person`
--

INSERT INTO `person` (`identifier`, `password`, `role`, `name`, `email`, `email_validated`, `date_of_birth`, `description`, `picture_path`, `created_at`) VALUES
('darren1', '$2a$10$Y7.kVj3aLFJUlA9VdFmW8e.xrNS1nQcF29yRQU5zMGSUE22Q3Uyq6', 'user', 'Darren1', 'darren1@email.com', 1, '2000-01-01', '', 'upload\\image\\profile\\darren1.jpg', '2025-05-25 14:56:06'),
('test', '$2a$10$PMmdIf03p/LzeSET6tQZUOpDBNDW1RGQ2Uhjvmllt/rJqalo7TPNa', 'user', 'test', 'test@email.com', 1, '2000-12-01', NULL, 'user', '2025-05-18 14:30:46');

-- --------------------------------------------------------

--
-- Table structure for table `post`
--

CREATE TABLE `post` (
  `id` bigint(20) NOT NULL,
  `author` varchar(50) NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `post`
--

INSERT INTO `post` (`id`, `author`, `title`, `content`, `created_at`) VALUES
(1, 'darren1', 'darren1 first post', 'lorem ipsum', '2025-06-08 13:23:07'),
(2, 'darren1', 'darren1 second post', 'dolor amet', '2025-06-08 13:23:56'),
(3, 'test', 'test first post', '1st post', '2025-06-08 13:24:48'),
(4, 'test', 'test second post', '2nd post', '2025-06-08 13:25:02'),
(5, 'test', '', '', '2025-06-08 13:28:49'),
(6, 'test', 't', 't', '2025-06-08 13:31:45');

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
-- Indexes for table `post`
--
ALTER TABLE `post`
  ADD PRIMARY KEY (`id`),
  ADD KEY `author` (`author`);

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
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `post`
--
ALTER TABLE `post`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

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

--
-- Constraints for table `post`
--
ALTER TABLE `post`
  ADD CONSTRAINT `post_ibfk_1` FOREIGN KEY (`author`) REFERENCES `person` (`identifier`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

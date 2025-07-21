-- -------------------------------------------------------------
-- TablePlus 6.6.8(632)
--
-- https://tableplus.com/
--
-- Database: technical_test
-- Generation Time: 2025-07-21 21:36:47.4580
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP TABLE IF EXISTS `histories`;
CREATE TABLE `histories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint(20) unsigned DEFAULT NULL,
  `table_name` varchar(100) DEFAULT NULL,
  `table_id` varchar(50) DEFAULT NULL,
  `data` text DEFAULT NULL,
  `type` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_histories_deleted_at` (`deleted_at`),
  KEY `idx_history` (`table_name`,`table_id`,`type`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `traditional_foods`;
CREATE TABLE `traditional_foods` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint(20) unsigned DEFAULT NULL,
  `updated_by_id` bigint(20) unsigned DEFAULT NULL,
  `deleted_by_id` bigint(20) unsigned DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `regional_origin` varchar(50) DEFAULT NULL,
  `main_ingredient` longtext DEFAULT NULL,
  `type` varchar(20) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_traditiona_food` (`name`,`regional_origin`),
  KEY `idx_traditional_foods_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint(20) unsigned DEFAULT NULL,
  `updated_by_id` bigint(20) unsigned DEFAULT NULL,
  `deleted_by_id` bigint(20) unsigned DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `role` varchar(15) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user` (`name`,`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

INSERT INTO `histories` (`id`, `created_at`, `updated_at`, `deleted_at`, `created_by_id`, `table_name`, `table_id`, `data`, `type`) VALUES
(1, '2025-07-20 11:41:45.758', '2025-07-20 11:41:45.758', NULL, 0, 'users', '', '{\"created_at\":\"2025-07-20T18:41:45.748+07:00\",\"created_by_id\":0,\"email\":\"muhfarid039@gmail.com\",\"id\":\"\",\"name\":\"Muhammad Farid Al Jabbar\",\"password\":\"$2a$10$zObgPd/b3qcNqUJXhnzseOvTpFY3yfczIJFcl1D3Z5ZQ3ZjWCj6vi\",\"updated_at\":\"2025-07-20T18:41:45.748+07:00\",\"updated_by_id\":0}', 'CREATE'),
(2, '2025-07-20 11:43:20.080', '2025-07-20 11:43:20.080', NULL, 0, 'users', '1', '{\"created_at\":\"2025-07-20T18:43:20.078+07:00\",\"created_by_id\":0,\"email\":\"muhfarid039@gmail.com\",\"id\":1,\"name\":\"Muhammad Farid Al Jabbar\",\"password\":\"$2a$10$uuoTN.V9qW1S7LsO4rpsju3E88Rpc/Ap.LsNfjDwibAhLAfLky6im\",\"updated_at\":\"2025-07-20T18:43:20.078+07:00\",\"updated_by_id\":0}', 'CREATE'),
(3, '2025-07-20 11:43:47.592', '2025-07-20 11:43:47.592', NULL, 0, 'users', '1', '{\"created_at\":\"2025-07-20T18:43:47.582+07:00\",\"created_by_id\":0,\"email\":\"muhfarid039@gmail.com\",\"id\":1,\"name\":\"Muhammad Farid Al Jabbar\",\"password\":\"$2a$10$XQ.ZeaFBtpwskmh5TPe49.Af0sbuKWX0IyUnyCBrCR3oAHaFn7GvK\",\"updated_at\":\"2025-07-20T18:43:47.582+07:00\",\"updated_by_id\":0}', 'CREATE'),
(4, '2025-07-20 13:17:39.573', '2025-07-20 13:17:39.573', NULL, 0, 'users', '1', '{\"created_at\":\"2025-07-20T20:17:39.561+07:00\",\"created_by_id\":0,\"email\":\"muhfarid039@gmail.com\",\"id\":1,\"name\":\"Muhammad Farid Al Jabbar\",\"password\":\"$2a$10$7WIOAj3YN5h6Xl6uRQdusuBIw3/.ftHx0TjfL0hWn3TQquJK33QZ2\",\"role\":\"Administrator\",\"updated_at\":\"2025-07-20T20:17:39.561+07:00\",\"updated_by_id\":0}', 'CREATE'),
(5, '2025-07-21 00:14:15.531', '2025-07-21 00:14:15.531', NULL, 0, 'users', '12', '{\"created_at\":\"2025-07-21T07:14:15.529+07:00\",\"created_by_id\":0,\"email\":\"client@gmail.com\",\"id\":12,\"name\":\"Client\",\"password\":\"$2a$10$MDSkmS1gcNse.VNeyw0twO06vOOfPG7ea5VMwAYwc4sUkLnm2hrNa\",\"role\":\"Client\",\"updated_at\":\"2025-07-21T07:14:15.529+07:00\",\"updated_by_id\":0}', 'CREATE'),
(6, '2025-07-21 00:48:44.176', '2025-07-21 00:48:44.176', NULL, 2207003, 'traditional_foods', '1', '{\"created_at\":\"2025-07-21T07:48:44.173+07:00\",\"created_by_id\":2207003,\"description\":\"Rendang sebagai salah satu masakan khas Indonesia tidak hanya dikenal karena rasa yang lezat, tetapi juga mengandung sejumlah manfaat kesehatan, terutama jika dipadukan dengan bahan-bahan alami yang bergizi. \",\"id\":1,\"main_ingredient\":\"1 kg daging sapi, 1 liter santan kental dari 3 butir kelapa (perasan pertama tanpa air), 550 gram kelapa parut, disangrai sampai kecokelatan, 5 lembar daun salam, 1 lembar daun kunyit, 10 lembar daun jeruk, 5 batang serai, 1/2 batang kayu manis, 3 butir cengkeh, 2 sdt garam, 1 buah kembang lawang.\",\"name\":\"Rendang\",\"regional_origin\":\"Padang\",\"type\":\"Lauk\",\"updated_at\":\"2025-07-21T07:48:44.173+07:00\",\"updated_by_id\":2207003}', 'CREATE'),
(7, '2025-07-21 00:52:12.586', '2025-07-21 00:52:12.586', NULL, 2207003, 'traditional_foods', '1', '{\"created_at\":\"2025-07-21T00:48:44.173Z\",\"created_by_id\":2207003,\"description\":\"Rendang sebagai salah satu masakan khas Indonesia tidak hanya dikenal karena rasa yang lezat, tetapi juga mengandung sejumlah manfaat kesehatan, terutama jika dipadukan dengan bahan-bahan alami yang bergizi. \",\"id\":1,\"main_ingredient\":\"1 kg daging sapi, 1 liter santan kental dari 3 butir kelapa (perasan pertama tanpa air), 550 gram kelapa parut, disangrai sampai kecokelatan, 5 lembar daun salam, 1 lembar daun kunyit, 10 lembar daun jeruk, 5 batang serai, 1/2 batang kayu manis, 3 butir cengkeh, 2 sdt garam, 1 buah kembang lawang.\",\"name\":\"Rendang\",\"regional_origin\":\"Padang\",\"type\":\"Lauk\",\"updated_at\":\"2025-07-21T00:52:12.575Z\",\"updated_by_id\":2207003}', 'UPDATE'),
(8, '2025-07-21 01:18:12.063', '2025-07-21 01:18:12.063', NULL, 2207003, 'traditional_foods', '1', '{\"created_at\":\"2025-07-21T00:48:44.173Z\",\"created_by_id\":2207003,\"description\":\"Rendang sebagai salah satu masakan khas Indonesia tidak hanya dikenal karena rasa yang lezat, tetapi juga mengandung sejumlah manfaat kesehatan, terutama jika dipadukan dengan bahan-bahan alami yang bergizi. \",\"id\":1,\"main_ingredient\":\"1 kg daging sapi, 1 liter santan kental dari 3 butir kelapa (perasan pertama tanpa air), 550 gram kelapa parut, disangrai sampai kecokelatan, 5 lembar daun salam, 1 lembar daun kunyit, 10 lembar daun jeruk, 5 batang serai, 1/2 batang kayu manis, 3 butir cengkeh, 2 sdt garam, 1 buah kembang lawang.\",\"name\":\"Rendang\",\"regional_origin\":\"Padang\",\"type\":\"Lauk\",\"updated_at\":\"2025-07-21T08:18:12.049+07:00\",\"updated_by_id\":2207003}', 'DELETE'),
(9, '2025-07-21 14:24:13.944', '2025-07-21 14:24:13.944', NULL, 2207003, 'traditional_foods', '2', '{\"created_at\":\"2025-07-21T21:24:13.924+07:00\",\"created_by_id\":2207003,\"description\":\"Rendang sebagai salah satu masakan khas Indonesia tidak hanya dikenal karena rasa yang lezat, tetapi juga mengandung sejumlah manfaat kesehatan, terutama jika dipadukan dengan bahan-bahan alami yang bergizi. \",\"id\":2,\"main_ingredient\":\"1 kg daging sapi, 1 liter santan kental dari 3 butir kelapa (perasan pertama tanpa air), 550 gram kelapa parut, disangrai sampai kecokelatan, 5 lembar daun salam, 1 lembar daun kunyit, 10 lembar daun jeruk, 5 batang serai, 1/2 batang kayu manis, 3 butir cengkeh, 2 sdt garam, 1 buah kembang lawang.\",\"name\":\"Rendang\",\"regional_origin\":\"Padang\",\"type\":\"Lauk\",\"updated_at\":\"2025-07-21T21:24:13.924+07:00\",\"updated_by_id\":2207003}', 'CREATE');

INSERT INTO `traditional_foods` (`id`, `created_at`, `updated_at`, `deleted_at`, `created_by_id`, `updated_by_id`, `deleted_by_id`, `name`, `regional_origin`, `main_ingredient`, `type`, `description`) VALUES
(2, '2025-07-21 14:24:13.924', '2025-07-21 14:24:13.924', NULL, 2207003, 2207003, NULL, 'Rendang', 'Padang', '1 kg daging sapi, 1 liter santan kental dari 3 butir kelapa (perasan pertama tanpa air), 550 gram kelapa parut, disangrai sampai kecokelatan, 5 lembar daun salam, 1 lembar daun kunyit, 10 lembar daun jeruk, 5 batang serai, 1/2 batang kayu manis, 3 butir cengkeh, 2 sdt garam, 1 buah kembang lawang.', 'Lauk', 'Rendang sebagai salah satu masakan khas Indonesia tidak hanya dikenal karena rasa yang lezat, tetapi juga mengandung sejumlah manfaat kesehatan, terutama jika dipadukan dengan bahan-bahan alami yang bergizi. ');

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `created_by_id`, `updated_by_id`, `deleted_by_id`, `name`, `email`, `role`, `password`) VALUES
(1, '2025-07-20 13:17:39.561', '2025-07-20 13:17:39.561', NULL, 0, 0, NULL, 'Muhammad Farid Al Jabbar', 'muhfarid039@gmail.com', 'Administrator', '$2a$10$7WIOAj3YN5h6Xl6uRQdusuBIw3/.ftHx0TjfL0hWn3TQquJK33QZ2'),
(12, '2025-07-21 00:14:15.529', '2025-07-21 00:14:15.529', NULL, 0, 0, NULL, 'Client', 'client@gmail.com', 'Client', '$2a$10$MDSkmS1gcNse.VNeyw0twO06vOOfPG7ea5VMwAYwc4sUkLnm2hrNa');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
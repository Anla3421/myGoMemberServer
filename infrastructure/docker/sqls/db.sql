-- Adminer 4.8.1 MySQL 8.0.25 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE `testdb` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `testdb`;

DROP TABLE IF EXISTS `drink`;
CREATE TABLE `drink` (
  `ID` int NOT NULL,
  `who` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `drink` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `sugar` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `ice` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;


DROP TABLE IF EXISTS `dummyUser`;
CREATE TABLE `dummyUser` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(45) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `password` varchar(45) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `jwt` varchar(256) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `dummyUser` (`id`, `account`, `password`, `jwt`, `update_at`) VALUES
(1,	'admin',	'21232f297a57a5a743894a0e4a801fc3',	'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50IjoiYWRtaW4iLCJleHAiOjE2MzkxNDU1NTQsImlzcyI6Ik1lIn0.VCp98bVMlPNYfrY9KiDIm7QRu-SfQ9WM7KmFZetNWT8',	'2021-12-24 07:19:03');

DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `amount` int NOT NULL,
  `price` int NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID_UNIQUE` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `goods` (`ID`, `name`, `amount`, `price`) VALUES
(1,	'蘋果',	100,	10),
(2,	'香蕉',	50,	50);

DROP TABLE IF EXISTS `movielist`;
CREATE TABLE `movielist` (
  `ID` int NOT NULL,
  `idre` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `moviename` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `movielist` (`ID`, `idre`, `moviename`) VALUES
(1,	'1292052',	'肖申克的救赎');

DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `ID` int unsigned NOT NULL AUTO_INCREMENT,
  `itemID` int NOT NULL,
  `amount` int NOT NULL,
  `price` int NOT NULL,
  `buyerID` int NOT NULL,
  `finish` enum('yes','no') COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT 'no',
  `createTime` timestamp NOT NULL,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `order` (`ID`, `itemID`, `amount`, `price`, `buyerID`, `finish`, `createTime`, `update_at`) VALUES
(7,	1,	10,	10000,	1,	'no',	'2021-12-24 02:29:13',	'2021-12-24 02:29:13'),
(8,	1,	10,	10000,	1,	'no',	'2021-12-24 15:23:07',	'2021-12-24 07:23:06'),
(9,	1,	10,	10000,	1,	'yes',	'2021-12-24 16:28:33',	'2021-12-24 08:28:32');

DROP TABLE IF EXISTS `page3`;
CREATE TABLE `page3` (
  `ID` int NOT NULL,
  `name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `text` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `input1` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `input2` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `page3` (`ID`, `name`, `text`, `input1`, `input2`) VALUES
(1,	'我是第一頁',	'測試用內文1',	NULL,	NULL),
(2,	'我是第二頁',	'測試用內文2',	NULL,	NULL),
(3,	'我是第三頁',	'測試用內文3',	NULL,	NULL);

DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `code` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci,
  `price` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`),
  KEY `idx_product_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `product` (`id`, `created_at`, `updated_at`, `deleted_at`, `code`, `price`) VALUES
(1,	'2021-03-29 12:28:35.568',	'2021-03-29 12:28:35.579',	NULL,	'F42',	200),
(2,	'2021-03-29 12:29:58.794',	'2021-03-29 12:29:58.801',	NULL,	'F42',	200),
(3,	'2021-03-29 12:30:58.603',	'2021-03-29 12:30:58.622',	NULL,	'F42',	200),
(4,	'2021-03-29 12:31:36.760',	'2021-03-29 12:31:36.768',	'2021-03-29 12:31:36.770',	'F42',	200),
(5,	'2021-03-29 12:31:48.726',	'2021-03-29 12:31:48.736',	'2021-03-29 12:34:08.422',	'F42',	200);

DROP TABLE IF EXISTS `weather`;
CREATE TABLE `weather` (
  `ID` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `Name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `Text` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `weather` (`ID`, `Name`, `Text`) VALUES
('1',	'台北',	'台北很冷,22度'),
('2',	'台中',	'台中普通,25度'),
('3',	'台南',	'台南很熱,29度'),
('4',	'高雄',	'高雄很熱,31度');

CREATE DATABASE `zmemberdb` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `zmemberdb`;

DROP TABLE IF EXISTS `logintime`;
CREATE TABLE `logintime` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `login_time` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `logintime` (`id`, `account`, `login_time`) VALUES
(1,	'admin',	1620713800),
(2,	'jared',	1620713846);

DROP TABLE IF EXISTS `member`;
CREATE TABLE `member` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `password` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `jwt` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `member` (`id`, `account`, `password`, `jwt`, `updated_at`) VALUES
(1,	'admin',	'21232f297a57a5a743894a0e4a801fc3',	'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50IjoiYWRtaW4iLCJleHAiOjE2MzkxNDU1NTQsImlzcyI6Ik1lIn0.VCp98bVMlPNYfrY9KiDIm7QRu-SfQ9WM7KmFZetNWT8',	'2021-05-11 07:07:26'),
(2,	'jared',	'b620e68b3bf4387bf7c43d51bd12910b',	NULL,	'2021-05-11 07:07:26'),
(3,	'derek',	'7815696ecbf1c96e6894b779456d330e',	NULL,	'2021-12-11 19:16:45');

CREATE DATABASE `zorderdb` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `zorderdb`;

DROP TABLE IF EXISTS `dummyUser`;
CREATE TABLE `dummyUser` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(45) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `password` varchar(45) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `jwt` varchar(256) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `dummyUser` (`id`, `account`, `password`, `jwt`, `update_at`) VALUES
(1,	'admin',	'21232f297a57a5a743894a0e4a801fc3',	'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50IjoiYWRtaW4iLCJleHAiOjE2MzkxNDU1NTQsImlzcyI6Ik1lIn0.VCp98bVMlPNYfrY9KiDIm7QRu-SfQ9WM7KmFZetNWT8',	'2021-12-24 07:19:03');

DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `amount` int NOT NULL,
  `price` int NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_520_ci NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID_UNIQUE` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `goods` (`ID`, `name`, `amount`, `price`, `description`) VALUES
(1,	'蘋果',	100,	10,	'只是個蘋果'),
(2,	'香蕉',	50,	50,	'只是個香蕉');

DROP TABLE IF EXISTS `orderDetail`;
CREATE TABLE `orderDetail` (
  `ID` int unsigned NOT NULL AUTO_INCREMENT,
  `orderID` int unsigned NOT NULL,
  `itemID` int NOT NULL,
  `amount` int NOT NULL,
  `price` int NOT NULL,
  `buyerID` int NOT NULL,
  `finish` enum('yes','no') COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT 'no',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `orderDetail` (`ID`, `orderID`, `itemID`, `amount`, `price`, `buyerID`, `finish`, `createTime`, `update_at`) VALUES
(1,	1,	1,	10,	10000,	1,	'no',	'2021-12-24 02:29:13',	'2021-12-24 02:29:13');

DROP TABLE IF EXISTS `orderGeneral`;
CREATE TABLE `orderGeneral` (
  `orderID` int NOT NULL AUTO_INCREMENT,
  `totalPrice` int NOT NULL,
  `buyerID` int NOT NULL,
  `finish` enum('yes','no') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'no',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`orderID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `orderGeneral` (`orderID`, `totalPrice`, `buyerID`, `finish`, `created_at`, `updated_at`) VALUES
(1,	99999,	1,	'no',	'2021-12-29 08:27:34',	'2021-12-29 08:27:34'),
(2,	99999,	2,	'no',	'2021-12-29 08:27:34',	'2021-12-29 08:27:34');

-- 2021-12-30 20:03:48

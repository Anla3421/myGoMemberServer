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


DROP TABLE IF EXISTS `movielist`;
CREATE TABLE `movielist` (
  `ID` int NOT NULL,
  `idre` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `moviename` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `movielist` (`ID`, `idre`, `moviename`) VALUES
(1,	'1292052',	'肖申克的救赎');

DROP TABLE IF EXISTS `new_table`;
CREATE TABLE `new_table` (
  `ID` int NOT NULL,
  `Name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `text` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

INSERT INTO `new_table` (`ID`, `Name`, `text`) VALUES
(53,	NULL,	NULL);

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
(2,	'jared',	1620713846),
(3,	'jared',	1620713900);

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
(1,	'admin',	'21232f297a57a5a743894a0e4a801fc3',	'indontneedthis',	'2021-05-11 07:07:26'),
(2,	'jared',	'b620e68b3bf4387bf7c43d51bd12910b',	NULL,	'2021-05-11 07:07:26');

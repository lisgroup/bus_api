CREATE DATABASE IF NOT EXISTS bus_api CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE bus_api;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

# 用户信息表，用于登录
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
     `id` int unsigned NOT NULL AUTO_INCREMENT,
     `identity` varchar(36) NOT NULL DEFAULT '',
     `name` varchar(60) NOT NULL DEFAULT '',
     `password` varchar(128) NOT NULL DEFAULT '',
     `email` varchar(100) NOT NULL DEFAULT '',
     `now_volume` int NOT NULL DEFAULT '0',
     `total_volume` int NOT NULL DEFAULT '0',
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL,
     PRIMARY KEY (`id`),
     UNIQUE KEY `uni_name` (`name`),
     UNIQUE KEY `uni_email` (`email`),
     UNIQUE KEY `uni_identity` (`identity`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb3 COMMENT='用户信息表';

SET FOREIGN_KEY_CHECKS = 1;
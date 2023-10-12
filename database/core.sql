CREATE DATABASE IF NOT EXISTS bus_api CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE bus_api;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

# 用户信息表，用于登录
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT,
    `identity`     varchar(36)      not null default '',

    `name`         varchar(60)      not null default '',
    `password`     varchar(128)      not null default '',
    `email`        varchar(100)     not null default '',
    `now_volume`   int(10)          not null default 0,
    `total_volume` int(10)          not null default 0,
    `created_at`   datetime                  DEFAULT NULL,
    `updated_at`   datetime                  DEFAULT NULL,
    `deleted_at`   datetime                  DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 comment '用户信息表';

SET FOREIGN_KEY_CHECKS = 1;
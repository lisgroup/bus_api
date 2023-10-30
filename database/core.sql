CREATE DATABASE IF NOT EXISTS bus_api CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE bus_api;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

# 用户信息表，用于登录
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`           int unsigned NOT NULL AUTO_INCREMENT,
    `identity`     varchar(36)  NOT NULL DEFAULT '',
    `username`     varchar(60)  NOT NULL DEFAULT '',
    `password`     varchar(128) NOT NULL DEFAULT '',
    `email`        varchar(100) NOT NULL DEFAULT '',
    `role`         varchar(100) NOT NULL DEFAULT '',
    `now_volume`   int          NOT NULL DEFAULT '0',
    `total_volume` int          NOT NULL DEFAULT '0',
    `created_at`   datetime              DEFAULT NULL,
    `updated_at`   datetime              DEFAULT NULL,
    `deleted_at`   datetime              DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_username` (`username`),
    UNIQUE KEY `uni_email` (`email`),
    UNIQUE KEY `uni_identity` (`identity`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户信息表';

DROP TABLE IF EXISTS `server_key`;
CREATE TABLE `server_key`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int          NOT NULL DEFAULT 0,
    `j_key`      varchar(60)  NOT NULL DEFAULT '',
    `created_at` datetime              DEFAULT NULL,
    `updated_at` datetime              DEFAULT NULL,
    `deleted_at` datetime              DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    UNIQUE KEY `uni_j_key` (`j_key`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb3 COMMENT ='serverJ秘钥';

DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`      int(11)          NOT NULL DEFAULT '0',
    `j_key`        varchar(60)      NOT NULL DEFAULT '',
    `cycle`        varchar(20)      NOT NULL DEFAULT '',
    `hour`         tinyint(4)       NOT NULL DEFAULT '0',
    `minute`       tinyint(4)       NOT NULL DEFAULT '0',
    `line_id`      varchar(60)      NOT NULL DEFAULT '',
    `line_name`    varchar(60)      NOT NULL DEFAULT '',
    `line_from_to` varchar(100)     NOT NULL DEFAULT '',
    `station_num`  varchar(10)      NOT NULL DEFAULT '',
    `station_id`   varchar(60)      NOT NULL DEFAULT '',
    `station_name` varchar(60)      NOT NULL DEFAULT '',
    `start_time`   time             NOT NULL COMMENT '开始时间',
    `end_time`     time             NOT NULL COMMENT '结束时间',
    `notice_time`  tinyint(4)       NOT NULL DEFAULT '1' COMMENT '通知次数',
    `created_at`   datetime                  DEFAULT NULL,
    `updated_at`   datetime                  DEFAULT NULL,
    `deleted_at`   datetime                  DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `j_key` (`j_key`) USING BTREE,
    KEY `idx_start_end_time` (`start_time`, `end_time`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4 COMMENT ='定时任务';

SET FOREIGN_KEY_CHECKS = 1;
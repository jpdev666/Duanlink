CREATE DATABASE IF NOT EXISTS duanlink;
USE duanlink;
-- 创建短链信息表 short_links
CREATE TABLE `short_links`
(
    `id` int(11)   NOT NULL AUTO_INCREMENT,
    `short_code`     varchar(32)   NOT NULL,
    `origin_url`  varchar(2048) NOT NULL,
    `created_at`     datetime      NOT NULL,
    `updated_at`     datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `short_links` ADD COLUMN `expires_at` datetime;
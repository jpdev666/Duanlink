CREATE
DATABASE IF NOT EXISTS duanlink;
USE
duanlink;
-- 创建短链信息表 short_links
CREATE TABLE `short_links`
(
    `short_link_id` varchar(32)   NOT NULL,
    `original_url`  varchar(2048) NOT NULL,
    `create_at`     datetime      NOT NULL,
    `expire_at`     datetime DEFAULT NULL,
    PRIMARY KEY (`short_link_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
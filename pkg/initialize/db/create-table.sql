-- +migrate Up

CREATE TABLE IF NOT EXISTS `md_user` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
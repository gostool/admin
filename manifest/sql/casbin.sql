DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
                               `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                               `ptype` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
                               `v0` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
                               `v1` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
                               `v2` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
                               `v3` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
                               `v4` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
                               `v5` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
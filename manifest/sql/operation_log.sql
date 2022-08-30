
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log` (
                                 `id` int(11) NOT NULL AUTO_INCREMENT,
                                 `user_id` int(11) DEFAULT NULL,
                                 `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                                 `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
                                 `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                                 `is_deleted` int NOT NULL COMMENT ' 数据的逻辑删除',
                                 `status` int(11) DEFAULT NULL COMMENT '状态',
                                 `extra` text,
                                 `ip` varchar(255) DEFAULT NULL COMMENT '请求ip',
                                 `path` varchar(255) DEFAULT NULL COMMENT '请求路径',
                                 `method` varchar(255) DEFAULT NULL COMMENT '请求方法',
                                 `request` text COMMENT '请求参数',
                                 `response` text COMMENT '响应内容',
                                 `agent` varchar(255) DEFAULT NULL COMMENT '代理',
                                 `latency` int(11) DEFAULT NULL COMMENT '延迟',
                                 `err_msg` varchar(255) DEFAULT NULL COMMENT '错误信息',
                                 `user_name` varchar(255) DEFAULT NULL COMMENT '用户名',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=644 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
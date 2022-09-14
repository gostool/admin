CREATE DATABASE IF NOT EXISTS godb
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_0900_ai_ci;
show databases;
use godb;
/*
 Navicat Premium Data Transfer

 Source Server         : local_ubuntu_server
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : 192.168.3.8:3306
 Source Schema         : godb

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 05/09/2022 18:02:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api` (
                       `id` int NOT NULL AUTO_INCREMENT,
                       `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                       `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
                       `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                       `is_deleted` int NOT NULL DEFAULT '0' COMMENT ' 数据的逻辑删除',
                       `group` varchar(255) NOT NULL COMMENT '分组',
                       `path` varchar(255) NOT NULL COMMENT '请求路径',
                       `detail` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT ' 简介',
                       `method` varchar(255) NOT NULL COMMENT '请求方式',
                       PRIMARY KEY (`id`),
                       KEY `group` (`group`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
                               `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                               `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
                               `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
                               `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
                               `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
                               `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
                               `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
                               `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=326 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                        `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
                        `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                        `is_deleted` int NOT NULL DEFAULT '0' COMMENT ' 数据的逻辑删除',
                        `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '名称',
                        `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '图标',
                        `keep_alive` int DEFAULT NULL COMMENT '是否缓存',
                        `default_menu` int DEFAULT NULL COMMENT '是否基础路由',
                        `hidden` int DEFAULT NULL,
                        `sort` int DEFAULT NULL,
                        `parent_id` int DEFAULT NULL,
                        `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                        `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                        `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for operation_log
-- ----------------------------
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `user_id` int DEFAULT NULL,
                                 `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                                 `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
                                 `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                                 `is_deleted` int NOT NULL DEFAULT '0' COMMENT ' 数据的逻辑删除',
                                 `status` int DEFAULT NULL COMMENT '状态',
                                 `extra` text,
                                 `ip` varchar(255) DEFAULT NULL COMMENT '请求ip',
                                 `path` varchar(255) DEFAULT NULL COMMENT '请求路径',
                                 `method` varchar(255) DEFAULT NULL COMMENT '请求方法',
                                 `request` text COMMENT '请求参数',
                                 `response` text COMMENT '响应内容',
                                 `agent` varchar(255) DEFAULT NULL COMMENT '代理',
                                 `latency` int DEFAULT NULL COMMENT '延迟',
                                 `err_msg` varchar(255) DEFAULT NULL COMMENT '错误信息',
                                 `user_name` varchar(255) DEFAULT NULL COMMENT '用户名',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=567 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                        `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
                        `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                        `is_deleted` int NOT NULL DEFAULT '0' COMMENT ' 数据的逻辑删除',
                        `name` varchar(255) NOT NULL COMMENT '角色名称',
                        `pid` int DEFAULT NULL COMMENT '父角色id',
                        `router` varchar(255) NOT NULL COMMENT '默认路由',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                             `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
                             `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                             `is_deleted` int NOT NULL DEFAULT '0' COMMENT ' 数据的逻辑删除',
                             `role_id` int DEFAULT NULL COMMENT '角色id',
                             `menu_id` int DEFAULT NULL COMMENT '菜单id',
                             PRIMARY KEY (`id`),
                             KEY `role_id` (`role_id`),
                             KEY `menu_id` (`menu_id`),
                             CONSTRAINT `role_menu_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                             CONSTRAINT `role_menu_ibfk_2` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for tpl
-- ----------------------------
DROP TABLE IF EXISTS `tpl`;
CREATE TABLE `tpl` (
                       `id` int NOT NULL AUTO_INCREMENT,
                       `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                       `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
                       `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                       `is_deleted` int NOT NULL DEFAULT '0' COMMENT ' 数据的逻辑删除',
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                        `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
                        `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                        `is_deleted` int NOT NULL DEFAULT '0' COMMENT ' 数据的逻辑删除',
                        `passport` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                        `password` varchar(255) NOT NULL,
                        `nickname` varchar(255) DEFAULT NULL,
                        `role_id` int NOT NULL DEFAULT '0',
                        `roleIds` varchar(1024) DEFAULT NULL COMMENT 'eg:[1, 2] 角色数组',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `name` (`passport`) USING BTREE,
                        KEY `role_id` (`role_id`),
                        CONSTRAINT `user_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;

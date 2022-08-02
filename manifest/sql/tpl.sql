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

 Date: 02/08/2022 20:32:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tpl
-- ----------------------------
DROP TABLE IF EXISTS `tpl`;
CREATE TABLE `tpl` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `is_deleted` int NOT NULL COMMENT ' 数据的逻辑删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;

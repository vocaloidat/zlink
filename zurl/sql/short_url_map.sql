/*
 Navicat Premium Dump SQL

 Source Server         : locahost_mysql
 Source Server Type    : MySQL
 Source Server Version : 90200 (9.2.0)
 Source Host           : localhost:3306
 Source Schema         : zlink

 Target Server Type    : MySQL
 Target Server Version : 90200 (9.2.0)
 File Encoding         : 65001

 Date: 28/03/2025 17:08:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for short_url_map
-- ----------------------------
DROP TABLE IF EXISTS `short_url_map`;
CREATE TABLE `short_url_map`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
  `is_del` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除：0正常1删除',
  `lurl` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '长链接',
  `md5` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '长链接MD5',
  `surl` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '短链接',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `md5`(`md5` ASC) USING BTREE,
  UNIQUE INDEX `surl`(`surl` ASC) USING BTREE,
  INDEX `is_del`(`is_del` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '长短链映射表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of short_url_map
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;

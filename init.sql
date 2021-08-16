/*
 Navicat Premium Data Transfer
 Source Server         : local--211
 Source Server Type    : MariaDB
 Source Server Version : 100327
 Target Server Type    : MariaDB
 Target Server Version : 100327
 File Encoding         : 65001

 Date: 09/06/2021 17:48:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
CREATE TABLE IF NOT EXISTS `account`  (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'acct id',
    `acct` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '帐号',
    `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '名字',
    `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT ' 密码',
    `access` int(11) UNSIGNED NULL DEFAULT 0 COMMENT '权限 bit',
    `valid` tinyint(8) UNSIGNED NULL DEFAULT 0 COMMENT '是否有效 0.无效 1有效',
    `last` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '最后登录ip',
    `ctm` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
    `utm` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '最后登录时间',
    `skin` tinyint(8) UNSIGNED NULL DEFAULT 0 COMMENT '皮肤编号',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
    AUTO_INCREMENT = 1
    CHARACTER SET = utf8mb4
    COLLATE = utf8mb4_general_ci
    COMMENT = '管理人员帐户表'
    ROW_FORMAT = Dynamic;

CREATE TABLE IF NOT EXISTS `action_log`  (
   `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
   `aid` int(11) NULL DEFAULT NULL COMMENT 'acct id',
   `opm` int(11) UNSIGNED NULL DEFAULT 0 COMMENT '操作主节点',
   `ops` int(11) UNSIGNED NULL DEFAULT 0 COMMENT '操作子节点',
   `val` bigint(16) UNSIGNED NULL DEFAULT NULL COMMENT '当前值',
   `chg` bigint(16) NULL DEFAULT NULL COMMENT '改变值',
   `par1` bigint(16) NULL DEFAULT NULL COMMENT '参数1',
   `par2` bigint(16) NULL DEFAULT NULL COMMENT '参数2',
   `par3` bigint(16) NULL DEFAULT NULL COMMENT '参数3',
   `par4` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '参数4',
   `ctm` bigint(16) NULL DEFAULT NULL COMMENT '创建时间',
   PRIMARY KEY (`id`) USING BTREE,
   INDEX `opm`(`opm`) USING BTREE,
   INDEX `ops`(`ops`) USING BTREE,
   INDEX `ctm`(`ctm`) USING BTREE,
   INDEX `aid`(`aid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;



-- ----------------------------
-- Records of account  ds@fjwe$%sadq15894K 对应 password_salt=ds@fjwe$%sadq15894K
-- ----------------------------
REPLACE INTO account (acct, name, password, access, valid, ctm)
VALUES ('admin', '管理员', MD5(CONCAT('admin', '123456', 'ds@fjwe$%sadq15894K')), 31, 1, UNIX_TIMESTAMP(NOW()));

SET FOREIGN_KEY_CHECKS = 1;

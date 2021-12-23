/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库159781
 Source Server Type    : MySQL
 Source Server Version : 80026
 Source Host           : localhost:3306
 Source Schema         : pinecms

 Target Server Type    : MySQL
 Target Server Version : 80026
 File Encoding         : 65001

 Date: 23/12/2021 16:53:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pinecms_admin
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_admin`;
CREATE TABLE `pinecms_admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `roles` json DEFAULT NULL,
  `encrypt` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `lastloginip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `lastlogintime` int DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `realname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_admin
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_admin` VALUES (1, 'admin', '14e1b600b1fd579f47433b88e8d85291', '[1]', '', '127.0.0.1', 0, '826466266@qq.com', 'xiusin', 'http://pinecms.oss-cn-beijing.aliyuncs.com/uploads/Ymd/F99iE8xG2D8tc9Nz.jpeg', '', '', 1);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_admin_role`;
CREATE TABLE `pinecms_admin_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `rolename` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `listorder` int NOT NULL,
  `disabled` tinyint unsigned NOT NULL DEFAULT '0',
  `menu_ids` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_admin_role
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_admin_role` VALUES (1, '超级管理员', '拥有系统最高权限', 3, 0, '[3,34,35,53,54,99,100,36,37,38,39,40,41,62,64,86,109,112,9,10,16,17,18,19,20,21,89,91,92,93,94,95,96,98,101,114,11,12,24,25,26,27,13,28,29,30,31,103,104,14,15,22,23,85,105,113,116,55,56,57,60,111,58,59,65,118,119,120,121,122,123,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,84,87,108,107,110,117,124]');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_admin_role_priv
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_admin_role_priv`;
CREATE TABLE `pinecms_admin_role_priv` (
  `role_id` int NOT NULL,
  `menu_id` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_admin_role_priv
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 3);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 34);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 35);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 53);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 54);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 99);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 100);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 36);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 37);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 38);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 39);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 40);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 41);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 62);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 64);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 86);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 109);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 112);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 9);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 10);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 16);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 17);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 18);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 19);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 20);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 21);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 89);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 91);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 92);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 93);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 94);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 95);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 96);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 98);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 101);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 114);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 11);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 12);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 24);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 25);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 26);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 27);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 13);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 28);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 29);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 30);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 31);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 103);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 104);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 14);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 15);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 22);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 23);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 85);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 105);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 113);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 116);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 55);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 56);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 57);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 60);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 111);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 58);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 59);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 65);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 118);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 119);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 120);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 121);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 122);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 123);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 67);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 68);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 69);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 70);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 71);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 72);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 73);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 74);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 75);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 76);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 77);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 78);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 79);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 80);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 81);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 82);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 84);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 87);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 108);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 107);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 110);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 117);
INSERT INTO `pinecms_admin_role_priv` VALUES (1, 124);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_advert
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_advert`;
CREATE TABLE `pinecms_advert` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `space_id` int DEFAULT NULL,
  `image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `link_url` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `listorder` mediumint DEFAULT '0',
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0' COMMENT '0=禁用,1=正常',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='广告表';

-- ----------------------------
-- Records of pinecms_advert
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_advert` VALUES (17, 'left2', 14, '/uploads/public/20200413/3VVwkb03OU.jpg', '', 0, '2020-04-13 18:31:19', '2021-04-13 18:31:19', 1);
INSERT INTO `pinecms_advert` VALUES (18, 'footer', 14, '/uploads/public/20200413/q92SSCjq0R.jpg', '', 0, '2020-04-13 18:31:37', '2021-04-13 18:31:37', 1);
INSERT INTO `pinecms_advert` VALUES (19, 'article1', 14, '/uploads/public/20200413/44twXycv3u.jpg', '', 0, '2020-04-13 18:31:56', '2021-04-13 18:31:56', 1);
INSERT INTO `pinecms_advert` VALUES (20, 'article2', 14, '/uploads/public/20200413/48PktdpOMG.jpg', '', 0, '2020-04-13 18:32:16', '2021-04-13 18:32:16', 1);
INSERT INTO `pinecms_advert` VALUES (21, 'right4', 14, '/uploads/public/20200413/1789ez9j0L.jpg', '', 0, '2020-04-13 18:32:43', '2021-04-13 18:32:43', 1);
INSERT INTO `pinecms_advert` VALUES (22, 'right2', 14, '/uploads/public/20200413/66OBu3Za08.jpg', '', 0, '2020-04-13 18:33:01', '2021-04-13 18:33:01', 1);
INSERT INTO `pinecms_advert` VALUES (23, 'right3', 14, '/uploads/public/20200413/3lnooWbZ4d.jpg', '', 0, '2020-04-13 18:33:44', '2021-04-13 18:33:44', 1);
INSERT INTO `pinecms_advert` VALUES (24, 'right1', 14, '/uploads/public/20200413/sopAEUX29h.jpg', '', 0, '2020-04-13 18:33:56', '2021-04-13 18:33:56', 1);
INSERT INTO `pinecms_advert` VALUES (25, 'top_all', 14, 'http://pinecms.oss-cn-beijing.aliyuncs.com/uploads/Ymd/976V98p6Y8oRW34y.jpeg', '', 0, '2020-04-13 18:34:25', '2021-04-13 18:34:25', 1);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_advert_space
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_advert_space`;
CREATE TABLE `pinecms_advert_space` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '广告位名称',
  `key` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '标识',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='广告位表';

-- ----------------------------
-- Records of pinecms_advert_space
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_advert_space` VALUES (14, '全站', 'full_site', '');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_articles
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_articles`;
CREATE TABLE `pinecms_articles` (
  `id` int NOT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `thumb` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `keywords` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `flag` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `tags` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `from_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `author` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `pubtime` datetime DEFAULT NULL,
  `status` int DEFAULT NULL,
  `visit_count` int DEFAULT NULL,
  `listorder` int DEFAULT NULL,
  `flash` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `catid` int NOT NULL,
  `mid` int NOT NULL,
  `created_time` datetime DEFAULT NULL,
  `updated_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_articles
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pinecms_attachment_type
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_attachment_type`;
CREATE TABLE `pinecms_attachment_type` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_attachment_type
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_attachment_type` VALUES (1, '文章图片', '2021-08-09 15:23:42');
INSERT INTO `pinecms_attachment_type` VALUES (2, '其他附件', '2021-08-09 15:29:08');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_attachments
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_attachments`;
CREATE TABLE `pinecms_attachments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `origin_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `size` int DEFAULT NULL,
  `upload_time` datetime DEFAULT NULL,
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `classify_id` int DEFAULT NULL COMMENT '归属分类ID',
  `md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附件的md5值',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `UQE_pinecms_attachments_md5` (`md5`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=180 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_attachments
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_attachments` VALUES (154, '5K2giheHM5.jpg', 'http://iriscms-test.oss-cn-beijing.aliyuncs.com/upload/public/20200324/5K2giheHM5.jpg', 'longad.jpg', 90292, '2020-03-24 18:51:37', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (155, '2nYZoj2iLI.jpg', 'http://iriscms-test.oss-cn-beijing.aliyuncs.com/upload/public/20200324/2nYZoj2iLI.jpg', 'ad.jpg', 155818, '2020-03-24 18:52:45', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (156, '40B5hc9v5I.jpg', 'http://iriscms-test.oss-cn-beijing.aliyuncs.com/upload/public/20200324/40B5hc9v5I.jpg', 'ad02.jpg', 116775, '2020-03-24 18:53:22', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (157, '4kyS4LoTIS.png', 'http://iriscms-test.oss-cn-beijing.aliyuncs.com/upload/public/20200324/4kyS4LoTIS.png', 'joinwx.png', 21736, '2020-03-24 18:54:17', 'image/png', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (158, 'y2LHji2RGI.jpg', 'http://iriscms-test.oss-cn-beijing.aliyuncs.com/upload/public/20200325/y2LHji2RGI.jpg', '3.jpg', 43351, '2020-03-25 16:58:59', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (169, 'H156oCxCpT.png', '/uploads/public/20200413/H156oCxCpT.png', 'center.png', 7960, '2020-04-13 10:30:32', 'image/png', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (170, 'JxvQHZJIg9.jpg', '/uploads/public/20200413/JxvQHZJIg9.jpg', 'left1.jpg', 43310, '2020-04-13 10:31:00', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (171, '3VVwkb03OU.jpg', '/uploads/public/20200413/3VVwkb03OU.jpg', 'left2.jpg', 43459, '2020-04-13 10:31:16', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (172, 'q92SSCjq0R.jpg', '/uploads/public/20200413/q92SSCjq0R.jpg', 'footer.jpg', 16774, '2020-04-13 10:31:33', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (173, '44twXycv3u.jpg', '/uploads/public/20200413/44twXycv3u.jpg', 'article_ad1.jpg', 34354, '2020-04-13 10:31:54', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (174, '48PktdpOMG.jpg', '/uploads/public/20200413/48PktdpOMG.jpg', 'article_ad2.jpg', 21965, '2020-04-13 10:32:13', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (175, '1789ez9j0L.jpg', '/uploads/public/20200413/1789ez9j0L.jpg', 'right4.jpg', 27499, '2020-04-13 10:32:40', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (176, '66OBu3Za08.jpg', '/uploads/public/20200413/66OBu3Za08.jpg', 'right2.jpg', 49343, '2020-04-13 10:32:57', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (177, '3lnooWbZ4d.jpg', '/uploads/public/20200413/3lnooWbZ4d.jpg', 'right3.jpg', 36814, '2020-04-13 10:33:37', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (178, 'sopAEUX29h.jpg', '/uploads/public/20200413/sopAEUX29h.jpg', 'right1.jpg', 27716, '2020-04-13 10:33:55', 'image/jpg', NULL, 1, NULL);
INSERT INTO `pinecms_attachments` VALUES (179, 'qq0g4e1Im6.png', '/uploads/public/20200413/qq0g4e1Im6.png', 'top.png', 8238, '2020-04-13 10:34:21', 'image/png', NULL, 1, NULL);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_casbin_rule`;
CREATE TABLE `pinecms_casbin_rule` (
  `p_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  KEY `IDX_pinecms_casbin_rule_v5` (`v5`) USING BTREE,
  KEY `IDX_pinecms_casbin_rule_p_type` (`p_type`) USING BTREE,
  KEY `IDX_pinecms_casbin_rule_v0` (`v0`) USING BTREE,
  KEY `IDX_pinecms_casbin_rule_v1` (`v1`) USING BTREE,
  KEY `IDX_pinecms_casbin_rule_v2` (`v2`) USING BTREE,
  KEY `IDX_pinecms_casbin_rule_v3` (`v3`) USING BTREE,
  KEY `IDX_pinecms_casbin_rule_v4` (`v4`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'content', 'top', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'content', 'right', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'content', 'index', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'content', 'news-list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'content', 'page', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'content', 'add', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'content', 'edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'category', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'category', 'view', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'category', 'add', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'category', 'edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'category', 'delete', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'category', 'order', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'model', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'model', 'add', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'model', 'edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'tags', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'table_field', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'setting', 'left', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'setting', 'site', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'menulist', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'menuview', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'menuadd', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'menuedit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'menudelete', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'menuorder', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'ad', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'ad', 'add', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'ad', 'edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'ad', 'delete', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'ad-space', 'add', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'ad-space', 'delete', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'ad-space', 'edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'setting', 'cache', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'dict', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'district', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'admin', 'left', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'admin', 'memberlist', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'admin', 'member-view', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'admin', 'member-add', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'admin', 'member-edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'admin', 'member-delete', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'role', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'role', 'add', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'role', 'edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'role', 'delete', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'department', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'position', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'index', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'loglist', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'logview', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'log-delete', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'tail', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'stat', 'data', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'errorLogList', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'pprof', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'user', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'member', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'member', 'info', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'user', 'edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'member_group', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'wechat', '', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'wechat', 'account', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'wechat', 'user', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'wechat', 'menu', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'wechat', 'rule', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'wechat', 'msg', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'wechat', 'qrcode', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'wechat', 'material', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'wechat', 'template', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'link', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'link', 'add', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'link', 'edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'link', 'delete', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'link', 'order', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'database', 'manager', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'database', 'backup', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'database', 'optimize', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'database', 'repair', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'database', 'backup-list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'assets-manager', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'assets-manager', 'add', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'assets-manager', 'edit', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'attachments', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'assets-manager', 'theme', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'plugin', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'task', 'list', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'task', 'script', '', '', '');
INSERT INTO `pinecms_casbin_rule` VALUES ('p', '1', 'system', 'document', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_category
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_category`;
CREATE TABLE `pinecms_category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` int NOT NULL,
  `parentid` int NOT NULL,
  `topid` int NOT NULL,
  `model_id` int NOT NULL,
  `catname` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `keywords` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `listorder` int NOT NULL,
  `ismenu` int NOT NULL,
  `list_tpl` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `detail_tpl` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `thumb` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `dir` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_category
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_category` VALUES (1, 0, 0, 0, 1, 'Win8之家', ' ', '', '', 30, 1, '', '', '', 'win8', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (2, 0, 0, 0, 1, '资讯', ' ', '', '', 30, 1, '', '', '', 'news', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (5, 0, 0, 0, 1, '苹果之家', ' ', '', '', 30, 1, '', '', '', 'apple', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (6, 0, 0, 0, 1, '谷歌之家', ' ', '', '', 30, 1, '', '', '', 'google', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (7, 0, 0, 0, 1, '壁纸之家', '壁纸|壁纸下载', '这是一个描述信息', '', 36, 1, '', '', '', 'bizhi', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (10, 0, 1, 1, 1, 'Win8快讯', ' ', '', '', 30, 1, '', '', '', 'win8news', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (11, 0, 1, 1, 1, 'Win8学院', ' ', '', '', 30, 1, '', '', '', 'win8xueyuan', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (12, 0, 1, 1, 1, 'Win8系统下载', ' ', '', '', 30, 1, '', '', '', 'win8xiazai', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (13, 0, 1, 1, 1, 'Win8软件下载', ' ', '', '', 30, 1, '', '', '', 'win8soft', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (14, 0, 1, 1, 1, 'Win8主题', ' ', '', '', 30, 1, '', '', '', 'win8zhuti', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (15, 0, 2, 2, 1, 'IT资讯', ' ', '', '', 30, 1, '', '', '', 'it', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (16, 0, 2, 2, 1, 'QQ之家', '  ', '', '', 30, 1, '', '', '', 'qq', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (17, 0, 2, 2, 1, 'IT之外', '  ', '', '', 30, 1, '', '', '', 'other', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (18, 1, 0, 0, 0, '投稿反馈', '', '', '', 0, 0, '', '', '', 'tougao', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (19, 1, 0, 0, 0, '加入我们', '', '', '', 0, 0, '', '', '', 'join', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (20, 1, 0, 0, 0, '联系我们', '', '', '', 0, 0, '', '', '', 'contact', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (21, 1, 0, 0, 0, '关于我们', '', '', '', 0, 0, '', '', '', 'about', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (23, 0, 0, 0, 12, '测试模型', '12313123', 'asdasd', '', 8, 1, '', '', '', 'aasdasd', NULL, NULL, '2021-08-10 16:58:32');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_category_priv
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_category_priv`;
CREATE TABLE `pinecms_category_priv` (
  `catid` int NOT NULL,
  `roleid` int NOT NULL,
  `is_admin` int NOT NULL,
  `action` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_category_priv
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pinecms_department
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_department`;
CREATE TABLE `pinecms_department` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `leader_name` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `leader_phone` varchar(35) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `listorder` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_department
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_department` VALUES (2, '技术部', '陈', '17610053500', '826466266@qq.com', 1, 12, NULL, 1);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_dict
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_dict`;
CREATE TABLE `pinecms_dict` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cid` int DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `value` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `UQE_pinecms_dict_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_dict
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_dict` VALUES (4, 10, '0-2000', '0-2000', 0, '2021-08-11 22:24:58', '2021-09-25 13:52:19', '价格区间', 1);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_dict_category
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_dict_category`;
CREATE TABLE `pinecms_dict_category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `UQE_pinecms_dict_category_key` (`key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_dict_category
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_dict_category` VALUES (10, 'prices', '价格', 1, '2021-08-11 22:24:50', '2021-09-25 13:49:38', '系统内价格调用标识');
INSERT INTO `pinecms_dict_category` VALUES (13, 'age', '年龄', 1, '2021-10-28 17:28:10', '2021-10-28 17:28:10', '');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_district
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_district`;
CREATE TABLE `pinecms_district` (
  `id` smallint DEFAULT NULL,
  `name` varchar(270) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `parent_id` smallint DEFAULT NULL,
  `initial` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `initials` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `pinyin` varchar(600) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `extra` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `suffix` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `code` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `area_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `order` tinyint DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_district
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pinecms_document_model
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_document_model`;
CREATE TABLE `pinecms_document_model` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `table` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `enabled` int DEFAULT NULL,
  `model_type` int DEFAULT NULL,
  `fe_tpl_index` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `fe_tpl_list` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `fe_tpl_detail` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `deleted_at` datetime DEFAULT NULL,
  `execed` int DEFAULT NULL,
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_document_model
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_document_model` VALUES (1, '普通文章', 'articles', 1, 1, '', 'asdasd.jet', 'list/index.jet', NULL, 1, '系统默认模型', NULL, NULL);
INSERT INTO `pinecms_document_model` VALUES (4, '招聘模型', 'zhaopin', 0, 1, '', '', '', NULL, 0, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model` VALUES (5, '商城模型', 'shop', 1, 1, '', 'list_shop.jet', 'index.jet', NULL, 0, '这是一个商城模型， 用于展示商城数据', NULL, NULL);
INSERT INTO `pinecms_document_model` VALUES (7, '新模型', 'new_article', 1, 0, '', 'index.jet', 'asdasd.jet', '2021-08-10 16:23:28', 0, 'asdasad', '0001-01-01 00:00:00', '0001-01-01 00:00:00');
INSERT INTO `pinecms_document_model` VALUES (8, 'asdasd', 'asdasd', 1, 0, '', 'asdasd.jet', 'index.jet', '2021-08-10 16:24:32', 0, 'asdasdasdasd', '0001-01-01 00:00:00', '0001-01-01 00:00:00');
INSERT INTO `pinecms_document_model` VALUES (9, 'asdasd', 'asdasd', 1, 0, '', 'asdasd.jet', 'index.jet', '2021-08-10 16:27:15', 0, 'asdsadsaasd', '0001-01-01 00:00:00', '0001-01-01 00:00:00');
INSERT INTO `pinecms_document_model` VALUES (10, 'asdasd', 'asdasdasd', 1, 0, '', 'index.jet', 'list/index.jet', '2021-08-10 16:28:32', 0, 'asdadasd', '0001-01-01 00:00:00', '0001-01-01 00:00:00');
INSERT INTO `pinecms_document_model` VALUES (11, 'adasd', 'asdasdasdasd', 1, 0, '', '', '', '2021-08-10 16:31:34', 0, '', '0001-01-01 00:00:00', '0001-01-01 00:00:00');
INSERT INTO `pinecms_document_model` VALUES (12, '测试模型', 'download', 1, 0, '', '', '', NULL, 0, '', '0001-01-01 00:00:00', '0001-01-01 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_document_model_dsl
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_document_model_dsl`;
CREATE TABLE `pinecms_document_model_dsl` (
  `id` int NOT NULL AUTO_INCREMENT,
  `mid` int NOT NULL,
  `form_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `listorder` int NOT NULL,
  `table_field` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `required` int DEFAULT NULL,
  `datasource` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `required_tips` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `validator` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `deleted_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `field_type` int NOT NULL,
  `default` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `is_system` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL COMMENT '状态 0=禁用 1=启用',
  `main_table_field` tinyint(1) DEFAULT NULL COMMENT '睡',
  `searchable` tinyint(1) DEFAULT NULL COMMENT '是否可搜索',
  `sortable` tinyint(1) DEFAULT NULL COMMENT '是否可排序',
  `span` tinyint DEFAULT NULL COMMENT '表单span宽度',
  `visible` tinyint(1) DEFAULT NULL COMMENT '是否表单可见',
  `list_visible` tinyint(1) DEFAULT NULL COMMENT '是否列表可见',
  `field_len` bigint DEFAULT NULL COMMENT '字段长度',
  `list_width` int DEFAULT NULL COMMENT '列表字段宽度',
  `component` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '自定义组件配置',
  `center` tinyint(1) DEFAULT NULL COMMENT '是否列表居中显示内容',
  `is_dict` tinyint(1) DEFAULT NULL,
  `dict_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典',
  `search_type` int DEFAULT NULL COMMENT '搜索类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=472 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_document_model_dsl
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_document_model_dsl` VALUES (1, 0, '标题', 1, 'title', 1, '', '', '', NULL, 1, '', NULL, '2021-08-02 15:54:52', NULL, 0, 0, 0, 1, 4, 0, 1, 0, 4, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (2, 0, '关键字', 3, 'keywords', 0, '', '', '', NULL, 1, '', NULL, '2021-08-02 09:54:03', NULL, 0, 0, 1, 0, 4, 0, 1, 0, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (3, 0, '摘要', 4, 'description', 0, '', '', '', NULL, 2, '', NULL, '2021-08-02 09:54:09', NULL, 0, 0, 0, 0, 4, 0, 1, 0, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (4, 0, '访问次数', 99, 'visit_count', 0, '', '', '', NULL, 9, '', NULL, '2021-08-02 15:31:38', NULL, 0, 0, 0, 1, 4, 0, 1, 0, 4, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (5, 0, '状态', 100, 'status', 1, '', '', '', NULL, 13, '50', NULL, '2021-08-02 15:49:49', NULL, 0, 0, 0, 1, 4, 0, 1, 0, 4, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (6, 0, '文档属性', 2, 'flag', 0, '', '', '', NULL, 16, '', NULL, '2021-08-02 09:53:54', NULL, 0, 0, 0, 0, 4, 0, 1, 0, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (7, 0, '排序', 50, 'listorder', 0, '', '', '', NULL, 9, '', NULL, '2021-08-02 15:31:23', NULL, 0, 0, 0, 0, 4, 0, 1, 0, 4, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (8, 0, '发布时间', 51, 'pubtime', 0, '', '', '', NULL, 14, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (9, 0, 'TAG标签', 5, 'tags', 1, '', '标签必须填写', '', NULL, 15, '', NULL, '2021-08-01 09:09:12', NULL, 1, 0, 1, 1, 4, 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (412, 1, '文章标题', 1, 'title', 0, '', '', '', NULL, 1, '', NULL, '2021-08-15 13:47:33', NULL, 1, 0, 0, 1, 24, 0, 1, 0, 150, '', 0, 0, '', 2);
INSERT INTO `pinecms_document_model_dsl` VALUES (413, 1, '缩略图', 2, 'thumb', 0, '', '', '', NULL, 11, '', NULL, '2021-08-14 16:24:06', NULL, 1, 0, 0, 0, 24, 0, 1, 0, 80, '', 1, 0, '', 1);
INSERT INTO `pinecms_document_model_dsl` VALUES (414, 1, '关键字', 2, 'keywords', 0, '', '', '', NULL, 1, '', NULL, '2021-08-14 16:24:12', NULL, 1, 0, 1, 0, 20, 0, 1, 0, 80, '', 0, 0, '', 1);
INSERT INTO `pinecms_document_model_dsl` VALUES (415, 1, '文档属性', 2, 'flag', 0, '', '', '', NULL, 16, '', NULL, '2021-08-14 16:24:24', NULL, 1, 0, 1, 0, 24, 0, 1, 0, 80, '', 0, 0, '', 1);
INSERT INTO `pinecms_document_model_dsl` VALUES (416, 1, '内容摘要', 3, 'description', 0, '', '', '', NULL, 2, '', NULL, '2021-08-14 16:28:02', NULL, 1, 0, 1, 0, 24, 0, 1, 0, 80, '', 0, 0, '', 1);
INSERT INTO `pinecms_document_model_dsl` VALUES (417, 1, 'TAG标签', 4, 'tags', 0, '', '', '', NULL, 15, '', NULL, '2021-08-14 16:28:10', NULL, 1, 0, 1, 0, 24, 0, 1, 0, 80, '', 0, 0, '', 1);
INSERT INTO `pinecms_document_model_dsl` VALUES (418, 1, '文章来源', 5, 'from_url', 0, '', '', '', NULL, 1, 'pinecms.git', NULL, '2021-08-14 16:28:20', NULL, 1, 0, 1, 0, 16, 0, 1, 0, 80, '', 0, 0, '', 1);
INSERT INTO `pinecms_document_model_dsl` VALUES (419, 1, '作者', 6, 'author', 0, '', '', '', NULL, 1, 'pinecms', NULL, '2021-08-14 16:28:26', NULL, 1, 0, 1, 0, 20, 0, 1, 0, 80, '', 0, 0, '', 1);
INSERT INTO `pinecms_document_model_dsl` VALUES (420, 1, '文章内容', 7, 'content', 1, '', '', '', NULL, 3, '', NULL, '2021-10-07 15:09:18', NULL, 1, 0, 1, 1, 24, 1, 0, 0, 140, '', 0, 0, '', 1);
INSERT INTO `pinecms_document_model_dsl` VALUES (421, 1, '发布时间', 8, 'pubtime', 0, '', '', '', NULL, 14, '', NULL, '2021-08-14 16:28:44', NULL, 1, 0, 0, 0, 16, 0, 1, 0, 80, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (422, 1, '状态', 10, 'status', 0, '', '', '', NULL, 13, '1', NULL, '2021-08-14 16:28:54', NULL, 1, 0, 0, 0, 8, 0, 1, 0, 80, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (423, 1, '访问次数', 11, 'visit_count', 0, '', '', '', NULL, 9, '', NULL, '2021-08-14 16:29:18', NULL, 1, 0, 0, 0, 4, 0, 1, 0, 80, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (424, 1, '排序', 12, 'listorder', 0, '', '', '', NULL, 9, '30', NULL, '2021-08-14 16:14:17', NULL, 1, 0, 0, 1, 4, 0, 1, 0, 80, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (425, 1, '幻灯片', 888, 'flash', 0, '', '', '', NULL, 11, '', NULL, '2021-08-14 16:29:28', NULL, 1, 0, 0, 0, 24, 0, 0, 0, 80, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (426, 5, '商品名称', 1, 'title', 0, '', '', '', NULL, 1, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (427, 5, '关键字', 2, 'keywords', 0, '', '', '', NULL, 1, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (428, 5, '文档属性', 2, 'flag', 0, '', '', '', NULL, 16, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (429, 5, '摘要', 3, 'description', 0, '', '', '', NULL, 2, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (430, 5, 'TAG标签', 4, 'tags', 0, '', '', '', NULL, 15, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (431, 5, '缩略图', 5, 'thumbs', 0, '', '', '', NULL, 12, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (432, 5, '价格', 6, 'price', 0, '', '请填写商品价格', '', NULL, 10, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (433, 5, '机身颜色', 7, 'color', 0, '[\"银色\", \"黑色\", \"白色\", \"皓月银\", \"其他\"]', '', '', NULL, 7, '其他', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (434, 5, '价格区间', 8, 'price_level', 0, '[\"0-1000\", \"1000-1699\",\"1700-2799\",\"2800-3500\",\"3500-10000\", \"10000以上\"]', '', '', NULL, 7, '0-1000', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (435, 5, '发布时间', 10, 'pubtime', 0, '[\"2020年\", \"2019年\",\"2018年\", \"2017年及以前\"]', '', '', NULL, 7, '2020年', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (436, 5, '内容详情', 11, 'content', 0, '', '', '', NULL, 3, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (437, 5, '访问次数', 99, 'visit_count', 0, '', '', '', NULL, 9, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (438, 5, '状态', 100, 'status', 0, '', '', '', NULL, 13, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (439, 5, '排序', 101, 'listorder', 1, '', '', '', NULL, 9, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (440, 4, '标题', 1, 'title', 0, '', '', '', NULL, 1, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (441, 4, '关键字', 2, 'keywords', 0, '', '', '', NULL, 1, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (442, 4, '文档属性', 2, 'flag', 0, '', '', '', NULL, 16, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (443, 4, '摘要', 3, 'description', 0, '', '', '', NULL, 2, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (444, 4, '工作地点', 5, 'address', 0, '', '', '', NULL, 1, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (445, 4, '学历要求', 6, 'education', 0, '', '', '', NULL, 1, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (446, 4, '薪资待遇', 7, 'salary', 0, '[\"0-3000元\", \"3001-5000\", \"5001-10000元\", \"10000以上\"]', '', '', NULL, 5, '0', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (447, 4, '工作性质', 8, 'job_type', 0, '[\"自由\", \"全职\", \"面议\"]', '', '', NULL, 5, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (448, 4, '工作年限', 9, 'working_years', 0, '[\"1-4年\", \"3-5年\", \"5年以上\", \"不限\"]', '', '', NULL, 5, '3', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (449, 4, '招聘人数', 10, 'total', 0, '', '', '', NULL, 9, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (450, 4, '招聘内容', 11, 'content', 0, '', '', '', NULL, 3, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (451, 4, '排序值', 50, 'listorder', 0, '', '', '', NULL, 9, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (452, 4, '访问次数', 99, 'visit_count', 0, '', '', '', NULL, 9, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (453, 4, '状态', 100, 'status', 0, '', '', '', NULL, 13, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (463, 12, '标题', 1, 'title', 1, '', '', '', NULL, 1, '', '2021-08-10 16:33:55', '2021-08-10 16:33:55', NULL, 0, 0, 0, 1, 4, 0, 1, 0, 4, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (464, 12, '关键字', 3, 'keywords', 0, '', '', '', NULL, 1, '', '2021-08-10 16:33:55', '2021-08-10 16:33:55', NULL, 0, 0, 1, 0, 4, 0, 1, 0, 0, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (465, 12, '摘要', 4, 'description', 0, '', '', '', NULL, 2, '', '2021-08-10 16:33:56', '2021-08-10 16:33:56', NULL, 0, 0, 0, 0, 4, 0, 1, 0, 0, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (466, 12, '访问次数', 99, 'visit_count', 0, '', '', '', NULL, 9, '', '2021-08-10 16:33:56', '2021-08-10 16:33:56', NULL, 0, 0, 0, 1, 4, 0, 1, 0, 4, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (467, 12, '状态', 100, 'status', 1, '', '', '', NULL, 13, '50', '2021-08-10 16:33:56', '2021-08-10 16:33:56', NULL, 0, 0, 0, 1, 4, 0, 1, 0, 4, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (468, 12, '文档属性', 2, 'flag', 0, '', '', '', NULL, 16, '', '2021-08-10 16:33:56', '2021-08-10 16:33:56', NULL, 0, 0, 0, 0, 4, 0, 1, 0, 0, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (469, 12, '排序', 50, 'listorder', 0, '', '', '', NULL, 9, '', '2021-08-10 16:33:56', '2021-08-10 16:33:56', NULL, 0, 0, 0, 0, 4, 0, 1, 0, 4, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (470, 12, '发布时间', 51, 'pubtime', 0, '', '', '', NULL, 14, '', '2021-08-10 16:33:56', '2021-08-10 16:33:56', NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, '', 0, 0, '', NULL);
INSERT INTO `pinecms_document_model_dsl` VALUES (471, 12, 'TAG标签', 5, 'tags', 1, '', '标签必须填写', '', NULL, 15, '', '2021-08-10 16:33:56', '2021-08-10 16:33:56', NULL, 1, 0, 1, 1, 4, 1, 1, 0, 0, '', 0, 0, '', NULL);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_document_model_field
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_document_model_field`;
CREATE TABLE `pinecms_document_model_field` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `type` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `list_comp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `form_comp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `props` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '属性配置',
  `html` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '自定义html',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_document_model_field
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_document_model_field` VALUES (1, '输入框', 'varchar', '常用字段，如文章标题、作者等都属于直接输入少量内容的文本，设置这个文本之后需要指定文本长度，默认为250，如果大于255则为text类型', NULL, 'el-input', NULL, NULL);
INSERT INTO `pinecms_document_model_field` VALUES (2, '多行输入框', 'text', '也是较为常用的字段类型，如个人简介、产品描述都可以使用多行文本进行存储', NULL, 'el-input', '{\"type\":\"textarea\",\"row\":4}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (3, '富文本编辑器', 'text', '编辑器编辑产生的html内容，用于比较复杂的内容形式, 可以认为是附带编辑器的多行文本', NULL, 'cl-editor-quill', '{\"height\": \"350px\"}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (4, '附件', 'varchar', '上传附件', NULL, 'cl-upload-space', '{\"accept\":\".txt,.doc,.docx,.xls,.xlsx,.csv\",\"limit\":10,\"multiple\":true,\"limit-size\":20}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (5, '下拉框', 'varchar', '下拉选择，一般用于如软件类型、语言类型等字段', NULL, 'el-select', NULL, NULL);
INSERT INTO `pinecms_document_model_field` VALUES (6, '联动类型', 'varchar', '一种数组形式的数据类型，请使用url接口方式提供', NULL, 'el-cascader', NULL, NULL);
INSERT INTO `pinecms_document_model_field` VALUES (7, '单选框', 'varchar', '平铺显示, 可以认为是下拉框的展开, 根据数据源展开为排列的组件', NULL, 'el-radio', NULL, NULL);
INSERT INTO `pinecms_document_model_field` VALUES (8, '多选框', 'varchar', '多选框, 平铺显示为多个选项,根据数据源展开为排列组件', NULL, 'cms-checkbox', NULL, NULL);
INSERT INTO `pinecms_document_model_field` VALUES (9, '整数类型', 'int', '常用字段, 仅能输入数字', NULL, 'el-input-number', '{\"step\":1, \"stepStrictly\": true}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (10, '浮点类型', 'float', '常用字段, 仅能输入浮点数(小数)', NULL, 'el-input-number', '{\"step\":0.01, \"stepStrictly\": true, precision: 2}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (11, '单图上传', 'varchar', '常用字段, 会生成一个单图上传框', '{name: \"el-image\", fit: \"contain\", style: {\"width\": \"40px\", \"height\": \"40px\"}}', 'cl-upload-space', '{\"accept\":\".png,.jpg,.bmp,.gif,.jpeg\",\"limit\":1,\"multiple\":false,\"limit-size\":20}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (12, '多图上传', 'varchar', '生成一个多图上传的组件', NULL, 'cl-upload-space', '{\"accept\":\".png,.jpg,.bmp,.gif,.jpeg\",\"limit\":10,\"multiple\":true,\"limit-size\":20}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (13, '开关按钮', 'tinyint', '用于做开关值的组件, 打开为1, 关闭为0', NULL, 'el-switch', '{\"active-value\": 1, \"inactive-value\": 0}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (14, '日历组件', 'datetime', '选择日期组件', NULL, 'el-date-picker', NULL, NULL);
INSERT INTO `pinecms_document_model_field` VALUES (15, '多选标签', 'varchar', '可以记录标签，并多选', NULL, 'input-tags', '{\"theme\":\"新标签\"}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (16, '文档属性', 'varchar', '标记文档属性，默认每个文档都存在该属性， 无需额外添加', NULL, 'cms-checkbox', '{\"options\":[{\"label\":\"头条\",\"key\":\"h\"},{\"label\":\"推荐\",\"key\":\"c\"},{\"label\":\"图片\",\"key\":\"p\"},{\"label\":\"幻灯\",\"key\":\"f\"},{\"label\":\"跳转\",\"key\":\"j\"},{\"label\":\"图文\",\"key\":\"a\"},{\"label\":\"加粗\",\"key\":\"b\"}]}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (17, 'Markdown', 'text', 'markdown编辑器', NULL, 'el-input', NULL, NULL);
INSERT INTO `pinecms_document_model_field` VALUES (18, 'Code编辑器', 'text', '代码编辑器', NULL, 'el-codemirror', NULL, NULL);
INSERT INTO `pinecms_document_model_field` VALUES (19, '评分组件', 'float', '一般用于评级选择, 可以自定义软件,评价等', NULL, 'el-rate', '{\"colors\": [\"#99A9BF\", \"#F7BA2A\", \"#FF9900\"]}', NULL);
INSERT INTO `pinecms_document_model_field` VALUES (20, '百度编辑器', 'text', '用于文章发布等字段', NULL, 'vue-ueditor-wrap', '{\"config\": \"@ueditorConf\"}', NULL);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_download
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_download`;
CREATE TABLE `pinecms_download` (
  `id` int NOT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `keywords` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `files` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `ext` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `size` double DEFAULT NULL,
  `downs` int DEFAULT NULL,
  `catid` int NOT NULL,
  `mid` int NOT NULL,
  `listorder` int NOT NULL,
  `visit_count` int NOT NULL,
  `status` int NOT NULL,
  `created_time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `updated_time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `deleted_time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_download
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_download` VALUES (1, 'sadas', 'dasd', 'asd', '<p>asdasdasd</p>', '[{\"title\":\"1jPeJ02eR9.gif\",\"url\":\"/upload/public/20200321/1jPeJ02eR9.gif\",\"type\":\"gif\"},{\"title\":\"ThECSqh65Q.zip\",\"url\":\"/upload/public/20200321/ThECSqh65Q.zip\",\"type\":\"zip\"}]', '11', 111, 1111, 6, 2, 0, 0, 1, '2020-03-22 12:10:40', '2020-03-22 12:14:58', '2020-03-28 15:43:33');
INSERT INTO `pinecms_download` VALUES (2, 'hello xiazai ', 'adad', 'asdsadad', '<p>asdasd</p>', '[{\"title\":\"hei_logo.png\",\"url\":\"/upload/public/20200321/s49kowaHVY.png\",\"type\":\"png\"},{\"title\":\"1.jpg\",\"url\":\"http://iriscms-test.oss-cn-beijing.aliyuncs.com/upload/public/20200324/p9zaDoQ1ZQ.jpg\",\"type\":\"jpg\"},{\"title\":\"2.jpg\",\"url\":\"http://iriscms-test.oss-cn-beijing.aliyuncs.com/upload/public/20200324/8DF0DlN4xQ.jpg\",\"type\":\"jpg\"},{\"title\":\"ad02.jpg\",\"url\":\"http://iriscms-test.oss-cn-beijing.aliyuncs.com/upload/public/20200324/40B5hc9v5I.jpg\",\"type\":\"jpg\"}]', 'aa', 13, 1313, 6, 2, 0, 0, 1, '2020-03-28 15:40:29', '2020-03-28 15:43:13', '2020-03-28 15:43:36');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_link
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_link`;
CREATE TABLE `pinecms_link` (
  `id` int NOT NULL AUTO_INCREMENT,
  `linktype` tinyint NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `logo` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `listorder` int NOT NULL,
  `passed` tinyint unsigned NOT NULL DEFAULT '0',
  `addtime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_link
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_link` VALUES (1, 0, '腾讯云', 'http://www.qcloud.com', 'http://pinecms.oss-cn-beijing.aliyuncs.com/uploads/Ymd/9e95355PEfIkWJ4B.jpeg', '国内常用的云端平台', 0, 1, '0001-01-01 00:00:00');
INSERT INTO `pinecms_link` VALUES (2, 0, '修心小站', 'http://www.xiusin.cn', 'http://pinecms.oss-cn-beijing.aliyuncs.com/uploads/Ymd/hSH73fTxkmw7O2i0.jpg', '个人网站', 0, 1, '0001-01-01 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_log
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_log`;
CREATE TABLE `pinecms_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `level` tinyint DEFAULT NULL COMMENT '日志类型',
  `uri` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '请求uri',
  `method` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '请求方法',
  `params` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '请求参数',
  `message` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '操作用户名',
  `stack` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '调用堆栈',
  `ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '操作IP',
  `time` datetime DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pinecms_member
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_member`;
CREATE TABLE `pinecms_member` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '账号',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '密码',
  `avatar` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '头像',
  `nickname` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '昵称',
  `integral` int DEFAULT NULL COMMENT '积分',
  `telphone` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '电话',
  `qq` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'QQ',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '个人简介',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `login_time` datetime DEFAULT NULL,
  `login_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `email` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '邮箱',
  `status` int DEFAULT NULL COMMENT '状态: 0=禁用 1=待验证 2=正常',
  `sex` tinyint DEFAULT NULL COMMENT '性别: 0=保密 1=男 2=女',
  `verify_token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '验证token',
  `group_id` int DEFAULT NULL COMMENT '分组ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_member
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_member` VALUES (1, 'xiusin', 'xiusin', 'http://pinecms.oss-cn-beijing.aliyuncs.com/uploads/Ymd/yLyM1VIZ8e6iax7k.jpeg', 'xiusin', 0, '16601313660', 'xiusin', 'xiusin', '2021-07-30 17:24:36', '2021-09-25 09:53:05', '0001-01-01 00:00:00', '', 'xiusin@qq.com', 2, 1, '', 1);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_member_group
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_member_group`;
CREATE TABLE `pinecms_member_group` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '名称',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '介绍',
  `status` int DEFAULT NULL COMMENT '状态: 0=禁用 1=正常',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `listorder` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_member_group
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_member_group` VALUES (1, '普通会员', '注册即可成为普通会员', 1, '2021-07-30 17:56:38', '2021-09-25 09:49:36', 1);
INSERT INTO `pinecms_member_group` VALUES (2, '中级会员', '客户签到达30天可升级为中级会员', 1, '2021-09-25 09:31:27', '2021-09-25 09:36:35', 2);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_menu
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_menu`;
CREATE TABLE `pinecms_menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `identification` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '单一权限标识',
  `parentid` int NOT NULL,
  `listorder` int NOT NULL,
  `display` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `type` int DEFAULT NULL COMMENT '类型 0：目录 1：菜单 2：按钮',
  `view_path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '页面路径',
  `path` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `icon` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `keep_alive` int DEFAULT NULL,
  `router` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `plugin_id` int DEFAULT '0' COMMENT '菜单来源插件',
  `perms` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '权限标识(绑定的权限标识, 比如edit,需要info权限)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `pinecms_menu_identification_uindex` (`identification`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=126 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_menu
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_menu` VALUES (3, '内容管理', NULL, 0, 3, '1', 0, NULL, NULL, 'icon-content', NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (9, '系统设置', NULL, 0, 1, '1', 0, NULL, NULL, 'icon-system', NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (10, '站点设置', NULL, 9, 1, '1', 1, 'cool/modules/base/views/setting.vue', NULL, 'icon-system', NULL, '/sys/setting', 0, 'base:sys:site:config');
INSERT INTO `pinecms_menu` VALUES (11, '权限管理', NULL, 0, 2, '1', 0, NULL, NULL, 'icon-dept', NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (12, '管理员管理', NULL, 11, 1, '1', 1, 'cool/modules/base/views/user.vue', NULL, 'icon-admin', NULL, '/sys/user', 0, 'base:sys:admin:list');
INSERT INTO `pinecms_menu` VALUES (13, '角色管理', NULL, 11, 2, '1', 1, 'cool/modules/base/views/role.vue', NULL, 'icon-role', NULL, '/sys/role', 0, 'base:sys:role:list');
INSERT INTO `pinecms_menu` VALUES (14, '日志监控', NULL, 0, 9, '1', 0, NULL, NULL, 'icon-log', NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (15, '操作日志', NULL, 14, 3, '1', 1, 'cool/modules/base/views/log.vue', NULL, 'icon-log-new', NULL, '/sys/log', 0, 'base:sys:opt:log:list');
INSERT INTO `pinecms_menu` VALUES (16, '菜单管理', NULL, 9, 2, '1', 1, 'cool/modules/base/views/menu.vue', NULL, 'icon-menu', NULL, '/sys/menu', 0, 'base:sys:menu:list');
INSERT INTO `pinecms_menu` VALUES (17, '查看菜单', NULL, 16, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (18, '添加菜单', NULL, 16, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (19, '修改菜单', NULL, 16, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (20, '删除菜单', NULL, 16, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (21, '菜单排序', NULL, 16, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (22, '查看日志', NULL, 15, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (23, '删除日志', NULL, 15, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (24, '查看管理员', NULL, 12, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (25, '添加管理员', NULL, 12, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (26, '编辑管理员', NULL, 12, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (27, '删除管理员', NULL, 12, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (28, '查看角色', NULL, 13, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (29, '添加角色', NULL, 13, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (30, '编辑角色', NULL, 13, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (31, '删除角色', NULL, 13, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (34, '发布管理', NULL, 3, 0, '1', 1, 'cool/modules/base/views/content.vue', NULL, 'icon-content', NULL, '/sys/content', 0, 'base:sys:content:list');
INSERT INTO `pinecms_menu` VALUES (35, '内容管理', NULL, 3, 0, '0', 0, 'cool/modules/base/views/content.vue', NULL, NULL, NULL, '/sys/content', 0, NULL);
INSERT INTO `pinecms_menu` VALUES (36, '栏目管理', NULL, 3, 0, '1', 1, 'cool/modules/base/views/category.vue', NULL, 'icon-menu', NULL, '/sys/category', 0, 'base:sys:category:list');
INSERT INTO `pinecms_menu` VALUES (37, '查看栏目', NULL, 36, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (38, '添加栏目', NULL, 36, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (39, '编辑栏目', NULL, 36, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (40, '删除栏目', NULL, 36, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (41, '栏目排序', NULL, 36, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (53, '新闻列表', NULL, 35, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (54, '分类单页', NULL, 35, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (55, '会员管理', NULL, 0, 5, '1', 0, NULL, NULL, 'icon-user', NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (56, '会员管理', NULL, 55, 0, '1', 1, 'cool/modules/base/views/member.vue', NULL, 'icon-user', NULL, '/sys/member', 0, 'base:sys:member:list');
INSERT INTO `pinecms_menu` VALUES (57, '会员信息', NULL, 56, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (58, '微信管理', NULL, 0, 7, '1', 0, NULL, NULL, 'icon-wechat', NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (59, '账号管理', NULL, 58, 1, '1', 1, 'cool/modules/wechat/views/account.vue', NULL, 'icon-wechat', NULL, '/sys/wechat/account', 0, 'base:sys:account:list');
INSERT INTO `pinecms_menu` VALUES (60, '编辑会员', NULL, 55, 0, '0', 1, NULL, NULL, NULL, NULL, NULL, 0, 'base:sys:');
INSERT INTO `pinecms_menu` VALUES (62, '模型管理', NULL, 3, 1, '1', 1, 'cool/modules/base/views/model.vue', NULL, 'icon-pic', NULL, '/sys/model', 0, 'base:sys:model:list');
INSERT INTO `pinecms_menu` VALUES (64, '添加模型', NULL, 62, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (65, '会员管理', NULL, 58, 2, '1', 1, 'cool/modules/wechat/views/user.vue', NULL, 'icon-user', NULL, '/sys/wechat/user', 0, 'base:sys:wechat:user:list');
INSERT INTO `pinecms_menu` VALUES (67, '友链管理', NULL, 0, 88, '1', 0, NULL, NULL, 'icon-link', NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (68, '友链管理', NULL, 67, 88, '1', 1, 'cool/modules/base/views/link.vue', NULL, 'icon-link', NULL, '/sys/link', 0, 'base:sys:link:list');
INSERT INTO `pinecms_menu` VALUES (69, '友链添加', NULL, 68, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (70, '友链编辑', NULL, 68, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (71, '友链删除', NULL, 68, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (72, '友链排序', NULL, 68, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (73, '数据库管理', NULL, 0, 16, '1', 0, NULL, NULL, 'icon-monitor', NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (74, '数据库管理', NULL, 73, 0, '1', 1, 'cool/modules/base/views/database_list.vue', NULL, 'icon-database', 0, '/sys/database/list', 0, 'base:sys:dbm:list');
INSERT INTO `pinecms_menu` VALUES (75, '数据库备份', NULL, 74, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (76, '数据库优化', NULL, 74, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (77, '数据库修复', NULL, 74, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (78, '备份列表', NULL, 73, 0, '1', 1, 'cool/modules/base/views/database_backup_list.vue', NULL, 'icon-list', NULL, '/sys/database/backup/list', 0, 'base:sys:dbm:backup:list');
INSERT INTO `pinecms_menu` VALUES (79, '资源管理', NULL, 0, 17, '1', 0, 'cool/modules/base/views/assets_manager.vue', NULL, 'icon-pic', NULL, '/sys/assets-manager', 0, NULL);
INSERT INTO `pinecms_menu` VALUES (80, '模板列表', NULL, 79, 0, '1', 1, 'cool/modules/base/views/assets_manager.vue', NULL, 'icon-template', NULL, '/sys/assets-manager', 0, 'base:sys:template:list');
INSERT INTO `pinecms_menu` VALUES (81, '添加资源', NULL, 80, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (82, '修改资源', NULL, 80, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (84, '附件列表', NULL, 79, 0, '1', 1, 'cool/modules/base/views/attachment.vue', NULL, 'icon-pic', NULL, '/sys/attachments', 0, 'base:sys:attachments:list');
INSERT INTO `pinecms_menu` VALUES (85, '系统监控', NULL, 14, 2, '1', 1, 'cool/modules/base/views/statsviz.vue', NULL, 'icon-workbench', 0, '/statsviz', 0, 'base:sys:statsviz');
INSERT INTO `pinecms_menu` VALUES (86, '修改模型', NULL, 62, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (87, '主题列表', NULL, 79, 0, '1', 1, 'cool/modules/base/views/theme.vue', NULL, 'icon-theme', NULL, '/sys/theme', 0, 'base:sys:theme:list');
INSERT INTO `pinecms_menu` VALUES (89, '广告管理', NULL, 9, 0, '1', 1, 'cool/modules/base/views/ad.vue', NULL, 'icon-ad', NULL, '/ad/list', 0, 'base:sys:ad:list');
INSERT INTO `pinecms_menu` VALUES (91, '添加', NULL, 89, 0, '1', 2, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (92, '修改', NULL, 89, 0, '1', 2, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (93, '删除', NULL, 89, 0, '1', 2, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (94, '添加', NULL, 89, 0, '1', 2, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (95, '删除', NULL, 89, 0, '1', 2, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (96, '编辑', NULL, 89, 0, '1', 2, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (98, '缓存管理', NULL, 9, 0, '0', 1, NULL, NULL, NULL, NULL, NULL, 0, 'base:sys:cache:list');
INSERT INTO `pinecms_menu` VALUES (99, '发布内容', NULL, 35, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (100, '修改内容', NULL, 35, 0, '1', 0, NULL, NULL, NULL, NULL, NULL, 0, NULL);
INSERT INTO `pinecms_menu` VALUES (101, '字典管理', NULL, 9, 0, '1', 1, 'cool/modules/base/views/dict.vue', NULL, 'icon-dict', 0, '/sys/dict', 0, 'base:sys:dict:list');
INSERT INTO `pinecms_menu` VALUES (103, '部门管理', NULL, 11, 0, '1', 1, 'cool/modules/base/views/department.vue', NULL, 'icon-dept', NULL, '/sys/department', 0, 'base:sys:department:list');
INSERT INTO `pinecms_menu` VALUES (104, '岗位管理', NULL, 11, 0, '1', 1, 'cool/modules/base/views/position.vue', NULL, 'icon-position', NULL, '/sys/position', 0, 'base:sys:position:list');
INSERT INTO `pinecms_menu` VALUES (105, '服务器状态', NULL, 14, 1, '1', 1, 'cool/modules/base/views/stat.vue', NULL, 'icon-status', NULL, '/sys/stat', 0, 'base:sys:stat:list');
INSERT INTO `pinecms_menu` VALUES (107, '任务管理', NULL, 108, 0, '1', 1, 'cool/modules/task/views/task.vue', NULL, 'icon-task', NULL, '/sys/task', 0, 'base:sys::task:list');
INSERT INTO `pinecms_menu` VALUES (108, '插件管理', NULL, 0, 99, '1', 0, '', NULL, 'icon-plugin', 0, '', 0, NULL);
INSERT INTO `pinecms_menu` VALUES (109, '标签管理', NULL, 3, 0, '1', 1, 'cool/modules/base/views/tags.vue', NULL, 'icon-tag', NULL, '/sys/tags', 0, 'base:sys:content:tag:list');
INSERT INTO `pinecms_menu` VALUES (110, '插件列表', NULL, 108, 0, '1', 1, 'cool/modules/base/views/plugin.vue', NULL, 'icon-plugin', NULL, '/sys/plugin', 0, 'base:sys:plugin:list');
INSERT INTO `pinecms_menu` VALUES (111, '分组管理', NULL, 55, 0, '1', 1, 'cool/modules/base/views/member_group.vue', NULL, 'icon-group', NULL, '/sys/member/group', 0, 'base:sys:admin:group:list');
INSERT INTO `pinecms_menu` VALUES (112, '字段管理', NULL, 3, 0, '0', 1, 'cool/modules/base/views/table.vue', NULL, NULL, NULL, '/sys/table/fields', 0, 'base:sys:model:table:field:list');
INSERT INTO `pinecms_menu` VALUES (113, '日志管理', NULL, 14, 4, '1', 1, 'cool/modules/base/views/errlog.vue', NULL, 'icon-log', 1, '/sys/errlog/list', 0, 'base:sys:log:error:list');
INSERT INTO `pinecms_menu` VALUES (114, '区域管理', NULL, 9, 0, '1', 1, 'cool/modules/base/views/district.vue', NULL, 'icon-district', NULL, '/sys/district/list', 0, 'base:sys:district:list');
INSERT INTO `pinecms_menu` VALUES (116, 'pprof', NULL, 14, 5, '1', 1, 'cool/modules/base/views/pprof.vue', NULL, 'icon-status', NULL, '/sys/pprof', 0, 'base:sys:pprof');
INSERT INTO `pinecms_menu` VALUES (117, '脚本管理', NULL, 108, 0, '1', 1, 'cool/modules/task/views/script.vue', NULL, 'icon-script', NULL, '/plugin/task/script', 0, 'base:sys:task:script:list');
INSERT INTO `pinecms_menu` VALUES (118, '微信菜单', NULL, 58, 3, '1', 1, 'cool/modules/wechat/views/menu.vue', NULL, 'icon-wechat-menu', NULL, '/sys/wechat/menu', 0, 'base:sys:wechat:menu');
INSERT INTO `pinecms_menu` VALUES (119, '回复规则', NULL, 58, 4, '1', 1, 'cool/modules/wechat/views/msg-reply-rule.vue', NULL, 'icon-wechat-rule', NULL, '/sys/wechat/msgReplyRule', 0, 'base:sys:wechat:msg:reply:list');
INSERT INTO `pinecms_menu` VALUES (120, '微信信息', NULL, 58, 5, '1', 1, 'cool/modules/wechat/views/msg.vue', NULL, 'icon-wechat-msg-log', NULL, '/sys/wechat/msg', 0, 'base:sys:wechat:msg:list');
INSERT INTO `pinecms_menu` VALUES (121, '带参二维码', NULL, 58, 6, '1', 1, 'cool/modules/wechat/views/qrcode.vue', NULL, 'icon-wechat-qrcode', NULL, '/sys/wechat/qrcode', 0, 'base:sys:wechat:qrcode:list');
INSERT INTO `pinecms_menu` VALUES (122, '素材管理', NULL, 58, 7, '1', 1, 'cool/modules/wechat/views/material.vue', NULL, 'icon-wechat-material', NULL, '/sys/wechat/material', 0, 'base:sys:wechat:material:list');
INSERT INTO `pinecms_menu` VALUES (123, '消息模板', NULL, 58, 8, '1', 1, 'cool/modules/wechat/views/template.vue', NULL, 'icon-wechat-temp-msg', NULL, '/sys/wechat/template', 0, 'base:sys:wechat:msg:template:list');
INSERT INTO `pinecms_menu` VALUES (124, '系统文档', NULL, 0, 9999, '1', 1, 'http://doc.xiusin.cn', NULL, 'icon-doc', NULL, '/sys/document', 0, 'base:sys:doc');
INSERT INTO `pinecms_menu` VALUES (125, '接口文档', NULL, 0, 10000, '1', 1, 'http://8.140.114.57:7000/opsli-boot/doc.html#/opsli%202.X/%E4%BB%A3%E7%A0%81%E7%94%9F%E6%88%90%E5%99%A8-%E6%97%A5%E5%BF%97/createUsingGET', '', 'icon-doc', NULL, '', 0, 'base:sys:api');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_page
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_page`;
CREATE TABLE `pinecms_page` (
  `id` int NOT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `keywords` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `updatetime` int NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_page
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_page` VALUES (5, '关于项目', '', '', '<p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　1、有时候，我们等的不是什么人、什么事，我们等的是<a href=\"http://www.duwenzhang.com/huati/shijian/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">时间</a>，等时间，让自己改变。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　2、如果<a href=\"http://www.duwenzhang.com/wenzhang/shenghuosuibi/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">生活</a>中，有一个人想把你踩在脚下，不要以为生活错待了你。或许，还有十个人想要把你踩在脚下，只是你的强大，让他们没有机会伸出脚来。不要抱怨这个世界弱肉强食，你逐渐会发现，它看起来很残酷，却十分公正。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　3、人总是各有苦衷和不甘平庸。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　4、躺在床上听着歌，忘了疼痛，忘了所有的一切。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　5、夜来皓月才当午，重帘悄悄无人语。深处麝烟长，卧时留薄妆。当年还自惜，往事那堪忆。花落月明残，锦衾知晓寒。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　6、不要等到黑夜降临才注意到星星的光芒，其实它们一直在那儿。不要等到<a href=\"http://www.duwenzhang.com/huati/gudu/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">孤独</a>才想起真正对你好的人，其实她们一直在那儿。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　7、我们如此用心追随，<a href=\"http://www.duwenzhang.com/huati/qipan/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">期盼</a>我们的盛世安宁。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　8、如果发短信息给一个人，他一直不回，不要再发了，没有这么卑微的<a href=\"http://www.duwenzhang.com/huati/dengdai/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">等待</a>；如果一个人开始怠慢你，请你<a href=\"http://www.duwenzhang.com/huati/likai/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">离开</a>他，不懂得<a href=\"http://www.duwenzhang.com/huati/zhenxi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">珍惜</a>你的人不要为之不舍。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　9、漫长的别离里，我只想做一件事：专职爱你。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　10、走了那么多弯路，终于回到了最想来的地方。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　11、如果宁缺毋滥的结果将是孤独终老 是否你还能从一而终。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　12、我<a href=\"http://www.duwenzhang.com/huati/nuli/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">努力</a>坚持不<a href=\"http://www.duwenzhang.com/huati/fangqi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">放弃</a>，把委屈通通都咽下去。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　13、你在我的<a href=\"http://www.duwenzhang.com/huati/huiyi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">回忆</a>里<a href=\"http://www.duwenzhang.com/huati/meihao/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">美好</a>的不像话，你在我的念想里灿烂的一塌糊涂。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　14、不要把自己的事情掏心掏肺地告诉别人，你知不知道，有些人，面前心连心，背后动脑筋。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　15、悬崖那么深，我终究是为了你跳了下去。</p>', 1583571992, NULL, NULL);
INSERT INTO `pinecms_page` VALUES (21, '关于我们', 'aboutus', 'aboutus', '<h2 style=\"margin: 0px; padding: 20px 0px; font-size: 20px; font-weight: normal; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; white-space: normal; background-color: rgb(255, 255, 255);\">[ 关于 ]</h2><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">IT之家是一个提供IT业界和数码产品资讯的泛科技媒体平台。我们快速精选泛科技新闻，分享即时的IT业界动态和紧跟潮流的数码产品资讯，提供给力的PC和手机技术文章、丰富的系统应用美化资源，以及享不尽的智能阅读。</p><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">IT之家还下辖Win7之家（win7china.com）、Win8之家（win8china.com）、IT之家社区（bbs.ithome.com）等科技资讯站点。</p><h2 style=\"margin: 0px; padding: 20px 0px; font-size: 20px; font-weight: normal; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; white-space: normal; background-color: rgb(255, 255, 255);\">[ 初衷 ]</h2><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">2007年，有了Vista，于是，便有了Vista之家。当我们沉醉于Vista的华丽时，恍惚间便过去了两年。顾盼左右，原来，有些孤单。</p><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">当我们沿着Windows 7的足迹继续前行，发现世界已经在变。iPhone、iPad、Android、Chrome OS先后入侵了我们的生活，手机、平板、电脑、智能电视原来一个都不能少。</p><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">漫漫长夜，禁不住的思索我们为什么而存在，拿什么来款待招呼每天固定来访并增加中的70多万朋友，所以，便有了它，IThome.com。IT之家，欢迎你。</p><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">有人的地方，就有江湖；有电脑手机平板的地方，就有IT之家。IT人的家园，IT界的烟火。</p><h2 style=\"margin: 0px; padding: 20px 0px; font-size: 20px; font-weight: normal; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; white-space: normal; background-color: rgb(255, 255, 255);\">[ 数据 ]</h2><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">每天30+万独立IP，每月3200+万浏览量用户平均浏览页面4.55个，平均访问时长11分钟新浪微博粉丝：65万+腾讯微博粉丝：35万+网易云阅读、鲜果、ZAKER、Flipboard、搜狐新闻客户端等平台覆盖：总计120万+通过魔方电脑大师、手机客户端等推送每天可覆盖人群：90万+</p><h2 style=\"margin: 0px; padding: 20px 0px; font-size: 20px; font-weight: normal; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; white-space: normal; background-color: rgb(255, 255, 255);\">[ 渠道 ]</h2><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">IT之家官网IT之家客户端（iOS、Android、WP、Win8）IT之家手机版（WAP版）IT之家新浪（腾讯）微博IT之家微信公共平台IT之家第三方订阅平台</p><h2 style=\"margin: 0px; padding: 20px 0px; font-size: 20px; font-weight: normal; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; white-space: normal; background-color: rgb(255, 255, 255);\">[ 用户 ]</h2><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">IT之家的读者集中在年轻新潮、有活力爱分享、善思考能创造、崇尚探索创新、敢于挑战关注未来的群体。男性占比89%年龄20 ~ 27岁占比76%典型读者：学生、上班族、IT从业者、科技爱好者、数码爱好者。</p><h2 style=\"margin: 0px; padding: 20px 0px; font-size: 20px; font-weight: normal; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; white-space: normal; background-color: rgb(255, 255, 255);\">[ 联系 ]</h2><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">联系地址：青岛市市北区银川西路67号动漫产业园E-217（邮编：266071）媒介联系：<a href=\"mailto:pr@ruanmei.com\" style=\"color: rgb(39, 42, 48); text-decoration-line: none; outline: none;\">pr@ruanmei.com</a>市场合作：<a href=\"mailto:xue@ruanmei.com\" style=\"color: rgb(39, 42, 48); text-decoration-line: none; outline: none;\">xue@ruanmei.com</a>内容投稿：<a href=\"mailto:tougao@ruanmei.com\" style=\"color: rgb(39, 42, 48); text-decoration-line: none; outline: none;\">tougao@ruanmei.com</a>广告垂询：0532-83662200其他业务：联系我们</p><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">IT之家同时还欢迎下列合作：</p><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">1、媒体合作、媒体联系；2、为门户、垂直媒体等网站供稿；3、内容投稿，并成为IT之家独立撰稿人或编辑。</p><p style=\"margin-top: 10px; margin-bottom: 10px; padding: 0px; line-height: 26px; color: rgb(39, 42, 48); font-family: &quot;Microsoft Yahei&quot;; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">请致电：0532-83662200，或联系QQ：884358</p><p><br></p>', 1586775045, NULL, '2021-09-12 11:53:37');
INSERT INTO `pinecms_page` VALUES (20, '联系我们', '联系我们', '联系我们', '<p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\"><strong>软件产品客服：</strong></p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\">QQ：884358<br/>邮箱：<a href=\"mailto:xue@ruanmei.com\" style=\"background-color: rgb(255, 255, 255); text-decoration-line: none; color: rgb(102, 102, 102); outline: none;\">8</a>84358@qq.com</p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\"><strong>广告 / 市场 / 媒介合作：</strong></p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\">QQ：884358<br/>电话：0532-83662200<br/>邮箱：<a href=\"mailto:xue@ruanmei.com\" style=\"text-decoration-line: none; color: rgb(102, 102, 102); outline: none;\">8</a>84358@qq.com</p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\"><strong>友情链接交换：</strong></p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\">QQ：884358&nbsp;<span style=\"color: rgb(119, 119, 119); font-family: 瀹嬩綋, arial, verdana, sans-serif; line-height: 25px; font-size: 16px;\">&nbsp;</span><a href=\"http://wpa.qq.com/msgrd?v=3&uin=884358&site=dedemao&menu=yes\" target=\"_blank\" style=\"color: rgb(26, 101, 182); text-decoration-line: none; font-family: 瀹嬩綋, arial, verdana, sans-serif; line-height: 25px;\"><img alt=\"点击这里给我发消息\" src=\"http://wpa.qq.com/pa?p=2:884358:41\" title=\"点击这里给我发消息\"/></a><br/><br/></p><p><br/></p>', 1586775083, NULL, NULL);
INSERT INTO `pinecms_page` VALUES (19, '加入我们', 'join', 'join', '<p><span style=\"display: block; width: 960px; padding: 10px 0px; font-size: 18px; color: rgb(0, 68, 136); border-bottom: 1px dotted rgb(188, 188, 188);\">C++软件工程师</span></p><p style=\"margin-top: 30px; margin-bottom: 20px; padding: 0px; line-height: 26px;\">1、精通C++进行Windows平台下程序开发；精通MFC（熟悉WTL者优先）；<br/>2、有至少1年以上C++项目开发经验；<br/>3、计算机或相关专业本科以上学历；<br/>4、有共享软件免费软件开发经验者尤佳。</p><p><span style=\"display: block; width: 960px; padding: 10px 0px; font-size: 18px; color: rgb(0, 68, 136); border-bottom: 1px dotted rgb(188, 188, 188);\">C#软件开发工程师（.net）</span></p><p style=\"margin-top: 30px; margin-bottom: 20px; padding: 0px; line-height: 26px;\">1、精通 C# .NET 进行Windows、Windows Phone平台下程序开发；<br/>2、有至少1年以上C#项目开发经验；<br/>3、计算机或相关专业本科以上学历；<br/>4、有共享软件免费软件开发经验者尤佳。</p><p><span style=\"display: block; width: 960px; padding: 10px 0px; font-size: 18px; color: rgb(0, 68, 136); border-bottom: 1px dotted rgb(188, 188, 188);\">Android安卓应用开发工程师</span></p><p style=\"margin-top: 30px; margin-bottom: 20px; padding: 0px; line-height: 26px;\">1、精通安卓平台下应用/游戏开发；<br/>2、有至少1年以上安卓平台Java项目开发经验；<br/>3、计算机或相关专业本科以上学历；</p><p><span style=\"display: block; width: 960px; padding: 10px 0px; font-size: 18px; color: rgb(0, 68, 136); border-bottom: 1px dotted rgb(188, 188, 188);\">高级UI设计师</span></p><p style=\"margin-top: 30px; margin-bottom: 20px; padding: 0px; line-height: 26px;\">1、PC客户端产品的界面视觉和交互设计；<br/>2、移动客户端产品（iOS、Android、WP）的界面视觉和交互设计；<br/>3、网站页面视觉设计及前端代码编写。<br/><br/>职位要求：<br/>1、一年以上相关从业经验；<br/>2、本科及以上学历，美术、设计相关专业毕业，熟悉并热爱互联网行业；<br/>3、熟悉用户体验设计的工作原理及开发流程，能够归纳产品经理提出的交互设计需求，并快速构建信息结构和交互流程；<br/>4、熟练掌握色彩、图形、布局在设计作品中的运用，能够把握整体的设计风格，完成项目所需的界面和图标设计；<br/>5、精通HTML、CSS、JavaScript、jQuery等Web前端语言，能够编写符合规范的、兼容各类浏览器的静态页面。<br/><br/>优势加分：<br/>1、良好的手绘表达能力。<br/>2、具备将构思通过Flash或AfterEffect等软件转化为演示的能力。 3、对心理学有所涉猎。</p><p><span style=\"display: block; width: 960px; padding: 10px 0px; font-size: 18px; color: rgb(0, 68, 136); border-bottom: 1px dotted rgb(188, 188, 188);\">软件UI/网站前端设计师</span></p><p style=\"margin-top: 30px; margin-bottom: 20px; padding: 0px; line-height: 26px;\">1、PC客户端产品的界面视觉和交互设计；<br/>2、移动客户端产品（iOS、Android、WP）的界面视觉和交互设计；<br/>3、网站页面视觉设计及前端代码编写。<br/><br/>职位要求：<br/>1、本科及以上学历，美术、设计相关专业毕业，熟悉并热爱互联网行业；<br/>2、熟悉用户体验设计的工作原理及开发流程；<br/>3、熟练掌握色彩、图形、布局在设计作品中的运用，完成项目所需的界面和图标设计；<br/>4、精通HTML、CSS、JavaScript、jQuery等Web前端语言，能够编写静态页面。<br/><br/>优势加分：<br/>1、良好的手绘表达能力。<br/>2、具备将构思通过Flash或AfterEffect等软件转化为演示的能力。</p><p><span style=\"display: block; width: 960px; padding: 10px 0px; font-size: 18px; color: rgb(0, 68, 136); border-bottom: 1px dotted rgb(188, 188, 188);\">php软件工程师</span></p><p style=\"margin-top: 30px; margin-bottom: 20px; padding: 0px; line-height: 26px;\">1、熟悉linux，熟悉nginx/apache之一的配置和优化；<br/>2、熟悉mysql数据库的开发与使用；<br/>3、本科以上学历，至少有2年以上php开发经验；<br/>4、有linux环境下c/c++/java开发经验优先。</p><p><br/></p>', 1586775118, NULL, NULL);
INSERT INTO `pinecms_page` VALUES (18, '投稿反馈', '投稿反馈', '投稿反馈', '<p><strong style=\"font-size: 14px; line-height: 26px;\">投稿：</strong></p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\">QQ：884358<br/>邮箱：<a href=\"mailto:xue@ruanmei.com\" style=\"background-color: rgb(255, 255, 255); text-decoration-line: none; color: rgb(102, 102, 102); outline: none;\">8</a>84358@qq.com</p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\"><strong>意见反馈：</strong></p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\">QQ：884358<br/>电话：0532-83662200<br/>邮箱：<a href=\"mailto:xue@ruanmei.com\" style=\"text-decoration-line: none; color: rgb(102, 102, 102); outline: none;\">8</a>84358@qq.com</p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\"><strong>投稿小提示：</strong></p><p style=\"white-space: normal; margin-top: 30px; margin-bottom: 20px; padding: 0px; font-size: 14px; line-height: 26px; color: rgb(102, 102, 102); font-family: 微软雅黑, Arial, Helvetica, sans-serif; background-color: rgb(253, 253, 253);\">1. 投稿内容可以为资讯、技巧、主题壁纸类或资源下载<br/>2. 欢迎原创作品，转载文章请您尽量注明文章来源出处<br/>3. 您填写的名称、网址将会显示，会有很好的宣传效果<br/>4. 本站编辑可能会对您的来稿进行适度编辑以适应显示<br/>5. 您的来稿会有个编辑期，我们会尽量第一时间审核&nbsp;<span style=\"color: rgb(119, 119, 119); font-family: 瀹嬩綋, arial, verdana, sans-serif; line-height: 25px; font-size: 16px;\">&nbsp;</span><br/><br/></p><p><br/></p>', 1586775142, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_plugin
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_plugin`;
CREATE TABLE `pinecms_plugin` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '插件名称',
  `author` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '作者',
  `contact` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '联系方式',
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '功能描述',
  `version` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '版本号',
  `enable` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '是否启用 0：否 1：是',
  `status` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '状态 0:缺少配置 1:可用 2: 配置错误 3:未知错误',
  `view` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '页面信息',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '插件本地路径',
  `sign` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '标志',
  `prefix` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '插件访问前缀',
  `config` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `UQE_pinecms_plugin_path` (`path`) USING BTREE,
  UNIQUE KEY `UQE_pinecms_plugin_sign` (`sign`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_plugin
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_plugin` VALUES (5, '任务管理', 'xiusin', '18818818818', '实现任务管理功能', 'dev 0.0.1', '0', '0', '[{\"component\":{\"attrs\":{\"placeholder\":\"阿里云accessKeyId\"},\"name\":\"el-input\"},\"label\":\"accessKeyId\",\"prop\":\"accessKeyId\",\"props\":{\"label-width\":\"130px\"},\"rules\":{\"message\":\"值不能为空\",\"required\":true}},{\"component\":{\"attrs\":{\"placeholder\":\"阿里云accessKeySecret\"},\"name\":\"el-input\"},\"label\":\"accessKeySecret\",\"prop\":\"accessKeySecret\",\"props\":{\"label-width\":\"130px\"},\"rules\":{\"message\":\"值不能为空\",\"required\":true}},{\"component\":{\"attrs\":{\"placeholder\":\"阿里云oss的bucket\"},\"name\":\"el-input\"},\"label\":\"bucket\",\"prop\":\"bucket\",\"props\":{\"label-width\":\"130px\"},\"rules\":{\"message\":\"值不能为空\",\"required\":true}},{\"component\":{\"attrs\":{\"placeholder\":\"阿里云oss的endpoint\"},\"name\":\"el-input\"},\"label\":\"endpoint\",\"prop\":\"endpoint\",\"props\":{\"label-width\":\"130px\"},\"rules\":{\"message\":\"值不能为空\",\"required\":true}},{\"component\":{\"attrs\":{\"placeholder\":\"阿里云oss的timeout\"},\"name\":\"el-input\"},\"label\":\"timeout\",\"prop\":\"timeout\",\"props\":{\"label-width\":\"130px\"},\"rules\":{\"message\":\"值不能为空\",\"required\":true},\"value\":\"3600s\"}]', '2021-07-22 14:09:35', '2021-07-23 15:25:57', 'plugins/task/task.so', '77975e7f-de8b-4f26-be90-38c24fcd7c7d', '', '{\"accessKeyId\":\"23124123123123\",\"accessKeySecret\":\"234234w123\",\"bucket\":\"234234qweqw\",\"endpoint\":\"234234234\",\"timeout\":\"3600s\"}');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_position
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_position`;
CREATE TABLE `pinecms_position` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `code` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `listorder` int DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `remark` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_position
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_position` VALUES (2, 'CEO', 'C0001', 0, 1, '总裁');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_request_log
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_request_log`;
CREATE TABLE `pinecms_request_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `params` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '参数',
  `uri` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '请求uri',
  `userid` int DEFAULT NULL COMMENT '操作用户ID',
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '操作用户名',
  `ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '操作IP',
  `time` datetime DEFAULT NULL COMMENT '操作时间',
  `method` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '请求方法',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_request_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pinecms_setting
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_setting`;
CREATE TABLE `pinecms_setting` (
  `id` int NOT NULL AUTO_INCREMENT,
  `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `group` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `default` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `form_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `editor` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `listorder` int DEFAULT NULL,
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `options` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_setting
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_setting` VALUES (1, 'EMAIL_USER', 'pinemater@163.com', '邮箱设置', '', '用户名', 'el-input', 96, '登录邮箱服务器账号', 'null');
INSERT INTO `pinecms_setting` VALUES (2, 'EMAIL_PWD', '', '邮箱设置', '', '密码', 'el-input', 95, '一般为授权码', 'null');
INSERT INTO `pinecms_setting` VALUES (3, 'EMAIL_SMTP', 'smtp.163.com', '邮箱设置', '', 'SMTP服务器', 'el-input', 100, 'SMTP服务器', 'null');
INSERT INTO `pinecms_setting` VALUES (4, 'EMAIL_EMAIL', 'pinemater@163.com', '邮箱设置', '', '邮箱地址', 'el-input', 98, '发件邮箱', 'null');
INSERT INTO `pinecms_setting` VALUES (5, 'SITE_ICP', '豫ICP备xxxxxxxx号', '前台设置', NULL, '备案号', 'el-input', 4, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (6, 'SITE_KEYWORDS', 'pine,pincms,gocms,cms,模板化框架', '前台设置', NULL, '关键字', 'el-input', 2, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (7, 'SITE_DESCRIPTION', 'pincms一个go语言的模板化CMS，支持类dedecms式的标签化调用。可以快速开发出企业网站，支持自定义文档模型', '前台设置', NULL, '描述', 'el-input', 3, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (8, 'SITE_TITLE', 'pincms一个go语言的模板化CMS', '前台设置', NULL, '站点标题', 'el-input', 1, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (10, 'SITE_OPEN', '关闭', '前台设置', '开启', '站点开启', '{\"type\":\"checkbox\", \"options\": {\"on\":\"开启\", \"off\":\"关闭\"}}', 0, '关闭后前端将显示为提示语', 'null');
INSERT INTO `pinecms_setting` VALUES (11, 'EMAIL_PORT', '465', '邮箱设置', '25', '端口', 'el-input', 99, '端口', 'null');
INSERT INTO `pinecms_setting` VALUES (14, 'SITE_PAGE_SIZE', '15', '前台设置', '25', '列表默认分页数', 'el-input', 5, '系统默认参数', 'null');
INSERT INTO `pinecms_setting` VALUES (16, 'UPLOAD_DIR', 'resources/assets/uploads', '存储配置', 'resources/assets/upload', '存储目录', 'el-input', 21, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (17, 'UPLOAD_ENGINE', 'oss存储', '存储配置', '本地存储', '存储引擎', 'el-select', 20, '存储类型，取值为： 本地存储，oss存储', 'null');
INSERT INTO `pinecms_setting` VALUES (18, 'UPLOAD_IMG_TYPES', 'jpg|jpeg|png|gif|bmp', '存储配置', 'jpg,jpeg,png,gif,bmp', '可上传图片类型', 'el-input', 22, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (19, 'UPLOAD_URL_PREFIX', '/uploads', '存储配置', 'upload', '地址前缀', 'el-input', 23, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (20, 'UPLOAD_DATABASE_PASS', '123456', '存储配置', '', '备份数据库密码', 'el-input', 24, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (21, 'OSS_HOST', 'http://pinecms.oss-cn-beijing.aliyuncs.com', 'OSS存储配置', '', 'HOST', 'el-input', 19, 'oss域名', 'null');
INSERT INTO `pinecms_setting` VALUES (22, 'OSS_ENDPOINT', '', 'OSS存储配置', NULL, 'ENDPOINT', 'el-input', 20, 'oss上传endpoint', NULL);
INSERT INTO `pinecms_setting` VALUES (23, 'OSS_KEYID', '', 'OSS存储配置', '', 'KEYID', 'el-input', 17, '', 'null');
INSERT INTO `pinecms_setting` VALUES (24, 'OSS_BUCKET', '', 'OSS存储配置', '', 'BUCKET', 'el-input', 18, 'oss桶名', 'null');
INSERT INTO `pinecms_setting` VALUES (25, 'OSS_KEYSECRET', '', 'OSS存储配置', '', 'SECRET', 'el-input', 16, '', 'null');
INSERT INTO `pinecms_setting` VALUES (26, 'SITE_STATIC_PAGE_DIR', 'resources/html', '前台设置', NULL, '静态页面地址', 'el-input', 6, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (27, 'SITE_COPYRIGHT', '版权归PineCMS所有', '前台设置', '', '版权信息', 'el-input', 7, '设置版权信息111', 'null');
INSERT INTO `pinecms_setting` VALUES (28, 'UPLOAD_ACCT_TYPES', 'zip|gz|rar|iso|doc|xsl|ppt|wps', '存储配置', 'zip|gz|rar|iso|doc|xsl|ppt|wps', '附件上传类型', 'el-input', 22, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (29, 'UPLOAD_MAX_SIZE', '20', '存储配置', '2', '上传大小(MB)', 'el-input', 25, '', 'null');
INSERT INTO `pinecms_setting` VALUES (30, 'EMAIL_ADMIN_EMAIL', 'xiusin.chen@gmail.com', '邮箱设置', '', '管理员邮箱', 'el-input', 94, '接收通知的邮箱，可用于系统通知', 'null');
INSERT INTO `pinecms_setting` VALUES (31, 'EMAIL_SEND_NAME', 'pinecms系统', '邮箱设置', '', '发件人名称', 'el-input', 97, '发送人信息', 'null');
INSERT INTO `pinecms_setting` VALUES (32, 'SITE_NAME', 'PINECMS', '前台设置', NULL, '站点名称', 'el-input', NULL, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (33, 'SITE_DEBUG', '开启', '前台设置', '开启', '动态渲染', '{\"type\":\"checkbox\", \"options\": {\"on\":\"开启\", \"off\":\"关闭\"}}', 0, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (34, 'BACKEND_PATH', '管理页面地址', '前台设置', '/backend/login', '管理地址', 'el-input', NULL, NULL, NULL);
INSERT INTO `pinecms_setting` VALUES (35, 'SITE_URL', 'http://localhost:2019', '前台设置', 'http://localhost:2019', '网站域名', 'el-input', NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_tags
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_tags`;
CREATE TABLE `pinecms_tags` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `ref_num` int DEFAULT NULL,
  `listorder` int DEFAULT NULL,
  `status` int DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `clicks` int DEFAULT NULL,
  `seo_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `seo_keywords` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `seo_description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_tags
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_tags` VALUES (2, '爆品', 0, 30, 1, NULL, NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_task_info
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_task_info`;
CREATE TABLE `pinecms_task_info` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `entity_id` int DEFAULT NULL COMMENT '任务ID',
  `repeat_conf` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '任务配置',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '任务名称',
  `cron` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'cron配置',
  `limit` int DEFAULT NULL COMMENT '最大执行次数 不传为无限次',
  `remark` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '备注',
  `status` int DEFAULT NULL COMMENT '状态 0:停止 1：运行',
  `start_date` datetime DEFAULT NULL COMMENT '开始时间',
  `end_date` datetime DEFAULT NULL COMMENT '结束时间',
  `data` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '数据',
  `service` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '执行的service实例ID',
  `type` int DEFAULT NULL COMMENT '类型 0:系统 1：用户',
  `next_run_time` datetime DEFAULT NULL COMMENT '下一次执行时间',
  `task_type` int DEFAULT NULL COMMENT '状态 0:cron 1：时间间隔',
  `error` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `every` int DEFAULT NULL COMMENT '间隔执行时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_task_info
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_task_info` VALUES (1, 0, '', '测试任务', '1 * * * * ? ', 0, '2183', 1, '2021-07-01 00:00:00', '2021-09-30 00:00:00', '', 'echo', 0, '2021-09-25 13:43:01', 0, '', NULL, 0);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_task_log
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_task_log`;
CREATE TABLE `pinecms_task_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `task_id` int DEFAULT NULL COMMENT '任务ID',
  `status` int DEFAULT NULL COMMENT '状态 0:失败 1：成功',
  `detail` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '详情',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `exec_time` int DEFAULT '0' COMMENT '执行时长',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_task_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pinecms_todo
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_todo`;
CREATE TABLE `pinecms_todo` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '1' COMMENT '字符串多选:1=外部链接,2=内部链接,3=通用链接',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '普通输入框',
  `introduce` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '普通多行输入框::cms-textarea',
  `listorder` int NOT NULL COMMENT '不可为空数字',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT 'tinyint单选:0=待审核,1=通过,2=拒绝:cms-radio',
  `put_date` date DEFAULT NULL COMMENT '日期',
  `put_datetime` datetime DEFAULT NULL COMMENT '时间日期',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间$end=end_time',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间被引用隐藏到代码区间选择器',
  `logo` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '单图上传',
  `logos` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '多图上传',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_todo
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pinecms_wechat_account
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_wechat_account`;
CREATE TABLE `pinecms_wechat_account` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `app_id` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'appid',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '公众号名称',
  `type` tinyint(1) DEFAULT NULL COMMENT '账号类型',
  `verified` tinyint(1) DEFAULT NULL COMMENT '认证状态',
  `secret` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `token` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `aes_key` varchar(43) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_wechat_account
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_wechat_account` VALUES (1, '2wxa3a66933a15b49ff', 'xiusin博客', 2, 0, '', '2v2fu051oaqdehtwdkbtr2ds4v2yc5qo', 'rSuIpBuw1OrYUNp4t0FWSZji6fMkrgWN8oET4FyUub4');
INSERT INTO `pinecms_wechat_account` VALUES (2, 'wxe43df03110f5981b1', '测试号', 1, 1, '', 'wtffjrorhohj0giffn53fdbelnbnoiq0', '');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_wechat_log
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_wechat_log`;
CREATE TABLE `pinecms_wechat_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `app_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `open_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `inout` tinyint(1) DEFAULT NULL,
  `msg_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `created_at` datetime DEFAULT NULL,
  `fans_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_wechat_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pinecms_wechat_material
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_wechat_material`;
CREATE TABLE `pinecms_wechat_material` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '媒体素材类型',
  `media_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '媒体ID',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `appid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'appid',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_wechat_material
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_wechat_material` VALUES (1, 'image', 'X8twhHsiiGmDV0b6iuJn1Ca4zJSJg6Y1TVyjhLeAYlc', 'https://mmbiz.qpic.cn/mmbiz_png/uicrGD717HA0TmJWH7Aiah5ZhTwyKJcKc9pib2lByeZ9Yug9iaDRak0GKxmEFybx1YnfYBibBymZynVTJmtcBoWYMdg/0?wx_fmt=png', '2021-08-23 15:12:03', 'wxe43df03110f5981b');
INSERT INTO `pinecms_wechat_material` VALUES (2, 'image', 'X8twhHsiiGmDV0b6iuJn1OELEj58TJxdOjlbuFwlxVI', 'https://mmbiz.qpic.cn/mmbiz_png/uicrGD717HA0TmJWH7Aiah5ZhTwyKJcKc9pib2lByeZ9Yug9iaDRak0GKxmEFybx1YnfYBibBymZynVTJmtcBoWYMdg/0?wx_fmt=png', '2021-08-23 15:11:46', 'wxe43df03110f5981b');
INSERT INTO `pinecms_wechat_material` VALUES (4, 'image', 'X8twhHsiiGmDV0b6iuJn1MTNh7uQIcNqTTEMruwNIkM', 'http://mmbiz.qpic.cn/mmbiz_jpg/uicrGD717HA1O5icwpWjOTOicJZaUoMKXFd9z7RDmKB5IKPWatMXBUh3pnU9mEcwfN2TNNRM5LHC29iaiayWvfKSKibQ/0?wx_fmt=jpeg', '2021-09-04 16:00:11', 'wxe43df03110f5981b');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_wechat_member
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_wechat_member`;
CREATE TABLE `pinecms_wechat_member` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `appid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sex` int DEFAULT NULL,
  `city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `province` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `headimgurl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `subscribe_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `subscribe` tinyint(1) DEFAULT NULL,
  `unionid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tagid_list` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `subscribe_scene` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `qr_scene_str` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `poster` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_wechat_member
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_wechat_member` VALUES (1, 'wxe43df03110f5981b', 'op-oIuBtwRFDZyIcceBy1H8wHwD4', '14414414414', '万峰', 1, '郑州', '河南', 'http://thirdwx.qlogo.cn/mmopen/98Nz5LFElxykASYo1gv8ZYzj9dfh3ebROTekIMHzIlezqyV3emEkDGLF7MEk8NGvgo1OPV9DOiadn3XBRLk8yfw/132', '2016-11-29 10:11:50', 1, '', '我是一头小毛驴', '[2]', 'ADD_SCENE_OTHERS', '', '');
INSERT INTO `pinecms_wechat_member` VALUES (2, 'wxe43df03110f5981b', 'op-oIuDMUtnZoipCW6PwOK7GSF5U', '', '刘俊', 1, '朝阳', '北京', 'http://thirdwx.qlogo.cn/mmopen/ajNVdqHZLLBW2ZYFSMhakTxs0JCBOp14DtCOYGqCor8iaGwWuoHvQ0ysvrvgQ2DveAFZUISnZf20UQpiaHSWY5Gw/132', '2016-12-05 09:54:07', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (3, 'wxe43df03110f5981b', 'op-oIuK-5cgb9iJFeL5nMMx5RgGY', '', '6', 2, '', '', 'http://thirdwx.qlogo.cn/mmopen/PiajxSqBRaEIKEZIOwAb3NeoCibuxS4WDUTnVy5IAkb07gpGsctm1Sun0qJRpyEkTFHnumhhgy9y92NhufN8oAUw/132', '2016-06-20 11:17:33', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (4, 'wxe43df03110f5981b', 'op-oIuPNRpZidfcZUNzhYCuZOM_k', '', 'Lisa', 0, '', '', 'http://thirdwx.qlogo.cn/mmopen/PiajxSqBRaEIzhBo8BPGyAcs5U0tia9xE1kpMuHzXsicuAkYaHibSPyL777YzRgyO6auYibqpoyqPaLn1zFMJf9nSaw/132', '2016-12-01 11:23:05', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (5, 'wxe43df03110f5981b', 'op-oIuPBWfyiyEa1c9PYIvZ1b8sg', '', '朱丽娟🌹', 2, '', '', 'http://thirdwx.qlogo.cn/mmopen/PiajxSqBRaEIvJ3VtK827ficpIicwBDAedEAQ6nRFICiaXoEmOarv12nIfh06ibIfBwakjyvWaOs3c7oFQbicVCQZvSw/132', '2016-12-05 10:53:16', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (6, 'wxe43df03110f5981b', 'op-oIuCz9dDBUfoL22BYwwdTMKXg', '', 'Why so serious', 1, '', '', 'http://thirdwx.qlogo.cn/mmopen/iavpLF1TxNoE5Bt9V12MGE6U7n6S0LOoCopoD6Ly6CEDvxtUYmtgribOK9EgxOkrgTrtqp6TwKiaH2d4FMQ1syAQQ/132', '2016-12-06 17:22:41', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (7, 'wxe43df03110f5981b', 'op-oIuElpNfFSfK2gw-funK3c9YU', '', '小米', 2, '郑州', '河南', 'http://thirdwx.qlogo.cn/mmopen/RVE8icjnX972lHaX3x9k3WIIaGRicwhOjY8Sy1wH72D0UBL48YS1jnUalSxYZkV2a9BbqenkvPtg8BSIl8djhgr5oC4KohP95l/132', '2016-09-23 18:10:17', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (8, 'wxe43df03110f5981b', 'op-oIuAF8JVmgedOIt6GdiQbaQ78', '', '乐', 1, '昌平', '北京', 'http://thirdwx.qlogo.cn/mmopen/RVE8icjnX972lHaX3x9k3WMQxt0HEHicdOYSHtSBsEhySOAicLxb6Sx0JjXUG6V6LxELQoA9rWOYU1bIU32Jj3nUVCftGypPHor/132', '2016-09-21 19:36:07', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (9, 'wxe43df03110f5981b', 'op-oIuMLJSaz33pr4dF-MV8PL5-A', '', '中原银行何龙翔160207', 1, '周口', '河南', 'http://thirdwx.qlogo.cn/mmopen/iavpLF1TxNoFGf6got8rFZr7jR1ibJ0U1uQfbU1I8yYG0ndr0joScpwib5Ny7NMpfg8cibWCG1XS6Bx047mektlWyaSusHp2B91ib/132', '2016-11-29 17:49:10', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (10, 'wxe43df03110f5981b', 'op-oIuGFLxGdEDIYDWJ9jkhTOsWE', '', '高同', 2, '海淀', '北京', 'http://thirdwx.qlogo.cn/mmopen/iavpLF1TxNoFGf6got8rFZoQJB5Xp2hOBicsUDvfdcqKar2jvFL1wAuISf78YsrJRVkgaCk1lRichpPKvb1mD3QECfVRicbNXNTp/132', '2016-11-29 16:50:49', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (11, 'wxe43df03110f5981b', 'op-oIuLy5Wx_YwNqZCK2otwNgjPg', '', '杨勇', 1, '广州', '广东', 'http://thirdwx.qlogo.cn/mmopen/iavpLF1TxNoFGf6got8rFZiarcL4yAa1nfjFQoXsEHlJHJm3iart2l4ekFK9ib3hehce0Pzodrgf8icRx47t8K3BtwNOIUkY1cQQW/132', '2016-09-22 13:10:40', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (12, 'wxe43df03110f5981b', 'op-oIuMcIjmjGaa68gG0DtwG8EME', '', '伟伟', 1, '石家庄', '河北', 'http://thirdwx.qlogo.cn/mmopen/98Nz5LFElxw3kRxPhFUenbxxybLaL7lln6xaGic2j3G0to5G3dtoaQ0dTNMO6aTsk69BK7Pm3H4EebKx7Q7T427wgz7U3gqbR/132', '2016-08-17 14:51:34', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (13, 'wxe43df03110f5981b', 'op-oIuE7Nrgf6bygx4w0LokpufX8', '', '亖分之3.二', 1, '', '', 'http://thirdwx.qlogo.cn/mmopen/hFQfVZf4UQufGpa3tFfacYnuqLlmBkYKEhEAgNMKRuZ9oBepfeRxyqwI4ZVkaReKHebzXVcmTmHro2TNQXT1qribEkQOfkC1B/132', '2016-11-29 09:38:08', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (14, 'wxe43df03110f5981b', 'op-oIuNfFkqV-9kCY1tDKrSfl3XQ', '', '刘佳琳lucky', 2, '郑州', '河南', 'http://thirdwx.qlogo.cn/mmopen/iavpLF1TxNoFGf6got8rFZhFtHOpicfzuaRDd7OTiaL2ACFSvYpS197mbxhO5myMqmRnQV4Tz5v06A6q3Wjp9FnwWNHo2uwN6pj/132', '2016-11-29 17:37:05', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (15, 'wxe43df03110f5981b', 'op-oIuE3jxDglIvaCkVBl3dwviM0', '', '丹', 2, '郑州', '河南', 'http://thirdwx.qlogo.cn/mmopen/oPd6uxrERPfZePoe6dDRvHFtyefnOibnVZSdrIaialyUNxqcnUMelNibBaXickc7M6NFxxxPFg6mG1JEbicpkicaQg0Qbs6KZQ2JBb/132', '2016-11-29 23:27:30', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (16, 'wxe43df03110f5981b', 'op-oIuMwrSfbDaPgdbupicBBLxcU', '', '嗳·僦說妳妠！', 2, '', '', 'http://thirdwx.qlogo.cn/mmopen/98Nz5LFElxw3kRxPhFUenVsYBuJVVZUgfqeTPKo14VewGjdmFfYgJmsLANzFoX5JMUKTAvjSt7QqOgFxHiaM68eXl1JEByczP/132', '2016-12-05 09:57:01', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (17, 'wxe43df03110f5981b', 'op-oIuFs_vSCfoAp4zELMXa9KEXc', '', '放牛哥', 1, '郑州', '河南', 'http://thirdwx.qlogo.cn/mmopen/iavpLF1TxNoFGf6got8rFZnPPBQlDAX75L3sBwxIm59PibQmYCObGh4eQoz9VbuuibYao6eUMCE52eqYYLv4slibpnpjjCPibE3UT/132', '2016-04-18 18:55:36', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (18, 'wxe43df03110f5981b', 'op-oIuOJypsSJaWaytBEWTz7BGVw', '', '龙凤吉祥', 1, '濮阳', '河南', 'http://thirdwx.qlogo.cn/mmopen/9Qc2hqj33RElW5YGiaPEpk0sSic96QsjqyUNtttag47ViaUW579Iu7VWkrJ8F8zWkp6p9tgFMQUBumIhzicY5x6l5yzmobkMibInO/132', '2016-11-29 15:57:10', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (19, 'wxe43df03110f5981b', 'op-oIuGZBKZm1IhUTpEpEJSeWKUc', '', '鹤', 2, '安庆', '安徽', 'http://thirdwx.qlogo.cn/mmopen/ajNVdqHZLLA5oW7x6Dv95HEqiaJiatBAlmCFtIt1U8mia43GSdfD2WuwkBt1xj80xxMtGBNLgLtKvoHHAzAz9bqiaG8uG5MDiaB85PHmDY4GV7lI/132', '2016-11-29 09:51:49', 1, '', '', 'null', 'ADD_SCENE_OTHERS', '', NULL);
INSERT INTO `pinecms_wechat_member` VALUES (20, 'wxe43df03110f5981b', 'op-oIuJztLwLOoNBX6hNoOzHFEws', '', '修心', 1, '郑州', '河南', 'http://thirdwx.qlogo.cn/mmopen/iavpLF1TxNoFGf6got8rFZt0UNBRwicG3ACrolLf6ribajAVqdQiaYq48SI1UZLuJg2JXEsyf2vjciaBlDsMp35stqF7kEveQZsvB/132', '2020-10-26 15:37:36', 1, '', '', 'null', 'ADD_SCENE_QR_CODE', '', NULL);
COMMIT;

-- ----------------------------
-- Table structure for pinecms_wechat_msg_reply_rule
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_wechat_msg_reply_rule`;
CREATE TABLE `pinecms_wechat_msg_reply_rule` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `app_id` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `rule_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `match_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `exact_match` tinyint(1) DEFAULT NULL COMMENT '是否精确匹配',
  `created_at` datetime DEFAULT NULL,
  `reply_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `reply_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `effect_time_start` time DEFAULT NULL,
  `effect_time_end` time DEFAULT NULL,
  `priority` int DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `appid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_wechat_msg_reply_rule
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_wechat_msg_reply_rule` VALUES (3, 'wxe43df03110f5981b', 'hello', 'hello', 1, '2021-08-22 22:13:03', 'text', 'hello world, hahhaha', 1, '', '22:12:40', '22:12:40', 0, '2021-08-28 16:33:56', 'wxe43df03110f5981b');
INSERT INTO `pinecms_wechat_msg_reply_rule` VALUES (4, NULL, 'hello1', 'hello1', 0, '2021-08-28 16:49:30', 'text', 'aaa', 1, '', '16:49:00', '23:49:00', 0, '2021-08-28 16:49:30', 'wxa3a66933a15b49ff');
INSERT INTO `pinecms_wechat_msg_reply_rule` VALUES (5, NULL, '123123', '123123', 1, '2021-08-28 16:50:13', 'news', '123123', 1, '', '16:50:03', '16:50:03', 0, '2021-08-28 16:50:13', 'wxe43df03110f5981b');
INSERT INTO `pinecms_wechat_msg_reply_rule` VALUES (6, NULL, '123123', '123123', 1, '2021-08-28 16:52:18', 'voice', '', 1, '', '16:52:08', '00:00:00', 0, '2021-08-28 16:52:18', '');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_wechat_msg_template
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_wechat_msg_template`;
CREATE TABLE `pinecms_wechat_msg_template` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `appid` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'appid',
  `template_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '模板ID',
  `primary_industry` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `deputy_industry` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '模板名称',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `data` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `mini_program` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `status` tinyint(1) DEFAULT NULL COMMENT '是否有效0=无效',
  `example` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '模板示例',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `title` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '模板标题',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of pinecms_wechat_msg_template
-- ----------------------------
BEGIN;
INSERT INTO `pinecms_wechat_msg_template` VALUES (1, 'wxe43df03110f5981b', '97_xBnPYcKDeGjUaq9BxDvUQneDAM4-mAEd8Jt0VRu4', '', '', '成为会员通知', '{{first.DATA}}\n昵称：{{keyword1.DATA}}\n手机：{{keyword2.DATA}}\n成为会员赠送积分：{{keyword3.DATA}}\n会员级别：{{keyword4.DATA}}\n会员卡号：{{keyword5.DATA}}\n{{remark.DATA}}', '[{\"color\":\"#000000\",\"name\":\"first\",\"value\":\"1\"},{\"color\":\"#000000\",\"name\":\"keyword1\",\"value\":\"2\"},{\"color\":\"#000000\",\"name\":\"keyword2\",\"value\":\"3\"},{\"color\":\"#000000\",\"name\":\"keyword3\",\"value\":\"4\"},{\"color\":\"#000000\",\"name\":\"keyword4\",\"value\":\"5\"},{\"color\":\"#000000\",\"name\":\"keyword5\",\"value\":\"6\"},{\"color\":\"#000000\",\"name\":\"remark\",\"value\":\"7\"}]', '', '{\"appid\":\"\",\"pagePath\":\"\"}', 0, '', '2021-08-24 18:09:01', '2021-08-24 18:09:01', '成为会员通知');
COMMIT;

-- ----------------------------
-- Table structure for pinecms_wechat_qrcode
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_wechat_qrcode`;
CREATE TABLE `pinecms_wechat_qrcode` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `appid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_temp` tinyint(1) DEFAULT NULL COMMENT '是否为临时二维码',
  `scene_str` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `ticket` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `expire_time` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of pinecms_wechat_qrcode
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

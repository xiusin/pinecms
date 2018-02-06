/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50547
 Source Host           : localhost:3306
 Source Schema         : iriscms

 Target Server Type    : MySQL
 Target Server Version : 50547
 File Encoding         : 65001

 Date: 06/02/2018 16:11:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for iriscms_admin
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_admin`;
CREATE TABLE `iriscms_admin`  (
  `userid` mediumint(6) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `roleid` smallint(5) NULL DEFAULT 0,
  `encrypt` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `lastloginip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `lastlogintime` int(10) UNSIGNED NULL DEFAULT 0,
  `email` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `realname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`userid`) USING BTREE,
  INDEX `username`(`username`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_admin
-- ----------------------------
INSERT INTO `iriscms_admin` VALUES (1, 'admin', '0087af20a551a8b804f89469534b7859', 1, 'qmRlFL', '127.0.0.1', 1474291850, 'chenchengbin92@gmail.com', 'mirchen.com');

-- ----------------------------
-- Table structure for iriscms_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_admin_role`;
CREATE TABLE `iriscms_admin_role`  (
  `roleid` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT,
  `rolename` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `listorder` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `disabled` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`roleid`) USING BTREE,
  INDEX `listorder`(`listorder`) USING BTREE,
  INDEX `disabled`(`disabled`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_admin_role
-- ----------------------------
INSERT INTO `iriscms_admin_role` VALUES (1, '超级管理员', '超级管理员', 0, 0);

-- ----------------------------
-- Table structure for iriscms_admin_role_priv
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_admin_role_priv`;
CREATE TABLE `iriscms_admin_role_priv`  (
  `roleid` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  `c` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `a` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  INDEX `roleid`(`roleid`, `c`, `a`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色操作权限配置表' ROW_FORMAT = Fixed;

-- ----------------------------
-- Records of iriscms_admin_role_priv
-- ----------------------------
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'category', 'add');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'category', 'delete');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'category', 'edit');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'category', 'list');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'category', 'order');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'category', 'view');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'content', 'index');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'content', 'news-list');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'content', 'page');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'content', 'right');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'content', 'top');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'left');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'logdelete');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'loglist');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'logview');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'menuadd');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'menudelete');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'menuedit');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'menulist');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'menuorder');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'menuview');
INSERT INTO `iriscms_admin_role_priv` VALUES (8, 'system', 'top');

-- ----------------------------
-- Table structure for iriscms_category
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_category`;
CREATE TABLE `iriscms_category`  (
  `catid` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `parentid` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `catname` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `url` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `listorder` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `ismenu` tinyint(1) UNSIGNED NOT NULL DEFAULT 1,
  `tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '后台列表',
  `home_tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前台列表模板',
  `content_tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前台内容页模板',
  `thumb` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`catid`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 26 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '内容分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for iriscms_category_priv
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_category_priv`;
CREATE TABLE `iriscms_category_priv`  (
  `catid` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `roleid` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `is_admin` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `action` char(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  INDEX `catid`(`catid`, `roleid`, `is_admin`, `action`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Fixed;

-- ----------------------------
-- Table structure for iriscms_link
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_link`;
CREATE TABLE `iriscms_link`  (
  `linkid` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT,
  `linktype` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `logo` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `introduce` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `listorder` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `passed` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `addtime` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`linkid`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '友情链接表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for iriscms_log
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_log`;
CREATE TABLE `iriscms_log`  (
  `logid` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `controller` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `action` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `querystring` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `userid` mediumint(8) UNSIGNED NOT NULL DEFAULT 0,
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`logid`) USING BTREE,
  INDEX `module`(`controller`, `action`) USING BTREE,
  INDEX `username`(`username`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '操作日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for iriscms_menu
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_menu`;
CREATE TABLE `iriscms_menu`  (
  `id` smallint(6) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `parentid` smallint(6) NOT NULL DEFAULT 0,
  `c` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `a` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `data` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `is_system` tinyint(1) NOT NULL DEFAULT 0,
  `listorder` smallint(6) UNSIGNED NOT NULL DEFAULT 0,
  `display` enum('1','0') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `listorder`(`listorder`) USING BTREE,
  INDEX `parentid`(`parentid`) USING BTREE,
  INDEX `module`(`c`, `a`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 55 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限表' ROW_FORMAT = Fixed;

-- ----------------------------
-- Records of iriscms_menu
-- ----------------------------
INSERT INTO `iriscms_menu` VALUES (1, '我的面板', 0, 'admin', 'public-top', '', 1, 1, '1');
INSERT INTO `iriscms_menu` VALUES (2, '系统设置', 0, 'setting', 'top', '', 0, 6, '1');
INSERT INTO `iriscms_menu` VALUES (3, '内容管理', 0, 'content', 'top', '', 0, 2, '1');
INSERT INTO `iriscms_menu` VALUES (54, '分类单页', 35, 'content', 'page', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (5, '后台管理', 0, 'system', 'top', '', 0, 5, '1');
INSERT INTO `iriscms_menu` VALUES (6, '个人信息', 1, 'admin', 'public-left', '', 1, 0, '1');
INSERT INTO `iriscms_menu` VALUES (7, '修改密码', 6, 'admin', 'public-editpwd', '', 1, 1, '1');
INSERT INTO `iriscms_menu` VALUES (8, '修改个人信息', 6, 'admin', 'public-editinfo', '', 1, 0, '1');
INSERT INTO `iriscms_menu` VALUES (9, '系统设置', 2, 'setting', 'left', '', 0, 1, '1');
INSERT INTO `iriscms_menu` VALUES (10, '站点设置', 9, 'setting', 'site', '', 0, 1, '1');
INSERT INTO `iriscms_menu` VALUES (11, '管理员设置', 2, 'admin', 'left', '', 0, 2, '1');
INSERT INTO `iriscms_menu` VALUES (12, '管理员管理', 11, 'admin', 'memberlist', '', 0, 1, '1');
INSERT INTO `iriscms_menu` VALUES (13, '角色管理', 11, 'admin', 'rolelist', '', 0, 2, '1');
INSERT INTO `iriscms_menu` VALUES (14, '后台管理', 5, 'system', 'left', '', 0, 1, '1');
INSERT INTO `iriscms_menu` VALUES (15, '日志管理', 14, 'system', 'loglist', '', 0, 1, '1');
INSERT INTO `iriscms_menu` VALUES (16, '菜单管理', 14, 'system', 'menulist', '', 0, 2, '1');
INSERT INTO `iriscms_menu` VALUES (17, '查看菜单', 16, 'system', 'menuview', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (18, '添加菜单', 16, 'system', 'menuadd', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (19, '修改菜单', 16, 'system', 'menuedit', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (20, '删除菜单', 16, 'system', 'menudelete', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (21, '菜单排序', 16, 'system', 'menuorder', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (22, '查看日志', 15, 'system', 'logview', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (23, '删除日志', 15, 'system', 'logdelete', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (24, '查看管理员', 12, 'admin', 'member-view', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (25, '添加管理员', 12, 'admin', 'member-add', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (26, '编辑管理员', 12, 'admin', 'member-edit', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (27, '删除管理员', 12, 'admin', 'member-delete', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (28, '查看角色', 13, 'admin', 'rolelist', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (29, '添加角色', 13, 'admin', 'role-add', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (30, '编辑角色', 13, 'admin', 'role-edit', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (31, '删除角色', 13, 'admin', 'role-delete', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (53, '新闻列表', 35, 'content', 'news-list', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (33, '权限设置', 13, 'admin', 'role-permission', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (34, '发布管理', 3, 'content', 'right', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (35, '内容管理', 34, 'content', 'index', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (36, '栏目管理', 34, 'category', 'list', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (37, '查看栏目', 36, 'category', 'view', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (38, '添加栏目', 36, 'category', 'add', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (39, '编辑栏目', 36, 'category', 'edit', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (40, '删除栏目', 36, 'category', 'delete', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (41, '栏目排序', 36, 'category', 'order', '', 0, 0, '1');

-- ----------------------------
-- Table structure for iriscms_news
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_news`;
CREATE TABLE `iriscms_news`  (
  `id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `catid` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `title` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `thumb` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `keywords` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `description` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `listorder` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  `status` tinyint(2) UNSIGNED NOT NULL DEFAULT 1,
  `username` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `inputtime` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `updatetime` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '//模板名称',
  `recommend` tinyint(1) NOT NULL DEFAULT 0 COMMENT '推荐',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`, `listorder`, `id`) USING BTREE,
  INDEX `listorder`(`catid`, `status`, `listorder`, `id`) USING BTREE,
  INDEX `catid`(`catid`, `status`, `id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '新闻表后期根据模型扩展' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for iriscms_page
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_page`;
CREATE TABLE `iriscms_page`  (
  `catid` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `title` varchar(160) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `keywords` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `updatetime` int(10) UNSIGNED NOT NULL DEFAULT 0,
  INDEX `catid`(`catid`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '单页内容表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for iriscms_setting
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_setting`;
CREATE TABLE `iriscms_setting`  (
  `key` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`key`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_setting
-- ----------------------------
INSERT INTO `iriscms_setting` VALUES ('EMAIL_USER', '');
INSERT INTO `iriscms_setting` VALUES ('EMAIL_SMTP', '');
INSERT INTO `iriscms_setting` VALUES ('EMAIL_PWD', '');
INSERT INTO `iriscms_setting` VALUES ('EMAIL_EMAIL', '');
INSERT INTO `iriscms_setting` VALUES ('SITE_ICP', '');
INSERT INTO `iriscms_setting` VALUES ('SITE_KEYWORDS', '');
INSERT INTO `iriscms_setting` VALUES ('SITE_DESCRIPTION', '');
INSERT INTO `iriscms_setting` VALUES ('SITE_TITLE', '');

-- ----------------------------
-- Table structure for iriscms_slide
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_slide`;
CREATE TABLE `iriscms_slide`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `name1` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '地址',
  `imgurl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图片地址',
  `wapimgurl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sigin` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图片标识',
  `sort` int(5) NOT NULL DEFAULT 1 COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '幻灯片' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

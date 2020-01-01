/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : iriscms

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 01/01/2020 19:50:36
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
  UNIQUE INDEX `UQE_iriscms_admin_username`(`username`) USING BTREE,
  INDEX `username`(`username`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_admin
-- ----------------------------
INSERT INTO `iriscms_admin` VALUES (1, 'admin', '0087af20a551a8b804f89469534b7859', 1, 'qmRlFL', '::1', 1474291850, 'chenchengbin92@gmail.com', 'mirchen.com');
INSERT INTO `iriscms_admin` VALUES (11, 'test', '834a33db060873a7a208617930edb29a', 9, '5u2G0w', '::1', 0, 'asdasd@asdasd.com1', 'ccc1');

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
  UNIQUE INDEX `UQE_iriscms_admin_role_rolename`(`rolename`) USING BTREE,
  INDEX `listorder`(`listorder`) USING BTREE,
  INDEX `disabled`(`disabled`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_admin_role
-- ----------------------------
INSERT INTO `iriscms_admin_role` VALUES (1, '超级管理员', '超级管理员', 0, 0);
INSERT INTO `iriscms_admin_role` VALUES (9, 'test', 'test', 0, 0);

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
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'left');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'log-delete');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'loglist');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'logview');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'menuadd');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'menudelete');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'menuedit');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'menulist');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'menuorder');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'menuview');
INSERT INTO `iriscms_admin_role_priv` VALUES (9, 'system', 'top');

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
  `tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '后台列表',
  `home_tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前台列表模板',
  `content_tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前台内容页模板',
  `thumb` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `tpl_prefix` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`catid`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 40 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '内容分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_category
-- ----------------------------
INSERT INTO `iriscms_category` VALUES (27, 0, 0, '免费视频', '', '', 2, 1, '', '', '', '', 'content_free_video_');
INSERT INTO `iriscms_category` VALUES (28, 0, 0, '免费图书', '', '', 3, 1, '', '', '', '', 'content_free_book_');
INSERT INTO `iriscms_category` VALUES (29, 0, 0, '付费视频', '', '', 4, 1, '', '', '', '', 'content_paid_video_');
INSERT INTO `iriscms_category` VALUES (30, 0, 0, '付费图书', '', '', 5, 1, '', '', '', '', 'content_paid_book_');
INSERT INTO `iriscms_category` VALUES (31, 0, 30, 'MySQL', '', '', 0, 1, '', '', '', '', '');
INSERT INTO `iriscms_category` VALUES (32, 0, 30, 'GO语言', '', '', 0, 1, '', '', '', '', '');
INSERT INTO `iriscms_category` VALUES (33, 0, 30, 'PYTHON', '', '', 0, 1, '', '', '', '', '');
INSERT INTO `iriscms_category` VALUES (34, 0, 29, '大数据', '', '', 0, 1, '', '', '', '', '');
INSERT INTO `iriscms_category` VALUES (35, 0, 29, 'PHP7', '', '', 0, 1, '', '', '', '', '');
INSERT INTO `iriscms_category` VALUES (36, 0, 29, 'NODEJS', '', '', 0, 1, '', '', '', '', '');
INSERT INTO `iriscms_category` VALUES (37, 0, 28, 'PYTHON', '', '', 0, 1, '', '', '', '', '');
INSERT INTO `iriscms_category` VALUES (38, 0, 27, 'PYTHON', '', '', 0, 1, '', '', '', '', '');
INSERT INTO `iriscms_category` VALUES (39, 0, 0, '软件下载', '', '', 1, 1, '', '', '', '', '');

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
-- Table structure for iriscms_content
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_content`;
CREATE TABLE `iriscms_content`  (
  `id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `catid` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `title` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `thumb` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `keywords` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `description` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `listorder` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  `status` tinyint(2) UNSIGNED NOT NULL DEFAULT 1,
  `recommend` tinyint(2) NULL DEFAULT NULL,
  `pwd_type` tinyint(2) NULL DEFAULT NULL,
  `money` tinyint(5) NULL DEFAULT NULL,
  `created_at` int(11) NULL DEFAULT NULL,
  `updated_at` int(11) NULL DEFAULT NULL,
  `deleted_at` int(11) NULL DEFAULT NULL,
  `source_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `source_pwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `catids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tags` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `userid` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`, `listorder`, `id`) USING BTREE,
  INDEX `listorder`(`catid`, `status`, `listorder`, `id`) USING BTREE,
  INDEX `catid`(`catid`, `status`, `id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '新闻表后期根据模型扩展' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_content
-- ----------------------------
INSERT INTO `iriscms_content` VALUES (1, 30, '项目列表页：导出复制的作品，一直exporting中；', '', '', '', '<p><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span><span style=\"white-space: normal;\">12312312312312312312312</span></p>', 0, 1, 0, 1, 50, 1548300082, 1548300082, 0, '', '123123', '', 'sss,ss', 0);

-- ----------------------------
-- Table structure for iriscms_document_model
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_document_model`;
CREATE TABLE `iriscms_document_model`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文档名称',
  `table` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '对应的表名',
  `enabled` tinyint(4) NULL DEFAULT 0 COMMENT '是否启用',
  `is_system` tinyint(4) NULL DEFAULT 0 COMMENT '是否为系统模型 无法删除',
  `model_type` tinyint(4) NULL DEFAULT 0 COMMENT '模型类型: 扩展模型 和 独立模型',
  `fe_tpl_index` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模型前端主页模板地址',
  `fe_tpl_list` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模型前端列表模板地址',
  `fe_tpl_detail` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模型前端详情模板地址',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文档模型用于存储自定义类型的文档内容' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for iriscms_document_model_dsl
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_document_model_dsl`;
CREATE TABLE `iriscms_document_model_dsl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mid` int(11) NOT NULL DEFAULT 0 COMMENT '模型id',
  `form_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段在表单内的名称',
  `html` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '字段html',
  `required` tinyint(4) NULL DEFAULT 0 COMMENT '是否必填',
  `datasource` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据源, 可以让下拉选项等高级功能有数据读取的源头,具体设计可以是提供列表函数类的',
  `required_tips` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '必填(选)提醒',
  `validator` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '验证器名称或内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for iriscms_document_model_field
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_document_model_field`;
CREATE TABLE `iriscms_document_model_field`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段名称',
  `type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段对应的数据类型',
  `desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字段描述',
  `html` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_document_model_field
-- ----------------------------
INSERT INTO `iriscms_document_model_field` VALUES (1, '单行文本', 'varchar', '常用字段，如文章标题、作者等都属于直接输入少量内容的文本，设置这个文本之后需要指定文本长度，默认为250，如果大于255则为text类型', '<input class=\"easyui-textbox\" {{attr}} style=\"width:300px\">');
INSERT INTO `iriscms_document_model_field` VALUES (2, '多行文本', 'text', '也是较为常用的字段类型，如个人简介、产品描述都可以使用多行文本进行存储', '<input class=\"easyui-textbox\" {{attr}} data-options=\"multiline:true\" style=\"height:80px; width: 300px;\" />');
INSERT INTO `iriscms_document_model_field` VALUES (3, 'HTML文本', 'text', '编辑器编辑产生的html内容，用于比较复杂的内容形式, 可以认为是附带编辑器的多行文本', '<editor />');
INSERT INTO `iriscms_document_model_field` VALUES (4, '附件', 'varshar', '前端表现为input[file]类型,可以上传并且返回一个相对地址', '<input class=\"easyui-filebox\" style=\"width:300px\">');
INSERT INTO `iriscms_document_model_field` VALUES (5, '下拉框', 'varchar', '下拉选择，一般用于如软件类型、语言类型等字段', '<input id=\"cc\" class=\"easyui-combobox\" {{attr}}  data-options=\"\'{{options}}\'\">');
INSERT INTO `iriscms_document_model_field` VALUES (6, '联动类型', 'varchar', '一种数组形式的数据类型，可以在系统后台联动类型管理中进行设置', 'select id=\"cc\" class=\"easyui-combotree\" {{attr}} style=\"width:200px;\" data-options=\"{{options}}\"></select>');
INSERT INTO `iriscms_document_model_field` VALUES (7, '单选框', 'varchar', '平铺显示, 可以认为是下拉框的展开', '<div>{{loop}}<input data-toggle=\"topjui-radiobutton\" {{attr}} value=\"{{value}}\"\">{{loopend}}</div>');
INSERT INTO `iriscms_document_model_field` VALUES (8, '多选框', 'varchar', '多选框, 平铺显示为多个选项', '<div>{{loop}}<input data-toggle=\"topjui-checkbox\" {{attr}} value=\"{{value}}\">{{loopend}} </div>');
INSERT INTO `iriscms_document_model_field` VALUES (9, '整数类型', 'int', '常用字段, 仅能输入数字', '<input type=\"text\" class=\"easyui-numberbox\" value=\"{{value}}\" {{attr}} data-options=\"min:0,precision:0\">');
INSERT INTO `iriscms_document_model_field` VALUES (10, '浮点类型', 'float', '常用字段, 仅能输入浮点数(小数)', '<input type=\"text\" class=\"easyui-numberbox\" value=\"{{value}}\" {{attr}} data-options=\"min:0,precision:4\">');
INSERT INTO `iriscms_document_model_field` VALUES (11, '单图上传', 'varchar', '常用字段, 会生成一个单图上传框', '<images />');
INSERT INTO `iriscms_document_model_field` VALUES (12, '多图上传', 'varchar', '生成一个多图上传的组件', '<mul-images />');

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
  `time` datetime(0) NOT NULL,
  PRIMARY KEY (`logid`) USING BTREE,
  INDEX `module`(`controller`, `action`) USING BTREE,
  INDEX `username`(`username`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 966 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '操作日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_log
-- ----------------------------
INSERT INTO `iriscms_log` VALUES (435, 'category', 'list', '/b/category/list?menuid=36&&_=1576660545834', 1, 'admin', '::1', '2019-12-18 17:15:45');
INSERT INTO `iriscms_log` VALUES (434, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-18 17:15:45');
INSERT INTO `iriscms_log` VALUES (433, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-18 17:15:45');
INSERT INTO `iriscms_log` VALUES (432, 'content', 'right', '/b/content/right?_=1576660545381', 1, 'admin', '::1', '2019-12-18 17:15:45');
INSERT INTO `iriscms_log` VALUES (431, 'content', 'index', '/b/content/index?menuid=35&&_=1576660545355', 1, 'admin', '::1', '2019-12-18 17:15:45');
INSERT INTO `iriscms_log` VALUES (430, 'category', 'list', '/b/category/list?menuid=36&&_=1576660540378', 1, 'admin', '::1', '2019-12-18 17:15:40');
INSERT INTO `iriscms_log` VALUES (429, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-18 17:15:38');
INSERT INTO `iriscms_log` VALUES (428, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-18 17:15:38');
INSERT INTO `iriscms_log` VALUES (427, 'content', 'right', '/b/content/right?_=1576660538305', 1, 'admin', '::1', '2019-12-18 17:15:38');
INSERT INTO `iriscms_log` VALUES (426, 'content', 'index', '/b/content/index?menuid=35&&_=1576660538271', 1, 'admin', '::1', '2019-12-18 17:15:38');
INSERT INTO `iriscms_log` VALUES (425, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-18 17:15:36');
INSERT INTO `iriscms_log` VALUES (424, 'category', 'list', '/b/category/list?menuid=36&&_=1576660536848', 1, 'admin', '::1', '2019-12-18 17:15:36');
INSERT INTO `iriscms_log` VALUES (423, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-18 17:15:34');
INSERT INTO `iriscms_log` VALUES (422, 'setting', 'site', '/b/setting/site?menuid=10&&_=1576660534221', 1, 'admin', '::1', '2019-12-18 17:15:34');
INSERT INTO `iriscms_log` VALUES (421, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:15:33');
INSERT INTO `iriscms_log` VALUES (420, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576660532932', 1, 'admin', '::1', '2019-12-18 17:15:32');
INSERT INTO `iriscms_log` VALUES (419, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:15:32');
INSERT INTO `iriscms_log` VALUES (418, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576660532403', 1, 'admin', '::1', '2019-12-18 17:15:32');
INSERT INTO `iriscms_log` VALUES (417, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-17 17:09:58');
INSERT INTO `iriscms_log` VALUES (416, 'setting', 'site', '/b/setting/site?menuid=10&&_=1576573798645', 1, 'admin', '::1', '2019-12-17 17:09:58');
INSERT INTO `iriscms_log` VALUES (415, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 17:09:57');
INSERT INTO `iriscms_log` VALUES (414, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576573797170', 1, 'admin', '::1', '2019-12-17 17:09:57');
INSERT INTO `iriscms_log` VALUES (413, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 17:09:56');
INSERT INTO `iriscms_log` VALUES (412, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576573796576', 1, 'admin', '::1', '2019-12-17 17:09:56');
INSERT INTO `iriscms_log` VALUES (411, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576573794352', 1, 'admin', '::1', '2019-12-17 17:09:54');
INSERT INTO `iriscms_log` VALUES (410, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '::1', '2019-12-17 17:09:53');
INSERT INTO `iriscms_log` VALUES (409, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576573793709', 1, 'admin', '::1', '2019-12-17 17:09:53');
INSERT INTO `iriscms_log` VALUES (408, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-17 17:09:53');
INSERT INTO `iriscms_log` VALUES (407, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576573792989', 1, 'admin', '::1', '2019-12-17 17:09:53');
INSERT INTO `iriscms_log` VALUES (406, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=35&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 17:09:51');
INSERT INTO `iriscms_log` VALUES (405, 'content', 'news-list', '/b/content/news-list?catid=35&_=1576573790739', 1, 'admin', '::1', '2019-12-17 17:09:50');
INSERT INTO `iriscms_log` VALUES (404, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=36&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 17:09:50');
INSERT INTO `iriscms_log` VALUES (403, 'content', 'news-list', '/b/content/news-list?catid=36&_=1576573790200', 1, 'admin', '::1', '2019-12-17 17:09:50');
INSERT INTO `iriscms_log` VALUES (402, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 17:09:49');
INSERT INTO `iriscms_log` VALUES (401, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 17:09:49');
INSERT INTO `iriscms_log` VALUES (400, 'content', 'right', '/b/content/right?_=1576573789175', 1, 'admin', '::1', '2019-12-17 17:09:49');
INSERT INTO `iriscms_log` VALUES (399, 'content', 'index', '/b/content/index?menuid=35&&_=1576573789128', 1, 'admin', '::1', '2019-12-17 17:09:49');
INSERT INTO `iriscms_log` VALUES (398, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-17 17:09:47');
INSERT INTO `iriscms_log` VALUES (397, 'category', 'list', '/b/category/list?menuid=36&&_=1576573786868', 1, 'admin', '::1', '2019-12-17 17:09:46');
INSERT INTO `iriscms_log` VALUES (396, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 17:09:45');
INSERT INTO `iriscms_log` VALUES (395, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 17:09:45');
INSERT INTO `iriscms_log` VALUES (394, 'content', 'right', '/b/content/right?_=1576573785687', 1, 'admin', '::1', '2019-12-17 17:09:45');
INSERT INTO `iriscms_log` VALUES (393, 'content', 'index', '/b/content/index?menuid=35&&_=1576573785653', 1, 'admin', '::1', '2019-12-17 17:09:45');
INSERT INTO `iriscms_log` VALUES (392, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576573392910', 11, 'test', '::1', '2019-12-17 17:03:12');
INSERT INTO `iriscms_log` VALUES (391, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 17:03:12');
INSERT INTO `iriscms_log` VALUES (390, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576573392322', 11, 'test', '::1', '2019-12-17 17:03:12');
INSERT INTO `iriscms_log` VALUES (389, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 17:03:11');
INSERT INTO `iriscms_log` VALUES (388, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576573391627', 11, 'test', '::1', '2019-12-17 17:03:11');
INSERT INTO `iriscms_log` VALUES (387, 'system', 'loglist', '/b/system/loglist?page=13&rows=20', 11, 'test', '::1', '2019-12-17 16:17:15');
INSERT INTO `iriscms_log` VALUES (386, 'system', 'loglist', '/b/system/loglist?page=12&rows=20', 11, 'test', '::1', '2019-12-17 16:17:14');
INSERT INTO `iriscms_log` VALUES (385, 'system', 'loglist', '/b/system/loglist?page=11&rows=20', 11, 'test', '::1', '2019-12-17 16:17:14');
INSERT INTO `iriscms_log` VALUES (384, 'system', 'loglist', '/b/system/loglist?page=10&rows=20', 11, 'test', '::1', '2019-12-17 16:17:14');
INSERT INTO `iriscms_log` VALUES (383, 'system', 'loglist', '/b/system/loglist?page=9&rows=20', 11, 'test', '::1', '2019-12-17 16:17:11');
INSERT INTO `iriscms_log` VALUES (382, 'system', 'loglist', '/b/system/loglist?page=8&rows=20', 11, 'test', '::1', '2019-12-17 16:17:11');
INSERT INTO `iriscms_log` VALUES (381, 'system', 'loglist', '/b/system/loglist?page=7&rows=20', 11, 'test', '::1', '2019-12-17 16:17:11');
INSERT INTO `iriscms_log` VALUES (380, 'system', 'loglist', '/b/system/loglist?page=6&rows=20', 11, 'test', '::1', '2019-12-17 16:17:11');
INSERT INTO `iriscms_log` VALUES (379, 'system', 'loglist', '/b/system/loglist?page=5&rows=20', 11, 'test', '::1', '2019-12-17 16:17:06');
INSERT INTO `iriscms_log` VALUES (378, 'system', 'loglist', '/b/system/loglist?page=5&rows=50', 11, 'test', '::1', '2019-12-17 16:17:04');
INSERT INTO `iriscms_log` VALUES (377, 'system', 'loglist', '/b/system/loglist?page=4&rows=50', 11, 'test', '::1', '2019-12-17 16:17:04');
INSERT INTO `iriscms_log` VALUES (376, 'system', 'loglist', '/b/system/loglist?page=3&rows=50', 11, 'test', '::1', '2019-12-17 16:17:03');
INSERT INTO `iriscms_log` VALUES (375, 'system', 'loglist', '/b/system/loglist?page=2&rows=50', 11, 'test', '::1', '2019-12-17 16:17:00');
INSERT INTO `iriscms_log` VALUES (374, 'system', 'loglist', '/b/system/loglist?page=1&rows=50', 11, 'test', '::1', '2019-12-17 16:16:58');
INSERT INTO `iriscms_log` VALUES (373, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 16:16:52');
INSERT INTO `iriscms_log` VALUES (372, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576570612835', 11, 'test', '::1', '2019-12-17 16:16:52');
INSERT INTO `iriscms_log` VALUES (371, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 16:16:52');
INSERT INTO `iriscms_log` VALUES (370, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576570611955', 11, 'test', '::1', '2019-12-17 16:16:51');
INSERT INTO `iriscms_log` VALUES (369, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576570582972', 11, 'test', '::1', '2019-12-17 16:16:22');
INSERT INTO `iriscms_log` VALUES (368, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 16:16:22');
INSERT INTO `iriscms_log` VALUES (367, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576570582358', 11, 'test', '::1', '2019-12-17 16:16:22');
INSERT INTO `iriscms_log` VALUES (366, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 16:16:22');
INSERT INTO `iriscms_log` VALUES (365, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576570581867', 11, 'test', '::1', '2019-12-17 16:16:21');
INSERT INTO `iriscms_log` VALUES (364, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 15:11:03');
INSERT INTO `iriscms_log` VALUES (363, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 15:10:47');
INSERT INTO `iriscms_log` VALUES (362, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576566646952', 11, 'test', '::1', '2019-12-17 15:10:46');
INSERT INTO `iriscms_log` VALUES (361, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 15:10:10');
INSERT INTO `iriscms_log` VALUES (360, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576566610010', 11, 'test', '::1', '2019-12-17 15:10:10');
INSERT INTO `iriscms_log` VALUES (359, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-17 15:10:02');
INSERT INTO `iriscms_log` VALUES (358, 'system', 'loglist', '/b/system/loglist?page=3&rows=10', 11, 'test', '::1', '2019-12-17 15:06:54');
INSERT INTO `iriscms_log` VALUES (357, 'system', 'loglist', '/b/system/loglist?page=2&rows=10', 11, 'test', '::1', '2019-12-17 15:06:52');
INSERT INTO `iriscms_log` VALUES (356, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-17 15:05:54');
INSERT INTO `iriscms_log` VALUES (355, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-17 15:05:52');
INSERT INTO `iriscms_log` VALUES (354, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576566352018', 11, 'test', '::1', '2019-12-17 15:05:52');
INSERT INTO `iriscms_log` VALUES (353, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 15:05:51');
INSERT INTO `iriscms_log` VALUES (352, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576566351112', 11, 'test', '::1', '2019-12-17 15:05:51');
INSERT INTO `iriscms_log` VALUES (351, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 15:05:50');
INSERT INTO `iriscms_log` VALUES (350, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576566350127', 11, 'test', '::1', '2019-12-17 15:05:50');
INSERT INTO `iriscms_log` VALUES (349, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 15:03:54');
INSERT INTO `iriscms_log` VALUES (348, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576566233912', 11, 'test', '::1', '2019-12-17 15:03:53');
INSERT INTO `iriscms_log` VALUES (347, 'system', 'loglist', '/b/system/loglist?page=7&rows=10', 11, 'test', '::1', '2019-12-17 15:03:52');
INSERT INTO `iriscms_log` VALUES (346, 'system', 'loglist', '/b/system/loglist?page=6&rows=10', 11, 'test', '::1', '2019-12-17 15:03:52');
INSERT INTO `iriscms_log` VALUES (345, 'system', 'loglist', '/b/system/loglist?page=5&rows=10', 11, 'test', '::1', '2019-12-17 15:03:52');
INSERT INTO `iriscms_log` VALUES (344, 'system', 'loglist', '/b/system/loglist?page=4&rows=10', 11, 'test', '::1', '2019-12-17 15:03:52');
INSERT INTO `iriscms_log` VALUES (343, 'system', 'loglist', '/b/system/loglist?page=3&rows=10', 11, 'test', '::1', '2019-12-17 15:03:51');
INSERT INTO `iriscms_log` VALUES (342, 'system', 'loglist', '/b/system/loglist?page=2&rows=10', 11, 'test', '::1', '2019-12-17 15:03:50');
INSERT INTO `iriscms_log` VALUES (341, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 15:01:31');
INSERT INTO `iriscms_log` VALUES (340, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576566090676', 11, 'test', '::1', '2019-12-17 15:01:30');
INSERT INTO `iriscms_log` VALUES (339, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 15:01:28');
INSERT INTO `iriscms_log` VALUES (338, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576566088116', 11, 'test', '::1', '2019-12-17 15:01:28');
INSERT INTO `iriscms_log` VALUES (337, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-17 15:01:20');
INSERT INTO `iriscms_log` VALUES (336, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-17 14:56:22');
INSERT INTO `iriscms_log` VALUES (335, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-17 14:56:19');
INSERT INTO `iriscms_log` VALUES (334, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:56:05');
INSERT INTO `iriscms_log` VALUES (333, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-17 14:55:52');
INSERT INTO `iriscms_log` VALUES (135, 'content', 'index', '/b/content/index?menuid=35&&_=1576550314690', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:34');
INSERT INTO `iriscms_log` VALUES (136, 'content', 'right', '/b/content/right?_=1576550314780', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:34');
INSERT INTO `iriscms_log` VALUES (137, 'content', 'right', '/b/content/right', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:34');
INSERT INTO `iriscms_log` VALUES (138, 'content', 'right', '/b/content/right', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:34');
INSERT INTO `iriscms_log` VALUES (139, 'category', 'list', '/b/category/list?menuid=36&&_=1576550315676', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:35');
INSERT INTO `iriscms_log` VALUES (140, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:35');
INSERT INTO `iriscms_log` VALUES (141, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576550318895', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:38');
INSERT INTO `iriscms_log` VALUES (142, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:39');
INSERT INTO `iriscms_log` VALUES (143, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576550319796', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:39');
INSERT INTO `iriscms_log` VALUES (144, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:40');
INSERT INTO `iriscms_log` VALUES (145, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576550322440', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:42');
INSERT INTO `iriscms_log` VALUES (146, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:42');
INSERT INTO `iriscms_log` VALUES (147, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576550323550', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:43');
INSERT INTO `iriscms_log` VALUES (148, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:43');
INSERT INTO `iriscms_log` VALUES (149, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576550324624', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:44');
INSERT INTO `iriscms_log` VALUES (150, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:44');
INSERT INTO `iriscms_log` VALUES (151, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576550325604', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:45');
INSERT INTO `iriscms_log` VALUES (152, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:45');
INSERT INTO `iriscms_log` VALUES (153, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576550327263', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:47');
INSERT INTO `iriscms_log` VALUES (154, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:47');
INSERT INTO `iriscms_log` VALUES (155, 'content', 'index', '/b/content/index?menuid=35&&_=1576550336530', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:56');
INSERT INTO `iriscms_log` VALUES (156, 'content', 'right', '/b/content/right?_=1576550336621', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:56');
INSERT INTO `iriscms_log` VALUES (157, 'content', 'right', '/b/content/right', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:56');
INSERT INTO `iriscms_log` VALUES (158, 'content', 'right', '/b/content/right', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:56');
INSERT INTO `iriscms_log` VALUES (159, 'category', 'list', '/b/category/list?menuid=36&&_=1576550337288', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:57');
INSERT INTO `iriscms_log` VALUES (160, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:57');
INSERT INTO `iriscms_log` VALUES (161, 'content', 'index', '/b/content/index?menuid=35&&_=1576550339277', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:59');
INSERT INTO `iriscms_log` VALUES (162, 'content', 'right', '/b/content/right?_=1576550339370', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:59');
INSERT INTO `iriscms_log` VALUES (163, 'content', 'right', '/b/content/right', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:59');
INSERT INTO `iriscms_log` VALUES (164, 'content', 'right', '/b/content/right', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:59');
INSERT INTO `iriscms_log` VALUES (165, 'category', 'list', '/b/category/list?menuid=36&&_=1576550339865', 1, 'admin', '127.0.0.1', '2019-12-17 10:38:59');
INSERT INTO `iriscms_log` VALUES (166, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '127.0.0.1', '2019-12-17 10:39:00');
INSERT INTO `iriscms_log` VALUES (167, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576550342172', 1, 'admin', '127.0.0.1', '2019-12-17 10:39:02');
INSERT INTO `iriscms_log` VALUES (168, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '127.0.0.1', '2019-12-17 10:39:02');
INSERT INTO `iriscms_log` VALUES (169, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576550343251', 1, 'admin', '127.0.0.1', '2019-12-17 10:39:03');
INSERT INTO `iriscms_log` VALUES (170, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '127.0.0.1', '2019-12-17 10:39:03');
INSERT INTO `iriscms_log` VALUES (171, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576550343901', 1, 'admin', '127.0.0.1', '2019-12-17 10:39:03');
INSERT INTO `iriscms_log` VALUES (172, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '127.0.0.1', '2019-12-17 10:39:04');
INSERT INTO `iriscms_log` VALUES (173, 'category', 'list', '/b/category/list?menuid=36&&_=1576563928708', 1, 'admin', '::1', '2019-12-17 14:25:28');
INSERT INTO `iriscms_log` VALUES (174, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-17 14:25:28');
INSERT INTO `iriscms_log` VALUES (175, 'content', 'index', '/b/content/index?menuid=35&&_=1576563929759', 1, 'admin', '::1', '2019-12-17 14:25:29');
INSERT INTO `iriscms_log` VALUES (176, 'content', 'right', '/b/content/right?_=1576563929798', 1, 'admin', '::1', '2019-12-17 14:25:29');
INSERT INTO `iriscms_log` VALUES (177, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:29');
INSERT INTO `iriscms_log` VALUES (178, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:29');
INSERT INTO `iriscms_log` VALUES (179, 'category', 'list', '/b/category/list?menuid=36&&_=1576563931143', 1, 'admin', '::1', '2019-12-17 14:25:31');
INSERT INTO `iriscms_log` VALUES (180, 'content', 'index', '/b/content/index?menuid=35&&_=1576563931661', 1, 'admin', '::1', '2019-12-17 14:25:31');
INSERT INTO `iriscms_log` VALUES (181, 'content', 'right', '/b/content/right?_=1576563931694', 1, 'admin', '::1', '2019-12-17 14:25:31');
INSERT INTO `iriscms_log` VALUES (182, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:31');
INSERT INTO `iriscms_log` VALUES (183, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:31');
INSERT INTO `iriscms_log` VALUES (184, 'category', 'list', '/b/category/list?menuid=36&&_=1576563932681', 1, 'admin', '::1', '2019-12-17 14:25:32');
INSERT INTO `iriscms_log` VALUES (185, 'content', 'index', '/b/content/index?menuid=35&&_=1576563933106', 1, 'admin', '::1', '2019-12-17 14:25:33');
INSERT INTO `iriscms_log` VALUES (186, 'content', 'right', '/b/content/right?_=1576563933139', 1, 'admin', '::1', '2019-12-17 14:25:33');
INSERT INTO `iriscms_log` VALUES (187, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:33');
INSERT INTO `iriscms_log` VALUES (188, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:33');
INSERT INTO `iriscms_log` VALUES (189, 'category', 'list', '/b/category/list?menuid=36&&_=1576563934001', 1, 'admin', '::1', '2019-12-17 14:25:34');
INSERT INTO `iriscms_log` VALUES (190, 'content', 'index', '/b/content/index?menuid=35&&_=1576563934628', 1, 'admin', '::1', '2019-12-17 14:25:34');
INSERT INTO `iriscms_log` VALUES (191, 'content', 'right', '/b/content/right?_=1576563934660', 1, 'admin', '::1', '2019-12-17 14:25:34');
INSERT INTO `iriscms_log` VALUES (192, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:34');
INSERT INTO `iriscms_log` VALUES (193, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:34');
INSERT INTO `iriscms_log` VALUES (194, 'category', 'list', '/b/category/list?menuid=36&&_=1576563935447', 1, 'admin', '::1', '2019-12-17 14:25:35');
INSERT INTO `iriscms_log` VALUES (195, 'content', 'index', '/b/content/index?menuid=35&&_=1576563936197', 1, 'admin', '::1', '2019-12-17 14:25:36');
INSERT INTO `iriscms_log` VALUES (196, 'content', 'right', '/b/content/right?_=1576563936228', 1, 'admin', '::1', '2019-12-17 14:25:36');
INSERT INTO `iriscms_log` VALUES (197, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:36');
INSERT INTO `iriscms_log` VALUES (198, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:25:36');
INSERT INTO `iriscms_log` VALUES (199, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576563939115', 1, 'admin', '::1', '2019-12-17 14:25:39');
INSERT INTO `iriscms_log` VALUES (200, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:25:39');
INSERT INTO `iriscms_log` VALUES (201, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576563952651', 1, 'admin', '::1', '2019-12-17 14:25:52');
INSERT INTO `iriscms_log` VALUES (202, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-17 14:25:52');
INSERT INTO `iriscms_log` VALUES (203, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576563953561', 1, 'admin', '::1', '2019-12-17 14:25:53');
INSERT INTO `iriscms_log` VALUES (204, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:25:53');
INSERT INTO `iriscms_log` VALUES (205, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576563961798', 1, 'admin', '::1', '2019-12-17 14:26:01');
INSERT INTO `iriscms_log` VALUES (206, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576563962301', 1, 'admin', '::1', '2019-12-17 14:26:02');
INSERT INTO `iriscms_log` VALUES (207, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576563964434', 1, 'admin', '::1', '2019-12-17 14:26:04');
INSERT INTO `iriscms_log` VALUES (208, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-17 14:26:04');
INSERT INTO `iriscms_log` VALUES (209, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576563965101', 1, 'admin', '::1', '2019-12-17 14:26:05');
INSERT INTO `iriscms_log` VALUES (210, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:26:05');
INSERT INTO `iriscms_log` VALUES (211, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576563984246', 1, 'admin', '::1', '2019-12-17 14:26:24');
INSERT INTO `iriscms_log` VALUES (212, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-17 14:26:24');
INSERT INTO `iriscms_log` VALUES (213, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576563994844', 1, 'admin', '::1', '2019-12-17 14:26:34');
INSERT INTO `iriscms_log` VALUES (214, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:26:34');
INSERT INTO `iriscms_log` VALUES (215, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576563995492', 1, 'admin', '::1', '2019-12-17 14:26:35');
INSERT INTO `iriscms_log` VALUES (216, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:26:35');
INSERT INTO `iriscms_log` VALUES (217, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576563996052', 1, 'admin', '::1', '2019-12-17 14:26:36');
INSERT INTO `iriscms_log` VALUES (218, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576563996560', 1, 'admin', '::1', '2019-12-17 14:26:36');
INSERT INTO `iriscms_log` VALUES (219, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576563997086', 1, 'admin', '::1', '2019-12-17 14:26:37');
INSERT INTO `iriscms_log` VALUES (220, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576563997564', 1, 'admin', '::1', '2019-12-17 14:26:37');
INSERT INTO `iriscms_log` VALUES (221, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576563998076', 1, 'admin', '::1', '2019-12-17 14:26:38');
INSERT INTO `iriscms_log` VALUES (222, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576563998767', 1, 'admin', '::1', '2019-12-17 14:26:38');
INSERT INTO `iriscms_log` VALUES (223, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576563999973', 1, 'admin', '::1', '2019-12-17 14:26:39');
INSERT INTO `iriscms_log` VALUES (224, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576564000400', 1, 'admin', '::1', '2019-12-17 14:26:40');
INSERT INTO `iriscms_log` VALUES (225, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576564001741', 1, 'admin', '::1', '2019-12-17 14:26:41');
INSERT INTO `iriscms_log` VALUES (226, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576564008778', 1, 'admin', '::1', '2019-12-17 14:26:48');
INSERT INTO `iriscms_log` VALUES (227, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-17 14:26:48');
INSERT INTO `iriscms_log` VALUES (228, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564009681', 1, 'admin', '::1', '2019-12-17 14:26:49');
INSERT INTO `iriscms_log` VALUES (229, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:26:49');
INSERT INTO `iriscms_log` VALUES (230, 'system', 'log-delete', '/b/system/log-delete', 1, 'admin', '::1', '2019-12-17 14:26:50');
INSERT INTO `iriscms_log` VALUES (231, 'system', 'log-delete', '/b/system/log-delete', 1, 'admin', '::1', '2019-12-17 14:26:55');
INSERT INTO `iriscms_log` VALUES (232, 'system', 'log-delete', '/b/system/log-delete', 1, 'admin', '::1', '2019-12-17 14:27:00');
INSERT INTO `iriscms_log` VALUES (233, 'content', 'index', '/b/content/index?menuid=35&&_=1576564280767', 1, 'admin', '::1', '2019-12-17 14:31:20');
INSERT INTO `iriscms_log` VALUES (234, 'content', 'right', '/b/content/right?_=1576564280800', 1, 'admin', '::1', '2019-12-17 14:31:20');
INSERT INTO `iriscms_log` VALUES (235, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:31:20');
INSERT INTO `iriscms_log` VALUES (236, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:31:20');
INSERT INTO `iriscms_log` VALUES (237, 'category', 'list', '/b/category/list?menuid=36&&_=1576564281398', 1, 'admin', '::1', '2019-12-17 14:31:21');
INSERT INTO `iriscms_log` VALUES (238, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-17 14:31:21');
INSERT INTO `iriscms_log` VALUES (239, 'content', 'index', '/b/content/index?menuid=35&&_=1576564282169', 1, 'admin', '::1', '2019-12-17 14:31:22');
INSERT INTO `iriscms_log` VALUES (240, 'content', 'right', '/b/content/right?_=1576564282201', 1, 'admin', '::1', '2019-12-17 14:31:22');
INSERT INTO `iriscms_log` VALUES (241, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:31:22');
INSERT INTO `iriscms_log` VALUES (242, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:31:22');
INSERT INTO `iriscms_log` VALUES (243, 'category', 'list', '/b/category/list?menuid=36&&_=1576564283896', 1, 'admin', '::1', '2019-12-17 14:31:23');
INSERT INTO `iriscms_log` VALUES (244, 'content', 'index', '/b/content/index?menuid=35&&_=1576564284305', 1, 'admin', '::1', '2019-12-17 14:31:24');
INSERT INTO `iriscms_log` VALUES (245, 'content', 'right', '/b/content/right?_=1576564284336', 1, 'admin', '::1', '2019-12-17 14:31:24');
INSERT INTO `iriscms_log` VALUES (246, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:31:24');
INSERT INTO `iriscms_log` VALUES (247, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-17 14:31:24');
INSERT INTO `iriscms_log` VALUES (248, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564287275', 1, 'admin', '::1', '2019-12-17 14:31:27');
INSERT INTO `iriscms_log` VALUES (249, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:31:27');
INSERT INTO `iriscms_log` VALUES (250, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576564288069', 1, 'admin', '::1', '2019-12-17 14:31:28');
INSERT INTO `iriscms_log` VALUES (251, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-17 14:31:28');
INSERT INTO `iriscms_log` VALUES (252, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576564290792', 1, 'admin', '::1', '2019-12-17 14:31:30');
INSERT INTO `iriscms_log` VALUES (253, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:31:30');
INSERT INTO `iriscms_log` VALUES (254, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576564291408', 1, 'admin', '::1', '2019-12-17 14:31:31');
INSERT INTO `iriscms_log` VALUES (255, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:31:31');
INSERT INTO `iriscms_log` VALUES (256, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576564291881', 1, 'admin', '::1', '2019-12-17 14:31:31');
INSERT INTO `iriscms_log` VALUES (257, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576564292348', 1, 'admin', '::1', '2019-12-17 14:31:32');
INSERT INTO `iriscms_log` VALUES (258, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576564292850', 1, 'admin', '::1', '2019-12-17 14:31:32');
INSERT INTO `iriscms_log` VALUES (259, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:32:59');
INSERT INTO `iriscms_log` VALUES (260, 'admin', 'member-add', '/b/admin/member-add?_=1576564385957', 1, 'admin', '::1', '2019-12-17 14:33:05');
INSERT INTO `iriscms_log` VALUES (261, 'admin', 'member-add', '/b/admin/member-add', 1, 'admin', '::1', '2019-12-17 14:33:25');
INSERT INTO `iriscms_log` VALUES (262, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:33:25');
INSERT INTO `iriscms_log` VALUES (263, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576564407670', 1, 'admin', '::1', '2019-12-17 14:33:27');
INSERT INTO `iriscms_log` VALUES (264, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:33:27');
INSERT INTO `iriscms_log` VALUES (265, 'admin', 'role-add', '/b/admin/role-add?_=1576564409015', 1, 'admin', '::1', '2019-12-17 14:33:29');
INSERT INTO `iriscms_log` VALUES (266, 'admin', 'role-add', '/b/admin/role-add', 1, 'admin', '::1', '2019-12-17 14:33:32');
INSERT INTO `iriscms_log` VALUES (267, 'admin', 'role-add', '/b/admin/role-add', 1, 'admin', '::1', '2019-12-17 14:33:39');
INSERT INTO `iriscms_log` VALUES (268, 'admin', 'role-add', '/b/admin/role-add', 1, 'admin', '::1', '2019-12-17 14:33:42');
INSERT INTO `iriscms_log` VALUES (269, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:33:42');
INSERT INTO `iriscms_log` VALUES (270, 'admin', 'role-permission', '/b/admin/role-permission?id=9&_=1576564425640', 1, 'admin', '::1', '2019-12-17 14:33:45');
INSERT INTO `iriscms_log` VALUES (271, 'admin', 'role-permission', '/b/admin/role-permission?id=9', 1, 'admin', '::1', '2019-12-17 14:33:45');
INSERT INTO `iriscms_log` VALUES (272, 'admin', 'role-permission', '/b/admin/role-permission?dosubmit=1&id=9', 1, 'admin', '::1', '2019-12-17 14:33:47');
INSERT INTO `iriscms_log` VALUES (273, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576564436343', 1, 'admin', '::1', '2019-12-17 14:33:56');
INSERT INTO `iriscms_log` VALUES (274, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:33:56');
INSERT INTO `iriscms_log` VALUES (275, 'admin', 'member-edit', '/b/admin/member-edit?id=11&_=1576564439223', 1, 'admin', '::1', '2019-12-17 14:33:59');
INSERT INTO `iriscms_log` VALUES (276, 'admin', 'member-edit', '/b/admin/member-edit?id=11', 1, 'admin', '::1', '2019-12-17 14:34:02');
INSERT INTO `iriscms_log` VALUES (277, 'admin', 'member-edit', '/b/admin/member-edit?id=11', 1, 'admin', '::1', '2019-12-17 14:34:05');
INSERT INTO `iriscms_log` VALUES (278, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576564452313', 1, 'admin', '::1', '2019-12-17 14:34:12');
INSERT INTO `iriscms_log` VALUES (279, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:34:12');
INSERT INTO `iriscms_log` VALUES (280, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576564453831', 1, 'admin', '::1', '2019-12-17 14:34:13');
INSERT INTO `iriscms_log` VALUES (281, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:34:14');
INSERT INTO `iriscms_log` VALUES (282, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576564454264', 1, 'admin', '::1', '2019-12-17 14:34:14');
INSERT INTO `iriscms_log` VALUES (283, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564472022', 11, 'test', '::1', '2019-12-17 14:34:32');
INSERT INTO `iriscms_log` VALUES (284, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:34:32');
INSERT INTO `iriscms_log` VALUES (285, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576564473675', 11, 'test', '::1', '2019-12-17 14:34:33');
INSERT INTO `iriscms_log` VALUES (286, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 14:34:33');
INSERT INTO `iriscms_log` VALUES (287, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564476531', 11, 'test', '::1', '2019-12-17 14:34:36');
INSERT INTO `iriscms_log` VALUES (288, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576564477448', 11, 'test', '::1', '2019-12-17 14:34:37');
INSERT INTO `iriscms_log` VALUES (289, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564511994', 11, 'test', '::1', '2019-12-17 14:35:12');
INSERT INTO `iriscms_log` VALUES (290, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:35:12');
INSERT INTO `iriscms_log` VALUES (291, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576564512523', 11, 'test', '::1', '2019-12-17 14:35:12');
INSERT INTO `iriscms_log` VALUES (292, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 14:35:12');
INSERT INTO `iriscms_log` VALUES (293, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564513138', 11, 'test', '::1', '2019-12-17 14:35:13');
INSERT INTO `iriscms_log` VALUES (294, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576564513525', 11, 'test', '::1', '2019-12-17 14:35:13');
INSERT INTO `iriscms_log` VALUES (295, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564513921', 11, 'test', '::1', '2019-12-17 14:35:13');
INSERT INTO `iriscms_log` VALUES (296, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576564514261', 11, 'test', '::1', '2019-12-17 14:35:14');
INSERT INTO `iriscms_log` VALUES (297, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564579863', 11, 'test', '::1', '2019-12-17 14:36:19');
INSERT INTO `iriscms_log` VALUES (298, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:36:19');
INSERT INTO `iriscms_log` VALUES (299, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576564580485', 11, 'test', '::1', '2019-12-17 14:36:20');
INSERT INTO `iriscms_log` VALUES (300, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 14:36:20');
INSERT INTO `iriscms_log` VALUES (301, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564581049', 11, 'test', '::1', '2019-12-17 14:36:21');
INSERT INTO `iriscms_log` VALUES (302, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576564581495', 11, 'test', '::1', '2019-12-17 14:36:21');
INSERT INTO `iriscms_log` VALUES (303, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576564633565', 11, 'test', '::1', '2019-12-17 14:37:13');
INSERT INTO `iriscms_log` VALUES (304, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:37:13');
INSERT INTO `iriscms_log` VALUES (305, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576565063582', 11, 'test', '::1', '2019-12-17 14:44:23');
INSERT INTO `iriscms_log` VALUES (306, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:44:23');
INSERT INTO `iriscms_log` VALUES (307, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576565064854', 11, 'test', '::1', '2019-12-17 14:44:24');
INSERT INTO `iriscms_log` VALUES (308, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 14:44:24');
INSERT INTO `iriscms_log` VALUES (309, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576565065345', 11, 'test', '::1', '2019-12-17 14:44:25');
INSERT INTO `iriscms_log` VALUES (310, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576565075222', 11, 'test', '::1', '2019-12-17 14:44:35');
INSERT INTO `iriscms_log` VALUES (311, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 14:44:35');
INSERT INTO `iriscms_log` VALUES (312, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576565077149', 11, 'test', '::1', '2019-12-17 14:44:37');
INSERT INTO `iriscms_log` VALUES (313, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:44:37');
INSERT INTO `iriscms_log` VALUES (314, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576565099598', 1, 'admin', '::1', '2019-12-17 14:44:59');
INSERT INTO `iriscms_log` VALUES (315, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:44:59');
INSERT INTO `iriscms_log` VALUES (316, 'admin', 'member-edit', '/b/admin/member-edit?id=11&_=1576565101674', 1, 'admin', '::1', '2019-12-17 14:45:01');
INSERT INTO `iriscms_log` VALUES (317, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576565104256', 1, 'admin', '::1', '2019-12-17 14:45:04');
INSERT INTO `iriscms_log` VALUES (318, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-17 14:45:04');
INSERT INTO `iriscms_log` VALUES (319, 'admin', 'role-permission', '/b/admin/role-permission?id=9&_=1576565105824', 1, 'admin', '::1', '2019-12-17 14:45:05');
INSERT INTO `iriscms_log` VALUES (320, 'admin', 'role-permission', '/b/admin/role-permission?id=9', 1, 'admin', '::1', '2019-12-17 14:45:05');
INSERT INTO `iriscms_log` VALUES (321, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576565143231', 11, 'test', '::1', '2019-12-17 14:45:43');
INSERT INTO `iriscms_log` VALUES (322, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 14:45:43');
INSERT INTO `iriscms_log` VALUES (323, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576565144798', 11, 'test', '::1', '2019-12-17 14:45:44');
INSERT INTO `iriscms_log` VALUES (324, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:45:45');
INSERT INTO `iriscms_log` VALUES (325, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576565255995', 11, 'test', '::1', '2019-12-17 14:47:36');
INSERT INTO `iriscms_log` VALUES (326, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:47:36');
INSERT INTO `iriscms_log` VALUES (327, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576565302315', 11, 'test', '::1', '2019-12-17 14:48:22');
INSERT INTO `iriscms_log` VALUES (328, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-17 14:48:22');
INSERT INTO `iriscms_log` VALUES (329, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576565303426', 11, 'test', '::1', '2019-12-17 14:48:23');
INSERT INTO `iriscms_log` VALUES (330, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-17 14:48:23');
INSERT INTO `iriscms_log` VALUES (331, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-17 14:51:36');
INSERT INTO `iriscms_log` VALUES (332, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-17 14:54:56');
INSERT INTO `iriscms_log` VALUES (436, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576660548323', 1, 'admin', '::1', '2019-12-18 17:15:48');
INSERT INTO `iriscms_log` VALUES (437, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:15:48');
INSERT INTO `iriscms_log` VALUES (438, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576660549182', 1, 'admin', '::1', '2019-12-18 17:15:49');
INSERT INTO `iriscms_log` VALUES (439, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-18 17:15:49');
INSERT INTO `iriscms_log` VALUES (440, 'setting', 'site', '/b/setting/site?menuid=10&&_=1576660562795', 1, 'admin', '::1', '2019-12-18 17:16:02');
INSERT INTO `iriscms_log` VALUES (441, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-18 17:16:02');
INSERT INTO `iriscms_log` VALUES (442, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576660564886', 1, 'admin', '::1', '2019-12-18 17:16:04');
INSERT INTO `iriscms_log` VALUES (443, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:16:05');
INSERT INTO `iriscms_log` VALUES (444, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576660565722', 1, 'admin', '::1', '2019-12-18 17:16:05');
INSERT INTO `iriscms_log` VALUES (445, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:16:05');
INSERT INTO `iriscms_log` VALUES (446, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576660569364', 1, 'admin', '::1', '2019-12-18 17:16:09');
INSERT INTO `iriscms_log` VALUES (447, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576660569826', 1, 'admin', '::1', '2019-12-18 17:16:09');
INSERT INTO `iriscms_log` VALUES (448, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576660603592', 1, 'admin', '::1', '2019-12-18 17:16:43');
INSERT INTO `iriscms_log` VALUES (449, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:16:43');
INSERT INTO `iriscms_log` VALUES (450, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576660606048', 1, 'admin', '::1', '2019-12-18 17:16:46');
INSERT INTO `iriscms_log` VALUES (451, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:16:46');
INSERT INTO `iriscms_log` VALUES (452, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576660611890', 1, 'admin', '::1', '2019-12-18 17:16:51');
INSERT INTO `iriscms_log` VALUES (453, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-18 17:16:51');
INSERT INTO `iriscms_log` VALUES (454, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-18 17:19:19');
INSERT INTO `iriscms_log` VALUES (455, 'content', 'index', '/b/content/index?menuid=35&&_=1576660778830', 1, 'admin', '::1', '2019-12-18 17:19:38');
INSERT INTO `iriscms_log` VALUES (456, 'content', 'right', '/b/content/right?_=1576660778857', 1, 'admin', '::1', '2019-12-18 17:19:38');
INSERT INTO `iriscms_log` VALUES (457, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-18 17:19:38');
INSERT INTO `iriscms_log` VALUES (458, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-18 17:19:38');
INSERT INTO `iriscms_log` VALUES (459, 'category', 'list', '/b/category/list?menuid=36&&_=1576660779525', 1, 'admin', '::1', '2019-12-18 17:19:39');
INSERT INTO `iriscms_log` VALUES (460, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-18 17:19:39');
INSERT INTO `iriscms_log` VALUES (461, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576660782181', 1, 'admin', '::1', '2019-12-18 17:19:42');
INSERT INTO `iriscms_log` VALUES (462, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-18 17:19:42');
INSERT INTO `iriscms_log` VALUES (463, 'system', 'menu-order', '/b/system/menu-order', 1, 'admin', '::1', '2019-12-18 17:19:50');
INSERT INTO `iriscms_log` VALUES (464, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-18 17:19:54');
INSERT INTO `iriscms_log` VALUES (465, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576660823268', 1, 'admin', '::1', '2019-12-18 17:20:23');
INSERT INTO `iriscms_log` VALUES (466, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:20:23');
INSERT INTO `iriscms_log` VALUES (467, 'user', 'info', '/b/user/info?menuid=57&&_=1576660956382', 1, 'admin', '::1', '2019-12-18 17:22:36');
INSERT INTO `iriscms_log` VALUES (468, 'user', 'list', '/b/user/list?menuid=56&&_=1576660985193', 1, 'admin', '::1', '2019-12-18 17:23:05');
INSERT INTO `iriscms_log` VALUES (469, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:23:05');
INSERT INTO `iriscms_log` VALUES (470, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576660990134', 1, 'admin', '::1', '2019-12-18 17:23:10');
INSERT INTO `iriscms_log` VALUES (471, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:23:10');
INSERT INTO `iriscms_log` VALUES (472, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576660990749', 1, 'admin', '::1', '2019-12-18 17:23:10');
INSERT INTO `iriscms_log` VALUES (473, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:23:10');
INSERT INTO `iriscms_log` VALUES (474, 'user', 'list', '/b/user/list?menuid=56&&_=1576660993210', 1, 'admin', '::1', '2019-12-18 17:23:13');
INSERT INTO `iriscms_log` VALUES (475, 'user', 'list', '/b/user/list?menuid=56&&_=1576661073587', 1, 'admin', '::1', '2019-12-18 17:24:33');
INSERT INTO `iriscms_log` VALUES (476, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:24:33');
INSERT INTO `iriscms_log` VALUES (477, 'wechat', 'userinfo', '/b/wechat/userinfo?menuid=59&&_=1576661077778', 1, 'admin', '::1', '2019-12-18 17:24:37');
INSERT INTO `iriscms_log` VALUES (478, 'user', 'list', '/b/user/list?menuid=56&&_=1576661083050', 1, 'admin', '::1', '2019-12-18 17:24:43');
INSERT INTO `iriscms_log` VALUES (479, 'wechat', 'userinfo', '/b/wechat/userinfo?menuid=59&&_=1576661165799', 1, 'admin', '::1', '2019-12-18 17:26:05');
INSERT INTO `iriscms_log` VALUES (480, 'user', 'list', '/b/user/list?menuid=56&&_=1576661168701', 1, 'admin', '::1', '2019-12-18 17:26:08');
INSERT INTO `iriscms_log` VALUES (481, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:26:08');
INSERT INTO `iriscms_log` VALUES (482, 'user', 'list', '/b/user/list?menuid=56&&_=1576661222465', 1, 'admin', '::1', '2019-12-18 17:27:02');
INSERT INTO `iriscms_log` VALUES (483, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:27:02');
INSERT INTO `iriscms_log` VALUES (484, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 17:27:04');
INSERT INTO `iriscms_log` VALUES (485, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 17:29:14');
INSERT INTO `iriscms_log` VALUES (486, 'user', 'list', '/b/user/list?menuid=56&&_=1576661372693', 1, 'admin', '::1', '2019-12-18 17:29:32');
INSERT INTO `iriscms_log` VALUES (487, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:29:32');
INSERT INTO `iriscms_log` VALUES (488, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 17:29:33');
INSERT INTO `iriscms_log` VALUES (489, 'user', 'list', '/b/user/list?menuid=56&&_=1576661402186', 1, 'admin', '::1', '2019-12-18 17:30:02');
INSERT INTO `iriscms_log` VALUES (490, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:30:02');
INSERT INTO `iriscms_log` VALUES (491, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 17:30:03');
INSERT INTO `iriscms_log` VALUES (492, 'user', 'list', '/b/user/list?menuid=56&&_=1576661571834', 1, 'admin', '::1', '2019-12-18 17:32:51');
INSERT INTO `iriscms_log` VALUES (493, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:32:52');
INSERT INTO `iriscms_log` VALUES (494, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 17:32:53');
INSERT INTO `iriscms_log` VALUES (495, 'user', 'list', '/b/user/list?menuid=56&&_=1576661705664', 1, 'admin', '::1', '2019-12-18 17:35:05');
INSERT INTO `iriscms_log` VALUES (496, 'user', 'list', '/b/user/list?menuid=56&&_=1576661710901', 1, 'admin', '::1', '2019-12-18 17:35:10');
INSERT INTO `iriscms_log` VALUES (497, 'user', 'list', '/b/user/list?menuid=56&&_=1576661772507', 1, 'admin', '::1', '2019-12-18 17:36:12');
INSERT INTO `iriscms_log` VALUES (498, 'user', 'list', '/b/user/list?menuid=56&&_=1576661952969', 1, 'admin', '::1', '2019-12-18 17:39:12');
INSERT INTO `iriscms_log` VALUES (499, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:39:13');
INSERT INTO `iriscms_log` VALUES (500, 'user', 'list', '/b/user/list?menuid=56&&_=1576661955539', 1, 'admin', '::1', '2019-12-18 17:39:15');
INSERT INTO `iriscms_log` VALUES (501, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 17:39:16');
INSERT INTO `iriscms_log` VALUES (502, 'user', 'edit', '/b/user/edit?id=id', 1, 'admin', '::1', '2019-12-18 17:50:04');
INSERT INTO `iriscms_log` VALUES (503, 'user', 'list', '/b/user/list?menuid=56&&_=1576662934373', 1, 'admin', '::1', '2019-12-18 17:55:34');
INSERT INTO `iriscms_log` VALUES (504, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 17:55:34');
INSERT INTO `iriscms_log` VALUES (505, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 17:55:36');
INSERT INTO `iriscms_log` VALUES (506, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 17:55:53');
INSERT INTO `iriscms_log` VALUES (507, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:03:55');
INSERT INTO `iriscms_log` VALUES (508, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:04:06');
INSERT INTO `iriscms_log` VALUES (509, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:04:07');
INSERT INTO `iriscms_log` VALUES (510, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:04:07');
INSERT INTO `iriscms_log` VALUES (511, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:04:07');
INSERT INTO `iriscms_log` VALUES (512, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:04:07');
INSERT INTO `iriscms_log` VALUES (513, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:04:08');
INSERT INTO `iriscms_log` VALUES (514, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:04:51');
INSERT INTO `iriscms_log` VALUES (515, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:06:31');
INSERT INTO `iriscms_log` VALUES (516, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:06:49');
INSERT INTO `iriscms_log` VALUES (517, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 18:07:13');
INSERT INTO `iriscms_log` VALUES (518, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:07:30');
INSERT INTO `iriscms_log` VALUES (519, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 18:08:41');
INSERT INTO `iriscms_log` VALUES (520, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:08:44');
INSERT INTO `iriscms_log` VALUES (521, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:08:48');
INSERT INTO `iriscms_log` VALUES (522, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:08:51');
INSERT INTO `iriscms_log` VALUES (523, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:09:27');
INSERT INTO `iriscms_log` VALUES (524, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 18:10:37');
INSERT INTO `iriscms_log` VALUES (525, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:10:41');
INSERT INTO `iriscms_log` VALUES (526, 'user', 'list', '/b/user/list?menuid=56&&_=1576663847414', 1, 'admin', '::1', '2019-12-18 18:10:47');
INSERT INTO `iriscms_log` VALUES (527, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 18:10:47');
INSERT INTO `iriscms_log` VALUES (528, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 18:10:49');
INSERT INTO `iriscms_log` VALUES (529, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:10:52');
INSERT INTO `iriscms_log` VALUES (530, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:11:05');
INSERT INTO `iriscms_log` VALUES (531, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:11:12');
INSERT INTO `iriscms_log` VALUES (532, 'user', 'list', '/b/user/list?menuid=56&&_=1576663912186', 1, 'admin', '::1', '2019-12-18 18:11:52');
INSERT INTO `iriscms_log` VALUES (533, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-18 18:11:52');
INSERT INTO `iriscms_log` VALUES (534, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 18:11:53');
INSERT INTO `iriscms_log` VALUES (535, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:11:58');
INSERT INTO `iriscms_log` VALUES (536, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:12:01');
INSERT INTO `iriscms_log` VALUES (537, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 18:12:32');
INSERT INTO `iriscms_log` VALUES (538, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:12:35');
INSERT INTO `iriscms_log` VALUES (539, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:12:58');
INSERT INTO `iriscms_log` VALUES (540, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 18:13:00');
INSERT INTO `iriscms_log` VALUES (541, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:13:04');
INSERT INTO `iriscms_log` VALUES (542, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 18:36:42');
INSERT INTO `iriscms_log` VALUES (543, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:36:44');
INSERT INTO `iriscms_log` VALUES (544, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:36:49');
INSERT INTO `iriscms_log` VALUES (545, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:36:57');
INSERT INTO `iriscms_log` VALUES (546, 'user', 'info', '/b/user/info?id=1', 1, 'admin', '::1', '2019-12-18 18:36:59');
INSERT INTO `iriscms_log` VALUES (547, 'user', 'edit', '/b/user/edit?id=1', 1, 'admin', '::1', '2019-12-18 18:44:29');
INSERT INTO `iriscms_log` VALUES (548, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576725212189', 11, 'test', '::1', '2019-12-19 11:13:32');
INSERT INTO `iriscms_log` VALUES (549, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-19 11:13:32');
INSERT INTO `iriscms_log` VALUES (550, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576725213083', 11, 'test', '::1', '2019-12-19 11:13:33');
INSERT INTO `iriscms_log` VALUES (551, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-19 11:13:33');
INSERT INTO `iriscms_log` VALUES (552, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576725214634', 11, 'test', '::1', '2019-12-19 11:13:34');
INSERT INTO `iriscms_log` VALUES (553, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576725226329', 11, 'test', '::1', '2019-12-19 11:13:46');
INSERT INTO `iriscms_log` VALUES (554, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-19 11:13:46');
INSERT INTO `iriscms_log` VALUES (555, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-19 11:13:47');
INSERT INTO `iriscms_log` VALUES (556, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-19 11:13:49');
INSERT INTO `iriscms_log` VALUES (557, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576725231739', 11, 'test', '::1', '2019-12-19 11:13:51');
INSERT INTO `iriscms_log` VALUES (558, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-19 11:13:51');
INSERT INTO `iriscms_log` VALUES (559, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576725330796', 11, 'test', '::1', '2019-12-19 11:15:30');
INSERT INTO `iriscms_log` VALUES (560, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-19 11:15:30');
INSERT INTO `iriscms_log` VALUES (561, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576725332130', 11, 'test', '::1', '2019-12-19 11:15:32');
INSERT INTO `iriscms_log` VALUES (562, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-19 11:15:32');
INSERT INTO `iriscms_log` VALUES (563, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576725341158', 11, 'test', '::1', '2019-12-19 11:15:41');
INSERT INTO `iriscms_log` VALUES (564, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-19 11:15:41');
INSERT INTO `iriscms_log` VALUES (565, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576725374735', 11, 'test', '::1', '2019-12-19 11:16:14');
INSERT INTO `iriscms_log` VALUES (566, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-19 11:16:14');
INSERT INTO `iriscms_log` VALUES (567, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576725376018', 11, 'test', '::1', '2019-12-19 11:16:16');
INSERT INTO `iriscms_log` VALUES (568, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-19 11:16:16');
INSERT INTO `iriscms_log` VALUES (569, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576725378527', 11, 'test', '::1', '2019-12-19 11:16:18');
INSERT INTO `iriscms_log` VALUES (570, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576725379971', 11, 'test', '::1', '2019-12-19 11:16:20');
INSERT INTO `iriscms_log` VALUES (571, 'system', 'log-delete', '/b/system/log-delete', 11, 'test', '::1', '2019-12-19 11:16:22');
INSERT INTO `iriscms_log` VALUES (572, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1576725448433', 11, 'test', '::1', '2019-12-19 11:17:28');
INSERT INTO `iriscms_log` VALUES (573, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 11, 'test', '::1', '2019-12-19 11:17:29');
INSERT INTO `iriscms_log` VALUES (574, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576725455432', 11, 'test', '::1', '2019-12-19 11:17:35');
INSERT INTO `iriscms_log` VALUES (575, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-19 11:17:35');
INSERT INTO `iriscms_log` VALUES (576, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576727598372', 11, 'test', '::1', '2019-12-19 11:53:18');
INSERT INTO `iriscms_log` VALUES (577, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 11, 'test', '::1', '2019-12-19 11:53:18');
INSERT INTO `iriscms_log` VALUES (578, 'user', 'list', '/b/user/list?menuid=56&&_=1576727630592', 1, 'admin', '::1', '2019-12-19 11:53:50');
INSERT INTO `iriscms_log` VALUES (579, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-19 11:53:51');
INSERT INTO `iriscms_log` VALUES (580, 'user', 'list', '/b/user/list?menuid=56&&_=1576727637169', 1, 'admin', '::1', '2019-12-19 11:53:57');
INSERT INTO `iriscms_log` VALUES (581, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576727639037', 1, 'admin', '::1', '2019-12-19 11:53:59');
INSERT INTO `iriscms_log` VALUES (582, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-19 11:53:59');
INSERT INTO `iriscms_log` VALUES (583, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576727639661', 1, 'admin', '::1', '2019-12-19 11:53:59');
INSERT INTO `iriscms_log` VALUES (584, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-19 11:53:59');
INSERT INTO `iriscms_log` VALUES (585, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576727640229', 1, 'admin', '::1', '2019-12-19 11:54:00');
INSERT INTO `iriscms_log` VALUES (586, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576727640814', 1, 'admin', '::1', '2019-12-19 11:54:00');
INSERT INTO `iriscms_log` VALUES (587, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576727641296', 1, 'admin', '::1', '2019-12-19 11:54:01');
INSERT INTO `iriscms_log` VALUES (588, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576727642100', 1, 'admin', '::1', '2019-12-19 11:54:02');
INSERT INTO `iriscms_log` VALUES (589, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1576727657251', 1, 'admin', '::1', '2019-12-19 11:54:17');
INSERT INTO `iriscms_log` VALUES (590, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-19 11:54:17');
INSERT INTO `iriscms_log` VALUES (591, 'user', 'list', '/b/user/list?menuid=56&&_=1576727660111', 1, 'admin', '::1', '2019-12-19 11:54:20');
INSERT INTO `iriscms_log` VALUES (592, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-19 11:54:20');
INSERT INTO `iriscms_log` VALUES (593, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1576727661435', 1, 'admin', '::1', '2019-12-19 11:54:21');
INSERT INTO `iriscms_log` VALUES (594, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-19 11:54:21');
INSERT INTO `iriscms_log` VALUES (595, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1576727662884', 1, 'admin', '::1', '2019-12-19 11:54:22');
INSERT INTO `iriscms_log` VALUES (596, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-19 11:54:23');
INSERT INTO `iriscms_log` VALUES (597, 'category', 'list', '/b/category/list?menuid=36&&_=1576729574844', 1, 'admin', '::1', '2019-12-19 12:26:14');
INSERT INTO `iriscms_log` VALUES (598, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-19 12:26:15');
INSERT INTO `iriscms_log` VALUES (599, 'user', 'list', '/b/user/list?menuid=56&&_=1577354021893', 1, 'admin', '::1', '2019-12-26 17:53:41');
INSERT INTO `iriscms_log` VALUES (600, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:53:41');
INSERT INTO `iriscms_log` VALUES (601, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1577354023829', 1, 'admin', '::1', '2019-12-26 17:53:43');
INSERT INTO `iriscms_log` VALUES (602, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:53:43');
INSERT INTO `iriscms_log` VALUES (603, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1577354024548', 1, 'admin', '::1', '2019-12-26 17:53:44');
INSERT INTO `iriscms_log` VALUES (604, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:53:44');
INSERT INTO `iriscms_log` VALUES (605, 'setting', 'site', '/b/setting/site?menuid=10&&_=1577354027099', 1, 'admin', '::1', '2019-12-26 17:53:47');
INSERT INTO `iriscms_log` VALUES (606, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-26 17:53:47');
INSERT INTO `iriscms_log` VALUES (607, 'wechat', 'userinfo', '/b/wechat/userinfo?menuid=59&&_=1577354029336', 1, 'admin', '::1', '2019-12-26 17:53:49');
INSERT INTO `iriscms_log` VALUES (608, 'category', 'list', '/b/category/list?menuid=36&&_=1577354036573', 1, 'admin', '::1', '2019-12-26 17:53:56');
INSERT INTO `iriscms_log` VALUES (609, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-26 17:53:56');
INSERT INTO `iriscms_log` VALUES (610, 'content', 'index', '/b/content/index?menuid=35&&_=1577354037417', 1, 'admin', '::1', '2019-12-26 17:53:57');
INSERT INTO `iriscms_log` VALUES (611, 'content', 'right', '/b/content/right?_=1577354037436', 1, 'admin', '::1', '2019-12-26 17:53:57');
INSERT INTO `iriscms_log` VALUES (612, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 17:53:57');
INSERT INTO `iriscms_log` VALUES (613, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 17:53:57');
INSERT INTO `iriscms_log` VALUES (614, 'content', 'news-list', '/b/content/news-list?catid=38&_=1577354041055', 1, 'admin', '::1', '2019-12-26 17:54:01');
INSERT INTO `iriscms_log` VALUES (615, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=38&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:01');
INSERT INTO `iriscms_log` VALUES (616, 'content', 'news-list', '/b/content/news-list?catid=37&_=1577354042129', 1, 'admin', '::1', '2019-12-26 17:54:02');
INSERT INTO `iriscms_log` VALUES (617, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=37&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:02');
INSERT INTO `iriscms_log` VALUES (618, 'content', 'news-list', '/b/content/news-list?catid=36&_=1577354043071', 1, 'admin', '::1', '2019-12-26 17:54:03');
INSERT INTO `iriscms_log` VALUES (619, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=36&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:03');
INSERT INTO `iriscms_log` VALUES (620, 'content', 'news-list', '/b/content/news-list?catid=35&_=1577354043887', 1, 'admin', '::1', '2019-12-26 17:54:03');
INSERT INTO `iriscms_log` VALUES (621, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=35&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:03');
INSERT INTO `iriscms_log` VALUES (622, 'content', 'news-list', '/b/content/news-list?catid=34&_=1577354044630', 1, 'admin', '::1', '2019-12-26 17:54:04');
INSERT INTO `iriscms_log` VALUES (623, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=34&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:04');
INSERT INTO `iriscms_log` VALUES (624, 'content', 'news-list', '/b/content/news-list?catid=32&_=1577354045476', 1, 'admin', '::1', '2019-12-26 17:54:05');
INSERT INTO `iriscms_log` VALUES (625, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=32&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:05');
INSERT INTO `iriscms_log` VALUES (626, 'content', 'news-list', '/b/content/news-list?catid=31&_=1577354046172', 1, 'admin', '::1', '2019-12-26 17:54:06');
INSERT INTO `iriscms_log` VALUES (627, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=31&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:06');
INSERT INTO `iriscms_log` VALUES (628, 'content', 'news-list', '/b/content/news-list?catid=33&_=1577354046802', 1, 'admin', '::1', '2019-12-26 17:54:06');
INSERT INTO `iriscms_log` VALUES (629, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=33&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:06');
INSERT INTO `iriscms_log` VALUES (630, 'content', 'news-list', '/b/content/news-list?catid=32&_=1577354047623', 1, 'admin', '::1', '2019-12-26 17:54:07');
INSERT INTO `iriscms_log` VALUES (631, 'content', 'news-list', '/b/content/news-list?catid=34&_=1577354048471', 1, 'admin', '::1', '2019-12-26 17:54:08');
INSERT INTO `iriscms_log` VALUES (632, 'content', 'news-list', '/b/content/news-list?catid=35&_=1577354049143', 1, 'admin', '::1', '2019-12-26 17:54:09');
INSERT INTO `iriscms_log` VALUES (633, 'content', 'news-list', '/b/content/news-list?catid=36&_=1577354050686', 1, 'admin', '::1', '2019-12-26 17:54:10');
INSERT INTO `iriscms_log` VALUES (634, 'category', 'list', '/b/category/list?menuid=36&&_=1577354085363', 1, 'admin', '::1', '2019-12-26 17:54:45');
INSERT INTO `iriscms_log` VALUES (635, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-26 17:54:45');
INSERT INTO `iriscms_log` VALUES (636, 'content', 'index', '/b/content/index?menuid=35&&_=1577354086340', 1, 'admin', '::1', '2019-12-26 17:54:46');
INSERT INTO `iriscms_log` VALUES (637, 'content', 'right', '/b/content/right?_=1577354086360', 1, 'admin', '::1', '2019-12-26 17:54:46');
INSERT INTO `iriscms_log` VALUES (638, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 17:54:46');
INSERT INTO `iriscms_log` VALUES (639, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 17:54:46');
INSERT INTO `iriscms_log` VALUES (640, 'category', 'list', '/b/category/list?menuid=36&&_=1577354087243', 1, 'admin', '::1', '2019-12-26 17:54:47');
INSERT INTO `iriscms_log` VALUES (641, 'content', 'index', '/b/content/index?menuid=35&&_=1577354087971', 1, 'admin', '::1', '2019-12-26 17:54:47');
INSERT INTO `iriscms_log` VALUES (642, 'content', 'right', '/b/content/right?_=1577354087991', 1, 'admin', '::1', '2019-12-26 17:54:48');
INSERT INTO `iriscms_log` VALUES (643, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 17:54:48');
INSERT INTO `iriscms_log` VALUES (644, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 17:54:48');
INSERT INTO `iriscms_log` VALUES (645, 'category', 'list', '/b/category/list?menuid=36&&_=1577354089027', 1, 'admin', '::1', '2019-12-26 17:54:49');
INSERT INTO `iriscms_log` VALUES (646, 'content', 'index', '/b/content/index?menuid=35&&_=1577354091931', 1, 'admin', '::1', '2019-12-26 17:54:51');
INSERT INTO `iriscms_log` VALUES (647, 'content', 'right', '/b/content/right?_=1577354091951', 1, 'admin', '::1', '2019-12-26 17:54:51');
INSERT INTO `iriscms_log` VALUES (648, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 17:54:51');
INSERT INTO `iriscms_log` VALUES (649, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 17:54:51');
INSERT INTO `iriscms_log` VALUES (650, 'content', 'news-list', '/b/content/news-list?catid=38&_=1577354093372', 1, 'admin', '::1', '2019-12-26 17:54:53');
INSERT INTO `iriscms_log` VALUES (651, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=38&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:53');
INSERT INTO `iriscms_log` VALUES (652, 'content', 'news-list', '/b/content/news-list?catid=35&_=1577354094757', 1, 'admin', '::1', '2019-12-26 17:54:54');
INSERT INTO `iriscms_log` VALUES (653, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=35&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:54');
INSERT INTO `iriscms_log` VALUES (654, 'content', 'news-list', '/b/content/news-list?catid=34&_=1577354095635', 1, 'admin', '::1', '2019-12-26 17:54:55');
INSERT INTO `iriscms_log` VALUES (655, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=34&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:54:55');
INSERT INTO `iriscms_log` VALUES (656, 'setting', 'site', '/b/setting/site?menuid=10&&_=1577354328045', 1, 'admin', '::1', '2019-12-26 17:58:48');
INSERT INTO `iriscms_log` VALUES (657, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-26 17:58:48');
INSERT INTO `iriscms_log` VALUES (658, 'user', 'list', '/b/user/list?menuid=56&&_=1577354339897', 1, 'admin', '::1', '2019-12-26 17:58:59');
INSERT INTO `iriscms_log` VALUES (659, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:58:59');
INSERT INTO `iriscms_log` VALUES (660, 'setting', 'site', '/b/setting/site?menuid=10&&_=1577354371390', 1, 'admin', '::1', '2019-12-26 17:59:31');
INSERT INTO `iriscms_log` VALUES (661, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-26 17:59:31');
INSERT INTO `iriscms_log` VALUES (662, 'wechat', 'userinfo', '/b/wechat/userinfo?menuid=59&&_=1577354374126', 1, 'admin', '::1', '2019-12-26 17:59:34');
INSERT INTO `iriscms_log` VALUES (663, 'user', 'list', '/b/user/list?menuid=56&&_=1577354376217', 1, 'admin', '::1', '2019-12-26 17:59:36');
INSERT INTO `iriscms_log` VALUES (664, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:59:36');
INSERT INTO `iriscms_log` VALUES (665, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1577354377375', 1, 'admin', '::1', '2019-12-26 17:59:37');
INSERT INTO `iriscms_log` VALUES (666, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:59:37');
INSERT INTO `iriscms_log` VALUES (667, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1577354377967', 1, 'admin', '::1', '2019-12-26 17:59:37');
INSERT INTO `iriscms_log` VALUES (668, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 17:59:38');
INSERT INTO `iriscms_log` VALUES (669, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1577354378377', 1, 'admin', '::1', '2019-12-26 17:59:38');
INSERT INTO `iriscms_log` VALUES (670, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1577354378849', 1, 'admin', '::1', '2019-12-26 17:59:38');
INSERT INTO `iriscms_log` VALUES (671, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1577354379232', 1, 'admin', '::1', '2019-12-26 17:59:39');
INSERT INTO `iriscms_log` VALUES (672, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1577354379560', 1, 'admin', '::1', '2019-12-26 17:59:39');
INSERT INTO `iriscms_log` VALUES (673, 'user', 'list', '/b/user/list?menuid=56&&_=1577354427719', 1, 'admin', '::1', '2019-12-26 18:00:27');
INSERT INTO `iriscms_log` VALUES (674, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 18:00:27');
INSERT INTO `iriscms_log` VALUES (675, 'setting', 'site', '/b/setting/site?menuid=10&&_=1577354428595', 1, 'admin', '::1', '2019-12-26 18:00:28');
INSERT INTO `iriscms_log` VALUES (676, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-26 18:00:28');
INSERT INTO `iriscms_log` VALUES (677, 'wechat', 'userinfo', '/b/wechat/userinfo?menuid=59&&_=1577354429872', 1, 'admin', '::1', '2019-12-26 18:00:29');
INSERT INTO `iriscms_log` VALUES (678, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1577354794313', 1, 'admin', '::1', '2019-12-26 18:06:34');
INSERT INTO `iriscms_log` VALUES (679, 'system', 'menulist', '/b/system/menulist?grid=treegrid', 1, 'admin', '::1', '2019-12-26 18:06:34');
INSERT INTO `iriscms_log` VALUES (680, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1577354795494', 1, 'admin', '::1', '2019-12-26 18:06:35');
INSERT INTO `iriscms_log` VALUES (681, 'system', 'loglist', '/b/system/loglist?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 18:06:35');
INSERT INTO `iriscms_log` VALUES (682, 'system', 'menulist', '/b/system/menulist?menuid=16&&_=1577354796655', 1, 'admin', '::1', '2019-12-26 18:06:36');
INSERT INTO `iriscms_log` VALUES (683, 'wechat', 'userinfo', '/b/wechat/userinfo?menuid=59&&_=1577354821681', 1, 'admin', '::1', '2019-12-26 18:07:01');
INSERT INTO `iriscms_log` VALUES (684, 'model', 'list', '/b/model/list?menuid=62&&_=1577355016724', 1, 'admin', '::1', '2019-12-26 18:10:16');
INSERT INTO `iriscms_log` VALUES (685, 'model', 'list', '/b/model/list?menuid=62&&_=1577355016724', 1, 'admin', '::1', '2019-12-26 18:11:11');
INSERT INTO `iriscms_log` VALUES (686, 'model', 'list', '/b/model/list?menuid=62&&_=1577355546995', 1, 'admin', '::1', '2019-12-26 18:19:07');
INSERT INTO `iriscms_log` VALUES (687, 'model', 'list', '/b/model/list?menuid=62&&_=1577355625299', 1, 'admin', '::1', '2019-12-26 18:20:25');
INSERT INTO `iriscms_log` VALUES (688, 'model', 'list', '/b/model/list?menuid=62&&_=1577355670360', 1, 'admin', '::1', '2019-12-26 18:21:10');
INSERT INTO `iriscms_log` VALUES (689, 'model', 'list', '/b/model/list?menuid=62&&_=1577355670360', 1, 'admin', '::1', '2019-12-26 18:22:20');
INSERT INTO `iriscms_log` VALUES (690, 'model', 'list', '/b/model/list?menuid=62&&_=1577362012455', 1, 'admin', '::1', '2019-12-26 20:06:52');
INSERT INTO `iriscms_log` VALUES (691, 'model', 'list', '/b/model/list?menuid=62&&_=1577362308478', 1, 'admin', '::1', '2019-12-26 20:11:48');
INSERT INTO `iriscms_log` VALUES (692, 'model', 'list', '/b/model/list?menuid=62&&_=1577362398342', 1, 'admin', '::1', '2019-12-26 20:13:18');
INSERT INTO `iriscms_log` VALUES (693, 'user', 'list', '/b/user/list?menuid=56&&_=1577362403335', 1, 'admin', '::1', '2019-12-26 20:13:23');
INSERT INTO `iriscms_log` VALUES (694, 'user', 'list', '/b/user/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:13:23');
INSERT INTO `iriscms_log` VALUES (695, 'model', 'list', '/b/model/list?menuid=62&&_=1577362407638', 1, 'admin', '::1', '2019-12-26 20:13:27');
INSERT INTO `iriscms_log` VALUES (696, 'model', 'list', '/b/model/list?menuid=62&&_=1577362434599', 1, 'admin', '::1', '2019-12-26 20:13:54');
INSERT INTO `iriscms_log` VALUES (697, 'model', 'list', '/b/model/list?menuid=62&&_=1577362559428', 1, 'admin', '::1', '2019-12-26 20:15:59');
INSERT INTO `iriscms_log` VALUES (698, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:15:59');
INSERT INTO `iriscms_log` VALUES (699, 'model', 'list', '/b/model/list?menuid=62&&_=1577362665473', 1, 'admin', '::1', '2019-12-26 20:17:45');
INSERT INTO `iriscms_log` VALUES (700, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:17:45');
INSERT INTO `iriscms_log` VALUES (701, 'model', 'list', '/b/model/list?page=2&rows=10', 1, 'admin', '::1', '2019-12-26 20:17:47');
INSERT INTO `iriscms_log` VALUES (702, 'model', 'list', '/b/model/list?page=3&rows=10', 1, 'admin', '::1', '2019-12-26 20:17:47');
INSERT INTO `iriscms_log` VALUES (703, 'model', 'list', '/b/model/list?page=1&rows=40', 1, 'admin', '::1', '2019-12-26 20:18:45');
INSERT INTO `iriscms_log` VALUES (704, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:18:46');
INSERT INTO `iriscms_log` VALUES (705, 'model', 'list', '/b/model/list?menuid=62&&_=1577362794145', 1, 'admin', '::1', '2019-12-26 20:19:54');
INSERT INTO `iriscms_log` VALUES (706, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:19:54');
INSERT INTO `iriscms_log` VALUES (707, 'model', 'list', '/b/model/list?page=2&rows=10', 1, 'admin', '::1', '2019-12-26 20:19:59');
INSERT INTO `iriscms_log` VALUES (708, 'model', 'list', '/b/model/list?page=3&rows=10', 1, 'admin', '::1', '2019-12-26 20:19:59');
INSERT INTO `iriscms_log` VALUES (709, 'model', 'list', '/b/model/list?menuid=62&&_=1577362834322', 1, 'admin', '::1', '2019-12-26 20:20:34');
INSERT INTO `iriscms_log` VALUES (710, 'model', 'list', '/b/model/list?menuid=62&&_=1577362881218', 1, 'admin', '::1', '2019-12-26 20:21:21');
INSERT INTO `iriscms_log` VALUES (711, 'model', 'list', '/b/model/list?menuid=62&&_=1577362935230', 1, 'admin', '127.0.0.1', '2019-12-26 20:22:16');
INSERT INTO `iriscms_log` VALUES (712, 'model', 'list', '/b/model/list?menuid=62&&_=1577362960587', 1, 'admin', '::1', '2019-12-26 20:22:40');
INSERT INTO `iriscms_log` VALUES (713, 'model', 'list', '/b/model/list?menuid=62&&_=1577363078943', 1, 'admin', '::1', '2019-12-26 20:24:38');
INSERT INTO `iriscms_log` VALUES (714, 'model', 'list', '/b/model/list?menuid=62&&_=1577363221102', 1, 'admin', '::1', '2019-12-26 20:27:01');
INSERT INTO `iriscms_log` VALUES (715, 'wechat', 'userinfo', '/b/wechat/userinfo?menuid=59&&_=1577363222289', 1, 'admin', '::1', '2019-12-26 20:27:02');
INSERT INTO `iriscms_log` VALUES (716, 'setting', 'site', '/b/setting/site?menuid=10&&_=1577363225374', 1, 'admin', '::1', '2019-12-26 20:27:05');
INSERT INTO `iriscms_log` VALUES (717, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-26 20:27:05');
INSERT INTO `iriscms_log` VALUES (718, 'model', 'list', '/b/model/list?menuid=62&&_=1577363226647', 1, 'admin', '::1', '2019-12-26 20:27:06');
INSERT INTO `iriscms_log` VALUES (719, 'model', 'list', '/b/model/list?menuid=62&&_=1577363351221', 1, 'admin', '::1', '2019-12-26 20:29:11');
INSERT INTO `iriscms_log` VALUES (720, 'model', 'list', '/b/model/list?menuid=62&&_=1577363357582', 1, 'admin', '::1', '2019-12-26 20:29:17');
INSERT INTO `iriscms_log` VALUES (721, 'model', 'list', '/b/model/list?menuid=62&&_=1577363384625', 1, 'admin', '::1', '2019-12-26 20:29:44');
INSERT INTO `iriscms_log` VALUES (722, 'model', 'list', '/b/model/list?menuid=62&&_=1577363441763', 1, 'admin', '::1', '2019-12-26 20:30:41');
INSERT INTO `iriscms_log` VALUES (723, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:30:41');
INSERT INTO `iriscms_log` VALUES (724, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:32:55');
INSERT INTO `iriscms_log` VALUES (725, 'model', 'list', '/b/model/list?menuid=62&&_=1577363595554', 1, 'admin', '::1', '2019-12-26 20:33:15');
INSERT INTO `iriscms_log` VALUES (726, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:33:15');
INSERT INTO `iriscms_log` VALUES (727, 'category', 'list', '/b/category/list?menuid=36&&_=1577363609812', 1, 'admin', '::1', '2019-12-26 20:33:29');
INSERT INTO `iriscms_log` VALUES (728, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-26 20:33:29');
INSERT INTO `iriscms_log` VALUES (729, 'content', 'index', '/b/content/index?menuid=35&&_=1577363610701', 1, 'admin', '::1', '2019-12-26 20:33:30');
INSERT INTO `iriscms_log` VALUES (730, 'content', 'right', '/b/content/right?_=1577363610753', 1, 'admin', '::1', '2019-12-26 20:33:30');
INSERT INTO `iriscms_log` VALUES (731, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 20:33:30');
INSERT INTO `iriscms_log` VALUES (732, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 20:33:30');
INSERT INTO `iriscms_log` VALUES (733, 'content', 'news-list', '/b/content/news-list?catid=38&_=1577363612077', 1, 'admin', '::1', '2019-12-26 20:33:32');
INSERT INTO `iriscms_log` VALUES (734, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=38&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:33:32');
INSERT INTO `iriscms_log` VALUES (735, 'content', 'news-list', '/b/content/news-list?catid=37&_=1577363613530', 1, 'admin', '::1', '2019-12-26 20:33:33');
INSERT INTO `iriscms_log` VALUES (736, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=37&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:33:33');
INSERT INTO `iriscms_log` VALUES (737, 'content', 'news-list', '/b/content/news-list?catid=36&_=1577363614695', 1, 'admin', '::1', '2019-12-26 20:33:34');
INSERT INTO `iriscms_log` VALUES (738, 'content', 'news-list', '/b/content/news-list?grid=datagrid&catid=36&page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:33:34');
INSERT INTO `iriscms_log` VALUES (739, 'content', 'add', '/b/content/add?catid=36', 1, 'admin', '::1', '2019-12-26 20:33:36');
INSERT INTO `iriscms_log` VALUES (740, 'model', 'list', '/b/model/list?menuid=62&&_=1577363692714', 1, 'admin', '::1', '2019-12-26 20:34:52');
INSERT INTO `iriscms_log` VALUES (741, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:34:52');
INSERT INTO `iriscms_log` VALUES (742, 'content', 'add', '/b/content/add?catid=', 1, 'admin', '::1', '2019-12-26 20:34:55');
INSERT INTO `iriscms_log` VALUES (743, 'model', 'list', '/b/model/list?menuid=62&&_=1577363724857', 1, 'admin', '::1', '2019-12-26 20:35:24');
INSERT INTO `iriscms_log` VALUES (744, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:35:24');
INSERT INTO `iriscms_log` VALUES (745, 'model', 'list', '/b/model/list?menuid=62&&_=1577363844372', 1, 'admin', '::1', '2019-12-26 20:37:24');
INSERT INTO `iriscms_log` VALUES (746, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:37:24');
INSERT INTO `iriscms_log` VALUES (747, 'model', 'list', '/b/model/list?menuid=62&&_=1577363882178', 1, 'admin', '::1', '2019-12-26 20:38:02');
INSERT INTO `iriscms_log` VALUES (748, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:38:02');
INSERT INTO `iriscms_log` VALUES (749, 'model', 'list', '/b/model/list?menuid=62&&_=1577363929812', 1, 'admin', '::1', '2019-12-26 20:38:49');
INSERT INTO `iriscms_log` VALUES (750, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:38:49');
INSERT INTO `iriscms_log` VALUES (751, 'model', 'list', '/b/model/list?menuid=62&&_=1577364013844', 1, 'admin', '::1', '2019-12-26 20:40:13');
INSERT INTO `iriscms_log` VALUES (752, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:40:14');
INSERT INTO `iriscms_log` VALUES (753, 'model', 'list', '/b/model/list?menuid=62&&_=1577364045260', 1, 'admin', '::1', '2019-12-26 20:40:45');
INSERT INTO `iriscms_log` VALUES (754, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:40:45');
INSERT INTO `iriscms_log` VALUES (755, 'model', 'list', '/b/model/list?menuid=62&&_=1577364082792', 1, 'admin', '::1', '2019-12-26 20:41:22');
INSERT INTO `iriscms_log` VALUES (756, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:41:22');
INSERT INTO `iriscms_log` VALUES (757, 'model', 'list', '/b/model/list?menuid=62&&_=1577364092088', 1, 'admin', '::1', '2019-12-26 20:41:32');
INSERT INTO `iriscms_log` VALUES (758, 'wechat', 'userinfo', '/b/wechat/userinfo?menuid=59&&_=1577364147224', 1, 'admin', '::1', '2019-12-26 20:42:27');
INSERT INTO `iriscms_log` VALUES (759, 'content', 'index', '/b/content/index?menuid=35&&_=1577364154042', 1, 'admin', '::1', '2019-12-26 20:42:34');
INSERT INTO `iriscms_log` VALUES (760, 'content', 'right', '/b/content/right?_=1577364154085', 1, 'admin', '::1', '2019-12-26 20:42:34');
INSERT INTO `iriscms_log` VALUES (761, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 20:42:34');
INSERT INTO `iriscms_log` VALUES (762, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 20:42:34');
INSERT INTO `iriscms_log` VALUES (763, 'category', 'list', '/b/category/list?menuid=36&&_=1577364154964', 1, 'admin', '::1', '2019-12-26 20:42:34');
INSERT INTO `iriscms_log` VALUES (764, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-26 20:42:35');
INSERT INTO `iriscms_log` VALUES (765, 'category', 'category-add', '/b/category/category-add?parentid=0&_=1577364156415', 1, 'admin', '::1', '2019-12-26 20:42:36');
INSERT INTO `iriscms_log` VALUES (766, 'category', 'category-select', '/b/category/category-select', 1, 'admin', '::1', '2019-12-26 20:42:36');
INSERT INTO `iriscms_log` VALUES (767, 'category', 'list', '/b/category/list?menuid=36&&_=1577364264106', 1, 'admin', '::1', '2019-12-26 20:44:24');
INSERT INTO `iriscms_log` VALUES (768, 'category', 'list', '/b/category/list?grid=treegrid', 1, 'admin', '::1', '2019-12-26 20:44:24');
INSERT INTO `iriscms_log` VALUES (769, 'content', 'index', '/b/content/index?menuid=35&&_=1577364267671', 1, 'admin', '::1', '2019-12-26 20:44:27');
INSERT INTO `iriscms_log` VALUES (770, 'content', 'right', '/b/content/right?_=1577364267719', 1, 'admin', '::1', '2019-12-26 20:44:27');
INSERT INTO `iriscms_log` VALUES (771, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 20:44:27');
INSERT INTO `iriscms_log` VALUES (772, 'content', 'right', '/b/content/right', 1, 'admin', '::1', '2019-12-26 20:44:27');
INSERT INTO `iriscms_log` VALUES (773, 'model', 'list', '/b/model/list?menuid=62&&_=1577364273250', 1, 'admin', '::1', '2019-12-26 20:44:33');
INSERT INTO `iriscms_log` VALUES (774, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:44:33');
INSERT INTO `iriscms_log` VALUES (775, 'model', 'add', '/b/model/add?_=1577364274576', 1, 'admin', '::1', '2019-12-26 20:44:34');
INSERT INTO `iriscms_log` VALUES (776, 'category', 'category-select', '/b/category/category-select', 1, 'admin', '::1', '2019-12-26 20:44:34');
INSERT INTO `iriscms_log` VALUES (777, 'model', 'list', '/b/model/list?menuid=62&&_=1577364589838', 1, 'admin', '::1', '2019-12-26 20:49:49');
INSERT INTO `iriscms_log` VALUES (778, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:49:49');
INSERT INTO `iriscms_log` VALUES (779, 'model', 'add', '/b/model/add?_=1577364591388', 1, 'admin', '::1', '2019-12-26 20:49:51');
INSERT INTO `iriscms_log` VALUES (780, 'model', 'list', '/b/model/list?menuid=62&&_=1577364605649', 1, 'admin', '::1', '2019-12-26 20:50:05');
INSERT INTO `iriscms_log` VALUES (781, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:50:05');
INSERT INTO `iriscms_log` VALUES (782, 'model', 'add', '/b/model/add?_=1577364606703', 1, 'admin', '::1', '2019-12-26 20:50:06');
INSERT INTO `iriscms_log` VALUES (783, 'model', 'list', '/b/model/list?menuid=62&&_=1577364609690', 1, 'admin', '::1', '2019-12-26 20:50:09');
INSERT INTO `iriscms_log` VALUES (784, 'model', 'add', '/b/model/add?_=1577364610634', 1, 'admin', '::1', '2019-12-26 20:50:10');
INSERT INTO `iriscms_log` VALUES (785, 'model', 'list', '/b/model/list?menuid=62&&_=1577364637005', 1, 'admin', '::1', '2019-12-26 20:50:37');
INSERT INTO `iriscms_log` VALUES (786, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:50:37');
INSERT INTO `iriscms_log` VALUES (787, 'model', 'add', '/b/model/add?_=1577364639207', 1, 'admin', '::1', '2019-12-26 20:50:39');
INSERT INTO `iriscms_log` VALUES (788, 'model', 'list', '/b/model/list?menuid=62&&_=1577364643281', 1, 'admin', '::1', '2019-12-26 20:50:43');
INSERT INTO `iriscms_log` VALUES (789, 'model', 'add', '/b/model/add?_=1577364644038', 1, 'admin', '::1', '2019-12-26 20:50:44');
INSERT INTO `iriscms_log` VALUES (790, 'model', 'list', '/b/model/list?menuid=62&&_=1577364714050', 1, 'admin', '::1', '2019-12-26 20:51:54');
INSERT INTO `iriscms_log` VALUES (791, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:51:54');
INSERT INTO `iriscms_log` VALUES (792, 'model', 'add', '/b/model/add?_=1577364715276', 1, 'admin', '::1', '2019-12-26 20:51:55');
INSERT INTO `iriscms_log` VALUES (793, 'model', 'list', '/b/model/list?menuid=62&&_=1577364861804', 1, 'admin', '::1', '2019-12-26 20:54:21');
INSERT INTO `iriscms_log` VALUES (794, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:54:21');
INSERT INTO `iriscms_log` VALUES (795, 'model', 'add', '/b/model/add?_=1577364862959', 1, 'admin', '::1', '2019-12-26 20:54:22');
INSERT INTO `iriscms_log` VALUES (796, 'model', 'list', '/b/model/list?menuid=62&&_=1577364994092', 1, 'admin', '::1', '2019-12-26 20:56:34');
INSERT INTO `iriscms_log` VALUES (797, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:56:34');
INSERT INTO `iriscms_log` VALUES (798, 'model', 'add', '/b/model/add?_=1577364995179', 1, 'admin', '::1', '2019-12-26 20:56:35');
INSERT INTO `iriscms_log` VALUES (799, 'model', 'list', '/b/model/list?menuid=62&&_=1577365040260', 1, 'admin', '::1', '2019-12-26 20:57:20');
INSERT INTO `iriscms_log` VALUES (800, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:57:20');
INSERT INTO `iriscms_log` VALUES (801, 'model', 'add', '/b/model/add?_=1577365041640', 1, 'admin', '::1', '2019-12-26 20:57:21');
INSERT INTO `iriscms_log` VALUES (802, 'model', 'list', '/b/model/list?menuid=62&&_=1577365176181', 1, 'admin', '::1', '2019-12-26 20:59:36');
INSERT INTO `iriscms_log` VALUES (803, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:59:36');
INSERT INTO `iriscms_log` VALUES (804, 'model', 'add', '/b/model/add?_=1577365177420', 1, 'admin', '::1', '2019-12-26 20:59:37');
INSERT INTO `iriscms_log` VALUES (805, 'model', 'list', '/b/model/list?menuid=62&&_=1577365187221', 1, 'admin', '::1', '2019-12-26 20:59:47');
INSERT INTO `iriscms_log` VALUES (806, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 20:59:47');
INSERT INTO `iriscms_log` VALUES (807, 'model', 'add', '/b/model/add?_=1577365189292', 1, 'admin', '::1', '2019-12-26 20:59:49');
INSERT INTO `iriscms_log` VALUES (808, 'model', 'list', '/b/model/list?menuid=62&&_=1577365243036', 1, 'admin', '::1', '2019-12-26 21:00:43');
INSERT INTO `iriscms_log` VALUES (809, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:00:43');
INSERT INTO `iriscms_log` VALUES (810, 'model', 'add', '/b/model/add?_=1577365244823', 1, 'admin', '::1', '2019-12-26 21:00:44');
INSERT INTO `iriscms_log` VALUES (811, 'model', 'list', '/b/model/list?menuid=62&&_=1577365272770', 1, 'admin', '::1', '2019-12-26 21:01:15');
INSERT INTO `iriscms_log` VALUES (812, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:01:15');
INSERT INTO `iriscms_log` VALUES (813, 'model', 'add', '/b/model/add?_=1577365277321', 1, 'admin', '::1', '2019-12-26 21:01:17');
INSERT INTO `iriscms_log` VALUES (814, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:04:37');
INSERT INTO `iriscms_log` VALUES (815, 'model', 'add', '/b/model/add?_=1577365479071', 1, 'admin', '::1', '2019-12-26 21:04:39');
INSERT INTO `iriscms_log` VALUES (816, 'model', 'list', '/b/model/list?menuid=62&&_=1577365585595', 1, 'admin', '::1', '2019-12-26 21:06:25');
INSERT INTO `iriscms_log` VALUES (817, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:06:25');
INSERT INTO `iriscms_log` VALUES (818, 'model', 'add', '/b/model/add?_=1577365586486', 1, 'admin', '::1', '2019-12-26 21:06:26');
INSERT INTO `iriscms_log` VALUES (819, 'model', 'list', '/b/model/list?menuid=62&&_=1577365666476', 1, 'admin', '::1', '2019-12-26 21:07:46');
INSERT INTO `iriscms_log` VALUES (820, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:07:46');
INSERT INTO `iriscms_log` VALUES (821, 'model', 'add', '/b/model/add?_=1577365667557', 1, 'admin', '::1', '2019-12-26 21:07:47');
INSERT INTO `iriscms_log` VALUES (822, 'model', 'list', '/b/model/list?menuid=62&&_=1577365743981', 1, 'admin', '::1', '2019-12-26 21:09:03');
INSERT INTO `iriscms_log` VALUES (823, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:09:04');
INSERT INTO `iriscms_log` VALUES (824, 'model', 'add', '/b/model/add?_=1577365745238', 1, 'admin', '::1', '2019-12-26 21:09:05');
INSERT INTO `iriscms_log` VALUES (825, 'model', 'list', '/b/model/list?menuid=62&&_=1577365806063', 1, 'admin', '::1', '2019-12-26 21:10:06');
INSERT INTO `iriscms_log` VALUES (826, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:10:06');
INSERT INTO `iriscms_log` VALUES (827, 'model', 'add', '/b/model/add?_=1577365807246', 1, 'admin', '::1', '2019-12-26 21:10:07');
INSERT INTO `iriscms_log` VALUES (828, 'model', 'list', '/b/model/list?menuid=62&&_=1577365863587', 1, 'admin', '::1', '2019-12-26 21:11:03');
INSERT INTO `iriscms_log` VALUES (829, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:11:03');
INSERT INTO `iriscms_log` VALUES (830, 'model', 'add', '/b/model/add?_=1577365864935', 1, 'admin', '::1', '2019-12-26 21:11:04');
INSERT INTO `iriscms_log` VALUES (831, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:14:36');
INSERT INTO `iriscms_log` VALUES (832, 'model', 'add', '/b/model/add?_=1577366077599', 1, 'admin', '::1', '2019-12-26 21:14:37');
INSERT INTO `iriscms_log` VALUES (833, 'model', 'list', '/b/model/list?menuid=62&&_=1577366097684', 1, 'admin', '::1', '2019-12-26 21:14:57');
INSERT INTO `iriscms_log` VALUES (834, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:14:57');
INSERT INTO `iriscms_log` VALUES (835, 'model', 'add', '/b/model/add?_=1577366098798', 1, 'admin', '::1', '2019-12-26 21:14:58');
INSERT INTO `iriscms_log` VALUES (836, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:18:44');
INSERT INTO `iriscms_log` VALUES (837, 'model', 'add', '/b/model/add?_=1577366325863', 1, 'admin', '::1', '2019-12-26 21:18:45');
INSERT INTO `iriscms_log` VALUES (838, 'model', 'list', '/b/model/list?menuid=62&&_=1577366347612', 1, 'admin', '::1', '2019-12-26 21:19:07');
INSERT INTO `iriscms_log` VALUES (839, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:19:07');
INSERT INTO `iriscms_log` VALUES (840, 'model', 'add', '/b/model/add?_=1577366348616', 1, 'admin', '::1', '2019-12-26 21:19:08');
INSERT INTO `iriscms_log` VALUES (841, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:22:53');
INSERT INTO `iriscms_log` VALUES (842, 'model', 'add', '/b/model/add?_=1577366574289', 1, 'admin', '::1', '2019-12-26 21:22:54');
INSERT INTO `iriscms_log` VALUES (843, 'model', 'list', '/b/model/list?menuid=62&&_=1577366626134', 1, 'admin', '::1', '2019-12-26 21:23:46');
INSERT INTO `iriscms_log` VALUES (844, 'model', 'list', '/b/model/list?page=1&rows=10', 1, 'admin', '::1', '2019-12-26 21:23:46');
INSERT INTO `iriscms_log` VALUES (845, 'model', 'add', '/b/model/add?_=1577366627400', 1, 'admin', '::1', '2019-12-26 21:23:47');
INSERT INTO `iriscms_log` VALUES (846, 'system', 'loglist', '/b/system/loglist?menuid=15&&_=1577500832349', 1, 'admin', '::1', '2019-12-28 10:40:37');
INSERT INTO `iriscms_log` VALUES (847, 'system', 'loglist', '/b/system/loglist?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:40:37');
INSERT INTO `iriscms_log` VALUES (848, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1577500832350', 1, 'admin', '::1', '2019-12-28 10:41:21');
INSERT INTO `iriscms_log` VALUES (849, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:41:22');
INSERT INTO `iriscms_log` VALUES (850, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1577500832351', 1, 'admin', '::1', '2019-12-28 10:41:23');
INSERT INTO `iriscms_log` VALUES (851, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:41:23');
INSERT INTO `iriscms_log` VALUES (852, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1577500832352', 1, 'admin', '::1', '2019-12-28 10:41:24');
INSERT INTO `iriscms_log` VALUES (853, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:41:24');
INSERT INTO `iriscms_log` VALUES (854, 'admin', 'rolelist', '/b/admin/rolelist?menuid=13&&_=1577500832353', 1, 'admin', '::1', '2019-12-28 10:41:25');
INSERT INTO `iriscms_log` VALUES (855, 'admin', 'rolelist', '/b/admin/rolelist?grid=datagrid&page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:41:25');
INSERT INTO `iriscms_log` VALUES (856, 'admin', 'memberlist', '/b/admin/memberlist?menuid=12&&_=1577500832354', 1, 'admin', '::1', '2019-12-28 10:41:25');
INSERT INTO `iriscms_log` VALUES (857, 'admin', 'memberlist', '/b/admin/memberlist?grid=datagrid&page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:41:25');
INSERT INTO `iriscms_log` VALUES (858, 'setting', 'site', '/b/setting/site?menuid=10&&_=1577500832355', 1, 'admin', '::1', '2019-12-28 10:41:27');
INSERT INTO `iriscms_log` VALUES (859, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-28 10:41:27');
INSERT INTO `iriscms_log` VALUES (860, 'model', 'list', '/b/model/list?menuid=62&&_=1577500832356', 1, 'admin', '::1', '2019-12-28 10:41:29');
INSERT INTO `iriscms_log` VALUES (861, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:41:29');
INSERT INTO `iriscms_log` VALUES (862, 'model', 'add', '/b/model/add?_=1577500832357', 1, 'admin', '::1', '2019-12-28 10:41:30');
INSERT INTO `iriscms_log` VALUES (863, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:45:20');
INSERT INTO `iriscms_log` VALUES (864, 'model', 'add', '/b/model/add?_=1577500832359', 1, 'admin', '::1', '2019-12-28 10:45:21');
INSERT INTO `iriscms_log` VALUES (865, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:47:50');
INSERT INTO `iriscms_log` VALUES (866, 'model', 'add', '/b/model/add?_=1577500832361', 1, 'admin', '::1', '2019-12-28 10:47:51');
INSERT INTO `iriscms_log` VALUES (867, 'model', 'list', '/b/model/list?menuid=62&&_=1577500832362', 1, 'admin', '::1', '2019-12-28 10:48:29');
INSERT INTO `iriscms_log` VALUES (868, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:48:29');
INSERT INTO `iriscms_log` VALUES (869, 'model', 'add', '/b/model/add?_=1577500832363', 1, 'admin', '::1', '2019-12-28 10:48:29');
INSERT INTO `iriscms_log` VALUES (870, 'setting', 'site', '/b/setting/site?menuid=10&&_=1577501359620', 1, 'admin', '::1', '2019-12-28 10:49:23');
INSERT INTO `iriscms_log` VALUES (871, 'setting', 'site', '/b/setting/site?grid=propertygrid', 1, 'admin', '::1', '2019-12-28 10:49:23');
INSERT INTO `iriscms_log` VALUES (872, 'model', 'list', '/b/model/list?menuid=62&&_=1577501359621', 1, 'admin', '::1', '2019-12-28 10:49:25');
INSERT INTO `iriscms_log` VALUES (873, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:49:25');
INSERT INTO `iriscms_log` VALUES (874, 'model', 'add', '/b/model/add?_=1577501359622', 1, 'admin', '::1', '2019-12-28 10:49:26');
INSERT INTO `iriscms_log` VALUES (875, 'model', 'list', '/b/model/list?menuid=62&&_=1577501404665', 1, 'admin', '::1', '2019-12-28 10:50:07');
INSERT INTO `iriscms_log` VALUES (876, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:50:07');
INSERT INTO `iriscms_log` VALUES (877, 'model', 'add', '/b/model/add?_=1577501404666', 1, 'admin', '::1', '2019-12-28 10:50:08');
INSERT INTO `iriscms_log` VALUES (878, 'model', 'list', '/b/model/list?menuid=62&&_=1577501404667', 1, 'admin', '::1', '2019-12-28 10:50:57');
INSERT INTO `iriscms_log` VALUES (879, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 10:50:57');
INSERT INTO `iriscms_log` VALUES (880, 'model', 'add', '/b/model/add?_=1577501404668', 1, 'admin', '::1', '2019-12-28 10:50:58');
INSERT INTO `iriscms_log` VALUES (881, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:07:49');
INSERT INTO `iriscms_log` VALUES (882, 'model', 'add', '/b/model/add?_=1577501404670', 1, 'admin', '::1', '2019-12-28 11:07:54');
INSERT INTO `iriscms_log` VALUES (883, 'model', 'list', '/b/model/list?menuid=62&&_=1577501404671', 1, 'admin', '::1', '2019-12-28 11:09:30');
INSERT INTO `iriscms_log` VALUES (884, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:09:30');
INSERT INTO `iriscms_log` VALUES (885, 'model', 'add', '/b/model/add?_=1577501404672', 1, 'admin', '::1', '2019-12-28 11:09:34');
INSERT INTO `iriscms_log` VALUES (886, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:11:37');
INSERT INTO `iriscms_log` VALUES (887, 'model', 'add', '/b/model/add?_=1577501404674', 1, 'admin', '::1', '2019-12-28 11:11:38');
INSERT INTO `iriscms_log` VALUES (888, 'model', 'list', '/b/model/list?menuid=62&&_=1577501404675', 1, 'admin', '::1', '2019-12-28 11:12:03');
INSERT INTO `iriscms_log` VALUES (889, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:12:03');
INSERT INTO `iriscms_log` VALUES (890, 'model', 'add', '/b/model/add?_=1577501404676', 1, 'admin', '::1', '2019-12-28 11:12:04');
INSERT INTO `iriscms_log` VALUES (891, 'model', 'list', '/b/model/list?menuid=62&&_=1577501404677', 1, 'admin', '::1', '2019-12-28 11:12:51');
INSERT INTO `iriscms_log` VALUES (892, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:12:51');
INSERT INTO `iriscms_log` VALUES (893, 'model', 'add', '/b/model/add?_=1577501404678', 1, 'admin', '::1', '2019-12-28 11:12:52');
INSERT INTO `iriscms_log` VALUES (894, 'model', 'list', '/b/model/list?menuid=62&&_=1577501404679', 1, 'admin', '::1', '2019-12-28 11:14:25');
INSERT INTO `iriscms_log` VALUES (895, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:14:25');
INSERT INTO `iriscms_log` VALUES (896, 'model', 'add', '/b/model/add?_=1577501404680', 1, 'admin', '::1', '2019-12-28 11:14:26');
INSERT INTO `iriscms_log` VALUES (897, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:16:43');
INSERT INTO `iriscms_log` VALUES (898, 'model', 'add', '/b/model/add?_=1577501404682', 1, 'admin', '::1', '2019-12-28 11:16:44');
INSERT INTO `iriscms_log` VALUES (899, 'model', 'list', '/b/model/list?menuid=62&&_=1577501404683', 1, 'admin', '::1', '2019-12-28 11:17:02');
INSERT INTO `iriscms_log` VALUES (900, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:17:02');
INSERT INTO `iriscms_log` VALUES (901, 'model', 'add', '/b/model/add?_=1577501404684', 1, 'admin', '::1', '2019-12-28 11:17:02');
INSERT INTO `iriscms_log` VALUES (902, 'model', 'list', '/b/model/list?menuid=62&&_=1577501404685', 1, 'admin', '::1', '2019-12-28 11:17:25');
INSERT INTO `iriscms_log` VALUES (903, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:17:25');
INSERT INTO `iriscms_log` VALUES (904, 'model', 'add', '/b/model/add?_=1577501404686', 1, 'admin', '::1', '2019-12-28 11:17:26');
INSERT INTO `iriscms_log` VALUES (905, 'model', 'list', '/b/model/list?menuid=62&&_=1577501404687', 1, 'admin', '::1', '2019-12-28 11:18:37');
INSERT INTO `iriscms_log` VALUES (906, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:18:37');
INSERT INTO `iriscms_log` VALUES (907, 'model', 'add', '/b/model/add?_=1577501404688', 1, 'admin', '::1', '2019-12-28 11:18:37');
INSERT INTO `iriscms_log` VALUES (908, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153477', 1, 'admin', '::1', '2019-12-28 11:19:18');
INSERT INTO `iriscms_log` VALUES (909, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:19:18');
INSERT INTO `iriscms_log` VALUES (910, 'model', 'add', '/b/model/add?_=1577503153478', 1, 'admin', '::1', '2019-12-28 11:19:19');
INSERT INTO `iriscms_log` VALUES (911, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153479', 1, 'admin', '::1', '2019-12-28 11:20:24');
INSERT INTO `iriscms_log` VALUES (912, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:20:24');
INSERT INTO `iriscms_log` VALUES (913, 'model', 'add', '/b/model/add?_=1577503153480', 1, 'admin', '::1', '2019-12-28 11:20:25');
INSERT INTO `iriscms_log` VALUES (914, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153481', 1, 'admin', '::1', '2019-12-28 11:21:21');
INSERT INTO `iriscms_log` VALUES (915, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:21:21');
INSERT INTO `iriscms_log` VALUES (916, 'model', 'add', '/b/model/add?_=1577503153482', 1, 'admin', '::1', '2019-12-28 11:21:22');
INSERT INTO `iriscms_log` VALUES (917, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153483', 1, 'admin', '::1', '2019-12-28 11:21:44');
INSERT INTO `iriscms_log` VALUES (918, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:21:44');
INSERT INTO `iriscms_log` VALUES (919, 'model', 'add', '/b/model/add?_=1577503153484', 1, 'admin', '::1', '2019-12-28 11:21:45');
INSERT INTO `iriscms_log` VALUES (920, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:24:09');
INSERT INTO `iriscms_log` VALUES (921, 'model', 'add', '/b/model/add?_=1577503153486', 1, 'admin', '::1', '2019-12-28 11:24:10');
INSERT INTO `iriscms_log` VALUES (922, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153487', 1, 'admin', '::1', '2019-12-28 11:25:32');
INSERT INTO `iriscms_log` VALUES (923, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:25:32');
INSERT INTO `iriscms_log` VALUES (924, 'model', 'add', '/b/model/add?_=1577503153488', 1, 'admin', '::1', '2019-12-28 11:25:32');
INSERT INTO `iriscms_log` VALUES (925, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153489', 1, 'admin', '::1', '2019-12-28 11:26:07');
INSERT INTO `iriscms_log` VALUES (926, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:26:07');
INSERT INTO `iriscms_log` VALUES (927, 'model', 'add', '/b/model/add?_=1577503153490', 1, 'admin', '::1', '2019-12-28 11:26:07');
INSERT INTO `iriscms_log` VALUES (928, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153491', 1, 'admin', '::1', '2019-12-28 11:26:21');
INSERT INTO `iriscms_log` VALUES (929, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:26:21');
INSERT INTO `iriscms_log` VALUES (930, 'model', 'add', '/b/model/add?_=1577503153492', 1, 'admin', '::1', '2019-12-28 11:26:22');
INSERT INTO `iriscms_log` VALUES (931, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153493', 1, 'admin', '::1', '2019-12-28 11:26:38');
INSERT INTO `iriscms_log` VALUES (932, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:26:38');
INSERT INTO `iriscms_log` VALUES (933, 'model', 'add', '/b/model/add?_=1577503153494', 1, 'admin', '::1', '2019-12-28 11:26:38');
INSERT INTO `iriscms_log` VALUES (934, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153495', 1, 'admin', '::1', '2019-12-28 11:27:09');
INSERT INTO `iriscms_log` VALUES (935, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:27:09');
INSERT INTO `iriscms_log` VALUES (936, 'model', 'add', '/b/model/add?_=1577503153496', 1, 'admin', '::1', '2019-12-28 11:27:09');
INSERT INTO `iriscms_log` VALUES (937, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153497', 1, 'admin', '::1', '2019-12-28 11:28:36');
INSERT INTO `iriscms_log` VALUES (938, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:28:37');
INSERT INTO `iriscms_log` VALUES (939, 'model', 'add', '/b/model/add?_=1577503153498', 1, 'admin', '::1', '2019-12-28 11:28:37');
INSERT INTO `iriscms_log` VALUES (940, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153499', 1, 'admin', '::1', '2019-12-28 11:29:17');
INSERT INTO `iriscms_log` VALUES (941, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:29:17');
INSERT INTO `iriscms_log` VALUES (942, 'model', 'add', '/b/model/add?_=1577503153500', 1, 'admin', '::1', '2019-12-28 11:29:18');
INSERT INTO `iriscms_log` VALUES (943, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153501', 1, 'admin', '::1', '2019-12-28 11:30:07');
INSERT INTO `iriscms_log` VALUES (944, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:30:07');
INSERT INTO `iriscms_log` VALUES (945, 'model', 'add', '/b/model/add?_=1577503153502', 1, 'admin', '::1', '2019-12-28 11:30:07');
INSERT INTO `iriscms_log` VALUES (946, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153503', 1, 'admin', '::1', '2019-12-28 11:30:54');
INSERT INTO `iriscms_log` VALUES (947, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:30:54');
INSERT INTO `iriscms_log` VALUES (948, 'model', 'add', '/b/model/add?_=1577503153504', 1, 'admin', '::1', '2019-12-28 11:30:54');
INSERT INTO `iriscms_log` VALUES (949, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153505', 1, 'admin', '::1', '2019-12-28 11:32:41');
INSERT INTO `iriscms_log` VALUES (950, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:32:41');
INSERT INTO `iriscms_log` VALUES (951, 'model', 'add', '/b/model/add?_=1577503153506', 1, 'admin', '::1', '2019-12-28 11:32:42');
INSERT INTO `iriscms_log` VALUES (952, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153507', 1, 'admin', '::1', '2019-12-28 11:33:27');
INSERT INTO `iriscms_log` VALUES (953, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:33:27');
INSERT INTO `iriscms_log` VALUES (954, 'model', 'add', '/b/model/add?_=1577503153508', 1, 'admin', '::1', '2019-12-28 11:33:27');
INSERT INTO `iriscms_log` VALUES (955, 'model', 'list', '/b/model/list?menuid=62&&_=1577503153509', 1, 'admin', '::1', '2019-12-28 11:34:04');
INSERT INTO `iriscms_log` VALUES (956, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:34:04');
INSERT INTO `iriscms_log` VALUES (957, 'model', 'add', '/b/model/add?_=1577503153510', 1, 'admin', '::1', '2019-12-28 11:34:05');
INSERT INTO `iriscms_log` VALUES (958, 'user', 'list', '/b/user/list?menuid=56&&_=1577504086299', 1, 'admin', '::1', '2019-12-28 11:34:48');
INSERT INTO `iriscms_log` VALUES (959, 'user', 'list', '/b/user/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:34:49');
INSERT INTO `iriscms_log` VALUES (960, 'model', 'list', '/b/model/list?menuid=62&&_=1577504086300', 1, 'admin', '::1', '2019-12-28 11:34:50');
INSERT INTO `iriscms_log` VALUES (961, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:34:50');
INSERT INTO `iriscms_log` VALUES (962, 'model', 'add', '/b/model/add?_=1577504086301', 1, 'admin', '::1', '2019-12-28 11:34:51');
INSERT INTO `iriscms_log` VALUES (963, 'model', 'list', '/b/model/list?menuid=62&&_=1577504086302', 1, 'admin', '::1', '2019-12-28 11:35:37');
INSERT INTO `iriscms_log` VALUES (964, 'model', 'list', '/b/model/list?page=1&rows=NaN', 1, 'admin', '::1', '2019-12-28 11:35:37');
INSERT INTO `iriscms_log` VALUES (965, 'model', 'add', '/b/model/add?_=1577504086303', 1, 'admin', '::1', '2019-12-28 11:35:38');

-- ----------------------------
-- Table structure for iriscms_member
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_member`;
CREATE TABLE `iriscms_member`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `integral` int(255) NULL DEFAULT NULL,
  `sale_integral` int(255) NULL DEFAULT NULL,
  `draw_account` varchar(0) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `telphone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `qq` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `enabled` tinyint(2) NOT NULL DEFAULT 0,
  `verify_token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of iriscms_member
-- ----------------------------
INSERT INTO `iriscms_member` VALUES (1, 'xiusin', '159781', '', '陈二皮', 1231211111, 0, '', '123123', '1111222', '', '2019-01-24 11:40:00', '2019-01-24 11:40:00', '159781@11.com', 0, '4b32a22c-5787-4d0b-98f2-ed5b0779bbcb');

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
) ENGINE = MyISAM AUTO_INCREMENT = 65 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限表' ROW_FORMAT = Fixed;

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
INSERT INTO `iriscms_menu` VALUES (23, '删除日志', 15, 'system', 'log-delete', '', 0, 0, '1');
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
INSERT INTO `iriscms_menu` VALUES (55, '会员管理', 2, 'user', 'list', '', 0, 5, '1');
INSERT INTO `iriscms_menu` VALUES (56, '会员列表', 55, 'user', 'list', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (57, '会员信息', 56, 'user', 'info', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (58, '微信管理', 2, 'wechat', 'userlist', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (59, '微信会员信息', 58, 'wechat', 'userinfo', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (60, '编辑会员', 55, 'user', 'edit', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (61, '模型管理', 2, 'model', 'index', '', 0, 0, '1');
INSERT INTO `iriscms_menu` VALUES (62, '模型列表', 61, 'model', 'list', '', 0, 1, '1');
INSERT INTO `iriscms_menu` VALUES (64, '添加模型', 62, 'model', 'add', '?menuid=64', 0, 0, '1');

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
  PRIMARY KEY (`key`) USING BTREE,
  UNIQUE INDEX `UQE_iriscms_setting_key`(`key`) USING BTREE
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
INSERT INTO `iriscms_setting` VALUES ('WX_TOKEN', '');
INSERT INTO `iriscms_setting` VALUES ('OSS_HOST', '');
INSERT INTO `iriscms_setting` VALUES ('SITE_OPEN', '开启');
INSERT INTO `iriscms_setting` VALUES ('EMAIL_PORT', '25');
INSERT INTO `iriscms_setting` VALUES ('WX_APPSECRET', '');
INSERT INTO `iriscms_setting` VALUES ('WX_AESKEY', '');
INSERT INTO `iriscms_setting` VALUES ('HPJ_APPID', '');
INSERT INTO `iriscms_setting` VALUES ('DATAGRID_PAGE_SIZE', '25');
INSERT INTO `iriscms_setting` VALUES ('WX_APPID', '');
INSERT INTO `iriscms_setting` VALUES ('OSS_ENDPOINT', '');
INSERT INTO `iriscms_setting` VALUES ('OSS_KEYID', '');
INSERT INTO `iriscms_setting` VALUES ('OSS_BUCKETNAME', '');
INSERT INTO `iriscms_setting` VALUES ('HPJ_APPSECRET', '');
INSERT INTO `iriscms_setting` VALUES ('OSS_KEYSECRET', '');

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

-- ----------------------------
-- Table structure for iriscms_wechat_member
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_wechat_member`;
CREATE TABLE `iriscms_wechat_member`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `mpid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sex` tinyint(2) NULL DEFAULT NULL,
  `headimgurl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `subscribe_scene` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for iriscms_wechat_message_log
-- ----------------------------
DROP TABLE IF EXISTS `iriscms_wechat_message_log`;
CREATE TABLE `iriscms_wechat_message_log`  (
  `logid` bigint(20) NOT NULL,
  `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`logid`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

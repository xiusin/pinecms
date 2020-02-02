#
# SQL Export
# Created by Querious (201067)
# Created: February 2, 2020 at 12:24:59 PM GMT+8
# Encoding: Unicode (UTF-8)
#


DROP DATABASE IF EXISTS `iriscms`;
CREATE DATABASE `iriscms` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_0900_ai_ci;
USE `iriscms`;




SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;


DROP TABLE IF EXISTS `iriscms_wechat_message_log`;
DROP TABLE IF EXISTS `iriscms_wechat_member`;
DROP TABLE IF EXISTS `iriscms_slide`;
DROP TABLE IF EXISTS `iriscms_setting`;
DROP TABLE IF EXISTS `iriscms_page`;
DROP TABLE IF EXISTS `iriscms_news`;
DROP TABLE IF EXISTS `iriscms_menu`;
DROP TABLE IF EXISTS `iriscms_member`;
DROP TABLE IF EXISTS `iriscms_log`;
DROP TABLE IF EXISTS `iriscms_link`;
DROP TABLE IF EXISTS `iriscms_document_model_field`;
DROP TABLE IF EXISTS `iriscms_document_model_dsl`;
DROP TABLE IF EXISTS `iriscms_document_model`;
DROP TABLE IF EXISTS `iriscms_content`;
DROP TABLE IF EXISTS `iriscms_category_priv`;
DROP TABLE IF EXISTS `iriscms_category`;
DROP TABLE IF EXISTS `iriscms_admin_role_priv`;
DROP TABLE IF EXISTS `iriscms_admin_role`;
DROP TABLE IF EXISTS `iriscms_admin`;


CREATE TABLE `iriscms_admin` (
  `userid` mediumint(6) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) DEFAULT NULL,
  `password` varchar(32) DEFAULT NULL,
  `roleid` smallint(5) DEFAULT '0',
  `encrypt` varchar(6) DEFAULT NULL,
  `lastloginip` varchar(15) DEFAULT NULL,
  `lastlogintime` int(10) unsigned DEFAULT '0',
  `email` varchar(40) DEFAULT NULL,
  `realname` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`userid`) USING BTREE,
  UNIQUE KEY `UQE_iriscms_admin_username` (`username`) USING BTREE,
  KEY `username` (`username`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='管理员表';


CREATE TABLE `iriscms_admin_role` (
  `roleid` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `rolename` varchar(50) NOT NULL,
  `description` text NOT NULL,
  `listorder` smallint(5) unsigned NOT NULL DEFAULT '0',
  `disabled` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`roleid`) USING BTREE,
  UNIQUE KEY `UQE_iriscms_admin_role_rolename` (`rolename`) USING BTREE,
  KEY `listorder` (`listorder`) USING BTREE,
  KEY `disabled` (`disabled`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='角色表';


CREATE TABLE `iriscms_admin_role_priv` (
  `roleid` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `c` char(20) NOT NULL,
  `a` char(20) NOT NULL,
  KEY `roleid` (`roleid`,`c`,`a`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED COMMENT='角色操作权限配置表';


CREATE TABLE `iriscms_category` (
  `catid` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '栏目类型',
  `parentid` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '父类ID',
  `catname` varchar(30) NOT NULL COMMENT '分类名称',
  `description` mediumtext NOT NULL COMMENT '描述',
  `url` varchar(100) NOT NULL COMMENT '链接地址, 仅栏目类型为2的时候可用',
  `listorder` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序号',
  `ismenu` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否为栏目, 最初设定是可以在前端展示',
  `index_tpl` varchar(255) NOT NULL DEFAULT '' COMMENT '后台列表',
  `list_tpl` varchar(255) NOT NULL COMMENT '前台列表模板',
  `detail_tpl` varchar(255) NOT NULL COMMENT '前台内容页模板',
  `thumb` varchar(50) NOT NULL COMMENT '栏目缩略图',
  `model_id` bigint(20) DEFAULT NULL COMMENT '模型id,仅栏目类型为0的时候可用',
  `tpl_prefix` varchar(255) DEFAULT NULL,
  `home_tpl` varchar(255) DEFAULT NULL,
  `content_tpl` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`catid`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='内容分类表';


CREATE TABLE `iriscms_category_priv` (
  `catid` smallint(5) unsigned NOT NULL DEFAULT '0',
  `roleid` smallint(5) unsigned NOT NULL DEFAULT '0',
  `is_admin` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `action` char(30) NOT NULL,
  KEY `catid` (`catid`,`roleid`,`is_admin`,`action`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED;


CREATE TABLE `iriscms_content` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `catid` smallint(5) unsigned NOT NULL DEFAULT '0',
  `title` varchar(80) NOT NULL DEFAULT '',
  `thumb` varchar(100) NOT NULL DEFAULT '',
  `keywords` char(40) NOT NULL DEFAULT '',
  `description` mediumtext NOT NULL,
  `content` mediumtext NOT NULL,
  `listorder` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '1',
  `recommend` tinyint(2) DEFAULT NULL,
  `pwd_type` tinyint(2) DEFAULT NULL,
  `money` tinyint(5) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `deleted_at` int(11) DEFAULT NULL,
  `source_url` varchar(255) DEFAULT NULL,
  `source_pwd` varchar(255) DEFAULT NULL,
  `catids` varchar(255) DEFAULT NULL,
  `tags` varchar(255) DEFAULT NULL,
  `userid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`,`listorder`,`id`) USING BTREE,
  KEY `listorder` (`catid`,`status`,`listorder`,`id`) USING BTREE,
  KEY `catid` (`catid`,`status`,`id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='新闻表后期根据模型扩展';


CREATE TABLE `iriscms_document_model` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文档名称',
  `table` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '对应的表名',
  `enabled` tinyint(4) DEFAULT '0' COMMENT '是否启用',
  `is_system` tinyint(4) DEFAULT '0' COMMENT '是否为系统模型 无法删除',
  `model_type` tinyint(4) DEFAULT '0' COMMENT '模型类型: 扩展模型 和 独立模型',
  `fe_tpl_index` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '模型前端主页模板地址',
  `fe_tpl_list` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '模型前端列表模板地址',
  `fe_tpl_detail` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '模型前端详情模板地址',
  `deleted_at` datetime DEFAULT NULL,
  `field_show_in_list` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '在后端列表页需要展示的字段以及字段应用的formatter函数.',
  `formatters` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '所有函数内容, 原样渲染到Html里',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='文档模型用于存储自定义类型的文档内容';


CREATE TABLE `iriscms_document_model_dsl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mid` int(11) NOT NULL DEFAULT '0' COMMENT '模型id',
  `form_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '字段在表单内的名称',
  `table_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `html` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '字段html',
  `required` tinyint(4) DEFAULT '0' COMMENT '是否必填',
  `datasource` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '数据源, 可以让下拉选项等高级功能有数据读取的源头,具体设计可以是提供列表函数类的',
  `required_tips` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '必填(选)提醒',
  `validator` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '验证器名称或内容',
  `deleted_at` datetime DEFAULT NULL,
  `field_type` int(11) NOT NULL DEFAULT '0',
  `default` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=418 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;


CREATE TABLE `iriscms_document_model_field` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '字段名称',
  `type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '字段对应的数据类型',
  `desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '字段描述',
  `html` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;


CREATE TABLE `iriscms_link` (
  `linkid` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `linktype` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `name` varchar(50) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  `logo` varchar(255) NOT NULL DEFAULT '',
  `introduce` text NOT NULL,
  `listorder` smallint(5) unsigned NOT NULL DEFAULT '0',
  `passed` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `addtime` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`linkid`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='友情链接表';


CREATE TABLE `iriscms_log` (
  `logid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `controller` varchar(15) NOT NULL,
  `action` varchar(20) NOT NULL,
  `querystring` mediumtext NOT NULL,
  `userid` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `username` varchar(20) NOT NULL,
  `ip` varchar(15) NOT NULL,
  `time` datetime NOT NULL,
  PRIMARY KEY (`logid`) USING BTREE,
  KEY `module` (`controller`,`action`) USING BTREE,
  KEY `username` (`username`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=43 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='操作日志表';


CREATE TABLE `iriscms_member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `integral` int(255) DEFAULT NULL,
  `sale_integral` int(255) DEFAULT NULL,
  `draw_account` varchar(0) DEFAULT NULL,
  `telphone` varchar(255) DEFAULT NULL,
  `qq` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `enabled` tinyint(2) NOT NULL DEFAULT '0',
  `verify_token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='系统配置表';


CREATE TABLE `iriscms_menu` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(40) NOT NULL DEFAULT '',
  `parentid` smallint(6) NOT NULL DEFAULT '0',
  `c` char(20) NOT NULL DEFAULT '',
  `a` char(20) NOT NULL DEFAULT '',
  `data` char(255) NOT NULL DEFAULT '',
  `is_system` tinyint(1) NOT NULL DEFAULT '0',
  `listorder` smallint(6) unsigned NOT NULL DEFAULT '0',
  `display` enum('1','0') NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `listorder` (`listorder`) USING BTREE,
  KEY `parentid` (`parentid`) USING BTREE,
  KEY `module` (`c`,`a`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=65 DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED COMMENT='权限表';


CREATE TABLE `iriscms_news` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `catid` smallint(5) unsigned NOT NULL DEFAULT '0',
  `title` varchar(80) NOT NULL DEFAULT '',
  `thumb` varchar(100) NOT NULL DEFAULT '',
  `keywords` char(40) NOT NULL DEFAULT '',
  `description` mediumtext NOT NULL,
  `content` mediumtext NOT NULL,
  `listorder` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '1',
  `username` char(20) NOT NULL,
  `inputtime` int(10) unsigned NOT NULL DEFAULT '0',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0',
  `tpl` varchar(255) NOT NULL COMMENT '//模板名称',
  `recommend` tinyint(1) NOT NULL DEFAULT '0' COMMENT '推荐',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`,`listorder`,`id`) USING BTREE,
  KEY `listorder` (`catid`,`status`,`listorder`,`id`) USING BTREE,
  KEY `catid` (`catid`,`status`,`id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='新闻表后期根据模型扩展';


CREATE TABLE `iriscms_page` (
  `catid` smallint(5) unsigned NOT NULL DEFAULT '0',
  `title` varchar(160) NOT NULL,
  `keywords` varchar(40) NOT NULL,
  `content` text NOT NULL,
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0',
  KEY `catid` (`catid`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='单页内容表';


CREATE TABLE `iriscms_setting` (
  `key` varchar(50) NOT NULL,
  `value` text,
  PRIMARY KEY (`key`) USING BTREE,
  UNIQUE KEY `UQE_iriscms_setting_key` (`key`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='系统配置表';


CREATE TABLE `iriscms_slide` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
  `name1` varchar(255) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL COMMENT '地址',
  `imgurl` varchar(255) NOT NULL COMMENT '图片地址',
  `wapimgurl` varchar(255) DEFAULT NULL,
  `sigin` varchar(255) NOT NULL COMMENT '图片标识',
  `sort` int(5) NOT NULL DEFAULT '1' COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='幻灯片';


CREATE TABLE `iriscms_wechat_member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(255) DEFAULT NULL,
  `mpid` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `sex` tinyint(2) DEFAULT NULL,
  `headimgurl` varchar(255) DEFAULT NULL,
  `subscribe_scene` varchar(255) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='系统配置表';


CREATE TABLE `iriscms_wechat_message_log` (
  `logid` bigint(20) NOT NULL,
  `content` varchar(255) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`logid`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;




SET FOREIGN_KEY_CHECKS = @PREVIOUS_FOREIGN_KEY_CHECKS;


SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;


LOCK TABLES `iriscms_admin` WRITE;
ALTER TABLE `iriscms_admin` DISABLE KEYS;
INSERT INTO `iriscms_admin` (`userid`, `username`, `password`, `roleid`, `encrypt`, `lastloginip`, `lastlogintime`, `email`, `realname`) VALUES 
	(1,'admin','0087af20a551a8b804f89469534b7859',1,'qmRlFL','127.0.0.1',1474291850,'chenchengbin92@gmail.com','mirchen.com'),
	(11,'test','834a33db060873a7a208617930edb29a',9,'5u2G0w','::1',0,'asdasd@asdasd.com1','ccc1');
ALTER TABLE `iriscms_admin` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_admin_role` WRITE;
ALTER TABLE `iriscms_admin_role` DISABLE KEYS;
INSERT INTO `iriscms_admin_role` (`roleid`, `rolename`, `description`, `listorder`, `disabled`) VALUES 
	(1,'超级管理员','超级管理员',0,0),
	(9,'test','test',0,0);
ALTER TABLE `iriscms_admin_role` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_admin_role_priv` WRITE;
ALTER TABLE `iriscms_admin_role_priv` DISABLE KEYS;
INSERT INTO `iriscms_admin_role_priv` (`roleid`, `c`, `a`) VALUES 
	(9,'system','left'),
	(9,'system','log-delete'),
	(9,'system','loglist'),
	(9,'system','logview'),
	(9,'system','menuadd'),
	(9,'system','menudelete'),
	(9,'system','menuedit'),
	(9,'system','menulist'),
	(9,'system','menuorder'),
	(9,'system','menuview'),
	(9,'system','top');
ALTER TABLE `iriscms_admin_role_priv` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_category` WRITE;
ALTER TABLE `iriscms_category` DISABLE KEYS;
INSERT INTO `iriscms_category` (`catid`, `type`, `parentid`, `catname`, `description`, `url`, `listorder`, `ismenu`, `index_tpl`, `list_tpl`, `detail_tpl`, `thumb`, `model_id`, `tpl_prefix`, `home_tpl`, `content_tpl`) VALUES 
	(2,0,0,'最新新闻','','',0,1,'','','','',3,NULL,NULL,NULL),
	(3,0,0,'系统模型演示','系统模型演示','',0,1,'','','','',5,NULL,NULL,NULL);
ALTER TABLE `iriscms_category` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_category_priv` WRITE;
ALTER TABLE `iriscms_category_priv` DISABLE KEYS;
ALTER TABLE `iriscms_category_priv` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_content` WRITE;
ALTER TABLE `iriscms_content` DISABLE KEYS;
ALTER TABLE `iriscms_content` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_document_model` WRITE;
ALTER TABLE `iriscms_document_model` DISABLE KEYS;
INSERT INTO `iriscms_document_model` (`id`, `name`, `table`, `enabled`, `is_system`, `model_type`, `fe_tpl_index`, `fe_tpl_list`, `fe_tpl_detail`, `deleted_at`, `field_show_in_list`, `formatters`) VALUES 
	(3,'标准模型','test',1,0,0,'resources/views/backend/content_index.html','resources/views/backend/category_edit.html','resources/views/backend/content_index.html',NULL,'{"checkbox":{"show":false,"formatter":""},"html":{"show":false,"formatter":""},"img":{"show":false,"formatter":""},"imgs":{"show":false,"formatter":""},"radio":{"show":false,"formatter":""},"rili":{"show":false,"formatter":""},"select":{"show":false,"formatter":""},"select1":{"show":false,"formatter":""},"select2":{"show":false,"formatter":""},"select3":{"show":false,"formatter":""},"switch":{"show":true,"formatter":""},"text":{"show":true,"formatter":""},"textarea":{"show":true,"formatter":""}}','asdasdasdasasdasdasdasdasds'),
	(5,'系统模型','articles',1,0,0,'','','',NULL,'{"click":{"show":true,"formatter":""},"content":{"show":false,"formatter":""},"description":{"show":true,"formatter":""},"from_url":{"show":false,"formatter":""},"keywords":{"show":true,"formatter":""},"publish_time":{"show":true,"formatter":""},"status":{"show":true,"formatter":""},"thumb":{"show":false,"formatter":""},"title":{"show":true,"formatter":""}}','');
ALTER TABLE `iriscms_document_model` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_document_model_dsl` WRITE;
ALTER TABLE `iriscms_document_model_dsl` DISABLE KEYS;
INSERT INTO `iriscms_document_model_dsl` (`id`, `mid`, `form_name`, `table_field`, `html`, `required`, `datasource`, `required_tips`, `validator`, `deleted_at`, `field_type`, `default`) VALUES 
	(396,3,'单行文本','text','<input class="easyui-textbox" {{attr}} style="width:300px">',1,'','请填写内容','',NULL,1,''),
	(397,3,'多行文本','textarea','<input class="easyui-textbox" {{attr}} style="height:80px; width: 300px;"  multiline />',1,'','','',NULL,2,''),
	(398,3,'HTML','html','<editor />',1,'','','',NULL,3,''),
	(399,3,'单图上传','img','<images />',1,'','','',NULL,11,''),
	(400,3,'多图上传','imgs','<mul-images />',1,'["张三", "李四", "王五"]','','',NULL,12,''),
	(401,3,'下拉框','select','<input class="easyui-combobox" {{attr}} style="width:300px;" />',1,'["张三", "李四", "王五"]','','',NULL,5,''),
	(402,3,'下拉框1','select1','<input class="easyui-combobox" {{attr}} style="width:300px;" />',1,'{"1":1111111,  "2":2222222}','','',NULL,5,''),
	(403,3,'下拉框2','select2','<input class="easyui-combobox" {{attr}} style="width:300px;" />',1,'{"1":"zhangsan", "2":"lisi"}','','',NULL,5,''),
	(404,3,'下拉框3','select3','<input class="easyui-combobox" {{attr}} style="width:300px;" />',1,'{"user1":"zhangsan", "user2":"lisi"}','','',NULL,5,''),
	(405,3,'开关按钮','switch','<input class="easyui-switchbutton" {{attr}} {{default}} />',1,'["是", "否"]','','',NULL,13,'否'),
	(406,3,'多选','checkbox','<input class="easyui-checkbox" {{attr}} value="{{value}}" {{default}} />',1,'["张三", "李四", "王五", "赵六"]','','',NULL,8,'李四|王五'),
	(407,3,'单选按钮','radio','<input class="easyui-radiobutton" {{attr}} value="{{value}}" {{default}} />',1,'["是", "否"]','','',NULL,7,'是'),
	(408,3,'日历组件','rili','<input class="easyui-datetimebox" style="width:300px" {{attr}} />',1,'','','',NULL,14,'');
ALTER TABLE `iriscms_document_model_dsl` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_document_model_field` WRITE;
ALTER TABLE `iriscms_document_model_field` DISABLE KEYS;
INSERT INTO `iriscms_document_model_field` (`id`, `name`, `type`, `desc`, `html`) VALUES 
	(1,'单行文本','varchar','常用字段，如文章标题、作者等都属于直接输入少量内容的文本，设置这个文本之后需要指定文本长度，默认为250，如果大于255则为text类型','<input class="easyui-textbox" {{attr}} style="width:300px">'),
	(2,'多行文本','text','也是较为常用的字段类型，如个人简介、产品描述都可以使用多行文本进行存储','<input class="easyui-textbox" {{attr}} style="height:80px; width: 300px;"  multiline />'),
	(3,'HTML文本','text','编辑器编辑产生的html内容，用于比较复杂的内容形式, 可以认为是附带编辑器的多行文本','<editor />'),
	(4,'附件','varshar','前端表现为input[file]类型,可以上传并且返回一个相对地址','<input class="easyui-filebox" style="width:300px">'),
	(5,'下拉框','varchar','下拉选择，一般用于如软件类型、语言类型等字段','<input class="easyui-combobox" {{attr}} style="width:300px;" />'),
	(6,'联动类型','varchar','一种数组形式的数据类型，请使用url接口方式提供','<select class="easyui-combotree" {{attr}} style="width:200px;">暂未实现</select>'),
	(7,'单选框','varchar','平铺显示, 可以认为是下拉框的展开, 根据数据源展开为排列的组件','<input class="easyui-radiobutton" {{attr}} value="{{value}}" {{default}} />'),
	(8,'多选框','varchar','多选框, 平铺显示为多个选项,根据数据源展开为排列组件','<input class="easyui-checkbox" {{attr}} value="{{value}}" {{default}} />'),
	(9,'整数类型','int','常用字段, 仅能输入数字','<input type="text" class="easyui-numberbox" value="{{value}}" {{attr}} />'),
	(10,'浮点类型','float','常用字段, 仅能输入浮点数(小数)','<input type="text" class="easyui-numberbox" value="{{value}}" {{attr}} />'),
	(11,'单图上传','varchar','常用字段, 会生成一个单图上传框','<images />'),
	(12,'多图上传','varchar','生成一个多图上传的组件','<mul-images />'),
	(13,'开关按钮','tinyint','用于做开关值的组件, 打开为1, 关闭为0','<input class="easyui-switchbutton" {{attr}} {{default}} />'),
	(14,'日历组件','datetime','选择日期组件','<input class="easyui-datetimebox" style="width:300px" {{attr}} />');
ALTER TABLE `iriscms_document_model_field` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_link` WRITE;
ALTER TABLE `iriscms_link` DISABLE KEYS;
ALTER TABLE `iriscms_link` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_log` WRITE;
ALTER TABLE `iriscms_log` DISABLE KEYS;
INSERT INTO `iriscms_log` (`logid`, `controller`, `action`, `querystring`, `userid`, `username`, `ip`, `time`) VALUES 
	(1,'category','list','/b/category/list?menuid=36&&_=1579053363466',1,'admin','127.0.0.1','2020-01-15 09:57:00'),
	(2,'category','list','/b/category/list?grid=treegrid',1,'admin','127.0.0.1','2020-01-15 09:57:00'),
	(3,'content','index','/b/content/index?menuid=35&&_=1579053363467',1,'admin','127.0.0.1','2020-01-15 09:57:00'),
	(4,'content','right','/b/content/right?_=1579053363468',1,'admin','127.0.0.1','2020-01-15 09:57:00'),
	(5,'content','right','/b/content/right',1,'admin','127.0.0.1','2020-01-15 09:57:00'),
	(6,'content','right','/b/content/right',1,'admin','127.0.0.1','2020-01-15 09:57:00'),
	(7,'category','list','/b/category/list?menuid=36&&_=1579053363470',1,'admin','127.0.0.1','2020-01-15 09:57:01'),
	(8,'category','list','/b/category/list?grid=treegrid',1,'admin','127.0.0.1','2020-01-15 09:57:01'),
	(9,'content','index','/b/content/index?menuid=35&&_=1579053363471',1,'admin','127.0.0.1','2020-01-15 09:57:03'),
	(10,'content','right','/b/content/right?_=1579053363472',1,'admin','127.0.0.1','2020-01-15 09:57:03'),
	(11,'content','right','/b/content/right',1,'admin','127.0.0.1','2020-01-15 09:57:03'),
	(12,'content','right','/b/content/right',1,'admin','127.0.0.1','2020-01-15 09:57:03'),
	(13,'model','list','/b/model/list?menuid=62&&_=1579053363474',1,'admin','127.0.0.1','2020-01-15 09:57:05'),
	(14,'model','list','/b/model/list?page=1&rows=20',1,'admin','127.0.0.1','2020-01-15 09:57:05'),
	(15,'model','list-field-show','/b/model/list-field-show?mid=5&_=1579053363475',1,'admin','127.0.0.1','2020-01-15 09:57:07'),
	(16,'model','list-field-show','/b/model/list-field-show?mid=5&_=1579053363476',1,'admin','127.0.0.1','2020-01-15 09:57:12'),
	(17,'model','edit','/b/model/edit?mid=5&_=1579053363477',1,'admin','127.0.0.1','2020-01-15 09:57:15'),
	(18,'model','edit','/b/model/edit',1,'admin','127.0.0.1','2020-01-15 10:01:26'),
	(19,'model','edit','/b/model/edit',1,'admin','127.0.0.1','2020-01-15 10:01:35'),
	(20,'model','list','/b/model/list?menuid=62&&_=1579053982048',1,'admin','127.0.0.1','2020-01-15 10:06:26'),
	(21,'model','list','/b/model/list?page=1&rows=20',1,'admin','127.0.0.1','2020-01-15 10:06:26'),
	(22,'model','add','/b/model/add?_=1579053982049',1,'admin','127.0.0.1','2020-01-15 10:06:26'),
	(23,'model','list','/b/model/list?menuid=62&&_=1579053982050',1,'admin','127.0.0.1','2020-01-15 10:08:08'),
	(24,'model','list','/b/model/list?page=1&rows=20',1,'admin','127.0.0.1','2020-01-15 10:08:08'),
	(25,'model','add','/b/model/add?_=1579053982051',1,'admin','127.0.0.1','2020-01-15 10:08:09'),
	(26,'model','list','/b/model/list?menuid=62&&_=1579053982052',1,'admin','127.0.0.1','2020-01-15 10:08:28'),
	(27,'model','list','/b/model/list?page=1&rows=20',1,'admin','127.0.0.1','2020-01-15 10:08:28'),
	(28,'model','list-field-show','/b/model/list-field-show?mid=3&_=1579053982053',1,'admin','127.0.0.1','2020-01-15 10:08:31'),
	(29,'model','list-field-show','/b/model/list-field-show?mid=5&_=1579053982054',1,'admin','127.0.0.1','2020-01-15 10:08:35'),
	(30,'model','list-field-show','/b/model/list-field-show?mid=5&_=1579053982055',1,'admin','127.0.0.1','2020-01-15 10:19:13'),
	(31,'model','edit','/b/model/edit?mid=5&_=1579053982056',1,'admin','127.0.0.1','2020-01-15 10:19:16'),
	(32,'model','edit','/b/model/edit',1,'admin','127.0.0.1','2020-01-15 10:26:27'),
	(33,'model','list','/b/model/list?menuid=62&&_=1579053982057',1,'admin','127.0.0.1','2020-01-15 10:28:04'),
	(34,'model','list','/b/model/list?page=1&rows=20',1,'admin','127.0.0.1','2020-01-15 10:28:05'),
	(35,'model','edit','/b/model/edit?mid=5&_=1579053982058',1,'admin','127.0.0.1','2020-01-15 10:28:11'),
	(36,'model','edit','/b/model/edit',1,'admin','127.0.0.1','2020-01-15 10:28:24'),
	(37,'wechat','userinfo','/b/wechat/userinfo?menuid=59&&_=1579055569342',1,'admin','127.0.0.1','2020-01-15 10:32:55'),
	(38,'model','list','/b/model/list?menuid=62&&_=1579055569343',1,'admin','127.0.0.1','2020-01-15 10:32:56'),
	(39,'model','list','/b/model/list?page=1&rows=20',1,'admin','127.0.0.1','2020-01-15 10:32:56'),
	(40,'model','list-field-show','/b/model/list-field-show?mid=5&_=1579055569344',1,'admin','127.0.0.1','2020-01-15 10:32:58'),
	(41,'model','edit','/b/model/edit?mid=5&_=1579055569345',1,'admin','127.0.0.1','2020-01-15 10:33:01'),
	(42,'model','edit','/b/model/edit',1,'admin','127.0.0.1','2020-01-15 10:33:14');
ALTER TABLE `iriscms_log` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_member` WRITE;
ALTER TABLE `iriscms_member` DISABLE KEYS;
INSERT INTO `iriscms_member` (`id`, `account`, `password`, `avatar`, `nickname`, `integral`, `sale_integral`, `draw_account`, `telphone`, `qq`, `description`, `created_at`, `updated_at`, `email`, `enabled`, `verify_token`) VALUES 
	(1,'xiusin','159781','','陈二皮',1231211111,0,'','123123','1111222','','2019-01-24 11:40:00','2019-01-24 11:40:00','159781@11.com',0,'4b32a22c-5787-4d0b-98f2-ed5b0779bbcb');
ALTER TABLE `iriscms_member` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_menu` WRITE;
ALTER TABLE `iriscms_menu` DISABLE KEYS;
INSERT INTO `iriscms_menu` (`id`, `name`, `parentid`, `c`, `a`, `data`, `is_system`, `listorder`, `display`) VALUES 
	(1,'我的面板',0,'admin','public-top','',1,1,'1'),
	(2,'系统设置',0,'setting','top','',0,6,'1'),
	(3,'内容管理',0,'content','top','',0,2,'1'),
	(54,'分类单页',35,'content','page','',0,0,'1'),
	(6,'个人信息',1,'admin','public-left','',1,0,'1'),
	(7,'修改密码',6,'admin','public-editpwd','',1,1,'1'),
	(8,'修改个人信息',6,'admin','public-editinfo','',1,0,'1'),
	(9,'系统设置',2,'setting','left','',0,1,'1'),
	(10,'站点设置',9,'setting','site','',0,1,'1'),
	(11,'管理员设置',2,'admin','left','',0,2,'1'),
	(12,'管理员管理',11,'admin','memberlist','',0,1,'1'),
	(13,'角色管理',11,'admin','rolelist','',0,2,'1'),
	(14,'后台管理',2,'system','left','',0,1,'1'),
	(15,'日志管理',14,'system','loglist','',0,1,'1'),
	(16,'菜单管理',14,'system','menulist','',0,2,'1'),
	(17,'查看菜单',16,'system','menuview','',0,0,'1'),
	(18,'添加菜单',16,'system','menuadd','',0,0,'1'),
	(19,'修改菜单',16,'system','menuedit','',0,0,'1'),
	(20,'删除菜单',16,'system','menudelete','',0,0,'1'),
	(21,'菜单排序',16,'system','menuorder','',0,0,'1'),
	(22,'查看日志',15,'system','logview','',0,0,'1'),
	(23,'删除日志',15,'system','log-delete','',0,0,'1'),
	(24,'查看管理员',12,'admin','member-view','',0,0,'1'),
	(25,'添加管理员',12,'admin','member-add','',0,0,'1'),
	(26,'编辑管理员',12,'admin','member-edit','',0,0,'1'),
	(27,'删除管理员',12,'admin','member-delete','',0,0,'1'),
	(28,'查看角色',13,'admin','rolelist','',0,0,'1'),
	(29,'添加角色',13,'admin','role-add','',0,0,'1'),
	(30,'编辑角色',13,'admin','role-edit','',0,0,'1'),
	(31,'删除角色',13,'admin','role-delete','',0,0,'1'),
	(53,'新闻列表',35,'content','news-list','',0,0,'1'),
	(33,'权限设置',13,'admin','role-permission','',0,0,'1'),
	(34,'发布管理',3,'content','right','',0,0,'1'),
	(35,'内容管理',34,'content','index','',0,0,'1'),
	(36,'栏目管理',34,'category','list','',0,0,'1'),
	(37,'查看栏目',36,'category','view','',0,0,'1'),
	(38,'添加栏目',36,'category','add','',0,0,'1'),
	(39,'编辑栏目',36,'category','edit','',0,0,'1'),
	(40,'删除栏目',36,'category','delete','',0,0,'1'),
	(41,'栏目排序',36,'category','order','',0,0,'1'),
	(55,'会员管理',2,'user','list','',0,5,'1'),
	(56,'会员列表',55,'user','list','',0,0,'1'),
	(57,'会员信息',56,'user','info','',0,0,'1'),
	(58,'微信管理',2,'wechat','userlist','',0,0,'1'),
	(59,'微信会员信息',58,'wechat','userinfo','',0,0,'1'),
	(60,'编辑会员',55,'user','edit','',0,0,'1'),
	(61,'模型管理',2,'model','index','',0,0,'1'),
	(62,'模型列表',61,'model','list','',0,1,'1'),
	(64,'添加模型',62,'model','add','?menuid=64',0,0,'1');
ALTER TABLE `iriscms_menu` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_news` WRITE;
ALTER TABLE `iriscms_news` DISABLE KEYS;
ALTER TABLE `iriscms_news` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_page` WRITE;
ALTER TABLE `iriscms_page` DISABLE KEYS;
ALTER TABLE `iriscms_page` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_setting` WRITE;
ALTER TABLE `iriscms_setting` DISABLE KEYS;
INSERT INTO `iriscms_setting` (`key`, `value`) VALUES 
	('EMAIL_USER',''),
	('EMAIL_SMTP',''),
	('EMAIL_PWD',''),
	('EMAIL_EMAIL',''),
	('SITE_ICP',''),
	('SITE_KEYWORDS',''),
	('SITE_DESCRIPTION',''),
	('SITE_TITLE',''),
	('WX_TOKEN',''),
	('OSS_HOST',''),
	('SITE_OPEN','开启'),
	('EMAIL_PORT','25'),
	('WX_APPSECRET',''),
	('WX_AESKEY',''),
	('HPJ_APPID',''),
	('DATAGRID_PAGE_SIZE','25'),
	('WX_APPID',''),
	('OSS_ENDPOINT',''),
	('OSS_KEYID',''),
	('OSS_BUCKETNAME',''),
	('HPJ_APPSECRET',''),
	('OSS_KEYSECRET','');
ALTER TABLE `iriscms_setting` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_slide` WRITE;
ALTER TABLE `iriscms_slide` DISABLE KEYS;
ALTER TABLE `iriscms_slide` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_wechat_member` WRITE;
ALTER TABLE `iriscms_wechat_member` DISABLE KEYS;
ALTER TABLE `iriscms_wechat_member` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_wechat_message_log` WRITE;
ALTER TABLE `iriscms_wechat_message_log` DISABLE KEYS;
ALTER TABLE `iriscms_wechat_message_log` ENABLE KEYS;
UNLOCK TABLES;




SET FOREIGN_KEY_CHECKS = @PREVIOUS_FOREIGN_KEY_CHECKS;



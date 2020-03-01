#
# SQL Export
# Created by Querious (201048)
# Created: 2020年3月1日 GMT+8 下午12:52:41
# Encoding: Unicode (UTF-8)
#


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
DROP TABLE IF EXISTS `iriscms_attachments`;
DROP TABLE IF EXISTS `iriscms_articles`;
DROP TABLE IF EXISTS `iriscms_admin_role_priv`;
DROP TABLE IF EXISTS `iriscms_admin_role`;
DROP TABLE IF EXISTS `iriscms_admin`;


CREATE TABLE `iriscms_admin` (
  `userid` mediumint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) DEFAULT NULL,
  `password` varchar(32) DEFAULT NULL,
  `roleid` smallint DEFAULT '0',
  `encrypt` varchar(6) DEFAULT NULL,
  `lastloginip` varchar(15) DEFAULT NULL,
  `lastlogintime` int unsigned DEFAULT '0',
  `email` varchar(40) DEFAULT NULL,
  `realname` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`userid`) USING BTREE,
  UNIQUE KEY `UQE_iriscms_admin_username` (`username`) USING BTREE,
  KEY `username` (`username`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='管理员表';


CREATE TABLE `iriscms_admin_role` (
  `roleid` tinyint unsigned NOT NULL AUTO_INCREMENT,
  `rolename` varchar(50) NOT NULL,
  `description` text NOT NULL,
  `listorder` smallint unsigned NOT NULL DEFAULT '0',
  `disabled` tinyint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`roleid`) USING BTREE,
  UNIQUE KEY `UQE_iriscms_admin_role_rolename` (`rolename`) USING BTREE,
  KEY `listorder` (`listorder`) USING BTREE,
  KEY `disabled` (`disabled`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='角色表';


CREATE TABLE `iriscms_admin_role_priv` (
  `roleid` tinyint unsigned NOT NULL DEFAULT '0',
  `c` char(20) NOT NULL,
  `a` char(20) NOT NULL,
  KEY `roleid` (`roleid`,`c`,`a`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED COMMENT='角色操作权限配置表';


CREATE TABLE `iriscms_articles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT NULL COMMENT '标题',
  `keywords` varchar(100) DEFAULT NULL COMMENT '关键字',
  `description` text COMMENT '描述',
  `thumb` varchar(100) DEFAULT NULL COMMENT '封面',
  `content` text COMMENT '内容',
  `from_website` varchar(100) DEFAULT NULL COMMENT '来源',
  `from_url` varchar(100) DEFAULT NULL COMMENT '来源地址',
  `tags` varchar(100) DEFAULT NULL COMMENT '标签',
  `published` varchar(100) DEFAULT NULL COMMENT '是否发布',
  `images` varchar(100) DEFAULT NULL COMMENT '图集',
  `catid` int unsigned NOT NULL DEFAULT '0' COMMENT '所属栏目ID',
  `mid` int unsigned NOT NULL DEFAULT '0' COMMENT '模型ID',
  `refid` int unsigned NOT NULL DEFAULT '0' COMMENT '模型关联的文章ID',
  `visit_count` int unsigned NOT NULL DEFAULT '0' COMMENT '访问次数',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `created_time` datetime DEFAULT NULL,
  `updated_time` datetime DEFAULT NULL,
  `deleted_time` datetime DEFAULT NULL,
  `listorder` int unsigned DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `refid` (`refid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;


CREATE TABLE `iriscms_attachments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL,
  `url` varchar(128) DEFAULT NULL,
  `origin_name` varchar(128) DEFAULT NULL,
  `size` int DEFAULT '0',
  `upload_time` datetime DEFAULT NULL,
  `type` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=121 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='附件表';


CREATE TABLE `iriscms_category` (
  `catid` smallint unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '栏目类型',
  `parentid` smallint unsigned NOT NULL DEFAULT '0' COMMENT '父类ID',
  `catname` varchar(30) NOT NULL COMMENT '分类名称',
  `description` mediumtext NOT NULL COMMENT '描述',
  `url` varchar(100) NOT NULL COMMENT '链接地址, 仅栏目类型为2的时候可用',
  `listorder` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序号',
  `ismenu` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否为栏目, 最初设定是可以在前端展示',
  `index_tpl` varchar(255) NOT NULL DEFAULT '' COMMENT '后台列表',
  `list_tpl` varchar(255) NOT NULL COMMENT '前台列表模板',
  `detail_tpl` varchar(255) NOT NULL COMMENT '前台内容页模板',
  `thumb` varchar(50) NOT NULL COMMENT '栏目缩略图',
  `model_id` bigint DEFAULT NULL COMMENT '模型id,仅栏目类型为0的时候可用',
  `tpl_prefix` varchar(255) DEFAULT NULL,
  `home_tpl` varchar(255) DEFAULT NULL,
  `content_tpl` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`catid`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=19 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='内容分类表';


CREATE TABLE `iriscms_category_priv` (
  `catid` smallint NOT NULL DEFAULT '0',
  `roleid` smallint NOT NULL DEFAULT '0',
  `is_admin` tinyint NOT NULL DEFAULT '0',
  `action` char(30) NOT NULL,
  KEY `catid` (`catid`,`roleid`,`is_admin`,`action`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED COMMENT='栏目权限表';


CREATE TABLE `iriscms_content` (
  `id` mediumint unsigned NOT NULL AUTO_INCREMENT,
  `catid` smallint unsigned NOT NULL DEFAULT '0',
  `title` varchar(80) NOT NULL DEFAULT '',
  `thumb` varchar(100) NOT NULL DEFAULT '',
  `keywords` char(40) NOT NULL DEFAULT '',
  `description` mediumtext NOT NULL,
  `content` mediumtext NOT NULL,
  `listorder` tinyint unsigned NOT NULL DEFAULT '0',
  `status` tinyint unsigned NOT NULL DEFAULT '1',
  `recommend` tinyint DEFAULT NULL,
  `pwd_type` tinyint DEFAULT NULL,
  `money` tinyint DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  `deleted_at` int DEFAULT NULL,
  `source_url` varchar(255) DEFAULT NULL,
  `source_pwd` varchar(255) DEFAULT NULL,
  `catids` varchar(255) DEFAULT NULL,
  `tags` varchar(255) DEFAULT NULL,
  `userid` int DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`,`listorder`,`id`) USING BTREE,
  KEY `listorder` (`catid`,`status`,`listorder`,`id`) USING BTREE,
  KEY `catid` (`catid`,`status`,`id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='新闻表后期根据模型扩展';


CREATE TABLE `iriscms_document_model` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL COMMENT '文档名称',
  `table` varchar(128) DEFAULT NULL COMMENT '对应的表名',
  `enabled` tinyint DEFAULT '0' COMMENT '是否启用',
  `model_type` tinyint DEFAULT '0' COMMENT '模型类型: 扩展模型 和 独立模型',
  `fe_tpl_index` varchar(128) DEFAULT NULL COMMENT '模型前端主页模板地址',
  `fe_tpl_list` varchar(128) DEFAULT NULL COMMENT '模型前端列表模板地址',
  `fe_tpl_detail` varchar(128) DEFAULT NULL COMMENT '模型前端详情模板地址',
  `deleted_at` datetime DEFAULT NULL,
  `field_show_in_list` text COMMENT '在后端列表页需要展示的字段以及字段应用的formatter函数.',
  `formatters` text COMMENT '所有函数内容, 原样渲染到Html里',
  `execed` tinyint DEFAULT '0' COMMENT '是否已经执行过改动',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='文档模型用于存储自定义类型的文档内容';


CREATE TABLE `iriscms_document_model_dsl` (
  `id` int NOT NULL AUTO_INCREMENT,
  `mid` int NOT NULL DEFAULT '0' COMMENT '模型id',
  `form_name` varchar(128) DEFAULT NULL COMMENT '字段在表单内的名称',
  `table_field` varchar(128) DEFAULT NULL,
  `html` text COMMENT '字段html',
  `required` tinyint DEFAULT '0' COMMENT '是否必填',
  `datasource` varchar(128) DEFAULT NULL COMMENT '数据源, 可以让下拉选项等高级功能有数据读取的源头,具体设计可以是提供列表函数类的',
  `required_tips` varchar(128) DEFAULT NULL COMMENT '必填(选)提醒',
  `validator` varchar(128) DEFAULT NULL COMMENT '验证器名称或内容',
  `deleted_at` datetime DEFAULT NULL,
  `field_type` int NOT NULL DEFAULT '0',
  `default` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=530 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='模型表单定义表';


CREATE TABLE `iriscms_document_model_field` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL COMMENT '字段名称',
  `type` varchar(128) DEFAULT NULL COMMENT '字段对应的数据类型',
  `desc` varchar(128) DEFAULT NULL COMMENT '字段描述',
  `html` text,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='模型表单组件定义表';


CREATE TABLE `iriscms_link` (
  `linkid` smallint NOT NULL AUTO_INCREMENT,
  `linktype` tinyint(1) NOT NULL DEFAULT '0',
  `name` varchar(50) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  `logo` varchar(255) NOT NULL DEFAULT '',
  `introduce` text NOT NULL,
  `listorder` smallint NOT NULL DEFAULT '0',
  `passed` tinyint(1) NOT NULL DEFAULT '0',
  `addtime` datetime DEFAULT NULL,
  PRIMARY KEY (`linkid`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='友情链接表';


CREATE TABLE `iriscms_log` (
  `logid` int NOT NULL AUTO_INCREMENT,
  `controller` varchar(15) NOT NULL,
  `action` varchar(20) NOT NULL,
  `querystring` mediumtext NOT NULL,
  `userid` mediumint NOT NULL DEFAULT '0',
  `username` varchar(20) NOT NULL,
  `ip` varchar(15) NOT NULL,
  `time` datetime NOT NULL,
  PRIMARY KEY (`logid`) USING BTREE,
  KEY `module` (`controller`,`action`) USING BTREE,
  KEY `username` (`username`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=710 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='操作日志表';


CREATE TABLE `iriscms_member` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `integral` int DEFAULT NULL,
  `sale_integral` int DEFAULT NULL,
  `draw_account` varchar(0) DEFAULT NULL,
  `telphone` varchar(255) DEFAULT NULL,
  `qq` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `enabled` tinyint NOT NULL DEFAULT '0',
  `verify_token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='系统配置表';


CREATE TABLE `iriscms_menu` (
  `id` smallint NOT NULL AUTO_INCREMENT,
  `name` char(40) NOT NULL DEFAULT '',
  `parentid` smallint NOT NULL DEFAULT '0',
  `c` char(20) NOT NULL DEFAULT '',
  `a` char(20) NOT NULL DEFAULT '',
  `data` char(255) NOT NULL DEFAULT '',
  `is_system` tinyint(1) NOT NULL DEFAULT '0',
  `listorder` smallint NOT NULL DEFAULT '0',
  `display` enum('1','0') NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `listorder` (`listorder`) USING BTREE,
  KEY `parentid` (`parentid`) USING BTREE,
  KEY `module` (`c`,`a`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=85 DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED COMMENT='权限菜单表';


CREATE TABLE `iriscms_news` (
  `id` mediumint NOT NULL AUTO_INCREMENT,
  `catid` smallint NOT NULL DEFAULT '0',
  `title` varchar(80) NOT NULL DEFAULT '',
  `thumb` varchar(100) NOT NULL DEFAULT '',
  `keywords` char(40) NOT NULL DEFAULT '',
  `description` mediumtext NOT NULL,
  `content` mediumtext NOT NULL,
  `listorder` tinyint NOT NULL DEFAULT '0',
  `status` tinyint NOT NULL DEFAULT '1',
  `username` char(20) NOT NULL,
  `inputtime` int NOT NULL DEFAULT '0',
  `updatetime` int NOT NULL DEFAULT '0',
  `tpl` varchar(255) NOT NULL COMMENT '//模板名称',
  `recommend` tinyint(1) NOT NULL DEFAULT '0' COMMENT '推荐',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`,`listorder`,`id`) USING BTREE,
  KEY `listorder` (`catid`,`status`,`listorder`,`id`) USING BTREE,
  KEY `catid` (`catid`,`status`,`id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='新闻表后期根据模型扩展';


CREATE TABLE `iriscms_page` (
  `catid` int NOT NULL DEFAULT '0',
  `title` varchar(160) NOT NULL,
  `keywords` varchar(40) NOT NULL,
  `content` text NOT NULL,
  `updatetime` int NOT NULL DEFAULT '0',
  KEY `catid` (`catid`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='单页内容表';


CREATE TABLE `iriscms_setting` (
  `key` varchar(50) NOT NULL,
  `value` text,
  `group` varchar(128) DEFAULT NULL,
  `default` varchar(128) DEFAULT NULL,
  `form_name` varchar(128) DEFAULT NULL,
  `editor` text,
  `listorder` int DEFAULT '0',
  PRIMARY KEY (`key`) USING BTREE,
  UNIQUE KEY `UQE_iriscms_setting_key` (`key`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='系统配置表';


CREATE TABLE `iriscms_slide` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
  `name1` varchar(255) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL COMMENT '地址',
  `imgurl` varchar(255) NOT NULL COMMENT '图片地址',
  `wapimgurl` varchar(255) DEFAULT NULL,
  `sigin` varchar(255) NOT NULL COMMENT '图片标识',
  `sort` int NOT NULL DEFAULT '1' COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='幻灯片';


CREATE TABLE `iriscms_wechat_member` (
  `id` int NOT NULL AUTO_INCREMENT,
  `openid` varchar(255) DEFAULT NULL,
  `mpid` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `sex` tinyint DEFAULT NULL,
  `headimgurl` varchar(255) DEFAULT NULL,
  `subscribe_scene` varchar(255) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='系统配置表';


CREATE TABLE `iriscms_wechat_message_log` (
  `logid` bigint NOT NULL DEFAULT '0',
  `content` varchar(255) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`logid`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='微信消息日志';




SET FOREIGN_KEY_CHECKS = @PREVIOUS_FOREIGN_KEY_CHECKS;


SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;


LOCK TABLES `iriscms_admin` WRITE;
ALTER TABLE `iriscms_admin` DISABLE KEYS;
INSERT INTO `iriscms_admin` (`userid`, `username`, `password`, `roleid`, `encrypt`, `lastloginip`, `lastlogintime`, `email`, `realname`) VALUES
	(1,'admin','5736a2a40f752bf2e82953702d25075b',1,'qmRlFL','::1',1474291850,'chenchengbin92111@gmail.com2','mirchen.com1'),
	(11,'test','834a33db060873a7a208617930edb29a',1,'5u2G0w','::1',0,'asdasd@asdasd.com1','ccc1');
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
	(9,'category','add'),
	(9,'category','delete'),
	(9,'category','edit'),
	(9,'category','list'),
	(9,'category','order'),
	(9,'category','view'),
	(9,'content','index'),
	(9,'content','news-list'),
	(9,'content','page'),
	(9,'content','right'),
	(9,'content','top'),
	(9,'system','left'),
	(9,'system','log-delete'),
	(9,'system','loglist'),
	(9,'system','logview'),
	(9,'system','menuadd'),
	(9,'system','menudelete'),
	(9,'system','menuedit'),
	(9,'system','menulist'),
	(9,'system','menuorder'),
	(9,'system','menuview');
ALTER TABLE `iriscms_admin_role_priv` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_articles` WRITE;
ALTER TABLE `iriscms_articles` DISABLE KEYS;
INSERT INTO `iriscms_articles` (`id`, `title`, `keywords`, `description`, `thumb`, `content`, `from_website`, `from_url`, `tags`, `published`, `images`, `catid`, `mid`, `refid`, `visit_count`, `status`, `created_time`, `updated_time`, `deleted_time`, `listorder`) VALUES
	(1,'ada','asd','asdasd','/upload/public/20200210/Qvj7ZlOvl7.jpg','<p>asdasdad</p>','asd','asd','asdad','是','/upload/public/20200210/Buz7l6Tj16.jpg',3,5,0,0,0,NULL,'2020-02-10 13:13:42','2020-02-11 13:05:28',345),
	(2,'ada','asd','asdasd','/upload/public/20200210/Qvj7ZlOvl7.jpg','<p>asdasdad</p>','asd','asd','asdad','是','/upload/public/20200210/Buz7l6Tj16.jpg',3,5,0,0,0,'2020-02-10 11:42:07','2020-02-10 12:39:59','2020-02-10 13:13:32',11),
	(3,'ada','asd','asdasd','/upload/public/20200210/Qvj7ZlOvl7.jpg','<p>asdasdad&quot;&quot;&nbsp; ``&nbsp; #</p>','asd33#','asd#','asdad``','是','/upload/public/20200210/Buz7l6Tj16.jpg',3,5,0,0,0,'2020-02-10 11:44:50','2020-02-10 12:39:59','2020-02-10 13:13:35',212),
	(4,'测试标题','asdasdad','asdasdasdasd','/upload/public/20200210/Qvj7ZlOvl7.jpg','<p>sdasd<img src="/upload/public/20200210/41nyeDFxxU.jpg" alt="41nyeDFxxU.jpg"/><br/></p>','asd33','asd','asdad,aasdasdasdasdasd,asdasdasd,asdasdasdasdasd,我去','是','/upload/public/20200210/Buz7l6Tj16.jpg',3,5,0,0,0,'2020-02-10 11:45:08','2020-02-10 19:44:45','2020-02-11 13:05:26',23),
	(5,'adasd','asdas','asdasd','/upload/public/20200211/MjvO7j6Pdd.jpg','<p>adadasd</p>','asd','http://www.baidu.com','打标签','是','/upload/public/20200211/OH7F7fV644.jpg',3,5,0,0,0,'2020-02-11 11:48:15','2020-02-11 11:55:28','2020-02-11 13:05:23',0);
ALTER TABLE `iriscms_articles` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_attachments` WRITE;
ALTER TABLE `iriscms_attachments` DISABLE KEYS;
INSERT INTO `iriscms_attachments` (`id`, `name`, `url`, `origin_name`, `size`, `upload_time`, `type`) VALUES
	(2,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(3,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(4,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(5,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(6,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(7,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(8,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(9,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(10,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(11,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(12,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(13,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(14,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(15,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(16,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(17,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(18,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(19,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(20,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(21,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(22,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(23,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(24,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(25,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(26,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(27,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(28,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(29,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(30,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(31,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(32,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(33,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(34,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(35,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(36,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(37,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(38,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(39,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(40,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(41,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(42,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(43,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(44,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(45,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(46,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(47,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(48,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(49,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(50,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(51,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(52,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(53,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(54,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(55,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(56,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(57,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(58,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(59,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(60,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(61,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(62,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(63,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(64,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(65,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(66,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(67,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(68,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(69,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(70,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(71,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(72,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(73,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(74,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(75,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(76,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(77,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(78,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(79,'HOT4zStb14.jpg','/upload/public/20200210/HOT4zStb14.jpg','64673db6d9.jpg',604762,'2020-02-10 17:09:21','img'),
	(80,'41nyeDFxxU.jpg','/upload/public/20200210/41nyeDFxxU.jpg','64673db6d9.jpg',218356,'2020-02-10 17:40:34','img'),
	(81,'Kdy4wdTWwO.jpg','/upload/public/20200210/Kdy4wdTWwO.jpg','64673db6d9.jpg',218356,'2020-02-10 17:42:23','img'),
	(82,'AQ7yfBTKC4.jpg','/upload/public/20200210/AQ7yfBTKC4.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:32','img'),
	(83,'eSId6e85oe.jpg','/upload/public/20200210/eSId6e85oe.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:47','img'),
	(84,'71okeV2pq8.jpg','/upload/public/20200210/71okeV2pq8.jpg','64673db6d9.jpg',218356,'2020-02-10 19:04:48','img'),
	(85,'g10BxfuPlf.png','/upload/public/20200210/g10BxfuPlf.png','K9k3WXVj3d.png',0,'2020-02-10 19:35:22','img'),
	(86,'E98Sy9Cb9T.png','/upload/public/20200210/E98Sy9Cb9T.png','N9DhuBERNF.png',0,'2020-02-10 19:38:10','img'),
	(87,'1efJjx17WH.jpg','/upload/public/20200210/1efJjx17WH.jpg','64673db6d9.jpg',604762,'2020-02-10 19:40:13','img'),
	(88,'J07ibEIWpi.jpg','/upload/public/20200211/J07ibEIWpi.jpg','64673db6d9.jpg',604762,'2020-02-11 11:26:57','img'),
	(89,'oyzQ8OsDTQ.jpg','/upload/public/20200211/oyzQ8OsDTQ.jpg','64673db6d9.jpg',604762,'2020-02-11 11:28:20','img'),
	(90,'yz4U6V0CXg.jpg','/upload/public/20200211/yz4U6V0CXg.jpg','64673db6d9.jpg',604762,'2020-02-11 11:29:55','img'),
	(91,'Ekf2kZ7kQS.jpg','/upload/public/20200211/Ekf2kZ7kQS.jpg','64673db6d9.jpg',604762,'2020-02-11 11:30:12','img'),
	(92,'Y3v355GWcz.jpg','/upload/public/20200211/Y3v355GWcz.jpg','64673db6d9.jpg',604762,'2020-02-11 11:30:54','img'),
	(93,'W2A95f6v1M.jpg','/upload/public/20200211/W2A95f6v1M.jpg','64673db6d9.jpg',604762,'2020-02-11 11:31:51','img'),
	(94,'tWXdAmSPLq.jpg','/upload/public/20200211/tWXdAmSPLq.jpg','64673db6d9.jpg',604762,'2020-02-11 11:32:30','img'),
	(95,'z2gvYbzb12.jpg','/upload/public/20200211/z2gvYbzb12.jpg','64673db6d9.jpg',604762,'2020-02-11 11:33:09','img'),
	(96,'PYl5gx1cs0.jpg','/upload/public/20200211/PYl5gx1cs0.jpg','64673db6d9.jpg',604762,'2020-02-11 11:34:17','img'),
	(97,'8ZI337YIgP.jpg','/upload/public/20200211/8ZI337YIgP.jpg','64673db6d9.jpg',604762,'2020-02-11 11:35:39','img'),
	(98,'MjvO7j6Pdd.jpg','/upload/public/20200211/MjvO7j6Pdd.jpg','64673db6d9.jpg',604762,'2020-02-11 11:48:01','img'),
	(99,'OH7F7fV644.jpg','/upload/public/20200211/OH7F7fV644.jpg','64673db6d9.jpg',604762,'2020-02-11 11:48:14','img'),
	(100,'86ZI22CG39.jpg','/upload/public/20200211/86ZI22CG39.jpg','64673db6d9.jpg',604762,'2020-02-11 12:12:48','img'),
	(101,'T6ia54P5Fb.jpg','/upload/public/20200211/T6ia54P5Fb.jpg','64673db6d9.jpg',604762,'2020-02-11 15:01:27','img'),
	(102,'TECwi9HrNX.jpg','/upload/public/20200211/TECwi9HrNX.jpg','64673db6d9.jpg',604762,'2020-02-11 21:07:09','img'),
	(103,'56ooFr2t27.jpg','/upload/public/20200212/56ooFr2t27.jpg','64673db6d9.jpg',604762,'2020-02-12 12:33:47','img'),
	(104,'gRl0i0cP7X.png','//public/20200212/gRl0i0cP7X.png','page.png',202109,'2020-02-12 19:36:23','img'),
	(105,'ixaGq0EChv.png','//public/20200212/ixaGq0EChv.png','page.png',202109,'2020-02-12 19:36:47','img'),
	(106,'xEAPFi0cYS.png','//public/20200212/xEAPFi0cYS.png','page.png',202109,'2020-02-12 19:37:09','img'),
	(107,'z9f5WOgTRa.png','public/20200212/z9f5WOgTRa.png','page.png',202109,'2020-02-12 19:47:42','img'),
	(108,'iwfz5POG9y.png','public/20200212/iwfz5POG9y.png','page.png',202109,'2020-02-12 19:48:23','img'),
	(109,'jy54Y6kQr6.png','public/20200212/jy54Y6kQr6.png','page.png',202109,'2020-02-12 19:49:21','img'),
	(110,'8enWIHKc7h.png','public/20200212/8enWIHKc7h.png','page.png',202109,'2020-02-12 19:50:06','img'),
	(111,'ww9yZHU57V.png','public/20200212/ww9yZHU57V.png','page.png',202109,'2020-02-12 19:50:28','img'),
	(112,'F215jUu1vJ.png','public/20200212/F215jUu1vJ.png','page.png',202109,'2020-02-12 19:50:49','img'),
	(113,'564u4p4rx4.png','/upload/public/20200212/564u4p4rx4.png','page.png',202109,'2020-02-12 19:51:43','img'),
	(114,'7CYWiFm5Bs.png','iriscms-test.oss-cn-beijing.aliyuncs.com/public/20200212/7CYWiFm5Bs.png','page.png',202109,'2020-02-12 19:57:41','img'),
	(115,'Qc0NY6ME36.png','http://iriscms-test.oss-cn-beijing.aliyuncs.com/public/20200212/Qc0NY6ME36.png','page.png',202109,'2020-02-12 19:58:23','img'),
	(116,'AYcL1slsZm.png','/upload/public/20200221/AYcL1slsZm.png','15d5e6410d49263.90261089.png',221324,'2020-02-21 11:59:07','img'),
	(117,'15yk028JIE.png','/upload/public/20200227/15yk028JIE.png','nopic.png',3137,'2020-02-27 14:05:38','img'),
	(118,'uc2x4oEbP0.png','/upload/public/20200227/uc2x4oEbP0.png','nopic.png',3137,'2020-02-27 14:06:48','img'),
	(119,'v3gR0WDhA6.png','/upload/public/20200227/v3gR0WDhA6.png','nopic.png',3137,'2020-02-27 14:07:01','img'),
	(120,'oZPnY55o76.png','/upload/public/20200227/oZPnY55o76.png','nopic.png',3137,'2020-02-27 14:07:20','img');
ALTER TABLE `iriscms_attachments` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_category` WRITE;
ALTER TABLE `iriscms_category` DISABLE KEYS;
INSERT INTO `iriscms_category` (`catid`, `type`, `parentid`, `catname`, `description`, `url`, `listorder`, `ismenu`, `index_tpl`, `list_tpl`, `detail_tpl`, `thumb`, `model_id`, `tpl_prefix`, `home_tpl`, `content_tpl`) VALUES
	(8,0,0,'栏目','','',1,1,'','','','',5,NULL,NULL,NULL),
	(18,2,8,'外部链接','','',0,1,'','','','',0,NULL,NULL,NULL),
	(17,0,8,'开启','','',2,1,'','','','',5,NULL,NULL,NULL);
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
INSERT INTO `iriscms_document_model` (`id`, `name`, `table`, `enabled`, `model_type`, `fe_tpl_index`, `fe_tpl_list`, `fe_tpl_detail`, `deleted_at`, `field_show_in_list`, `formatters`, `execed`) VALUES
	(5,'系统模型','articles',1,0,'resources/views/backend/admin_editpwd.html','resources/views/backend/category_edit.html','resources/views/backend/index_index.html',NULL,'{"content":{"show":false,"formatter":""},"description":{"show":false,"formatter":""},"from_url":{"show":false,"formatter":""},"from_website":{"show":false,"formatter":""},"images":{"show":false,"formatter":""},"keywords":{"show":true,"formatter":""},"published":{"show":false,"formatter":""},"tags":{"show":false,"formatter":""},"thumb":{"show":false,"formatter":""},"title":{"show":true,"formatter":""}}','function titleFormatter(title) {return title + "111";};',0);
ALTER TABLE `iriscms_document_model` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_document_model_dsl` WRITE;
ALTER TABLE `iriscms_document_model_dsl` DISABLE KEYS;
INSERT INTO `iriscms_document_model_dsl` (`id`, `mid`, `form_name`, `table_field`, `html`, `required`, `datasource`, `required_tips`, `validator`, `deleted_at`, `field_type`, `default`) VALUES
	(520,5,'标题','title','<input class="easyui-textbox" {{attr}} value="{{value}}" style="width:300px">',1,'','','',NULL,1,''),
	(521,5,'关键字','keywords','<input class="easyui-textbox" {{attr}} value="{{value}}" style="width:300px">',0,'','','',NULL,1,''),
	(522,5,'描述','description','<input class="easyui-textbox" {{attr}} value="{{value}}" style="height:80px; width: 300px;"  multiline />',0,'','','',NULL,2,''),
	(523,5,'封面','thumb','<images />',1,'','封面图片必须上传','',NULL,11,''),
	(524,5,'内容','content','<editor />',1,'','','',NULL,3,''),
	(525,5,'来源','from_website','<input class="easyui-textbox" {{attr}} value="{{value}}" style="width:300px">',0,'','','',NULL,1,''),
	(526,5,'来源地址','from_url','<input class="easyui-textbox" {{attr}} value="{{value}}" style="width:300px">',0,'','','\'url\'',NULL,1,''),
	(527,5,'标签','tags','<tags />',0,'','','',NULL,15,''),
	(528,5,'是否发布','published','<input class="easyui-switchbutton" value="{{value}}" {{attr}} {{default}} />',0,'["是", "否"]','','',NULL,13,'是'),
	(529,5,'图集','images','<mul-images />',1,'','','',NULL,12,'');
ALTER TABLE `iriscms_document_model_dsl` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_document_model_field` WRITE;
ALTER TABLE `iriscms_document_model_field` DISABLE KEYS;
INSERT INTO `iriscms_document_model_field` (`id`, `name`, `type`, `desc`, `html`) VALUES
	(1,'单行文本','varchar','常用字段，如文章标题、作者等都属于直接输入少量内容的文本，设置这个文本之后需要指定文本长度，默认为250，如果大于255则为text类型','<input class="easyui-textbox" {{attr}} value="{{value}}" style="width:300px">'),
	(2,'多行文本','text','也是较为常用的字段类型，如个人简介、产品描述都可以使用多行文本进行存储','<input class="easyui-textbox" {{attr}} value="{{value}}" style="height:80px; width: 300px;"  multiline />'),
	(3,'HTML文本','text','编辑器编辑产生的html内容，用于比较复杂的内容形式, 可以认为是附带编辑器的多行文本','<editor />'),
	(4,'附件','varshar','前端表现为input[file]类型,可以上传并且返回一个相对地址','<input class="easyui-filebox" value="{{value}}" style="width:300px">'),
	(5,'下拉框','varchar','下拉选择，一般用于如软件类型、语言类型等字段','<input class="easyui-combobox" {{attr}} style="width:300px;" value="{{value}}" />'),
	(6,'联动类型','varchar','一种数组形式的数据类型，请使用url接口方式提供','<select class="easyui-combotree" {{attr}} style="width:200px;">暂未实现</select>'),
	(7,'单选框','varchar','平铺显示, 可以认为是下拉框的展开, 根据数据源展开为排列的组件','<input class="easyui-radiobutton" {{attr}} value="{{value}}" {{default}} />'),
	(8,'多选框','varchar','多选框, 平铺显示为多个选项,根据数据源展开为排列组件','<input class="easyui-checkbox" {{attr}} value="{{value}}" {{default}} />'),
	(9,'整数类型','int','常用字段, 仅能输入数字','<input type="text" class="easyui-numberbox" value="{{value}}" {{attr}} />'),
	(10,'浮点类型','float','常用字段, 仅能输入浮点数(小数)','<input type="text" class="easyui-numberbox" value="{{value}}" {{attr}} />'),
	(11,'单图上传','varchar','常用字段, 会生成一个单图上传框','<images />'),
	(12,'多图上传','varchar','生成一个多图上传的组件','<mul-images />'),
	(13,'开关按钮','tinyint','用于做开关值的组件, 打开为1, 关闭为0','<input class="easyui-switchbutton" value="{{value}}" {{attr}} {{default}} />'),
	(14,'日历组件','datetime','选择日期组件','<input class="easyui-datetimebox" style="width:300px" value="{{value}}" {{attr}} />'),
	(15,'多选标签','varchar','可以记录标签，并多选','<tags />');
ALTER TABLE `iriscms_document_model_field` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_link` WRITE;
ALTER TABLE `iriscms_link` DISABLE KEYS;
INSERT INTO `iriscms_link` (`linkid`, `linktype`, `name`, `url`, `logo`, `introduce`, `listorder`, `passed`, `addtime`) VALUES
	(5,0,'aaaa','https://www.baidu.com','','',1,1,'2020-02-29 21:19:40'),
	(6,0,'aaaa','http://www.baidu.com','','',2,1,'2020-02-29 21:19:30');
ALTER TABLE `iriscms_link` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_log` WRITE;
ALTER TABLE `iriscms_log` DISABLE KEYS;
ALTER TABLE `iriscms_log` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_member` WRITE;
ALTER TABLE `iriscms_member` DISABLE KEYS;
INSERT INTO `iriscms_member` (`id`, `account`, `password`, `avatar`, `nickname`, `integral`, `sale_integral`, `draw_account`, `telphone`, `qq`, `description`, `created_at`, `updated_at`, `email`, `enabled`, `verify_token`) VALUES
	(1,'xiusin','159781','','陈二皮',1231211111,0,'','123123','1111222','','2019-01-24 11:40:00','2019-01-24 11:40:00','159781@11.com',1,'4b32a22c-5787-4d0b-98f2-ed5b0779bbcb');
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
	(9,'系统设置',2,'setting','left','',0,2,'1'),
	(10,'站点设置',9,'setting','site','',0,1,'1'),
	(11,'管理员设置',2,'admin','left','',0,4,'1'),
	(12,'管理员管理',11,'admin','memberlist','',0,1,'1'),
	(13,'角色管理',11,'admin','rolelist','',0,2,'1'),
	(14,'后台管理',2,'system','left','',0,1,'1'),
	(15,'日志管理',14,'system','loglist','',0,1,'0'),
	(16,'菜单管理',9,'system','menulist','',0,2,'1'),
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
	(35,'内容管理',36,'content','index','',0,0,'1'),
	(36,'栏目管理',34,'category','list','',0,0,'1'),
	(37,'查看栏目',36,'category','view','',0,0,'1'),
	(38,'添加栏目',36,'category','add','',0,0,'1'),
	(39,'编辑栏目',36,'category','edit','',0,0,'1'),
	(40,'删除栏目',36,'category','delete','',0,0,'1'),
	(41,'栏目排序',36,'category','order','',0,0,'1'),
	(55,'会员管理',2,'user','list','',0,5,'1'),
	(56,'会员列表',55,'user','list','',0,0,'1'),
	(57,'会员信息',56,'user','info','',0,0,'1'),
	(58,'微信管理',2,'wechat','userlist','',0,7,'1'),
	(59,'微信会员信息',58,'wechat','userinfo','',0,0,'0'),
	(60,'编辑会员',55,'user','edit','',0,0,'0'),
	(66,'友链管理',2,'link','list','',0,0,'1'),
	(62,'模型管理',9,'model','list','',0,1,'1'),
	(64,'添加模型',62,'model','add','',0,0,'1'),
	(65,'微信会员列表',58,'wechat','userlist','',0,0,'1'),
	(67,'友链管理',2,'link','list','',0,88,'1'),
	(68,'友链管理',67,'link','list','',0,88,'1'),
	(69,'友链添加',68,'link','add','',0,0,'1'),
	(70,'友链编辑',68,'link','edit','',0,0,'1'),
	(71,'友链删除',68,'link','delete','',0,0,'1'),
	(72,'友链排序',68,'link','order','',0,0,'1'),
	(73,'数据库管理',2,'database','manager','',0,3,'1'),
	(74,'数据库管理',73,'database','manager','',0,0,'1'),
	(75,'数据库备份',74,'database','backup','',0,0,'1'),
	(76,'数据库优化',74,'database','optimize','',0,0,'1'),
	(77,'数据库修复',74,'database','repair','',0,0,'1'),
	(78,'备份列表',73,'database','backup-list','',0,0,'1'),
	(79,'资源管理',2,'assets-manager','list','',0,3,'1'),
	(80,'资源列表',79,'assets-manager','list','',0,0,'1'),
	(81,'添加资源',80,'assets-manager','add','',0,0,'1'),
	(82,'修改资源',80,'assets-manager','edit','',0,0,'1'),
	(83,'附件管理',2,'attachments','list','',0,0,'0'),
	(84,'附件列表',79,'attachments','list','',0,0,'1');
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
INSERT INTO `iriscms_setting` (`key`, `value`, `group`, `default`, `form_name`, `editor`, `listorder`) VALUES
	('EMAIL_USER','','邮箱设置',NULL,'用户名','text',6),
	('EMAIL_PWD','','邮箱设置',NULL,'密码','text',7),
	('EMAIL_SMTP','','邮箱设置',NULL,'SMTP服务器','text',8),
	('EMAIL_EMAIL','','邮箱设置',NULL,'邮箱地址','text',9),
	('SITE_ICP','','前台设置',NULL,'备案号','text',4),
	('SITE_KEYWORDS','iriscms','前台设置',NULL,'关键字','text',2),
	('SITE_DESCRIPTION','iriscms','前台设置',NULL,'描述','text',3),
	('SITE_TITLE','iriscms','前台设置',NULL,'站点标题','text',1),
	('WX_TOKEN','','微信配置',NULL,'TOKEN','text',13),
	('SITE_OPEN','开启','前台设置','开启','站点开启','{"type":"checkbox", "options": {"on":"开启", "off":"关闭"}}',0),
	('EMAIL_PORT','25','邮箱设置','25','端口','text',10),
	('WX_APPSECRET','','微信配置',NULL,'APPSECTET','text',12),
	('WX_AESKEY','','微信配置',NULL,'AESKEY','text',14),
	('DATAGRID_PAGE_SIZE','25','前台设置','25','列表默认分页数','text',5),
	('WX_APPID','','微信配置',NULL,'APPID','text',11),
	('UPLOAD_DIR','resources/assets/upload','存储配置','resources/assets/upload','存储目录','text',21),
	('UPLOAD_ENGINE','本地存储','存储配置','local','存储引擎','{"type":"combogrid","options":{"idField":"key","textField":"key","fitColumns":true,"columns":[[{"field":"key","title":"存储引擎","width":120}]],"data":[{"key":"本地存储","value":"local"},{"key":"OSS存储","value":"oss"}]}}',20),
	('UPLOAD_IMG_TYPES','jpg,jpeg,png,gif,bmp','存储配置','jpg,jpeg,png,gif,bmp','可上传图片类型','text',22),
	('UPLOAD_URL_PREFIX','/upload','存储配置','upload','地址前缀','text',23),
	('UPLOAD_DATABASE_PASS','123456','存储配置','','备份数据库密码','text',0),
	('DATABASE_AUTO_BACKUP_TIME','02','存储配置','02','自动备份时间(仅小时)','text',0);
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



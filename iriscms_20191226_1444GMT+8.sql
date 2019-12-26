#
# SQL Export
# Created by Querious (201054)
# Created: December 26, 2019 at 2:45:03 PM GMT+8
# Encoding: Unicode (UTF-8)
#


CREATE DATABASE IF NOT EXISTS `iriscms` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
USE `iriscms`;




SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;


SET FOREIGN_KEY_CHECKS = @PREVIOUS_FOREIGN_KEY_CHECKS;


SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;


LOCK TABLES `iriscms_admin` WRITE;
ALTER TABLE `iriscms_admin` DISABLE KEYS;
INSERT INTO `iriscms_admin` (`userid`, `username`, `password`, `roleid`, `encrypt`, `lastloginip`, `lastlogintime`, `email`, `realname`) VALUES 
	(1,'admin','0087af20a551a8b804f89469534b7859',1,'qmRlFL','::1',1474291850,'chenchengbin92@gmail.com','mirchen.com'),
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
INSERT INTO `iriscms_category` (`catid`, `type`, `parentid`, `catname`, `description`, `url`, `listorder`, `ismenu`, `tpl`, `home_tpl`, `content_tpl`, `thumb`, `tpl_prefix`) VALUES 
	(27,0,0,'免费视频','','',2,1,'','','','','content_free_video_'),
	(28,0,0,'免费图书','','',3,1,'','','','','content_free_book_'),
	(29,0,0,'付费视频','','',4,1,'','','','','content_paid_video_'),
	(30,0,0,'付费图书','','',5,1,'','','','','content_paid_book_'),
	(31,0,30,'MySQL','','',0,1,'','','','',''),
	(32,0,30,'GO语言','','',0,1,'','','','',''),
	(33,0,30,'PYTHON','','',0,1,'','','','',''),
	(34,0,29,'大数据','','',0,1,'','','','',''),
	(35,0,29,'PHP7','','',0,1,'','','','',''),
	(36,0,29,'NODEJS','','',0,1,'','','','',''),
	(37,0,28,'PYTHON','','',0,1,'','','','',''),
	(38,0,27,'PYTHON','','',0,1,'','','','',''),
	(39,0,0,'软件下载','','',1,1,'','','','','');
ALTER TABLE `iriscms_category` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_category_priv` WRITE;
ALTER TABLE `iriscms_category_priv` DISABLE KEYS;
ALTER TABLE `iriscms_category_priv` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_content` WRITE;
ALTER TABLE `iriscms_content` DISABLE KEYS;
INSERT INTO `iriscms_content` (`id`, `catid`, `title`, `thumb`, `keywords`, `description`, `content`, `listorder`, `status`, `recommend`, `pwd_type`, `money`, `created_at`, `updated_at`, `deleted_at`, `source_url`, `source_pwd`, `catids`, `tags`, `userid`) VALUES 
	(1,30,'项目列表页：导出复制的作品，一直exporting中；','','','','<p><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span><span style="white-space: normal;">12312312312312312312312</span></p>',0,1,0,1,50,1548300082,1548300082,0,'','123123','','sss,ss',0);
ALTER TABLE `iriscms_content` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_link` WRITE;
ALTER TABLE `iriscms_link` DISABLE KEYS;
ALTER TABLE `iriscms_link` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `iriscms_log` WRITE;
ALTER TABLE `iriscms_log` DISABLE KEYS;
INSERT INTO `iriscms_log` (`logid`, `controller`, `action`, `querystring`, `userid`, `username`, `ip`, `time`) VALUES 
	(435,'category','list','/b/category/list?menuid=36&&_=1576660545834',1,'admin','::1','2019-12-18 17:15:45'),
	(434,'content','right','/b/content/right',1,'admin','::1','2019-12-18 17:15:45'),
	(433,'content','right','/b/content/right',1,'admin','::1','2019-12-18 17:15:45'),
	(432,'content','right','/b/content/right?_=1576660545381',1,'admin','::1','2019-12-18 17:15:45'),
	(431,'content','index','/b/content/index?menuid=35&&_=1576660545355',1,'admin','::1','2019-12-18 17:15:45'),
	(430,'category','list','/b/category/list?menuid=36&&_=1576660540378',1,'admin','::1','2019-12-18 17:15:40'),
	(429,'content','right','/b/content/right',1,'admin','::1','2019-12-18 17:15:38'),
	(428,'content','right','/b/content/right',1,'admin','::1','2019-12-18 17:15:38'),
	(427,'content','right','/b/content/right?_=1576660538305',1,'admin','::1','2019-12-18 17:15:38'),
	(426,'content','index','/b/content/index?menuid=35&&_=1576660538271',1,'admin','::1','2019-12-18 17:15:38'),
	(425,'category','list','/b/category/list?grid=treegrid',1,'admin','::1','2019-12-18 17:15:36'),
	(424,'category','list','/b/category/list?menuid=36&&_=1576660536848',1,'admin','::1','2019-12-18 17:15:36'),
	(423,'setting','site','/b/setting/site?grid=propertygrid',1,'admin','::1','2019-12-18 17:15:34'),
	(422,'setting','site','/b/setting/site?menuid=10&&_=1576660534221',1,'admin','::1','2019-12-18 17:15:34'),
	(421,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-18 17:15:33'),
	(420,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576660532932',1,'admin','::1','2019-12-18 17:15:32'),
	(419,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-18 17:15:32'),
	(418,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576660532403',1,'admin','::1','2019-12-18 17:15:32'),
	(417,'setting','site','/b/setting/site?grid=propertygrid',1,'admin','::1','2019-12-17 17:09:58'),
	(416,'setting','site','/b/setting/site?menuid=10&&_=1576573798645',1,'admin','::1','2019-12-17 17:09:58'),
	(415,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 17:09:57'),
	(414,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576573797170',1,'admin','::1','2019-12-17 17:09:57'),
	(413,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 17:09:56'),
	(412,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576573796576',1,'admin','::1','2019-12-17 17:09:56'),
	(411,'system','menulist','/b/system/menulist?menuid=16&&_=1576573794352',1,'admin','::1','2019-12-17 17:09:54'),
	(410,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','::1','2019-12-17 17:09:53'),
	(409,'system','loglist','/b/system/loglist?menuid=15&&_=1576573793709',1,'admin','::1','2019-12-17 17:09:53'),
	(408,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-17 17:09:53'),
	(407,'system','menulist','/b/system/menulist?menuid=16&&_=1576573792989',1,'admin','::1','2019-12-17 17:09:53'),
	(406,'content','news-list','/b/content/news-list?grid=datagrid&catid=35&page=1&rows=10',1,'admin','::1','2019-12-17 17:09:51'),
	(405,'content','news-list','/b/content/news-list?catid=35&_=1576573790739',1,'admin','::1','2019-12-17 17:09:50'),
	(404,'content','news-list','/b/content/news-list?grid=datagrid&catid=36&page=1&rows=10',1,'admin','::1','2019-12-17 17:09:50'),
	(403,'content','news-list','/b/content/news-list?catid=36&_=1576573790200',1,'admin','::1','2019-12-17 17:09:50'),
	(402,'content','right','/b/content/right',1,'admin','::1','2019-12-17 17:09:49'),
	(401,'content','right','/b/content/right',1,'admin','::1','2019-12-17 17:09:49'),
	(400,'content','right','/b/content/right?_=1576573789175',1,'admin','::1','2019-12-17 17:09:49'),
	(399,'content','index','/b/content/index?menuid=35&&_=1576573789128',1,'admin','::1','2019-12-17 17:09:49'),
	(398,'category','list','/b/category/list?grid=treegrid',1,'admin','::1','2019-12-17 17:09:47'),
	(397,'category','list','/b/category/list?menuid=36&&_=1576573786868',1,'admin','::1','2019-12-17 17:09:46'),
	(396,'content','right','/b/content/right',1,'admin','::1','2019-12-17 17:09:45'),
	(395,'content','right','/b/content/right',1,'admin','::1','2019-12-17 17:09:45'),
	(394,'content','right','/b/content/right?_=1576573785687',1,'admin','::1','2019-12-17 17:09:45'),
	(393,'content','index','/b/content/index?menuid=35&&_=1576573785653',1,'admin','::1','2019-12-17 17:09:45'),
	(392,'system','menulist','/b/system/menulist?menuid=16&&_=1576573392910',11,'test','::1','2019-12-17 17:03:12'),
	(391,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 17:03:12'),
	(390,'system','loglist','/b/system/loglist?menuid=15&&_=1576573392322',11,'test','::1','2019-12-17 17:03:12'),
	(389,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 17:03:11'),
	(388,'system','menulist','/b/system/menulist?menuid=16&&_=1576573391627',11,'test','::1','2019-12-17 17:03:11'),
	(387,'system','loglist','/b/system/loglist?page=13&rows=20',11,'test','::1','2019-12-17 16:17:15'),
	(386,'system','loglist','/b/system/loglist?page=12&rows=20',11,'test','::1','2019-12-17 16:17:14'),
	(385,'system','loglist','/b/system/loglist?page=11&rows=20',11,'test','::1','2019-12-17 16:17:14'),
	(384,'system','loglist','/b/system/loglist?page=10&rows=20',11,'test','::1','2019-12-17 16:17:14'),
	(383,'system','loglist','/b/system/loglist?page=9&rows=20',11,'test','::1','2019-12-17 16:17:11'),
	(382,'system','loglist','/b/system/loglist?page=8&rows=20',11,'test','::1','2019-12-17 16:17:11'),
	(381,'system','loglist','/b/system/loglist?page=7&rows=20',11,'test','::1','2019-12-17 16:17:11'),
	(380,'system','loglist','/b/system/loglist?page=6&rows=20',11,'test','::1','2019-12-17 16:17:11'),
	(379,'system','loglist','/b/system/loglist?page=5&rows=20',11,'test','::1','2019-12-17 16:17:06'),
	(378,'system','loglist','/b/system/loglist?page=5&rows=50',11,'test','::1','2019-12-17 16:17:04'),
	(377,'system','loglist','/b/system/loglist?page=4&rows=50',11,'test','::1','2019-12-17 16:17:04'),
	(376,'system','loglist','/b/system/loglist?page=3&rows=50',11,'test','::1','2019-12-17 16:17:03'),
	(375,'system','loglist','/b/system/loglist?page=2&rows=50',11,'test','::1','2019-12-17 16:17:00'),
	(374,'system','loglist','/b/system/loglist?page=1&rows=50',11,'test','::1','2019-12-17 16:16:58'),
	(373,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 16:16:52'),
	(372,'system','loglist','/b/system/loglist?menuid=15&&_=1576570612835',11,'test','::1','2019-12-17 16:16:52'),
	(371,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 16:16:52'),
	(370,'system','menulist','/b/system/menulist?menuid=16&&_=1576570611955',11,'test','::1','2019-12-17 16:16:51'),
	(369,'system','menulist','/b/system/menulist?menuid=16&&_=1576570582972',11,'test','::1','2019-12-17 16:16:22'),
	(368,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 16:16:22'),
	(367,'system','loglist','/b/system/loglist?menuid=15&&_=1576570582358',11,'test','::1','2019-12-17 16:16:22'),
	(366,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 16:16:22'),
	(365,'system','menulist','/b/system/menulist?menuid=16&&_=1576570581867',11,'test','::1','2019-12-17 16:16:21'),
	(364,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 15:11:03'),
	(363,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 15:10:47'),
	(362,'system','menulist','/b/system/menulist?menuid=16&&_=1576566646952',11,'test','::1','2019-12-17 15:10:46'),
	(361,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 15:10:10'),
	(360,'system','menulist','/b/system/menulist?menuid=16&&_=1576566610010',11,'test','::1','2019-12-17 15:10:10'),
	(359,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-17 15:10:02'),
	(358,'system','loglist','/b/system/loglist?page=3&rows=10',11,'test','::1','2019-12-17 15:06:54'),
	(357,'system','loglist','/b/system/loglist?page=2&rows=10',11,'test','::1','2019-12-17 15:06:52'),
	(356,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-17 15:05:54'),
	(355,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-17 15:05:52'),
	(354,'system','loglist','/b/system/loglist?menuid=15&&_=1576566352018',11,'test','::1','2019-12-17 15:05:52'),
	(353,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 15:05:51'),
	(352,'system','menulist','/b/system/menulist?menuid=16&&_=1576566351112',11,'test','::1','2019-12-17 15:05:51'),
	(351,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 15:05:50'),
	(350,'system','loglist','/b/system/loglist?menuid=15&&_=1576566350127',11,'test','::1','2019-12-17 15:05:50'),
	(349,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 15:03:54'),
	(348,'system','menulist','/b/system/menulist?menuid=16&&_=1576566233912',11,'test','::1','2019-12-17 15:03:53'),
	(347,'system','loglist','/b/system/loglist?page=7&rows=10',11,'test','::1','2019-12-17 15:03:52'),
	(346,'system','loglist','/b/system/loglist?page=6&rows=10',11,'test','::1','2019-12-17 15:03:52'),
	(345,'system','loglist','/b/system/loglist?page=5&rows=10',11,'test','::1','2019-12-17 15:03:52'),
	(344,'system','loglist','/b/system/loglist?page=4&rows=10',11,'test','::1','2019-12-17 15:03:52'),
	(343,'system','loglist','/b/system/loglist?page=3&rows=10',11,'test','::1','2019-12-17 15:03:51'),
	(342,'system','loglist','/b/system/loglist?page=2&rows=10',11,'test','::1','2019-12-17 15:03:50'),
	(341,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 15:01:31'),
	(340,'system','loglist','/b/system/loglist?menuid=15&&_=1576566090676',11,'test','::1','2019-12-17 15:01:30'),
	(339,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 15:01:28'),
	(338,'system','menulist','/b/system/menulist?menuid=16&&_=1576566088116',11,'test','::1','2019-12-17 15:01:28'),
	(337,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-17 15:01:20'),
	(336,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-17 14:56:22'),
	(335,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-17 14:56:19'),
	(334,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:56:05'),
	(333,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-17 14:55:52'),
	(135,'content','index','/b/content/index?menuid=35&&_=1576550314690',1,'admin','127.0.0.1','2019-12-17 10:38:34'),
	(136,'content','right','/b/content/right?_=1576550314780',1,'admin','127.0.0.1','2019-12-17 10:38:34'),
	(137,'content','right','/b/content/right',1,'admin','127.0.0.1','2019-12-17 10:38:34'),
	(138,'content','right','/b/content/right',1,'admin','127.0.0.1','2019-12-17 10:38:34'),
	(139,'category','list','/b/category/list?menuid=36&&_=1576550315676',1,'admin','127.0.0.1','2019-12-17 10:38:35'),
	(140,'category','list','/b/category/list?grid=treegrid',1,'admin','127.0.0.1','2019-12-17 10:38:35'),
	(141,'system','menulist','/b/system/menulist?menuid=16&&_=1576550318895',1,'admin','127.0.0.1','2019-12-17 10:38:38'),
	(142,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','127.0.0.1','2019-12-17 10:38:39'),
	(143,'system','loglist','/b/system/loglist?menuid=15&&_=1576550319796',1,'admin','127.0.0.1','2019-12-17 10:38:39'),
	(144,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','127.0.0.1','2019-12-17 10:38:40'),
	(145,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576550322440',1,'admin','127.0.0.1','2019-12-17 10:38:42'),
	(146,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','127.0.0.1','2019-12-17 10:38:42'),
	(147,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576550323550',1,'admin','127.0.0.1','2019-12-17 10:38:43'),
	(148,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','127.0.0.1','2019-12-17 10:38:43'),
	(149,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576550324624',1,'admin','127.0.0.1','2019-12-17 10:38:44'),
	(150,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','127.0.0.1','2019-12-17 10:38:44'),
	(151,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576550325604',1,'admin','127.0.0.1','2019-12-17 10:38:45'),
	(152,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','127.0.0.1','2019-12-17 10:38:45'),
	(153,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576550327263',1,'admin','127.0.0.1','2019-12-17 10:38:47'),
	(154,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','127.0.0.1','2019-12-17 10:38:47'),
	(155,'content','index','/b/content/index?menuid=35&&_=1576550336530',1,'admin','127.0.0.1','2019-12-17 10:38:56'),
	(156,'content','right','/b/content/right?_=1576550336621',1,'admin','127.0.0.1','2019-12-17 10:38:56'),
	(157,'content','right','/b/content/right',1,'admin','127.0.0.1','2019-12-17 10:38:56'),
	(158,'content','right','/b/content/right',1,'admin','127.0.0.1','2019-12-17 10:38:56'),
	(159,'category','list','/b/category/list?menuid=36&&_=1576550337288',1,'admin','127.0.0.1','2019-12-17 10:38:57'),
	(160,'category','list','/b/category/list?grid=treegrid',1,'admin','127.0.0.1','2019-12-17 10:38:57'),
	(161,'content','index','/b/content/index?menuid=35&&_=1576550339277',1,'admin','127.0.0.1','2019-12-17 10:38:59'),
	(162,'content','right','/b/content/right?_=1576550339370',1,'admin','127.0.0.1','2019-12-17 10:38:59'),
	(163,'content','right','/b/content/right',1,'admin','127.0.0.1','2019-12-17 10:38:59'),
	(164,'content','right','/b/content/right',1,'admin','127.0.0.1','2019-12-17 10:38:59'),
	(165,'category','list','/b/category/list?menuid=36&&_=1576550339865',1,'admin','127.0.0.1','2019-12-17 10:38:59'),
	(166,'category','list','/b/category/list?grid=treegrid',1,'admin','127.0.0.1','2019-12-17 10:39:00'),
	(167,'system','loglist','/b/system/loglist?menuid=15&&_=1576550342172',1,'admin','127.0.0.1','2019-12-17 10:39:02'),
	(168,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','127.0.0.1','2019-12-17 10:39:02'),
	(169,'system','menulist','/b/system/menulist?menuid=16&&_=1576550343251',1,'admin','127.0.0.1','2019-12-17 10:39:03'),
	(170,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','127.0.0.1','2019-12-17 10:39:03'),
	(171,'system','loglist','/b/system/loglist?menuid=15&&_=1576550343901',1,'admin','127.0.0.1','2019-12-17 10:39:03'),
	(172,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','127.0.0.1','2019-12-17 10:39:04'),
	(173,'category','list','/b/category/list?menuid=36&&_=1576563928708',1,'admin','::1','2019-12-17 14:25:28'),
	(174,'category','list','/b/category/list?grid=treegrid',1,'admin','::1','2019-12-17 14:25:28'),
	(175,'content','index','/b/content/index?menuid=35&&_=1576563929759',1,'admin','::1','2019-12-17 14:25:29'),
	(176,'content','right','/b/content/right?_=1576563929798',1,'admin','::1','2019-12-17 14:25:29'),
	(177,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:29'),
	(178,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:29'),
	(179,'category','list','/b/category/list?menuid=36&&_=1576563931143',1,'admin','::1','2019-12-17 14:25:31'),
	(180,'content','index','/b/content/index?menuid=35&&_=1576563931661',1,'admin','::1','2019-12-17 14:25:31'),
	(181,'content','right','/b/content/right?_=1576563931694',1,'admin','::1','2019-12-17 14:25:31'),
	(182,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:31'),
	(183,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:31'),
	(184,'category','list','/b/category/list?menuid=36&&_=1576563932681',1,'admin','::1','2019-12-17 14:25:32'),
	(185,'content','index','/b/content/index?menuid=35&&_=1576563933106',1,'admin','::1','2019-12-17 14:25:33'),
	(186,'content','right','/b/content/right?_=1576563933139',1,'admin','::1','2019-12-17 14:25:33'),
	(187,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:33'),
	(188,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:33'),
	(189,'category','list','/b/category/list?menuid=36&&_=1576563934001',1,'admin','::1','2019-12-17 14:25:34'),
	(190,'content','index','/b/content/index?menuid=35&&_=1576563934628',1,'admin','::1','2019-12-17 14:25:34'),
	(191,'content','right','/b/content/right?_=1576563934660',1,'admin','::1','2019-12-17 14:25:34'),
	(192,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:34'),
	(193,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:34'),
	(194,'category','list','/b/category/list?menuid=36&&_=1576563935447',1,'admin','::1','2019-12-17 14:25:35'),
	(195,'content','index','/b/content/index?menuid=35&&_=1576563936197',1,'admin','::1','2019-12-17 14:25:36'),
	(196,'content','right','/b/content/right?_=1576563936228',1,'admin','::1','2019-12-17 14:25:36'),
	(197,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:36'),
	(198,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:25:36'),
	(199,'system','loglist','/b/system/loglist?menuid=15&&_=1576563939115',1,'admin','::1','2019-12-17 14:25:39'),
	(200,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','::1','2019-12-17 14:25:39'),
	(201,'system','menulist','/b/system/menulist?menuid=16&&_=1576563952651',1,'admin','::1','2019-12-17 14:25:52'),
	(202,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-17 14:25:52'),
	(203,'system','loglist','/b/system/loglist?menuid=15&&_=1576563953561',1,'admin','::1','2019-12-17 14:25:53'),
	(204,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','::1','2019-12-17 14:25:53'),
	(205,'system','menulist','/b/system/menulist?menuid=16&&_=1576563961798',1,'admin','::1','2019-12-17 14:26:01'),
	(206,'system','loglist','/b/system/loglist?menuid=15&&_=1576563962301',1,'admin','::1','2019-12-17 14:26:02'),
	(207,'system','menulist','/b/system/menulist?menuid=16&&_=1576563964434',1,'admin','::1','2019-12-17 14:26:04'),
	(208,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-17 14:26:04'),
	(209,'system','loglist','/b/system/loglist?menuid=15&&_=1576563965101',1,'admin','::1','2019-12-17 14:26:05'),
	(210,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','::1','2019-12-17 14:26:05'),
	(211,'system','menulist','/b/system/menulist?menuid=16&&_=1576563984246',1,'admin','::1','2019-12-17 14:26:24'),
	(212,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-17 14:26:24'),
	(213,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576563994844',1,'admin','::1','2019-12-17 14:26:34'),
	(214,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:26:34'),
	(215,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576563995492',1,'admin','::1','2019-12-17 14:26:35'),
	(216,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:26:35'),
	(217,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576563996052',1,'admin','::1','2019-12-17 14:26:36'),
	(218,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576563996560',1,'admin','::1','2019-12-17 14:26:36'),
	(219,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576563997086',1,'admin','::1','2019-12-17 14:26:37'),
	(220,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576563997564',1,'admin','::1','2019-12-17 14:26:37'),
	(221,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576563998076',1,'admin','::1','2019-12-17 14:26:38'),
	(222,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576563998767',1,'admin','::1','2019-12-17 14:26:38'),
	(223,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576563999973',1,'admin','::1','2019-12-17 14:26:39'),
	(224,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576564000400',1,'admin','::1','2019-12-17 14:26:40'),
	(225,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576564001741',1,'admin','::1','2019-12-17 14:26:41'),
	(226,'system','menulist','/b/system/menulist?menuid=16&&_=1576564008778',1,'admin','::1','2019-12-17 14:26:48'),
	(227,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-17 14:26:48'),
	(228,'system','loglist','/b/system/loglist?menuid=15&&_=1576564009681',1,'admin','::1','2019-12-17 14:26:49'),
	(229,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','::1','2019-12-17 14:26:49'),
	(230,'system','log-delete','/b/system/log-delete',1,'admin','::1','2019-12-17 14:26:50'),
	(231,'system','log-delete','/b/system/log-delete',1,'admin','::1','2019-12-17 14:26:55'),
	(232,'system','log-delete','/b/system/log-delete',1,'admin','::1','2019-12-17 14:27:00'),
	(233,'content','index','/b/content/index?menuid=35&&_=1576564280767',1,'admin','::1','2019-12-17 14:31:20'),
	(234,'content','right','/b/content/right?_=1576564280800',1,'admin','::1','2019-12-17 14:31:20'),
	(235,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:31:20'),
	(236,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:31:20'),
	(237,'category','list','/b/category/list?menuid=36&&_=1576564281398',1,'admin','::1','2019-12-17 14:31:21'),
	(238,'category','list','/b/category/list?grid=treegrid',1,'admin','::1','2019-12-17 14:31:21'),
	(239,'content','index','/b/content/index?menuid=35&&_=1576564282169',1,'admin','::1','2019-12-17 14:31:22'),
	(240,'content','right','/b/content/right?_=1576564282201',1,'admin','::1','2019-12-17 14:31:22'),
	(241,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:31:22'),
	(242,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:31:22'),
	(243,'category','list','/b/category/list?menuid=36&&_=1576564283896',1,'admin','::1','2019-12-17 14:31:23'),
	(244,'content','index','/b/content/index?menuid=35&&_=1576564284305',1,'admin','::1','2019-12-17 14:31:24'),
	(245,'content','right','/b/content/right?_=1576564284336',1,'admin','::1','2019-12-17 14:31:24'),
	(246,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:31:24'),
	(247,'content','right','/b/content/right',1,'admin','::1','2019-12-17 14:31:24'),
	(248,'system','loglist','/b/system/loglist?menuid=15&&_=1576564287275',1,'admin','::1','2019-12-17 14:31:27'),
	(249,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','::1','2019-12-17 14:31:27'),
	(250,'system','menulist','/b/system/menulist?menuid=16&&_=1576564288069',1,'admin','::1','2019-12-17 14:31:28'),
	(251,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-17 14:31:28'),
	(252,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576564290792',1,'admin','::1','2019-12-17 14:31:30'),
	(253,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:31:30'),
	(254,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576564291408',1,'admin','::1','2019-12-17 14:31:31'),
	(255,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:31:31'),
	(256,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576564291881',1,'admin','::1','2019-12-17 14:31:31'),
	(257,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576564292348',1,'admin','::1','2019-12-17 14:31:32'),
	(258,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576564292850',1,'admin','::1','2019-12-17 14:31:32'),
	(259,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:32:59'),
	(260,'admin','member-add','/b/admin/member-add?_=1576564385957',1,'admin','::1','2019-12-17 14:33:05'),
	(261,'admin','member-add','/b/admin/member-add',1,'admin','::1','2019-12-17 14:33:25'),
	(262,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:33:25'),
	(263,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576564407670',1,'admin','::1','2019-12-17 14:33:27'),
	(264,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:33:27'),
	(265,'admin','role-add','/b/admin/role-add?_=1576564409015',1,'admin','::1','2019-12-17 14:33:29'),
	(266,'admin','role-add','/b/admin/role-add',1,'admin','::1','2019-12-17 14:33:32'),
	(267,'admin','role-add','/b/admin/role-add',1,'admin','::1','2019-12-17 14:33:39'),
	(268,'admin','role-add','/b/admin/role-add',1,'admin','::1','2019-12-17 14:33:42'),
	(269,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:33:42'),
	(270,'admin','role-permission','/b/admin/role-permission?id=9&_=1576564425640',1,'admin','::1','2019-12-17 14:33:45'),
	(271,'admin','role-permission','/b/admin/role-permission?id=9',1,'admin','::1','2019-12-17 14:33:45'),
	(272,'admin','role-permission','/b/admin/role-permission?dosubmit=1&id=9',1,'admin','::1','2019-12-17 14:33:47'),
	(273,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576564436343',1,'admin','::1','2019-12-17 14:33:56'),
	(274,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:33:56'),
	(275,'admin','member-edit','/b/admin/member-edit?id=11&_=1576564439223',1,'admin','::1','2019-12-17 14:33:59'),
	(276,'admin','member-edit','/b/admin/member-edit?id=11',1,'admin','::1','2019-12-17 14:34:02'),
	(277,'admin','member-edit','/b/admin/member-edit?id=11',1,'admin','::1','2019-12-17 14:34:05'),
	(278,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576564452313',1,'admin','::1','2019-12-17 14:34:12'),
	(279,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:34:12'),
	(280,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576564453831',1,'admin','::1','2019-12-17 14:34:13'),
	(281,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:34:14'),
	(282,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576564454264',1,'admin','::1','2019-12-17 14:34:14'),
	(283,'system','loglist','/b/system/loglist?menuid=15&&_=1576564472022',11,'test','::1','2019-12-17 14:34:32'),
	(284,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:34:32'),
	(285,'system','menulist','/b/system/menulist?menuid=16&&_=1576564473675',11,'test','::1','2019-12-17 14:34:33'),
	(286,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 14:34:33'),
	(287,'system','loglist','/b/system/loglist?menuid=15&&_=1576564476531',11,'test','::1','2019-12-17 14:34:36'),
	(288,'system','menulist','/b/system/menulist?menuid=16&&_=1576564477448',11,'test','::1','2019-12-17 14:34:37'),
	(289,'system','loglist','/b/system/loglist?menuid=15&&_=1576564511994',11,'test','::1','2019-12-17 14:35:12'),
	(290,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:35:12'),
	(291,'system','menulist','/b/system/menulist?menuid=16&&_=1576564512523',11,'test','::1','2019-12-17 14:35:12'),
	(292,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 14:35:12'),
	(293,'system','loglist','/b/system/loglist?menuid=15&&_=1576564513138',11,'test','::1','2019-12-17 14:35:13'),
	(294,'system','menulist','/b/system/menulist?menuid=16&&_=1576564513525',11,'test','::1','2019-12-17 14:35:13'),
	(295,'system','loglist','/b/system/loglist?menuid=15&&_=1576564513921',11,'test','::1','2019-12-17 14:35:13'),
	(296,'system','menulist','/b/system/menulist?menuid=16&&_=1576564514261',11,'test','::1','2019-12-17 14:35:14'),
	(297,'system','loglist','/b/system/loglist?menuid=15&&_=1576564579863',11,'test','::1','2019-12-17 14:36:19'),
	(298,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:36:19'),
	(299,'system','menulist','/b/system/menulist?menuid=16&&_=1576564580485',11,'test','::1','2019-12-17 14:36:20'),
	(300,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 14:36:20'),
	(301,'system','loglist','/b/system/loglist?menuid=15&&_=1576564581049',11,'test','::1','2019-12-17 14:36:21'),
	(302,'system','menulist','/b/system/menulist?menuid=16&&_=1576564581495',11,'test','::1','2019-12-17 14:36:21'),
	(303,'system','loglist','/b/system/loglist?menuid=15&&_=1576564633565',11,'test','::1','2019-12-17 14:37:13'),
	(304,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:37:13'),
	(305,'system','loglist','/b/system/loglist?menuid=15&&_=1576565063582',11,'test','::1','2019-12-17 14:44:23'),
	(306,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:44:23'),
	(307,'system','menulist','/b/system/menulist?menuid=16&&_=1576565064854',11,'test','::1','2019-12-17 14:44:24'),
	(308,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 14:44:24'),
	(309,'system','loglist','/b/system/loglist?menuid=15&&_=1576565065345',11,'test','::1','2019-12-17 14:44:25'),
	(310,'system','menulist','/b/system/menulist?menuid=16&&_=1576565075222',11,'test','::1','2019-12-17 14:44:35'),
	(311,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 14:44:35'),
	(312,'system','loglist','/b/system/loglist?menuid=15&&_=1576565077149',11,'test','::1','2019-12-17 14:44:37'),
	(313,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:44:37'),
	(314,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576565099598',1,'admin','::1','2019-12-17 14:44:59'),
	(315,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:44:59'),
	(316,'admin','member-edit','/b/admin/member-edit?id=11&_=1576565101674',1,'admin','::1','2019-12-17 14:45:01'),
	(317,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576565104256',1,'admin','::1','2019-12-17 14:45:04'),
	(318,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-17 14:45:04'),
	(319,'admin','role-permission','/b/admin/role-permission?id=9&_=1576565105824',1,'admin','::1','2019-12-17 14:45:05'),
	(320,'admin','role-permission','/b/admin/role-permission?id=9',1,'admin','::1','2019-12-17 14:45:05'),
	(321,'system','menulist','/b/system/menulist?menuid=16&&_=1576565143231',11,'test','::1','2019-12-17 14:45:43'),
	(322,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 14:45:43'),
	(323,'system','loglist','/b/system/loglist?menuid=15&&_=1576565144798',11,'test','::1','2019-12-17 14:45:44'),
	(324,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:45:45'),
	(325,'system','loglist','/b/system/loglist?menuid=15&&_=1576565255995',11,'test','::1','2019-12-17 14:47:36'),
	(326,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:47:36'),
	(327,'system','menulist','/b/system/menulist?menuid=16&&_=1576565302315',11,'test','::1','2019-12-17 14:48:22'),
	(328,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-17 14:48:22'),
	(329,'system','loglist','/b/system/loglist?menuid=15&&_=1576565303426',11,'test','::1','2019-12-17 14:48:23'),
	(330,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-17 14:48:23'),
	(331,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-17 14:51:36'),
	(332,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-17 14:54:56'),
	(436,'system','loglist','/b/system/loglist?menuid=15&&_=1576660548323',1,'admin','::1','2019-12-18 17:15:48'),
	(437,'system','loglist','/b/system/loglist?page=1&rows=10',1,'admin','::1','2019-12-18 17:15:48'),
	(438,'system','menulist','/b/system/menulist?menuid=16&&_=1576660549182',1,'admin','::1','2019-12-18 17:15:49'),
	(439,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-18 17:15:49'),
	(440,'setting','site','/b/setting/site?menuid=10&&_=1576660562795',1,'admin','::1','2019-12-18 17:16:02'),
	(441,'setting','site','/b/setting/site?grid=propertygrid',1,'admin','::1','2019-12-18 17:16:02'),
	(442,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576660564886',1,'admin','::1','2019-12-18 17:16:04'),
	(443,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-18 17:16:05'),
	(444,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576660565722',1,'admin','::1','2019-12-18 17:16:05'),
	(445,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-18 17:16:05'),
	(446,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576660569364',1,'admin','::1','2019-12-18 17:16:09'),
	(447,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576660569826',1,'admin','::1','2019-12-18 17:16:09'),
	(448,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576660603592',1,'admin','::1','2019-12-18 17:16:43'),
	(449,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-18 17:16:43'),
	(450,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576660606048',1,'admin','::1','2019-12-18 17:16:46'),
	(451,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-18 17:16:46'),
	(452,'system','menulist','/b/system/menulist?menuid=16&&_=1576660611890',1,'admin','::1','2019-12-18 17:16:51'),
	(453,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-18 17:16:51'),
	(454,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-18 17:19:19'),
	(455,'content','index','/b/content/index?menuid=35&&_=1576660778830',1,'admin','::1','2019-12-18 17:19:38'),
	(456,'content','right','/b/content/right?_=1576660778857',1,'admin','::1','2019-12-18 17:19:38'),
	(457,'content','right','/b/content/right',1,'admin','::1','2019-12-18 17:19:38'),
	(458,'content','right','/b/content/right',1,'admin','::1','2019-12-18 17:19:38'),
	(459,'category','list','/b/category/list?menuid=36&&_=1576660779525',1,'admin','::1','2019-12-18 17:19:39'),
	(460,'category','list','/b/category/list?grid=treegrid',1,'admin','::1','2019-12-18 17:19:39'),
	(461,'system','menulist','/b/system/menulist?menuid=16&&_=1576660782181',1,'admin','::1','2019-12-18 17:19:42'),
	(462,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-18 17:19:42'),
	(463,'system','menu-order','/b/system/menu-order',1,'admin','::1','2019-12-18 17:19:50'),
	(464,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-18 17:19:54'),
	(465,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576660823268',1,'admin','::1','2019-12-18 17:20:23'),
	(466,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-18 17:20:23'),
	(467,'user','info','/b/user/info?menuid=57&&_=1576660956382',1,'admin','::1','2019-12-18 17:22:36'),
	(468,'user','list','/b/user/list?menuid=56&&_=1576660985193',1,'admin','::1','2019-12-18 17:23:05'),
	(469,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 17:23:05'),
	(470,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576660990134',1,'admin','::1','2019-12-18 17:23:10'),
	(471,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-18 17:23:10'),
	(472,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576660990749',1,'admin','::1','2019-12-18 17:23:10'),
	(473,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-18 17:23:10'),
	(474,'user','list','/b/user/list?menuid=56&&_=1576660993210',1,'admin','::1','2019-12-18 17:23:13'),
	(475,'user','list','/b/user/list?menuid=56&&_=1576661073587',1,'admin','::1','2019-12-18 17:24:33'),
	(476,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 17:24:33'),
	(477,'wechat','userinfo','/b/wechat/userinfo?menuid=59&&_=1576661077778',1,'admin','::1','2019-12-18 17:24:37'),
	(478,'user','list','/b/user/list?menuid=56&&_=1576661083050',1,'admin','::1','2019-12-18 17:24:43'),
	(479,'wechat','userinfo','/b/wechat/userinfo?menuid=59&&_=1576661165799',1,'admin','::1','2019-12-18 17:26:05'),
	(480,'user','list','/b/user/list?menuid=56&&_=1576661168701',1,'admin','::1','2019-12-18 17:26:08'),
	(481,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 17:26:08'),
	(482,'user','list','/b/user/list?menuid=56&&_=1576661222465',1,'admin','::1','2019-12-18 17:27:02'),
	(483,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 17:27:02'),
	(484,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 17:27:04'),
	(485,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 17:29:14'),
	(486,'user','list','/b/user/list?menuid=56&&_=1576661372693',1,'admin','::1','2019-12-18 17:29:32'),
	(487,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 17:29:32'),
	(488,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 17:29:33'),
	(489,'user','list','/b/user/list?menuid=56&&_=1576661402186',1,'admin','::1','2019-12-18 17:30:02'),
	(490,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 17:30:02'),
	(491,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 17:30:03'),
	(492,'user','list','/b/user/list?menuid=56&&_=1576661571834',1,'admin','::1','2019-12-18 17:32:51'),
	(493,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 17:32:52'),
	(494,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 17:32:53'),
	(495,'user','list','/b/user/list?menuid=56&&_=1576661705664',1,'admin','::1','2019-12-18 17:35:05'),
	(496,'user','list','/b/user/list?menuid=56&&_=1576661710901',1,'admin','::1','2019-12-18 17:35:10'),
	(497,'user','list','/b/user/list?menuid=56&&_=1576661772507',1,'admin','::1','2019-12-18 17:36:12'),
	(498,'user','list','/b/user/list?menuid=56&&_=1576661952969',1,'admin','::1','2019-12-18 17:39:12'),
	(499,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 17:39:13'),
	(500,'user','list','/b/user/list?menuid=56&&_=1576661955539',1,'admin','::1','2019-12-18 17:39:15'),
	(501,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 17:39:16'),
	(502,'user','edit','/b/user/edit?id=id',1,'admin','::1','2019-12-18 17:50:04'),
	(503,'user','list','/b/user/list?menuid=56&&_=1576662934373',1,'admin','::1','2019-12-18 17:55:34'),
	(504,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 17:55:34'),
	(505,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 17:55:36'),
	(506,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 17:55:53'),
	(507,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:03:55'),
	(508,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:04:06'),
	(509,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:04:07'),
	(510,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:04:07'),
	(511,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:04:07'),
	(512,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:04:07'),
	(513,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:04:08'),
	(514,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:04:51'),
	(515,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:06:31'),
	(516,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:06:49'),
	(517,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 18:07:13'),
	(518,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:07:30'),
	(519,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 18:08:41'),
	(520,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:08:44'),
	(521,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:08:48'),
	(522,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:08:51'),
	(523,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:09:27'),
	(524,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 18:10:37'),
	(525,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:10:41'),
	(526,'user','list','/b/user/list?menuid=56&&_=1576663847414',1,'admin','::1','2019-12-18 18:10:47'),
	(527,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 18:10:47'),
	(528,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 18:10:49'),
	(529,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:10:52'),
	(530,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:11:05'),
	(531,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:11:12'),
	(532,'user','list','/b/user/list?menuid=56&&_=1576663912186',1,'admin','::1','2019-12-18 18:11:52'),
	(533,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-18 18:11:52'),
	(534,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 18:11:53'),
	(535,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:11:58'),
	(536,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:12:01'),
	(537,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 18:12:32'),
	(538,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:12:35'),
	(539,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:12:58'),
	(540,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 18:13:00'),
	(541,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:13:04'),
	(542,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 18:36:42'),
	(543,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:36:44'),
	(544,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:36:49'),
	(545,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:36:57'),
	(546,'user','info','/b/user/info?id=1',1,'admin','::1','2019-12-18 18:36:59'),
	(547,'user','edit','/b/user/edit?id=1',1,'admin','::1','2019-12-18 18:44:29'),
	(548,'system','menulist','/b/system/menulist?menuid=16&&_=1576725212189',11,'test','::1','2019-12-19 11:13:32'),
	(549,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-19 11:13:32'),
	(550,'system','loglist','/b/system/loglist?menuid=15&&_=1576725213083',11,'test','::1','2019-12-19 11:13:33'),
	(551,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-19 11:13:33'),
	(552,'system','menulist','/b/system/menulist?menuid=16&&_=1576725214634',11,'test','::1','2019-12-19 11:13:34'),
	(553,'system','loglist','/b/system/loglist?menuid=15&&_=1576725226329',11,'test','::1','2019-12-19 11:13:46'),
	(554,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-19 11:13:46'),
	(555,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-19 11:13:47'),
	(556,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-19 11:13:49'),
	(557,'system','menulist','/b/system/menulist?menuid=16&&_=1576725231739',11,'test','::1','2019-12-19 11:13:51'),
	(558,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-19 11:13:51'),
	(559,'system','menulist','/b/system/menulist?menuid=16&&_=1576725330796',11,'test','::1','2019-12-19 11:15:30'),
	(560,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-19 11:15:30'),
	(561,'system','loglist','/b/system/loglist?menuid=15&&_=1576725332130',11,'test','::1','2019-12-19 11:15:32'),
	(562,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-19 11:15:32'),
	(563,'system','menulist','/b/system/menulist?menuid=16&&_=1576725341158',11,'test','::1','2019-12-19 11:15:41'),
	(564,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-19 11:15:41'),
	(565,'system','menulist','/b/system/menulist?menuid=16&&_=1576725374735',11,'test','::1','2019-12-19 11:16:14'),
	(566,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-19 11:16:14'),
	(567,'system','loglist','/b/system/loglist?menuid=15&&_=1576725376018',11,'test','::1','2019-12-19 11:16:16'),
	(568,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-19 11:16:16'),
	(569,'system','menulist','/b/system/menulist?menuid=16&&_=1576725378527',11,'test','::1','2019-12-19 11:16:18'),
	(570,'system','loglist','/b/system/loglist?menuid=15&&_=1576725379971',11,'test','::1','2019-12-19 11:16:20'),
	(571,'system','log-delete','/b/system/log-delete',11,'test','::1','2019-12-19 11:16:22'),
	(572,'system','loglist','/b/system/loglist?menuid=15&&_=1576725448433',11,'test','::1','2019-12-19 11:17:28'),
	(573,'system','loglist','/b/system/loglist?page=1&rows=10',11,'test','::1','2019-12-19 11:17:29'),
	(574,'system','menulist','/b/system/menulist?menuid=16&&_=1576725455432',11,'test','::1','2019-12-19 11:17:35'),
	(575,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-19 11:17:35'),
	(576,'system','menulist','/b/system/menulist?menuid=16&&_=1576727598372',11,'test','::1','2019-12-19 11:53:18'),
	(577,'system','menulist','/b/system/menulist?grid=treegrid',11,'test','::1','2019-12-19 11:53:18'),
	(578,'user','list','/b/user/list?menuid=56&&_=1576727630592',1,'admin','::1','2019-12-19 11:53:50'),
	(579,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-19 11:53:51'),
	(580,'user','list','/b/user/list?menuid=56&&_=1576727637169',1,'admin','::1','2019-12-19 11:53:57'),
	(581,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576727639037',1,'admin','::1','2019-12-19 11:53:59'),
	(582,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-19 11:53:59'),
	(583,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576727639661',1,'admin','::1','2019-12-19 11:53:59'),
	(584,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-19 11:53:59'),
	(585,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576727640229',1,'admin','::1','2019-12-19 11:54:00'),
	(586,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576727640814',1,'admin','::1','2019-12-19 11:54:00'),
	(587,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576727641296',1,'admin','::1','2019-12-19 11:54:01'),
	(588,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576727642100',1,'admin','::1','2019-12-19 11:54:02'),
	(589,'system','menulist','/b/system/menulist?menuid=16&&_=1576727657251',1,'admin','::1','2019-12-19 11:54:17'),
	(590,'system','menulist','/b/system/menulist?grid=treegrid',1,'admin','::1','2019-12-19 11:54:17'),
	(591,'user','list','/b/user/list?menuid=56&&_=1576727660111',1,'admin','::1','2019-12-19 11:54:20'),
	(592,'user','list','/b/user/list?page=1&rows=10',1,'admin','::1','2019-12-19 11:54:20'),
	(593,'admin','memberlist','/b/admin/memberlist?menuid=12&&_=1576727661435',1,'admin','::1','2019-12-19 11:54:21'),
	(594,'admin','memberlist','/b/admin/memberlist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-19 11:54:21'),
	(595,'admin','rolelist','/b/admin/rolelist?menuid=13&&_=1576727662884',1,'admin','::1','2019-12-19 11:54:22'),
	(596,'admin','rolelist','/b/admin/rolelist?grid=datagrid&page=1&rows=10',1,'admin','::1','2019-12-19 11:54:23'),
	(597,'category','list','/b/category/list?menuid=36&&_=1576729574844',1,'admin','::1','2019-12-19 12:26:14'),
	(598,'category','list','/b/category/list?grid=treegrid',1,'admin','::1','2019-12-19 12:26:15');
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
	(5,'后台管理',0,'system','top','',0,5,'1'),
	(6,'个人信息',1,'admin','public-left','',1,0,'1'),
	(7,'修改密码',6,'admin','public-editpwd','',1,1,'1'),
	(8,'修改个人信息',6,'admin','public-editinfo','',1,0,'1'),
	(9,'系统设置',2,'setting','left','',0,1,'1'),
	(10,'站点设置',9,'setting','site','',0,1,'1'),
	(11,'管理员设置',2,'admin','left','',0,2,'1'),
	(12,'管理员管理',11,'admin','memberlist','',0,1,'1'),
	(13,'角色管理',11,'admin','rolelist','',0,2,'1'),
	(14,'后台管理',5,'system','left','',0,1,'1'),
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
	(60,'编辑会员',55,'user','edit','',0,0,'1');
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



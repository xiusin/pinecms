/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3306
 Source Schema         : pinecms

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 21/03/2020 18:31:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pinecms_admin
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_admin`;
CREATE TABLE `pinecms_admin`  (
  `userid` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `roleid` smallint(6) NULL DEFAULT 0,
  `encrypt` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `lastloginip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `lastlogintime` int(10) UNSIGNED NULL DEFAULT 0,
  `email` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `realname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  PRIMARY KEY (`userid`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_admin
-- ----------------------------
INSERT INTO `pinecms_admin` VALUES (1, 'admin', '5736a2a40f752bf2e82953702d25075b', 1, 'qmRlFL', '::1', 1474291850, 'chenchengbin92@gmail.com2', '陈成彬（xiusin）');
INSERT INTO `pinecms_admin` VALUES (11, 'test', '834a33db060873a7a208617930edb29a', 1, '5u2G0w', '::1', 0, 'asdasd@asdasd.com1', 'ccc1');

-- ----------------------------
-- Table structure for pinecms_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_admin_role`;
CREATE TABLE `pinecms_admin_role`  (
  `roleid` tinyint(3) UNSIGNED NOT NULL AUTO_INCREMENT,
  `rolename` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `listorder` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `disabled` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`roleid`) USING BTREE,
  UNIQUE INDEX `UQE_iriscms_admin_role_rolename`(`rolename`) USING BTREE,
  INDEX `listorder`(`listorder`) USING BTREE,
  INDEX `disabled`(`disabled`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_admin_role
-- ----------------------------
INSERT INTO `pinecms_admin_role` VALUES (1, '超级管理员', '超级管理员', 0, 0);

-- ----------------------------
-- Table structure for pinecms_admin_role_priv
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_admin_role_priv`;
CREATE TABLE `pinecms_admin_role_priv`  (
  `roleid` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  `c` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `a` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  INDEX `roleid`(`roleid`, `c`, `a`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色操作权限配置表' ROW_FORMAT = Fixed;

-- ----------------------------
-- Records of pinecms_admin_role_priv
-- ----------------------------
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'category', 'add');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'category', 'delete');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'category', 'edit');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'category', 'list');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'category', 'order');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'category', 'view');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'content', 'index');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'content', 'news-list');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'content', 'page');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'content', 'right');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'content', 'top');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'left');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'log-delete');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'loglist');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'logview');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'menuadd');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'menudelete');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'menuedit');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'menulist');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'menuorder');
INSERT INTO `pinecms_admin_role_priv` VALUES (9, 'system', 'menuview');

-- ----------------------------
-- Table structure for pinecms_advert
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_advert`;
CREATE TABLE `pinecms_advert`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `space_id` bigint(20) NULL DEFAULT NULL,
  `image` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `link_url` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `listorder` int(11) NULL DEFAULT NULL,
  `start_time` datetime(0) NULL DEFAULT NULL,
  `end_time` datetime(0) NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '广告表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_advert
-- ----------------------------
INSERT INTO `pinecms_advert` VALUES (2, 'aaaa', 3, '/upload/public/20200309/BFZ5z5e5R9.jpg', 'https://www.baidu.com', 30, '2020-03-04 20:49:26', '2022-03-17 20:49:30', 1);
INSERT INTO `pinecms_advert` VALUES (3, '个人二维码', 4, '/upload/public/20200311/st1Cu5DfW6.png', '', 10, '2020-03-04 21:32:16', '2020-03-19 21:32:21', 1);

-- ----------------------------
-- Table structure for pinecms_advert_space
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_advert_space`;
CREATE TABLE `pinecms_advert_space`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '广告位表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_advert_space
-- ----------------------------
INSERT INTO `pinecms_advert_space` VALUES (1, '111首页Banner');
INSERT INTO `pinecms_advert_space` VALUES (3, '首页Banner1');
INSERT INTO `pinecms_advert_space` VALUES (4, '二维码位置');
INSERT INTO `pinecms_advert_space` VALUES (5, 'QQ二维码');

-- ----------------------------
-- Table structure for pinecms_articles
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_articles`;
CREATE TABLE `pinecms_articles`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文章标题',
  `tags` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'TAG标签',
  `thumb` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '缩略图',
  `from_url` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文章来源',
  `author` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '作者',
  `keywords` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关键字',
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容摘要',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '文章内容',
  `pubtime` datetime(0) NULL DEFAULT NULL COMMENT '发布时间',
  `catid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '所属栏目ID',
  `mid` int(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型ID',
  `refid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型关联的文章ID',
  `listorder` int(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序值',
  `visit_count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '访问次数',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `created_time` datetime(0) NULL DEFAULT NULL,
  `updated_time` datetime(0) NULL DEFAULT NULL,
  `deleted_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `refid`(`refid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统内置模型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_articles
-- ----------------------------
INSERT INTO `pinecms_articles` VALUES (1, '山有木兮木有枝，心悦君兮君不知', '古文,越人歌', '/upload/public/20200306/zebRQfaG71.png', '古诗文网', '出自先秦的《越人歌》', '今夕何夕,王子,君子,双关,诗人,楚语,襄成君,同舟,心悦,鄂君,君兮君,两句', '　据刘向《说苑·善说》记载：春秋时代，楚王母弟鄂君子皙在河中游玩，钟鼓齐鸣。摇船者是位越人，趁乐声刚停，便抱双桨用越语唱了一支歌。鄂君子皙听不懂，叫人翻译成楚语。就是上面的歌谣。歌中唱出了越人对子皙的那种深沉真挚的爱恋之情，歌词 声义双关，委婉动听。是中国最早的译诗，也是古代楚越文化交融的结晶和见证。它对楚辞创作有着直接的影响作用。（选自《先秦诗文精华》 人民文学出版社2000.1版）', '<p>　　据刘向《说苑·善说》记载：春秋时代，楚王母弟鄂君子皙在河中游玩，钟鼓齐鸣。摇船者是位越人，趁乐声刚停，便抱双桨用越语唱了一支歌。鄂君子皙听不懂，叫人翻译成楚语。就是上面的歌谣。歌中唱出了越人对子皙的那种深沉真挚的爱恋之情，歌词 声义双关，委婉动听。是中国最早的译诗，也是古代楚越文化交融的结晶和见证。它对楚辞创作有着直接的影响作用。（选自《先秦诗文精华》 人民文学出版社2000.1版）</p><p><br/></p><p>　　故事讲的是楚国襄成君册封受爵那天，身着华服伫立河边。楚大夫庄辛经过，见了他心中欢喜，于是上前行礼，想要握他的手。襄成君忿其越礼之举，不予理睬。于是庄辛洗了手，给襄成君讲述了楚国鄂君的故事：</p><p><br/></p><p>　　鄂君子皙是楚王的弟弟，坐船出游，有爱慕他的越人船夫抱着船桨对他唱歌。歌声悠扬缠绵，委婉动听，打动了鄂君，当即让人翻译成楚语，这便有了《越人歌》之词。鄂君明白歌意后，非但没有生气，还走过去拥抱船夫，给他盖上绣花被，愿与之同床共寝。</p><p><br/></p><p>　　庄辛进而问襄成君：鄂君身份高贵仍可以与越人船夫交欢尽意，我为何不可以握你的手呢？襄成君当真答应了他的请求，将手递给了他。</p><p><br/></p><p>　　据刘向《说苑·善说》记载：春秋时代，楚王母弟鄂君子皙在河中游玩，钟鼓齐鸣。摇船者是位越人，趁乐声刚停，便抱双桨用越语唱了一支歌。原文汉字注音为“滥兮抃草滥予昌枑泽予昌州州州焉乎秦胥胥缦予乎昭澶秦逾渗惿随河湖”。</p><p><br/></p><p>　　鄂君子皙听不懂，叫人翻译成楚语。译文为“今夕何夕兮，搴州中流。今夕何夕兮，得与王子同舟。蒙羞被好兮，不訾诟耻。心几烦而不绝兮，得知王子。山有木兮木有枝，心悦君兮君不知”。鄂君子皙在听完译文后万分感动，与越人缠绵一夜，“行而拥之，举绣被而覆之”。</p><p><br/></p><p>　　歌中唱出了越人对子皙的那种深沉真挚的爱恋之情，歌词 声义双关，委婉动听。是我国最早的译诗，也是古代楚越文化交融的结晶和见证。它对楚辞创作有着直接的影响作用。其中“山有木兮木有枝，心悦君兮君不知”一句最为经典，后来楚辞中的“沅有芷兮澧有兰，思公子兮未敢言”，被认为是借鉴了其“兴”的修辞手法。</p><p><br/></p><p>　　起首两句“今夕何夕兮搴洲中流，今日何日兮得与王子同舟”，“洲”，当从《北堂书钞》卷一O六引作“舟”。“搴洲中流”即在河中荡舟之意。这是记事，记叙了这天晚上荡舟河中，又有幸能与王子同舟这样一件事。在这里，诗人用了十分情感化的“今夕何夕兮”、“今日何日兮”的句式。“今夕”、“今日”本来已经是很明确的时间概念，还要重复追问“今夕何夕”、“今日何日”，这表明诗人内心的激动无比，意绪已不复平静有序而变得紊乱无序，难以控抑。这种句式及其变化以后常为诗人所取用，著名的如宋张孝祥《念奴娇·过洞庭》的末两句“扣舷独啸，不知今夕何夕”。</p><p><br/></p><p>　　进入诗的中间两句行文用字和章法都明显地由相对平易转为比较艰涩了。这是诗人在非常感情化的叙事完毕之后转入了理性地对自己的心情进行描述。“蒙羞被好兮不訾诟耻，心几烦而不绝兮得知王子”，是说我十分惭愧承蒙王子您的错爱，王子的知遇之恩令我心绪荡漾。</p><p><br/></p><p>　　最后两句是诗人在非常情感化的叙事和理性描述自己心情之后的情感抒发，此时的诗人已经将激动紊乱的意绪梳平，因此这种情感抒发十分艺术化，用字平易而意蕴深长，余韵袅袅。“山有木兮木有枝”是一个比兴句，既以“山有木”、“木有枝”兴起下面一句的“心悦君”、“君不知”，又以“枝”谐音比喻“知”。在自然界，山上有树树上有枝，顺理成章；但在人间社会，自己对别人的感情深浅归根到底却只有自己知道，许多时候你会觉得自己对别人的感情难以完全表达，因此越人唱出了这样的歌词。而借“枝”与“知”的谐音双关关系做文章的比兴手法，也是《诗经》所惯用的。这种谐音双关对后代的诗歌如南朝乐府民歌《子夜歌》等恐怕不无影响。而“山有木兮木有枝，心悦君兮君不知”二句，与《九歌·湘夫人》中“沅有茝兮醴有兰，思公子兮未敢言”二句相仿佛，也可见出此楚译《越人歌》深受楚声的影响。虽然今人所读到的《越人歌》是翻译作品，但仍可这样说：《越人歌》的艺术成就表明，两千多年前，古越族的文学已经达到了相当高的水平。</p><p><br/></p>', '2020-03-19 15:19:32', 1, 1, 0, 2, 0, 1, '2020-03-06 15:19:35', '2020-03-06 16:29:19', NULL);
INSERT INTO `pinecms_articles` VALUES (2, '将进酒', '李白,古文', '', '古文岛', '李白的弟弟', '诗人,狂放,诗情,千金,人生,将进酒,万古愁,开篇,全篇,曹植,夸张,得意', '', '<p>　　《将进酒》原是汉乐府短箫铙歌的曲调，标题的意思为“劝酒歌”，内容多是咏唱喝酒放歌之事。这首诗是诗人当时和友人岑勋在嵩山另一老友元丹丘的颍阳山居作客，作者正值仕途遇挫之际，所以借酒兴诗，来了一次酣畅淋漓的抒发。在这首诗里，李白“借题发挥”，借酒消愁，感叹人生易老，抒发了自己怀才不遇的心情。</p><p><br/></p><p>　　这首诗十分形象的体现了李白桀骜不驯的性格：对自己充满自信、孤高自傲、热情豪放，“天生我材必有用”、“人生得意须尽欢”。全诗气势豪迈，感情豪放，言语流畅，具有极强的感染力。李白咏酒的诗歌非常能体现他的个性，思想内容深沉，艺术表现成熟。《将进酒》即为其代表作。</p><p><br/></p><p>　　诗歌发端就是两组排比长句，如挟天风海雨向读者迎面扑来。“君不见，黄河之水天上来，奔腾到海不复回”，颍阳去黄河不远，登高纵目，故借以起兴。黄河源远流长，落差极大，如从天而降，一泻千里，东走大海。如此波澜壮阔的现象，必定不是肉眼能够看到的，作者是幻想的，言语带有夸张。上句写大河之来，势不可挡；下句写大河之去，势不可回。一涨一消，构成舒卷往复的咏叹味，是短促的单句（如“黄河落天走东海”）所没有的。紧接着，“君不见，高堂明镜悲白发，朝如青丝暮成雪”，恰似一波未平、一波又起。如果说前二句为空间范畴的夸张，这二句则是时间范畴的夸张。悲叹人生苦短，而又不直言，却说“高堂明镜悲白发”，一种搔首顾影、徒呼奈何的神态宛如画出。将人生由青春到老的全过程说成“朝”“暮”之事，把原本就短暂的说得更为短暂，与前两句把原本壮阔的说得更为壮阔，是“反向”的夸张。开篇“以河之水一去不复返喻人生易逝”，“以黄河的伟大永恒形出生命的渺小脆弱”。这个开端可谓悲感至极，却又不堕纤弱，可以说是巨人式的感伤，具有惊心动魄的艺术力量，同时也是由长句排比开篇的气势感造成的。这种开篇的方法作者经常用，比如“弃我去者，昨日之日不可留；乱我心者，今日之日多烦忧”（《宣城谢朓楼饯别校书叔云》），沈德潜说：“此种格调，太白从心化出”，可见其颇具创造性。此诗两作“君不见”的呼告（一般乐府诗只是篇首或篇末偶尔用），又使诗句感情色彩大增。所谓大开大阖者，此可谓大开。</p><p><br/></p><p>　　“人生得意须尽欢”，这似乎是宣扬及时行乐的思想，然而只不过是现象而已。诗人“得意”过没有？“凤凰初下紫泥诏，谒帝称觞登御筵”（《玉壶吟》）似乎得意过；然而那不过是一场幻影。“弹剑作歌奏苦声，曳裾王门不称情”（《行路难三首·其二》）又似乎并没有得意，有的是失望与愤慨，但并不就此消沉。诗人于是用乐观好强的口吻肯定人生，肯定自我：“天生我材必有用”，这是一个令人鼓掌赞叹的好句子。“有用”而且“必”，非常的自信，简直像是人的价值宣言，而这个人“我”是需要大写的。于是，从貌似消极的现象中透露出了深藏其内的一种怀才不遇而又渴望入世的积极的态度。正是“长风破浪会有时”，应为这样的未来痛饮高歌，破费又算得了什么！</p><p><br/></p><p>　　“千金散尽还复来！”这又是一个高度自信的惊人之句，能驱使金钱而不为金钱所使，这足以令所有凡夫俗子们咋舌。诗如其人，想诗人“曩昔东游维扬，不逾一年，散金三十馀万”（《上安州裴长史书》），是何等的豪举。所以此句是深蕴在骨子里的豪情，绝非装腔作势者可以得其万分之一。与此气派相当，作者描绘了一场盛筵，那决不是“菜要一碟乎，两碟乎？酒要一壶乎，两壶乎？”而是整头整头地“烹羊宰牛”，不喝上“三百杯”决不罢休。多痛快的筵宴，又是多么豪壮的诗句！至此，狂放之情趋于高潮，诗的旋律加快。“岑夫子，丹丘生，将进酒，杯莫停！”几个短句忽然加入，不但使诗歌节奏富于变化， 而且使我们似乎听到了诗人在席上频频地劝酒。既是生逢知己，又是酒逢对手，不但“忘形到尔汝”，诗人甚至忘了是在写诗，笔下之诗似乎还原为生活，他还要“与君歌一曲，请君为我倾耳听”。以下八句就是诗中之歌了，这纯粹是神来之笔。</p><p><br/></p><p>　　“钟鼓馔玉”即富贵生活（富贵人家吃饭时鸣钟列鼎，食物精美如玉），可诗人却认为这“不足贵”，并放言“但愿长醉不复醒”。诗情至此，便分明由狂放转而为愤激。这里不仅是酒后吐狂言，而且是酒后吐真言了。以“我”天生有用之才，本当位极卿相，飞黄腾达，然而“大道如青天，我独不得出”（《行路难三首·其二》）。说富贵“不足贵”，乃是出于愤慨。以下“古来圣贤皆寂寞”二句亦属愤语。诗人曾喟叹“自言管葛竟谁许”，说古人“寂寞”，其实也表现出了自己的“寂寞”，所以才愿长醉不醒了。这里，诗人是用古人的酒杯，浇自己的块垒。说到“唯有饮者留其名”，便举出“陈王”曹植作代表。并化用其《名都篇》“归来宴平乐，美酒斗十千”之句。古来酒徒很多，而为何偏举“陈王”，这又与李白一向自命不凡分不开，他心目中树为榜样的都是谢安这些高级人物，而这类人物当中，“陈王”曹植与酒联系得比较多。这样写便有了气派，与前文极度自信的口吻一贯。再者，“陈王”曹植于曹丕、曹睿两朝备受猜忌，有志难展，也激起诗人的同情。一提“古来圣贤”，二提“陈王”曹植，满满的不平之气。此诗开始似乎只涉及人生感慨，而不染指政治色彩，其实全篇饱含了一种深广的忧愤和对自我的信念。诗情之所以悲而不伤，悲而能壮，即根源在此。</p><p><br/></p><p>　　刚露一点深衷，又说回酒了，而且看起来酒兴更高了。以下诗情再入狂放，而且愈来愈狂。“主人何为言少钱”，既照应“千金散尽”句，又故作跌宕，引出最后一番豪言壮语：即便千金散尽，也不惜将名贵宝物“五花马”（毛色作五花纹的良马）、“千金裘”（昂贵的皮衣）用来换美酒，图个一醉方休。这结尾之妙，不仅在于“呼儿”“与尔”，口气甚大；而且具有一种作者一时可能觉察不到的将宾作主的任诞情态。须知诗人不过是被友招饮的客人，此刻他却高踞一席，气使颐指，提议典裘当马，令人不知谁是“主人”，浪漫色彩极浓。快人快语，非不拘形迹的豪迈知交断不能出此。诗情至此狂放至极，令人嗟叹咏歌，直欲“手之舞之，足之蹈之”。情犹未已，诗已告终，突然又迸出一句“与尔同销万古愁”，与开篇之“悲”关合，而“万古愁”的含义更其深沉。这“白云从空，随风变灭”的结尾，显见诗人奔涌跌宕的感情激流。通观全篇，真是大起大落，非如椽巨笔不办。</p><p><br/></p><p>　　《将进酒》篇幅不算长，却五音繁会，气象不凡。它笔酣墨饱，情极悲愤而作狂放，语极豪纵而又沉着。全篇具有震动古今的气势与力量，这诚然与夸张手法不无关系，比如诗中屡用巨额数字（“千金”、“三百杯”、“斗酒十千”、“千金裘”、“万古愁”等等）表现豪迈诗情，同时，又不给人空洞浮夸感，其根源就在于它那充实深厚的内在感情，那潜在酒话底下如波涛汹涌的郁怒情绪。此外，全篇大起大落，诗情忽翕忽张，由悲转乐、转狂放、转愤激、再转狂放、最后结穴于“万古愁”，回应篇首，如大河奔流，有气势，亦有曲折，纵横捭阖，力能扛鼎。其歌中有歌的包孕写法，又有鬼斧神工、“绝去笔墨畦径”之妙，既非鑱刻能学，又非率尔可到。通篇以七言为主，而又以三、五言句“破”之，极参差错综之致；诗句以散行为主，又以短小的对仗语点染（如“岑夫子，丹丘生”，“五花马，千金裘”），节奏疾徐尽变，奔放而不流易。《唐诗别裁》谓“读李诗者于雄快之中，得其深远宕逸之神，才是谪仙人面目”，此篇足以当之</p><p><br/></p>', '2020-03-23 16:23:22', 1, 1, 0, 1, 0, 1, '2020-03-06 16:25:41', '2020-03-06 16:29:19', NULL);
INSERT INTO `pinecms_articles` VALUES (3, '一曲《梁祝》久萦怀', '梁祝,散文', '', '短文学网', '李胜波', '梁祝,音乐,越剧,故事,民间传说,小提琴,表演,爱情,形式,钢琴伴奏,袁雪芬,欣赏', '', '<p style=\"margin-top: 0px; margin-bottom: 25px; padding: 0px; line-height: 2.8em; text-indent: 28px; color: rgb(102, 102, 102); font-family: Arial; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">江南人对《梁祝》的<a href=\"https://www.duanwenxue.com/qinggan/gushi/\" style=\"margin: 0px; padding: 0px; color: rgb(102, 102, 102); outline: none; text-decoration-line: none;\">故事</a>，可谓妇孺皆知耳熟能详。每当听到这悠扬委婉、如诉如泣的旋律时，我无不为之如痴如醉。</p><p style=\"margin-top: 0px; margin-bottom: 25px; padding: 0px; line-height: 2.8em; text-indent: 28px; color: rgb(102, 102, 102); font-family: Arial; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">《梁祝》作为我国四大民间传说之一，采用现实主义和浪漫主义相结合的手法，讴歌了人间忠贞不渝的<a href=\"https://www.duanwenxue.com/qinggan/aiqing/\" style=\"margin: 0px; padding: 0px; color: rgb(102, 102, 102); outline: none; text-decoration-line: none;\">爱情</a>，同时也鞭挞了传统的卫道士。这一民间传说故事相传越千年，经久不衰。究其原因，除了其自身的<a href=\"https://www.duanwenxue.com/shanggan/ganrengushi/\" style=\"margin: 0px; padding: 0px; color: rgb(102, 102, 102); outline: none; text-decoration-line: none;\">感人故事</a>情节和爱情主题外，也与其表现手法和形式的多种多样分不开。与其他民间传说故事相比，《梁祝》的表现形式更具独特性，而且它并不固步自封，还随着时代的发展而发展。</p><p style=\"margin-top: 0px; margin-bottom: 25px; padding: 0px; line-height: 2.8em; text-indent: 28px; color: rgb(102, 102, 102); font-family: Arial; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">相传一千多年前，它的主要表现形式为民间传说故事。每当农闲或夏季纳凉时分，由上一代人讲故事给下一代人听，一代代口耳相传，用文字的形式记录下来的几乎很少。所以，它的传承带有很大的不完整性。后来渐渐衍变成为民间说唱和话本形式，流传于乡间和市井。到了清末民初，从事地方戏的民间艺人，把它搬上了草台班子，流动在乡间的简陋舞台上。当时都是以折子戏的形式来表演，取其一折，突出情节，言简意赅，通俗易懂，深受百姓欢迎。其中最早的大概是流传于江浙一带的越剧《梁祝哀史》，由早期的越剧表演艺术家施银花、马樟花等主演。后来又经袁雪芬、傅全香和范瑞娟等表演艺术家的不断加工塑造，特别是由袁雪芬对越剧改革、创新和发展之后，使《梁祝》的故事更加完整，把它放置在反封建礼教的历史大环境中来表现人间爱情，具有强烈的审美性和鲜明的时代性，还被拍摄成新中国第一部彩色影片传到国外，被喻为“东方的《罗密欧与朱丽叶》”，成为新中国文化交流的纽带和使者。</p><p style=\"margin-top: 0px; margin-bottom: 25px; padding: 0px; line-height: 2.8em; text-indent: 28px; color: rgb(102, 102, 102); font-family: Arial; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">《梁祝》不仅成为越剧的<a href=\"https://www.duanwenxue.com/yuju/\" style=\"margin: 0px; padding: 0px; color: rgb(102, 102, 102); outline: none; text-decoration-line: none;\">经典</a>剧目，还被搬上了其他的戏剧舞台，如京剧、黄梅戏等，还有大型芭蕾舞剧。最值得一提的是，由陈钢和何占豪的经典合作，使这一原来只属于民族艺术的爱情主题故事，与西洋的小提琴相结合来演译，大大拓展了它的表演空间，成为不同年龄、不同语言、不同文化和不同民族的听众都能接受、雅俗共赏的世界音乐经典。以钢琴伴奏、小提琴独奏来表演《梁祝》，已经成为中外音乐史上的传奇。俞丽拿和盛中国等一批首演此曲的音乐家也因此享誉海内外。如今表演这一题材的音乐形式还有琵琶、二胡、古筝、笛子和钢琴等。</p><p style=\"margin-top: 0px; margin-bottom: 25px; padding: 0px; line-height: 2.8em; text-indent: 28px; color: rgb(102, 102, 102); font-family: Arial; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">在其他方面，一些导演把《梁祝》搬上了银幕和银屏，各种版本应有尽有，梁山伯和祝英台甚至还会武功和轻功，能上天入地。</p><p style=\"margin-top: 0px; margin-bottom: 25px; padding: 0px; line-height: 2.8em; text-indent: 28px; color: rgb(102, 102, 102); font-family: Arial; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">于我而言，最爱的还是钢琴伴奏小提琴独奏曲《梁祝》。它用音乐给人留下了很大的想象和<a href=\"https://www.duanwenxue.com/qinggan/meiwen/\" style=\"margin: 0px; padding: 0px; color: rgb(102, 102, 102); outline: none; text-decoration-line: none;\">欣赏</a>空间。有时候觉得听音乐未必要会乐器、懂乐理、识谱子。只要有<a href=\"https://www.duanwenxue.com/jingdian/shenghuo/\" style=\"margin: 0px; padding: 0px; color: rgb(102, 102, 102); outline: none; text-decoration-line: none;\">生活</a>、有感受、有悟性，人人皆是欣赏者。当你在品尝生活经历的同时，去聆听和欣赏音乐，你的感受会与众不同，甚至会融入角色，身临其境，同音乐中的形象同命运共生死。当钢琴以轻缓的音符营造出一种夜深人静的氛围时，蕴含丰富情愫的小提琴声从演奏者的指间徐徐流淌出来，两个主人公的形象交替出现，渐远渐近，慢慢地向你走来，让你靠近他们，倾听他们由衷的诉说……</p><p style=\"margin-top: 0px; margin-bottom: 25px; padding: 0px; line-height: 2.8em; text-indent: 28px; color: rgb(102, 102, 102); font-family: Arial; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">其实，在倾听过程中，你自己也在用心与音乐作着无声的交流和沟通，倾吐着你的心迹，排解着你的压抑和郁闷。此时此刻，音乐所表达的意境和主题、所塑造的形象，不就是你在现实中所追求的那种意境吗?</p><p style=\"margin-top: 0px; margin-bottom: 25px; padding: 0px; line-height: 2.8em; text-indent: 28px; color: rgb(102, 102, 102); font-family: Arial; font-size: 14px; white-space: normal; background-color: rgb(255, 255, 255);\">欣赏音乐，需要的不是结果，而是过程。一曲《梁祝》久萦怀，其实你也身在其中。</p><p><br/></p>', '2020-03-24 11:55:22', 1, 1, 0, 0, 0, 1, '2020-03-07 11:55:25', '2020-03-07 12:10:02', NULL);
INSERT INTO `pinecms_articles` VALUES (4, '时光不老去 我们再相聚', '', '', '', '', '同窗,相聚,同学,人生,岁月,四十七年,纯真,依然,珍惜,一份,青春,有太多', '', '<p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">&nbsp;时光荏苒，白驹过隙，一晃就是四十七年，我们重相聚。那年我们<a href=\"http://www.duwenzhang.com/huati/likai/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">离开</a>学校，年级的毕业典礼，简单而乏味，没有热烈的拥抱告别、没有拍照留影。离开<a href=\"http://www.duwenzhang.com/wenzhang/xiaoyuanwenzhang/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">校园</a>时，我班的同学要求班主任到照相馆，一定要留个影。镁光灯一闪，同学们那副稚嫩、那般<a href=\"http://www.duwenzhang.com/huati/qingchun/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">青春</a>、那股活力模样就定格在毕业照里。如今，经历了风雨的洗礼，岁月的磨砺，我们很多人都已两鬓如霜、褶皱满面，不再有年少时的青春<a href=\"http://www.duwenzhang.com/huati/langman/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">浪漫</a>，也不再有年轻时的蓬勃朝气，但是在我们中间却贮就了一副<a href=\"http://www.duwenzhang.com/huati/chengshu/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">成熟</a>和稳重，多了一份经世的阅历。岁月如歌，人间沧桑，我们经历了许多许多，许许多多的事情也离我们远去。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　 金色的学生年代，有太多的<a href=\"http://www.duwenzhang.com/huati/gandong/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">感动</a>，有太多的追忆。今天再相聚，我们沿着当年上学所走的路，回母校寻觅青春的足迹。四十七载变迁还是斑驳了痕迹，那教室、那走廊、那礼堂、那操场、那林荫道上，都经历了岁月的洗礼，还有那几棵参天的梧桐树，只能永远记忆在脑海里。校园依旧，物已原非，只有当年学子们的笑声和<a href=\"http://www.duwenzhang.com/huati/laoshi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">老师</a>的谆谆教诲仍在耳边响起！真想用画笔描摹出它的绚丽和深意，但却怎么也画不出了<a href=\"http://www.duwenzhang.com/huati/tongnian/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">童年</a>时代滚铁圈的小伙伴和跳皮筋小<a href=\"http://www.duwenzhang.com/huati/nvhai/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">女孩</a>的身影；画不出跟自己划三八线的小同学那副严肃样和打小报告的臭屁孩那满脸的诡异；画不出嘟嘴埋怨老师管束太严厉的小脸；画不出那借来你还带有余温的那支钢笔；画不出互相帮助时的纯真友爱；画不出那不用做作业时的惬意；画不出操场上追逐的天真浪漫；画不出对那个令自己心动好久的小女孩略带羞涩的表情………！莫怪画笔萎靡，只是里面蕴藏着太多的有趣。这一切都还历历在目，深深地篆刻在脑海里，依然是这样的清晰。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　 离开学校，迈向<a href=\"http://www.duwenzhang.com/huati/shehui/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">社会</a>，我们为<a href=\"http://www.duwenzhang.com/huati/zhuiqiu/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">追求</a><a href=\"http://www.duwenzhang.com/huati/meihao/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">美好</a>的<a href=\"http://www.duwenzhang.com/wenzhang/renshengzheli/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">人生</a>各奔了东西。虽然，我们经历了上山下乡、参军，而后又在各自的工作岗位上奔忙，为家人劳作不息。一路走来，人生的路上并不平坦，甚至还有些崎岖，但我们并没有气馁，依然迈着坚实的步子没有停息！虽然，我们很多人没有跨进过<a href=\"http://www.duwenzhang.com/wenzhang/xiaoyuanwenzhang/daxueshenghuo/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">大学</a>的校门，但我们依旧<a href=\"http://www.duwenzhang.com/huati/nuli/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">努力</a>学习，在社会这个大课堂里，我们学到了大学课程无法涉猎的<a href=\"http://www.duwenzhang.com/wenzhang/jingdianwenzhang/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">经典</a>，并读懂了人生的<a href=\"http://www.duwenzhang.com/huati/xingfu/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">幸福</a>与艰辛！虽然，同窗的你我在改革开放的大潮中，没能成为叱咤风云的大款、现代的土豪，但我们并没有因物质而志短，我们依然保持着旺盛的精神气！虽然，同窗的你我没能混到高官、得到厚禄，但依然胸怀宽广，心地坦然，保持着一个美好的<a href=\"http://www.duwenzhang.com/huati/xinling/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">心灵</a>！始终给自己的心灵一份安慰，给<a href=\"http://www.duwenzhang.com/huati/shengming/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">生命</a>一份真实，给自己一份感激。现在我们都老了，回首看，<a href=\"http://www.duwenzhang.com/huati/pingdan/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">平淡</a>悠闲地欣赏着我们曾经浇灌的<a href=\"http://www.duwenzhang.com/huati/qipan/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">期盼</a>、追求的<a href=\"http://www.duwenzhang.com/huati/mengxiang/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">梦想</a>和付过的艰辛。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　 今天的欢聚，侃谈人生事，世态不炎凉，述说着往昔。谈笑中有人坦言自己四十七年前的劣迹；有人坦言自己曾对某个女同学的着迷。嘻嘻哈哈地，那怕是一次一次张狂的历险，如今听起来也很有意义。往昔的<a href=\"http://www.duwenzhang.com/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">故事</a>，最真，最纯，最痴，许多的粗糙和鲁莽都成了<a href=\"http://www.duwenzhang.com/wenzhang/shenghuosuibi/qushi/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">趣事</a><a href=\"http://www.duwenzhang.com/huati/huiyi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">回忆</a>。一年又一年，服饰在变，一年又一年，容颜在变，没变的却是对彼此的挂念之心。今天相聚时，最多的是问侯、最多的是感慨、最多的还是儿时的那段记忆。这同学的<a href=\"http://www.duwenzhang.com/huati/youyi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">友谊</a>，一片纯真之情，一颗真挚之心，都体现在相逢之间，这同学的情意，来自纯真年代，都储存在芝麻般微小的记忆里。成年人的言辞没有羞怯，在幽默与揶揄的调配中，早已把人生的酸甜苦辣叙尽。如今，你还是你，我还是我，我们各自又有新的事情：孙儿的吃饭、穿衣、接送、学习，还有每天跳广场舞，还要远途旅行。哟，唠唠叨叨怎么有这么多说不完的事，青春不老去，我们还有很多很多要做的事情。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　来也匆匆，去也匆匆，相聚是一种奢求，然，那份同窗之情，依然恒温窖藏于心。打开一坛陈酿之酒，它显得很香、很甜、很浓、很纯。这香甜的酒，绵柔如饴，沁人心脾，醉到了每个人的心，这同学间<a href=\"http://www.duwenzhang.com/huati/zhenqing/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">真情</a>话语，早已让人满<a href=\"http://www.duwenzhang.com/huati/yanlei/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">眼泪</a>痕。醉了，醉了，情愿！呵！情愿醉它千年不醒！在这茫茫人海里，我们何其有幸能够成为同窗？成为学友？是<a href=\"http://www.duwenzhang.com/huati/yuanfen/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">缘分</a>！那些从小萌生的友谊<a href=\"http://www.duwenzhang.com/huati/ganqing/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">感情</a>，始终是我们心灵的支柱，不管是浓是淡，是远是近，它一直都令我们<a href=\"http://www.duwenzhang.com/huati/zhenxi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">珍惜</a>，珍惜共同走过的那段岁月，珍惜我们在一起<a href=\"http://www.duwenzhang.com/wenzhang/shenghuosuibi/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">生活</a>的点点滴滴。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　啊！欢聚毕竟是短暂的，虽潋滟起沉淀的<a href=\"http://www.duwenzhang.com/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">情感</a>畅谈不尽，我们各处五湖四海、南北东西。挽留！可终究挽不住你我的去留，大家各自还有许多的事情。然“流水不因石而阻，情谊不因远而疏”，我们笃信，这次四十七年分别的再聚首，同窗间的情谊将会愈加深厚，同时我们也期盼：保重！再过十年、二十年，三十年，青春不老去，我们再相聚！待那时杖国高年的同窗重聚，那自然情趣更浓，别有一番深意！</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　 坚信你会来的！而我也一定会与大家一起来相聚！</p>', '2020-03-24 17:49:10', 4, 1, 0, 0, 0, 1, '2020-03-07 17:49:12', NULL, NULL);

-- ----------------------------
-- Table structure for pinecms_attachments
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_attachments`;
CREATE TABLE `pinecms_attachments`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `url` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `origin_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `size` int(11) NULL DEFAULT 0,
  `upload_time` datetime(0) NULL DEFAULT NULL,
  `type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 150 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '附件表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_attachments
-- ----------------------------
INSERT INTO `pinecms_attachments` VALUES (3, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (4, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (5, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (6, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (7, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (8, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (9, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (10, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (11, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (12, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (13, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (14, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (15, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (16, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (17, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (18, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (19, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (20, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (21, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (22, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (23, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (24, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (25, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (26, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (27, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (28, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (29, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (30, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (31, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (32, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (33, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (34, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (35, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (36, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (37, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (38, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (39, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (40, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (41, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (42, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (43, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (44, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (45, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (46, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (47, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (48, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (49, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (50, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (51, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (52, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (53, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (54, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (55, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (56, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (57, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (58, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (59, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (60, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (61, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (62, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (63, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (64, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (65, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (66, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (67, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (68, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (69, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (70, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (71, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (72, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (73, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (74, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (75, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (76, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (77, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (78, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (79, 'HOT4zStb14.jpg', '/upload/public/20200210/HOT4zStb14.jpg', '64673db6d9.jpg', 604762, '2020-02-10 17:09:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (80, '41nyeDFxxU.jpg', '/upload/public/20200210/41nyeDFxxU.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:40:34', 'img');
INSERT INTO `pinecms_attachments` VALUES (81, 'Kdy4wdTWwO.jpg', '/upload/public/20200210/Kdy4wdTWwO.jpg', '64673db6d9.jpg', 218356, '2020-02-10 17:42:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (82, 'AQ7yfBTKC4.jpg', '/upload/public/20200210/AQ7yfBTKC4.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (83, 'eSId6e85oe.jpg', '/upload/public/20200210/eSId6e85oe.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (84, '71okeV2pq8.jpg', '/upload/public/20200210/71okeV2pq8.jpg', '64673db6d9.jpg', 218356, '2020-02-10 19:04:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (85, 'g10BxfuPlf.png', '/upload/public/20200210/g10BxfuPlf.png', 'K9k3WXVj3d.png', 0, '2020-02-10 19:35:22', 'img');
INSERT INTO `pinecms_attachments` VALUES (86, 'E98Sy9Cb9T.png', '/upload/public/20200210/E98Sy9Cb9T.png', 'N9DhuBERNF.png', 0, '2020-02-10 19:38:10', 'img');
INSERT INTO `pinecms_attachments` VALUES (87, '1efJjx17WH.jpg', '/upload/public/20200210/1efJjx17WH.jpg', '64673db6d9.jpg', 604762, '2020-02-10 19:40:13', 'img');
INSERT INTO `pinecms_attachments` VALUES (88, 'J07ibEIWpi.jpg', '/upload/public/20200211/J07ibEIWpi.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:26:57', 'img');
INSERT INTO `pinecms_attachments` VALUES (89, 'oyzQ8OsDTQ.jpg', '/upload/public/20200211/oyzQ8OsDTQ.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:28:20', 'img');
INSERT INTO `pinecms_attachments` VALUES (90, 'yz4U6V0CXg.jpg', '/upload/public/20200211/yz4U6V0CXg.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:29:55', 'img');
INSERT INTO `pinecms_attachments` VALUES (91, 'Ekf2kZ7kQS.jpg', '/upload/public/20200211/Ekf2kZ7kQS.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:30:12', 'img');
INSERT INTO `pinecms_attachments` VALUES (92, 'Y3v355GWcz.jpg', '/upload/public/20200211/Y3v355GWcz.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:30:54', 'img');
INSERT INTO `pinecms_attachments` VALUES (93, 'W2A95f6v1M.jpg', '/upload/public/20200211/W2A95f6v1M.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:31:51', 'img');
INSERT INTO `pinecms_attachments` VALUES (94, 'tWXdAmSPLq.jpg', '/upload/public/20200211/tWXdAmSPLq.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:32:30', 'img');
INSERT INTO `pinecms_attachments` VALUES (95, 'z2gvYbzb12.jpg', '/upload/public/20200211/z2gvYbzb12.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:33:09', 'img');
INSERT INTO `pinecms_attachments` VALUES (96, 'PYl5gx1cs0.jpg', '/upload/public/20200211/PYl5gx1cs0.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:34:17', 'img');
INSERT INTO `pinecms_attachments` VALUES (97, '8ZI337YIgP.jpg', '/upload/public/20200211/8ZI337YIgP.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:35:39', 'img');
INSERT INTO `pinecms_attachments` VALUES (98, 'MjvO7j6Pdd.jpg', '/upload/public/20200211/MjvO7j6Pdd.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:48:01', 'img');
INSERT INTO `pinecms_attachments` VALUES (99, 'OH7F7fV644.jpg', '/upload/public/20200211/OH7F7fV644.jpg', '64673db6d9.jpg', 604762, '2020-02-11 11:48:14', 'img');
INSERT INTO `pinecms_attachments` VALUES (100, '86ZI22CG39.jpg', '/upload/public/20200211/86ZI22CG39.jpg', '64673db6d9.jpg', 604762, '2020-02-11 12:12:48', 'img');
INSERT INTO `pinecms_attachments` VALUES (101, 'T6ia54P5Fb.jpg', '/upload/public/20200211/T6ia54P5Fb.jpg', '64673db6d9.jpg', 604762, '2020-02-11 15:01:27', 'img');
INSERT INTO `pinecms_attachments` VALUES (102, 'TECwi9HrNX.jpg', '/upload/public/20200211/TECwi9HrNX.jpg', '64673db6d9.jpg', 604762, '2020-02-11 21:07:09', 'img');
INSERT INTO `pinecms_attachments` VALUES (103, '56ooFr2t27.jpg', '/upload/public/20200212/56ooFr2t27.jpg', '64673db6d9.jpg', 604762, '2020-02-12 12:33:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (104, 'gRl0i0cP7X.png', '//public/20200212/gRl0i0cP7X.png', 'page.png', 202109, '2020-02-12 19:36:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (105, 'ixaGq0EChv.png', '//public/20200212/ixaGq0EChv.png', 'page.png', 202109, '2020-02-12 19:36:47', 'img');
INSERT INTO `pinecms_attachments` VALUES (107, 'z9f5WOgTRa.png', 'public/20200212/z9f5WOgTRa.png', 'page.png', 202109, '2020-02-12 19:47:42', 'img');
INSERT INTO `pinecms_attachments` VALUES (108, 'iwfz5POG9y.png', 'public/20200212/iwfz5POG9y.png', 'page.png', 202109, '2020-02-12 19:48:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (109, 'jy54Y6kQr6.png', 'public/20200212/jy54Y6kQr6.png', 'page.png', 202109, '2020-02-12 19:49:21', 'img');
INSERT INTO `pinecms_attachments` VALUES (110, '8enWIHKc7h.png', 'public/20200212/8enWIHKc7h.png', 'page.png', 202109, '2020-02-12 19:50:06', 'img');
INSERT INTO `pinecms_attachments` VALUES (113, '564u4p4rx4.png', '/upload/public/20200212/564u4p4rx4.png', 'page.png', 202109, '2020-02-12 19:51:43', 'img');
INSERT INTO `pinecms_attachments` VALUES (115, 'Qc0NY6ME36.png', 'http://iriscms-test.oss-cn-beijing.aliyuncs.com/public/20200212/Qc0NY6ME36.png', 'page.png', 202109, '2020-02-12 19:58:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (116, 'AYcL1slsZm.png', '/upload/public/20200221/AYcL1slsZm.png', '15d5e6410d49263.90261089.png', 221324, '2020-02-21 11:59:07', 'img');
INSERT INTO `pinecms_attachments` VALUES (121, 'OBXu6mIm5i.jpg', '/upload/public/20200303/OBXu6mIm5i.jpg', 'banker.jpg', 88751, '2020-03-03 11:54:04', 'img');
INSERT INTO `pinecms_attachments` VALUES (122, '963rCk18MP.jpg', '/upload/public/20200303/963rCk18MP.jpg', 'banker.jpg', 88751, '2020-03-03 11:54:23', 'img');
INSERT INTO `pinecms_attachments` VALUES (123, 'kHvGB7UMjv.jpg', '/upload/public/20200303/kHvGB7UMjv.jpg', 'banker.jpg', 88751, '2020-03-03 11:56:01', 'img');
INSERT INTO `pinecms_attachments` VALUES (124, '7I2RO0j8VU.jpg', '/upload/public/20200303/7I2RO0j8VU.jpg', 'banker.jpg', 88751, '2020-03-03 11:56:17', 'img');
INSERT INTO `pinecms_attachments` VALUES (125, 'CY9yw7B9G2.jpg', '/upload/public/20200303/CY9yw7B9G2.jpg', 'banker.jpg', 88751, '2020-03-03 14:36:45', 'img');
INSERT INTO `pinecms_attachments` VALUES (126, '1EZ28y5M4Q.jpg', '/upload/public/20200303/1EZ28y5M4Q.jpg', 'banker.jpg', 88751, '2020-03-03 14:36:49', 'img');
INSERT INTO `pinecms_attachments` VALUES (127, 'zebRQfaG71.png', '/upload/public/20200306/zebRQfaG71.png', 'appdown2.png', 9820, '2020-03-06 15:59:56', 'img');
INSERT INTO `pinecms_attachments` VALUES (128, 'BFZ5z5e5R9.jpg', '/upload/public/20200309/BFZ5z5e5R9.jpg', 'banker.jpg', 88751, '2020-03-09 10:27:22', 'img');
INSERT INTO `pinecms_attachments` VALUES (129, '3Yx7CiKeB5.jpg', '/upload/public/20200309/3Yx7CiKeB5.jpg', 'banker.jpg', 88751, '2020-03-09 10:29:37', 'img');
INSERT INTO `pinecms_attachments` VALUES (130, 'k0SXAg1oNW.jpg', '/upload/public/20200309/k0SXAg1oNW.jpg', 'banker.jpg', 88751, '2020-03-09 10:29:51', 'img');
INSERT INTO `pinecms_attachments` VALUES (133, 'CaB53Ew4oF.png', '/upload/public/20200309/CaB53Ew4oF.png', 'ICMP.png', 7404, '2020-03-09 10:41:00', 'img');
INSERT INTO `pinecms_attachments` VALUES (134, 'BXxEyKCP0o.png', '/upload/public/20200310/BXxEyKCP0o.png', '动漫.png', 1028236, '2020-03-10 20:13:32', 'img');
INSERT INTO `pinecms_attachments` VALUES (135, 'F9CEecV0a4.png', '/upload/public/20200310/F9CEecV0a4.png', '动漫.png', 1028236, '2020-03-10 20:21:35', 'img');
INSERT INTO `pinecms_attachments` VALUES (136, 'st1Cu5DfW6.png', '/upload/public/20200311/st1Cu5DfW6.png', 'self_weixin_erweima.png', 18424, '2020-03-11 21:32:10', 'img');
INSERT INTO `pinecms_attachments` VALUES (137, 'HXhhLGDwf8.png', '/upload/public/20200312/HXhhLGDwf8.png', '1.png', 19521, '2020-03-12 17:13:01', 'img');
INSERT INTO `pinecms_attachments` VALUES (138, 'R33eGyVe0v.png', '/upload/public/20200312/R33eGyVe0v.png', '1.png', 19521, '2020-03-12 17:13:01', 'img');
INSERT INTO `pinecms_attachments` VALUES (139, 'v3kPak3gS5.png', '/upload/public/20200312/v3kPak3gS5.png', 'default_man.png', 10811, '2020-03-12 17:20:28', 'img');
INSERT INTO `pinecms_attachments` VALUES (140, '3wHh8hrjzZ.jpg', '/upload/public/20200321/3wHh8hrjzZ.jpg', '542.Carousel1-potassium-1024w-1366h@2x_ipad.jpg', 7287249, '2020-03-21 12:06:58', 'img');
INSERT INTO `pinecms_attachments` VALUES (141, 'Tog81s99Q2.png', '/upload/public/20200321/Tog81s99Q2.png', 'modul.png', 15459, '2020-03-21 12:10:09', 'img');
INSERT INTO `pinecms_attachments` VALUES (142, 'u9dP9AoSmo.png', '/upload/public/20200321/u9dP9AoSmo.png', 'cross.png', 438, '2020-03-21 12:11:03', 'img');
INSERT INTO `pinecms_attachments` VALUES (143, '41Q6FIwac4.gif', '/upload/public/20200321/41Q6FIwac4.gif', 'logo.gif', 3427, '2020-03-21 12:17:58', 'file');
INSERT INTO `pinecms_attachments` VALUES (144, 's49kowaHVY.png', '/upload/public/20200321/s49kowaHVY.png', 'hei_logo.png', 1737, '2020-03-21 12:18:41', 'file');
INSERT INTO `pinecms_attachments` VALUES (145, 'RZ5138jE7i.png', '/upload/public/20200321/RZ5138jE7i.png', 'hei_logo.png', 1737, '2020-03-21 12:19:48', 'file');
INSERT INTO `pinecms_attachments` VALUES (146, 'fMu0Kbd99o.png', '/upload/public/20200321/fMu0Kbd99o.png', 'hei_logo.png', 1737, '2020-03-21 12:20:15', 'file');
INSERT INTO `pinecms_attachments` VALUES (147, '89ac9Vv86P.png', '/upload/public/20200321/89ac9Vv86P.png', 'modul.png', 15459, '2020-03-21 12:23:17', 'file');
INSERT INTO `pinecms_attachments` VALUES (148, '1jPeJ02eR9.gif', '/upload/public/20200321/1jPeJ02eR9.gif', 'logo.gif', 3427, '2020-03-21 14:07:39', 'file');
INSERT INTO `pinecms_attachments` VALUES (149, 'ThECSqh65Q.zip', '/upload/public/20200321/ThECSqh65Q.zip', '归档.zip', 24543728, '2020-03-21 14:43:10', 'file');

-- ----------------------------
-- Table structure for pinecms_category
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_category`;
CREATE TABLE `pinecms_category`  (
  `catid` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '栏目类型',
  `parentid` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父类ID',
  `catname` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '分类名称',
  `description` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '描述',
  `model_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '模型id,仅栏目类型为0的时候可用',
  `manager_content_router` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `edit_content_router` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `add_content_router` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `url` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '链接地址, 仅栏目类型为2的时候可用',
  `listorder` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序号',
  `ismenu` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否为栏目, 最初设定是可以在前端展示',
  `index_tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '后台列表',
  `list_tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前台列表模板',
  `detail_tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前台内容页模板',
  `thumb` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '栏目缩略图',
  `tpl_prefix` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `home_tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `content_tpl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`catid`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '内容分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_category
-- ----------------------------
INSERT INTO `pinecms_category` VALUES (1, 0, 0, '随意速记', '发布杂乱文章内容', 1, '', '', '', '', 1, 1, '', '', '', '', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (2, 0, 0, '最新动态', '记录项目最新情况', 1, '', '', '', '', 2, 1, '', '', '', '', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (3, 0, 1, '项目速记', '关于项目的速记内容', 1, '', '', '', '', 0, 1, '', '', '', '', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (4, 0, 1, '工作速记', '工作内容相关速记', 1, '', '', '', '', 0, 1, '', '', '', '', NULL, NULL, NULL);
INSERT INTO `pinecms_category` VALUES (5, 1, 0, '关于项目', '项目介绍页面', 0, '', '', '', '', 0, 1, '', '', '', '', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for pinecms_category_priv
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_category_priv`;
CREATE TABLE `pinecms_category_priv`  (
  `catid` smallint(6) NOT NULL DEFAULT 0,
  `roleid` smallint(6) NOT NULL DEFAULT 0,
  `is_admin` tinyint(4) NOT NULL DEFAULT 0,
  `action` char(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  INDEX `catid`(`catid`, `roleid`, `is_admin`, `action`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '栏目权限表' ROW_FORMAT = Fixed;

-- ----------------------------
-- Table structure for pinecms_content
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_content`;
CREATE TABLE `pinecms_content`  (
  `id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `catid` smallint(5) UNSIGNED NOT NULL DEFAULT 0,
  `title` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `thumb` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `keywords` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `description` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `listorder` tinyint(3) UNSIGNED NOT NULL DEFAULT 0,
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1,
  `recommend` tinyint(4) NULL DEFAULT NULL,
  `pwd_type` tinyint(4) NULL DEFAULT NULL,
  `money` tinyint(4) NULL DEFAULT NULL,
  `created_at` int(11) NULL DEFAULT NULL,
  `updated_at` int(11) NULL DEFAULT NULL,
  `deleted_at` int(11) NULL DEFAULT NULL,
  `source_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `source_pwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `catids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tags` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `userid` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '新闻表后期根据模型扩展' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pinecms_document_model
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_document_model`;
CREATE TABLE `pinecms_document_model`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文档名称',
  `table` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '对应的表名',
  `enabled` tinyint(4) NULL DEFAULT 0 COMMENT '是否启用',
  `model_type` tinyint(4) NULL DEFAULT 0 COMMENT '模型类型: 扩展模型 和 独立模型',
  `fe_tpl_index` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '模型前端主页模板地址',
  `fe_tpl_list` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '模型前端列表模板地址',
  `fe_tpl_detail` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '模型前端详情模板地址',
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `field_show_in_list` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '在后端列表页需要展示的字段以及字段应用的formatter函数.',
  `formatters` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '所有函数内容, 原样渲染到Html里',
  `execed` tinyint(4) NULL DEFAULT 0 COMMENT '是否已经执行过改动',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文档模型用于存储自定义类型的文档内容' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_document_model
-- ----------------------------
INSERT INTO `pinecms_document_model` VALUES (1, '普通文章', 'articles', 1, 1, 'resources/views/frontend/index.jet', '', '', NULL, '', '', 0);
INSERT INTO `pinecms_document_model` VALUES (2, '下载模型', 'download', 1, 0, '', '', '', NULL, '', '', 0);

-- ----------------------------
-- Table structure for pinecms_document_model_dsl
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_document_model_dsl`;
CREATE TABLE `pinecms_document_model_dsl`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mid` int(11) NOT NULL DEFAULT 0 COMMENT '模型id',
  `form_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字段在表单内的名称',
  `listorder` int(11) UNSIGNED NOT NULL DEFAULT 0,
  `table_field` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `html` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '字段html',
  `required` tinyint(4) NULL DEFAULT 0 COMMENT '是否必填',
  `datasource` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据源, 可以让下拉选项等高级功能有数据读取的源头,具体设计可以是提供列表函数类的',
  `required_tips` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '必填(选)提醒',
  `validator` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '验证器名称或内容',
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `field_type` int(11) NOT NULL DEFAULT 0,
  `default` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 106 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '模型表单定义表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_document_model_dsl
-- ----------------------------
INSERT INTO `pinecms_document_model_dsl` VALUES (46, 0, '标题', 1, 'title', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (47, 0, '关键字', 1, 'keywords', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (48, 0, '摘要', 1, 'description', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"height:80px; width: 300px;\"  multiline />', 0, '', '', '', NULL, 2, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (49, 1, '文章标题', 1, 'title', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (50, 1, 'TAG标签', 1, 'tags', '<tags />', 0, '', '', '', NULL, 15, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (51, 1, '缩略图', 1, 'thumb', '<images />', 0, '', '', '', NULL, 11, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (52, 1, '文章来源', 1, 'from_url', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (53, 1, '作者', 1, 'author', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (54, 1, '关键字', 1, 'keywords', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (55, 1, '内容摘要', 1, 'description', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"height:80px; width: 300px;\"  multiline />', 0, '', '', '', NULL, 2, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (56, 1, '文章内容', 1, 'content', '<editor />', 1, '', '', '', NULL, 3, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (57, 1, '发布时间', 1, 'pubtime', '<input class=\"easyui-datetimebox\" style=\"width:300px\" value=\"{{value}}\" {{attr}} />', 0, '', '', '', NULL, 14, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (98, 2, '标题', 1, 'title', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (99, 2, '关键字', 2, 'keywords', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (100, 2, '描述', 3, 'description', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (101, 2, '内容', 4, 'content', '<editor />', 0, '', '', '', NULL, 3, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (102, 2, '上传文件', 5, 'files', '<input {{attr}} type=\"text\" readonly value=\"{{value}}\" style=\"display:none;\"/><button class=\"btn btn-default\" type=\"button\" onclick=\"fromUEFileUploader(this, 0)\">选择文件</button><div class=\'easy-uploader\'><ul class=\"list\"></ul></div>', 0, '', '', '', NULL, 4, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (103, 2, '文件类型', 6, 'ext', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">', 0, '', '', '', NULL, 1, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (104, 2, '文件大小', 7, 'size', '<input type=\"text\" class=\"easyui-numberbox\" value=\"{{value}}\" {{attr}} />', 0, '', '', '', NULL, 10, '');
INSERT INTO `pinecms_document_model_dsl` VALUES (105, 2, '下载次数', 8, 'downs', '<input type=\"text\" class=\"easyui-numberbox\" value=\"{{value}}\" {{attr}} />', 0, '', '', '', NULL, 9, '');

-- ----------------------------
-- Table structure for pinecms_document_model_field
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_document_model_field`;
CREATE TABLE `pinecms_document_model_field`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字段名称',
  `type` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字段对应的数据类型',
  `desc` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字段描述',
  `html` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '模型表单组件定义表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_document_model_field
-- ----------------------------
INSERT INTO `pinecms_document_model_field` VALUES (1, '单行文本', 'varchar', '常用字段，如文章标题、作者等都属于直接输入少量内容的文本，设置这个文本之后需要指定文本长度，默认为250，如果大于255则为text类型', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"width:300px\">');
INSERT INTO `pinecms_document_model_field` VALUES (2, '多行文本', 'text', '也是较为常用的字段类型，如个人简介、产品描述都可以使用多行文本进行存储', '<input class=\"easyui-textbox\" {{attr}} value=\"{{value}}\" style=\"height:80px; width: 300px;\"  multiline />');
INSERT INTO `pinecms_document_model_field` VALUES (3, 'HTML文本', 'text', '编辑器编辑产生的html内容，用于比较复杂的内容形式, 可以认为是附带编辑器的多行文本', '<editor />');
INSERT INTO `pinecms_document_model_field` VALUES (4, '附件', 'varshar', '前端表现为input[file]类型,可以上传并且返回一个相对地址', '<input {{attr}} type=\"text\" readonly value=\"{{value}}\" style=\"display:none;\"/><button class=\"btn btn-default\" type=\"button\" onclick=\"fromUEFileUploader(this, 0)\">选择文件</button><div class=\'easy-uploader\'><ul class=\"list\"></ul></div>');
INSERT INTO `pinecms_document_model_field` VALUES (5, '下拉框', 'varchar', '下拉选择，一般用于如软件类型、语言类型等字段', '<input class=\"easyui-combobox\" {{attr}} style=\"width:300px;\" value=\"{{value}}\" />');
INSERT INTO `pinecms_document_model_field` VALUES (6, '联动类型', 'varchar', '一种数组形式的数据类型，请使用url接口方式提供', '<select class=\"easyui-combotree\" {{attr}} style=\"width:200px;\">暂未实现</select>');
INSERT INTO `pinecms_document_model_field` VALUES (7, '单选框', 'varchar', '平铺显示, 可以认为是下拉框的展开, 根据数据源展开为排列的组件', '<input class=\"easyui-radiobutton\" {{attr}} value=\"{{value}}\" {{default}} />');
INSERT INTO `pinecms_document_model_field` VALUES (8, '多选框', 'varchar', '多选框, 平铺显示为多个选项,根据数据源展开为排列组件', '<input class=\"easyui-checkbox\" {{attr}} value=\"{{value}}\" {{default}} />');
INSERT INTO `pinecms_document_model_field` VALUES (9, '整数类型', 'int', '常用字段, 仅能输入数字', '<input type=\"text\" class=\"easyui-numberbox\" value=\"{{value}}\" {{attr}} />');
INSERT INTO `pinecms_document_model_field` VALUES (10, '浮点类型', 'float', '常用字段, 仅能输入浮点数(小数)', '<input type=\"text\" class=\"easyui-numberbox\" value=\"{{value}}\" {{attr}} />');
INSERT INTO `pinecms_document_model_field` VALUES (11, '单图上传', 'varchar', '常用字段, 会生成一个单图上传框', '<images />');
INSERT INTO `pinecms_document_model_field` VALUES (12, '多图上传', 'varchar', '生成一个多图上传的组件', '<mul-images />');
INSERT INTO `pinecms_document_model_field` VALUES (13, '开关按钮', 'tinyint', '用于做开关值的组件, 打开为1, 关闭为0', '<input class=\"easyui-switchbutton\" value=\"{{value}}\" {{attr}} {{default}} />');
INSERT INTO `pinecms_document_model_field` VALUES (14, '日历组件', 'datetime', '选择日期组件', '<input class=\"easyui-datetimebox\" style=\"width:300px\" value=\"{{value}}\" {{attr}} />');
INSERT INTO `pinecms_document_model_field` VALUES (15, '多选标签', 'varchar', '可以记录标签，并多选', '<tags />');

-- ----------------------------
-- Table structure for pinecms_link
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_link`;
CREATE TABLE `pinecms_link`  (
  `linkid` smallint(6) NOT NULL AUTO_INCREMENT,
  `linktype` tinyint(1) NOT NULL DEFAULT 0,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `logo` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `introduce` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `listorder` smallint(6) NOT NULL DEFAULT 0,
  `passed` tinyint(1) NOT NULL DEFAULT 0,
  `addtime` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`linkid`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '友情链接表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_link
-- ----------------------------
INSERT INTO `pinecms_link` VALUES (5, 0, '百度', 'https://www.baidu.com', 'https://www.baidu.com/img/bd_logo1.png?where=super', '', 1, 1, '2020-02-29 21:19:40');
INSERT INTO `pinecms_link` VALUES (6, 0, '新浪', 'http://www.baidu.com', 'https://i1.sinaimg.cn/dy/deco/2013/0329/logo/LOGO_1x.png', '', 2, 1, '2020-02-29 21:19:30');

-- ----------------------------
-- Table structure for pinecms_log
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_log`;
CREATE TABLE `pinecms_log`  (
  `logid` int(11) NOT NULL AUTO_INCREMENT,
  `controller` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `action` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `querystring` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `userid` mediumint(9) NOT NULL DEFAULT 0,
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `time` datetime(0) NOT NULL,
  PRIMARY KEY (`logid`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 710 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '操作日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pinecms_member
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_member`;
CREATE TABLE `pinecms_member`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `integral` int(11) NULL DEFAULT NULL,
  `sale_integral` int(11) NULL DEFAULT NULL,
  `draw_account` varchar(0) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `telphone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `qq` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `enabled` tinyint(4) NOT NULL DEFAULT 0,
  `verify_token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_member
-- ----------------------------
INSERT INTO `pinecms_member` VALUES (1, 'xiusin', '159781', '', '陈二皮', 1231211111, 0, '', '123123', '1111222', '', '2019-01-24 11:40:00', '2019-01-24 11:40:00', '159781@11.com', 1, '4b32a22c-5787-4d0b-98f2-ed5b0779bbcb');

-- ----------------------------
-- Table structure for pinecms_menu
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_menu`;
CREATE TABLE `pinecms_menu`  (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `name` char(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `parentid` smallint(6) NOT NULL DEFAULT 0,
  `c` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `a` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `data` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `is_system` tinyint(1) NOT NULL DEFAULT 0,
  `listorder` smallint(6) NOT NULL DEFAULT 0,
  `display` enum('1','0') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 98 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限菜单表' ROW_FORMAT = Fixed;

-- ----------------------------
-- Records of pinecms_menu
-- ----------------------------
INSERT INTO `pinecms_menu` VALUES (1, '我的面板', 0, 'admin', 'public-top', '', 1, 1, '1');
INSERT INTO `pinecms_menu` VALUES (3, '内容管理', 1, 'content', 'top', '', 0, 2, '1');
INSERT INTO `pinecms_menu` VALUES (54, '分类单页', 35, 'content', 'page', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (6, '个人信息', 1, 'admin', 'public-left', '', 1, 0, '1');
INSERT INTO `pinecms_menu` VALUES (7, '修改密码', 6, 'admin', 'public-editpwd', '', 1, 1, '1');
INSERT INTO `pinecms_menu` VALUES (8, '修改个人信息', 6, 'admin', 'public-editinfo', '', 1, 0, '1');
INSERT INTO `pinecms_menu` VALUES (9, '系统设置', 1, 'setting', 'left', '', 0, 2, '1');
INSERT INTO `pinecms_menu` VALUES (10, '站点设置', 9, 'setting', 'site', '', 0, 1, '1');
INSERT INTO `pinecms_menu` VALUES (11, '管理员设置', 1, 'admin', 'left', '', 0, 4, '1');
INSERT INTO `pinecms_menu` VALUES (12, '管理员管理', 11, 'admin', 'memberlist', '', 0, 1, '1');
INSERT INTO `pinecms_menu` VALUES (13, '角色管理', 11, 'admin', 'rolelist', '', 0, 2, '1');
INSERT INTO `pinecms_menu` VALUES (14, '日志管理', 1, 'system', 'loglist', '', 0, 1, '1');
INSERT INTO `pinecms_menu` VALUES (15, '操作日志', 14, 'system', 'loglist', '', 0, 1, '1');
INSERT INTO `pinecms_menu` VALUES (16, '菜单管理', 9, 'system', 'menulist', '', 0, 2, '1');
INSERT INTO `pinecms_menu` VALUES (17, '查看菜单', 16, 'system', 'menuview', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (18, '添加菜单', 16, 'system', 'menuadd', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (19, '修改菜单', 16, 'system', 'menuedit', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (20, '删除菜单', 16, 'system', 'menudelete', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (21, '菜单排序', 16, 'system', 'menuorder', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (22, '查看日志', 15, 'system', 'logview', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (23, '删除日志', 15, 'system', 'log-delete', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (24, '查看管理员', 12, 'admin', 'member-view', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (25, '添加管理员', 12, 'admin', 'member-add', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (26, '编辑管理员', 12, 'admin', 'member-edit', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (27, '删除管理员', 12, 'admin', 'member-delete', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (28, '查看角色', 13, 'admin', 'rolelist', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (29, '添加角色', 13, 'admin', 'role-add', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (30, '编辑角色', 13, 'admin', 'role-edit', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (31, '删除角色', 13, 'admin', 'role-delete', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (53, '新闻列表', 35, 'content', 'news-list', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (33, '权限设置', 13, 'admin', 'role-permission', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (34, '发布管理', 3, 'content', 'right', '', 0, 0, '0');
INSERT INTO `pinecms_menu` VALUES (35, '内容管理', 36, 'content', 'index', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (36, '栏目管理', 3, 'category', 'list', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (37, '查看栏目', 36, 'category', 'view', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (38, '添加栏目', 36, 'category', 'add', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (39, '编辑栏目', 36, 'category', 'edit', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (40, '删除栏目', 36, 'category', 'delete', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (41, '栏目排序', 36, 'category', 'order', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (55, '会员管理', 1, 'user', 'list', '', 0, 5, '0');
INSERT INTO `pinecms_menu` VALUES (56, '会员列表', 55, 'user', 'list', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (57, '会员信息', 56, 'user', 'info', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (58, '微信管理', 1, 'wechat', 'userlist', '', 0, 7, '0');
INSERT INTO `pinecms_menu` VALUES (59, '微信会员信息', 58, 'wechat', 'userinfo', '', 0, 0, '0');
INSERT INTO `pinecms_menu` VALUES (60, '编辑会员', 55, 'user', 'edit', '', 0, 0, '0');
INSERT INTO `pinecms_menu` VALUES (62, '模型管理', 9, 'model', 'list', '', 0, 1, '1');
INSERT INTO `pinecms_menu` VALUES (64, '添加模型', 62, 'model', 'add', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (65, '微信会员列表', 58, 'wechat', 'userlist', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (67, '友链管理', 1, 'link', 'list', '', 0, 88, '1');
INSERT INTO `pinecms_menu` VALUES (68, '友链管理', 67, 'link', 'list', '', 0, 88, '1');
INSERT INTO `pinecms_menu` VALUES (69, '友链添加', 68, 'link', 'add', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (70, '友链编辑', 68, 'link', 'edit', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (71, '友链删除', 68, 'link', 'delete', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (72, '友链排序', 68, 'link', 'order', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (73, '数据库管理', 1, 'database', 'manager', '', 0, 3, '1');
INSERT INTO `pinecms_menu` VALUES (74, '数据库管理', 73, 'database', 'manager', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (75, '数据库备份', 74, 'database', 'backup', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (76, '数据库优化', 74, 'database', 'optimize', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (77, '数据库修复', 74, 'database', 'repair', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (78, '备份列表', 73, 'database', 'backup-list', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (79, '资源管理', 1, 'assets-manager', 'list', '', 0, 3, '1');
INSERT INTO `pinecms_menu` VALUES (80, '资源列表', 79, 'assets-manager', 'list', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (81, '添加资源', 80, 'assets-manager', 'add', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (82, '修改资源', 80, 'assets-manager', 'edit', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (84, '附件列表', 79, 'attachments', 'list', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (85, '实时日志', 14, 'system', 'tail', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (86, '修改模型', 62, 'model', 'edit', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (87, '主题列表', 79, 'assets-manager', 'theme', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (88, '广告管理', 1, 'ad', 'list', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (89, '广告列表', 88, 'ad', 'list', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (90, '广告位', 88, 'ad-space', 'list', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (91, '添加广告', 89, 'ad', 'add', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (92, '修改广告', 89, 'ad', 'edit', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (93, '删除广告', 89, 'ad', 'delete', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (94, '添加广告位', 90, 'ad-space', 'add', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (95, '删除广告位', 90, 'ad-space', 'delete', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (96, '编辑广告位', 90, 'ad-space', 'edit', '', 0, 0, '1');
INSERT INTO `pinecms_menu` VALUES (97, 'TODO', 0, 'public', 'todos', '', 0, 0, '0');

-- ----------------------------
-- Table structure for pinecms_page
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_page`;
CREATE TABLE `pinecms_page`  (
  `catid` int(11) NOT NULL DEFAULT 0,
  `title` varchar(160) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `keywords` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `updatetime` int(11) NOT NULL DEFAULT 0
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '单页内容表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_page
-- ----------------------------
INSERT INTO `pinecms_page` VALUES (5, '关于项目', '', '', '<p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　1、有时候，我们等的不是什么人、什么事，我们等的是<a href=\"http://www.duwenzhang.com/huati/shijian/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">时间</a>，等时间，让自己改变。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　2、如果<a href=\"http://www.duwenzhang.com/wenzhang/shenghuosuibi/\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">生活</a>中，有一个人想把你踩在脚下，不要以为生活错待了你。或许，还有十个人想要把你踩在脚下，只是你的强大，让他们没有机会伸出脚来。不要抱怨这个世界弱肉强食，你逐渐会发现，它看起来很残酷，却十分公正。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　3、人总是各有苦衷和不甘平庸。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　4、躺在床上听着歌，忘了疼痛，忘了所有的一切。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　5、夜来皓月才当午，重帘悄悄无人语。深处麝烟长，卧时留薄妆。当年还自惜，往事那堪忆。花落月明残，锦衾知晓寒。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　6、不要等到黑夜降临才注意到星星的光芒，其实它们一直在那儿。不要等到<a href=\"http://www.duwenzhang.com/huati/gudu/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">孤独</a>才想起真正对你好的人，其实她们一直在那儿。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　7、我们如此用心追随，<a href=\"http://www.duwenzhang.com/huati/qipan/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">期盼</a>我们的盛世安宁。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　8、如果发短信息给一个人，他一直不回，不要再发了，没有这么卑微的<a href=\"http://www.duwenzhang.com/huati/dengdai/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">等待</a>；如果一个人开始怠慢你，请你<a href=\"http://www.duwenzhang.com/huati/likai/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">离开</a>他，不懂得<a href=\"http://www.duwenzhang.com/huati/zhenxi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">珍惜</a>你的人不要为之不舍。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　9、漫长的别离里，我只想做一件事：专职爱你。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　10、走了那么多弯路，终于回到了最想来的地方。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　11、如果宁缺毋滥的结果将是孤独终老 是否你还能从一而终。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　12、我<a href=\"http://www.duwenzhang.com/huati/nuli/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">努力</a>坚持不<a href=\"http://www.duwenzhang.com/huati/fangqi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">放弃</a>，把委屈通通都咽下去。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　13、你在我的<a href=\"http://www.duwenzhang.com/huati/huiyi/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">回忆</a>里<a href=\"http://www.duwenzhang.com/huati/meihao/index1.html\" style=\"color: rgb(51, 51, 51); text-decoration-line: none;\">美好</a>的不像话，你在我的念想里灿烂的一塌糊涂。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　14、不要把自己的事情掏心掏肺地告诉别人，你知不知道，有些人，面前心连心，背后动脑筋。</p><p style=\"font-family: &quot;PingFang SC&quot;; font-size: 14px; white-space: normal;\">　　15、悬崖那么深，我终究是为了你跳了下去。</p>', 1583571992);

-- ----------------------------
-- Table structure for pinecms_setting
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_setting`;
CREATE TABLE `pinecms_setting`  (
  `key` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `value` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `group` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `default` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `form_name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `editor` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `listorder` int(11) NULL DEFAULT 0,
  PRIMARY KEY (`key`) USING BTREE,
  UNIQUE INDEX `UQE_iriscms_setting_key`(`key`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pinecms_setting
-- ----------------------------
INSERT INTO `pinecms_setting` VALUES ('EMAIL_USER', '', '邮箱设置', NULL, '用户名', 'text', 6);
INSERT INTO `pinecms_setting` VALUES ('EMAIL_PWD', '', '邮箱设置', NULL, '密码', 'text', 7);
INSERT INTO `pinecms_setting` VALUES ('EMAIL_SMTP', '', '邮箱设置', NULL, 'SMTP服务器', 'text', 8);
INSERT INTO `pinecms_setting` VALUES ('EMAIL_EMAIL', '', '邮箱设置', NULL, '邮箱地址', 'text', 9);
INSERT INTO `pinecms_setting` VALUES ('SITE_ICP', '', '前台设置', NULL, '备案号', 'text', 4);
INSERT INTO `pinecms_setting` VALUES ('SITE_KEYWORDS', 'pine,pincms,gocms,cms,模板化框架', '前台设置', NULL, '关键字', 'text', 2);
INSERT INTO `pinecms_setting` VALUES ('SITE_DESCRIPTION', 'pincms一个go语言的模板化CMS，支持类dedecms式的标签化调用。可以快速开发出企业网站，支持自定义文档模型', '前台设置', NULL, '描述', 'text', 3);
INSERT INTO `pinecms_setting` VALUES ('SITE_TITLE', 'pincms一个go语言的模板化CMS', '前台设置', NULL, '站点标题', 'text', 1);
INSERT INTO `pinecms_setting` VALUES ('WX_TOKEN', '', '微信配置', NULL, 'TOKEN', 'text', 13);
INSERT INTO `pinecms_setting` VALUES ('SITE_OPEN', '开启', '前台设置', '开启', '站点开启', '{\"type\":\"checkbox\", \"options\": {\"on\":\"开启\", \"off\":\"关闭\"}}', 0);
INSERT INTO `pinecms_setting` VALUES ('EMAIL_PORT', '25', '邮箱设置', '25', '端口', 'text', 10);
INSERT INTO `pinecms_setting` VALUES ('WX_APPSECRET', '', '微信配置', NULL, 'APPSECTET', 'text', 12);
INSERT INTO `pinecms_setting` VALUES ('WX_AESKEY', '', '微信配置', NULL, 'AESKEY', 'text', 14);
INSERT INTO `pinecms_setting` VALUES ('DATAGRID_PAGE_SIZE', '25', '前台设置', '25', '列表默认分页数', 'text', 5);
INSERT INTO `pinecms_setting` VALUES ('WX_APPID', '', '微信配置', NULL, 'APPID', 'text', 11);
INSERT INTO `pinecms_setting` VALUES ('UPLOAD_DIR', 'resources/assets/upload', '存储配置', 'resources/assets/upload', '存储目录', 'text', 21);
INSERT INTO `pinecms_setting` VALUES ('UPLOAD_ENGINE', '本地存储', '存储配置', 'local', '存储引擎', '{\"type\":\"combogrid\",\"options\":{\"idField\":\"key\",\"textField\":\"key\",\"fitColumns\":true,\"columns\":[[{\"field\":\"key\",\"title\":\"存储引擎\",\"width\":120}]],\"data\":[{\"key\":\"本地存储\",\"value\":\"local\"},{\"key\":\"OSS存储\",\"value\":\"oss\"}]}}', 20);
INSERT INTO `pinecms_setting` VALUES ('UPLOAD_IMG_TYPES', 'jpg,jpeg,png,gif,bmp', '存储配置', 'jpg,jpeg,png,gif,bmp', '可上传图片类型', 'text', 22);
INSERT INTO `pinecms_setting` VALUES ('UPLOAD_URL_PREFIX', '/upload', '存储配置', 'upload', '地址前缀', 'text', 23);
INSERT INTO `pinecms_setting` VALUES ('UPLOAD_DATABASE_PASS', '123456', '存储配置', '', '备份数据库密码', 'text', 24);

-- ----------------------------
-- Table structure for pinecms_wechat_member
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_wechat_member`;
CREATE TABLE `pinecms_wechat_member`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `mpid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sex` tinyint(4) NULL DEFAULT NULL,
  `headimgurl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `subscribe_scene` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pinecms_wechat_message_log
-- ----------------------------
DROP TABLE IF EXISTS `pinecms_wechat_message_log`;
CREATE TABLE `pinecms_wechat_message_log`  (
  `logid` bigint(20) NOT NULL DEFAULT 0,
  `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`logid`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '微信消息日志' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

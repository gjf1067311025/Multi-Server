/*
Navicat MySQL Data Transfer

Source Database       : starling

Target Server Type    : MYSQL
Target Server Version : 50639
File Encoding         : 65001

*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for starling_article
-- ----------------------------

DROP TABLE IF EXISTS `starling_basic_info`;
create table `starling_basic_info`(
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `create_time` TIMESTAMP NULL default NULL,
  `update_time` TIMESTAMP NULL default NULL,
  `key` VARCHAR(255) default '',
  `chinese_text` VARCHAR(255) default '',
  `english_text` VARCHAR(255) default '',
  `is_del`TINYINT(3) default 0 not null,
  `remark` VARCHAR(255) default '',
  `is_machine_translate` TINYINT(2) DEFAULT '1',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='中英文文案管理';
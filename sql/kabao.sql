/*
Navicat MySQL Data Transfer

Source Server         : kabao
Source Server Version : 50549
Source Host           : 172.98.201.182:3306
Source Database       : kabao

Target Server Type    : MYSQL
Target Server Version : 50549
File Encoding         : 65001

Date: 2016-05-30 22:45:09
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `categoryid` int(11) NOT NULL AUTO_INCREMENT,
  `categoryname` varchar(64) NOT NULL,
  `parentid` int(11) NOT NULL DEFAULT '-1' COMMENT '父分类ID',
  PRIMARY KEY (`categoryid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `commentid` int(11) NOT NULL,
  `cardid` int(11) NOT NULL,
  `userid` int(11) NOT NULL,
  `content` varchar(255) DEFAULT NULL,
  `createtime` datetime NOT NULL,
  `lastupdatetime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`commentid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for commoncard
-- ----------------------------
DROP TABLE IF EXISTS `commoncard`;
CREATE TABLE `commoncard` (
  `ccardid` int(11) NOT NULL,
  `ccardtitle` varchar(64) NOT NULL,
  `ccardtype` varchar(64) DEFAULT NULL,
  `usage` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0仅仅积分，1打折，2满减，3满赠,4红利价',
  `discount` tinyint(4) DEFAULT '100' COMMENT '折扣',
  `over` int(11) DEFAULT NULL COMMENT '满多少',
  `bonus` int(11) DEFAULT NULL COMMENT '满多少减多少，或者赠多少。',
  `vipprice` int(11) DEFAULT NULL,
  `cardbackgroud` varchar(255) DEFAULT NULL,
  `hit` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`ccardid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of commoncard
-- ----------------------------

-- ----------------------------
-- Table structure for token
-- ----------------------------
DROP TABLE IF EXISTS `token`;
CREATE TABLE `token` (
  `tokenid` int(11) NOT NULL AUTO_INCREMENT,
  `userid` int(11) NOT NULL,
  `token` varchar(64) NOT NULL,
  `type` tinyint(4) NOT NULL DEFAULT '0',
  `isactive` tinyint(4) NOT NULL DEFAULT '1',
  `expiretime` datetime NOT NULL,
  `createat` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`tokenid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of token
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `userid` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(32) DEFAULT NULL,
  `password` varchar(64) NOT NULL COMMENT '两次md5',
  `salt` varchar(32) NOT NULL COMMENT '加盐',
  `phone` char(11) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `sex` tinyint(4) NOT NULL DEFAULT '1' COMMENT '0女，1男',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像url',
  `realname` varchar(64) DEFAULT NULL COMMENT '真实姓名',
  `slogan` varchar(255) DEFAULT NULL COMMENT '个人签名',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0普通用户，1企业用户',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0正常，1被封',
  `createtime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `lastupdatetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------

-- ----------------------------
-- Table structure for vipcard
-- ----------------------------
DROP TABLE IF EXISTS `vipcard`;
CREATE TABLE `vipcard` (
  `cardid` int(11) NOT NULL AUTO_INCREMENT,
  `userid` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0会员卡，1优惠券',
  `cardtitle` varchar(64) DEFAULT NULL,
  `cardtype` varchar(255) DEFAULT NULL COMMENT '金卡，银卡等',
  `cardno` varchar(64) DEFAULT NULL,
  `cardcode` varchar(64) DEFAULT NULL COMMENT '1开头一维码，2开头二维码',
  `cardname` varchar(64) DEFAULT NULL COMMENT '持卡人姓名',
  `cardphone` char(16) DEFAULT NULL,
  `carddescription` varchar(255) DEFAULT NULL,
  `cardbackgroud` varchar(255) DEFAULT NULL COMMENT '背景图片url',
  `categoryid` int(11) DEFAULT NULL,
  `usage` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0仅仅积分，1打折，2满减，3满赠,4红利价',
  `discount` tinyint(4) DEFAULT '100' COMMENT '折扣',
  `over` int(11) DEFAULT NULL COMMENT '满多少',
  `bonus` int(11) DEFAULT NULL COMMENT '满多少减多少，或者赠多少。',
  `vipprice` int(11) DEFAULT NULL,
  `starttime` datetime DEFAULT NULL COMMENT '主要针对优惠券',
  `expiretime` datetime DEFAULT NULL COMMENT '过期,主要针对优惠券',
  `shared` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0未分享，1分享',
  `sharedname` varchar(64) DEFAULT NULL COMMENT '保护隐私，允许用户修改',
  `sharedphone` varchar(16) DEFAULT NULL COMMENT '保护隐私，允许用户修改',
  `like` int(11) NOT NULL DEFAULT '0' COMMENT '赞的数量',
  `dislike` int(11) NOT NULL DEFAULT '0' COMMENT '负面',
  `tags` varchar(255) DEFAULT NULL COMMENT '标签，网友评论关键字，以|隔开',
  `createtime` datetime NOT NULL,
  `lastupdatetime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`cardid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of vipcard
-- ----------------------------

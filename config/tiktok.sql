/*
 Navicat Premium Data Transfer

 Source Server         : panker
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : localhost:3306
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 29/01/2023 18:41:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `user_id` int NOT NULL COMMENT '评论者id',
  `video_id` int NOT NULL COMMENT '视频id',
  `action_type` int NOT NULL COMMENT '评论的状态（存在1，删除0）',
  `content` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '评论内容',
  `created_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of comments
-- ----------------------------

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int NOT NULL COMMENT '用户id',
  `follower_id` int NOT NULL COMMENT '关注的用户的id',
  `action_type` int NOT NULL COMMENT '关注的状态（关注1，取关0）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of follows
-- ----------------------------

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int NOT NULL COMMENT '点赞的用户的id',
  `video_id` int NOT NULL COMMENT '被点赞的视频的id',
  `action_type` int NOT NULL COMMENT '赞的状态（点赞1，取消0）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of likes
-- ----------------------------

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id主键',
  `from_user_id` int NOT NULL COMMENT '发送消息的用户id',
  `to_user_id` int NOT NULL COMMENT '接收消息的用户id',
  `content` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '消息内容',
  `create_time` datetime NOT NULL COMMENT '发送消息的时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of messages
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '密码',
  `salt` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '加密盐-生成密码用',
  `create_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (2, 'panker', '123456', 'kkk', '0000-00-00 00:00:00');
INSERT INTO `users` VALUES (3, 'test', 'test', '2333', '2023-01-29 11:27:14');
INSERT INTO `users` VALUES (4, 'test', '123456', '2333', '2023-01-29 12:00:55');
INSERT INTO `users` VALUES (5, 'yangsa', 'yangsa', '2333', '2023-01-29 12:35:10');

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `user_id` int NOT NULL COMMENT '作者id',
  `play_url` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '播放地址',
  `cover_url` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '封面地址',
  `title` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '视频标题',
  `created_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建实践',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of videos
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;

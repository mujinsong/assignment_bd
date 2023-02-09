/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : localhost:3306
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 09/02/2023 22:39:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `user_id` int(0) NOT NULL COMMENT '评论者id',
  `video_id` int(0) NOT NULL COMMENT '视频id',
  `action_type` int(0) NOT NULL COMMENT '评论的状态（存在1，删除0）',
  `content` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '评论内容',
  `created_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(0) NOT NULL COMMENT '用户id',
  `follower_id` int(0) NOT NULL COMMENT '关注的用户的id',
  `action_type` int(0) NOT NULL COMMENT '关注的状态（关注1，取关0）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(0) NOT NULL COMMENT '点赞的用户的id',
  `video_id` int(0) NOT NULL COMMENT '被点赞的视频的id',
  `action_type` int(0) NOT NULL COMMENT '赞的状态（点赞1，取消0）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of likes
-- ----------------------------
INSERT INTO `likes` VALUES (2, 1, 1, 1);
INSERT INTO `likes` VALUES (3, 1, 2, 1);
INSERT INTO `likes` VALUES (4, 1, 3, 1);
INSERT INTO `likes` VALUES (5, 1, 4, 1);
INSERT INTO `likes` VALUES (6, 2, 1, 1);
INSERT INTO `likes` VALUES (7, 3, 2, 1);
INSERT INTO `likes` VALUES (8, 3, 1, 1);

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT 'id主键',
  `from_user_id` int(0) NOT NULL COMMENT '发送消息的用户id',
  `to_user_id` int(0) NOT NULL COMMENT '接收消息的用户id',
  `content` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '消息内容',
  `create_at` datetime(0) NOT NULL COMMENT '发送消息的时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(32) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '密码',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户的昵称',
  `salt` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '加密盐-生成密码用',
  `create_at` datetime(0) NOT NULL COMMENT '创建时间',
  `follow_count` int unsigned NOT NULL COMMENT '用户的关注数',
  `follower_count` int unsigned NOT NULL COMMENT '用户的粉丝数',
  PRIMARY KEY (`id`, `username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '123456536', 'a938e25ba23662689d247f4e5d93ec72', '文艺的博比·摩尔', 'Q9QD5TxFla', '2023-02-09 22:31:48', 0, 0);
INSERT INTO `users` VALUES (2, '123456', 'a7d177d69df0a881bc6b7d3e931fc67d', '自然的乔布斯', 'Rly29QTChm', '2023-02-09 22:38:47', 0, 0);
INSERT INTO `users` VALUES (3, '12345', '7668ee8afbfea24752efb6e0b4dfec89', '列夫·雅辛掐指一算', 'Gx7rV2ft3Z', '2023-02-09 22:39:05', 0, 0);

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `user_id` int(0) NOT NULL COMMENT '作者id',
  `play_url` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '播放地址',
  `cover_url` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '封面地址',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频标题',
  `created_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `favorite_count` int unsigned NOT NULL COMMENT '喜欢数目',
  `comment_count` int unsigned NOT NULL COMMENT '评论数目',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, 3, 'http://81.68.91.70/video/1.mp4', 'http://81.68.91.70/image/1.jpg', '填充的视频,避免视频列表是空的', '2023-02-09 22:39:14', 3, 0);
INSERT INTO `videos` VALUES (2, 3, 'http://81.68.91.70/video/2.mp4', 'http://81.68.91.70/image/2.jpg', '测试视频,无封面', '2023-02-09 22:39:12', 2, 0);
INSERT INTO `videos` VALUES (3, 3, 'http://81.68.91.70/video/3.mp4', 'http://81.68.91.70/image/3.jpg', '测试视频流完结,视频为外链', '2023-02-09 22:31:58', 1, 0);
INSERT INTO `videos` VALUES (4, 1, 'http://192.168.1.4:8888/video/u9QHtMhxm3.mp4', 'http://192.168.1.4:8888/image/u9QHtMhxm3.jpg', '新的视频', '2023-02-09 22:32:48', 1, 0);

SET FOREIGN_KEY_CHECKS = 1;

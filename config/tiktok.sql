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

 Date: 12/02/2023 21:05:48
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
                             `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
                             `create_date` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, 2, 1, 1, '你好', '2023-02-12 20:16:50');
INSERT INTO `comments` VALUES (2, 3, 1, 1, '哈哈哈', '2023-02-12 20:17:14');
INSERT INTO `comments` VALUES (3, 3, 2, 1, '哈哈哈', '2023-02-12 20:17:18');

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
-- Records of follows
-- ----------------------------
INSERT INTO `follows` VALUES (3, 1, 4, 1);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of likes
-- ----------------------------
INSERT INTO `likes` VALUES (1, 2, 1, 2);
INSERT INTO `likes` VALUES (2, 2, 2, 1);
INSERT INTO `likes` VALUES (3, 2, 3, 1);
INSERT INTO `likes` VALUES (4, 3, 1, 1);
INSERT INTO `likes` VALUES (5, 3, 2, 1);
INSERT INTO `likes` VALUES (6, 4, 1, 1);
INSERT INTO `likes` VALUES (7, 4, 2, 1);

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
                          `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户的昵称',
                          `salt` varchar(10) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '加密盐-生成密码用',
                          `create_at` datetime(0) NOT NULL COMMENT '创建时间',
                          `follow_count` int unsigned NOT NULL COMMENT '用户的关注数',
                          `follower_count` int unsigned NOT NULL COMMENT '用户的粉丝数',
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '123', 'b06f38dbd9e602a4141239ee345db793', '超帅的科比', 'LAyFE5lWSx', '2023-02-12 20:13:31', 0, 1);
INSERT INTO `users` VALUES (2, '1234', '0fb340fc04e34496b9322d22c22b6061', '稳重的姆巴佩', 'ZyaLXXLdAz', '2023-02-12 20:16:38', 0, 0);
INSERT INTO `users` VALUES (3, '12345', 'd697399dde92930603042ece5ec09a9d', '肯尼横扫千军', 'ZwFEGB2Pan', '2023-02-12 20:17:03', 0, 0);
INSERT INTO `users` VALUES (4, '123456', '4aef24125c30dca39cbe50dd3b12ed14', '天神下凡的巴乔', 'rcIP2QV8EE', '2023-02-12 20:17:34', 1, 0);

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
                           `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '视频id',
                           `user_id` int(0) NOT NULL COMMENT '作者id',
                           `play_url` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '播放地址',
                           `cover_url` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '封面地址',
                           `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频标题',
                           `created_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
                           `favorite_count` int unsigned NOT NULL COMMENT '喜欢数目',
                           `comment_count` int unsigned NOT NULL COMMENT '评论数目',
                           PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, 1, 'http://81.68.91.70/video/1.mp4', 'http://81.68.91.70/image/1.jpg', '填充的视频,避免视频列表是空的', '2023-02-12 20:17:38', 2, 2);
INSERT INTO `videos` VALUES (2, 1, 'http://81.68.91.70/video/2.mp4', 'http://81.68.91.70/image/2.jpg', '测试视频,无封面', '2023-02-12 20:17:39', 3, 1);
INSERT INTO `videos` VALUES (3, 1, 'http://81.68.91.70/video/3.mp4', 'http://81.68.91.70/image/3.jpg', '测试视频流完结,视频为外链', '2023-02-12 20:16:45', 1, 0);

SET FOREIGN_KEY_CHECKS = 1;

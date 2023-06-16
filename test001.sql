/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : test001

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 16/06/2023 16:44:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int NOT NULL COMMENT '谁发的消息,关联到user表id',
  `category` tinyint NOT NULL DEFAULT 0 COMMENT '消息类型,10表示私聊,11表示群聊,,0是单聊',
  `target_id` int NOT NULL COMMENT '如果是好友就关联到user表的id,如果是群聊就关联到community表id',
  `media` tinyint NOT NULL DEFAULT 0 COMMENT '消息按照什么样式展示',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '消息内容体',
  `pic` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '图片预览地址',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '服务的url',
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `deleted_at` timestamp(6) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '消息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (1, 1, 1, 1, 1, '哈哈哈', '11', '11', '233', '2023-06-13 14:11:45.709000', '2023-06-13 14:11:45.709000', NULL);

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `is_del` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除,1表示删除,0表示正常',
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `order_sn` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `trade_no` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '支付交易号(微信支付的交易号)',
  `pay_method` tinyint NOT NULL DEFAULT 0 COMMENT '支付方式0是微信,1是支付宝',
  `user_id` int NOT NULL DEFAULT 0 COMMENT '交易的用户ID(小程序)',
  `activity_id` int NOT NULL DEFAULT 0 COMMENT '订单关联活动的ID',
  `activity_end_time` timestamp NULL DEFAULT NULL COMMENT '活动结束时间',
  `paid_amount` decimal(13, 2) NOT NULL DEFAULT 0.00 COMMENT '支付金额(总金额)',
  `start_paid_at` timestamp NULL DEFAULT NULL COMMENT '开始支付时间',
  `end_paid_at` timestamp NULL DEFAULT NULL COMMENT '支付成功时间',
  `status` tinyint NULL DEFAULT 0 COMMENT '支付状态,0表示未支付,1表示已支付,2表示支付失败,3表示取消订单,4表示退款,5表示核销,6表示在支付等待支付回调',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注信息',
  `cancel_order_reason` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '取消订单原因',
  `sponsor_id` int NOT NULL DEFAULT 0 COMMENT '订单关联主办方ID',
  `activity_type_id` int NOT NULL COMMENT '活动类型id',
  `ticket_total` int NULL DEFAULT NULL COMMENT '该订单下票数量',
  `deleted_at` timestamp(6) NULL DEFAULT NULL COMMENT '软删除时间',
  `activity_start_time` timestamp NULL DEFAULT NULL COMMENT '活动开始时间',
  `refund_status` int NULL DEFAULT NULL COMMENT '退款状态',
  `refund_ticket_total` int NULL DEFAULT 0 COMMENT '退票数量',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1368 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '用户名',
  `age` int UNSIGNED NOT NULL COMMENT '年龄',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '密码',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '软删除时间',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '手机号码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '张无忌', 25, '123456', '2023-06-13 08:23:23', '2023-06-13 08:25:55', '2023-06-13 08:25:55', '13888888888');
INSERT INTO `user` VALUES (2, '周芷若', 18, '654321', '2023-06-13 08:23:23', '2023-06-13 08:23:23', NULL, '13999999999');
INSERT INTO `user` VALUES (3, '赵敏', 23, 'abcabc', '2023-06-13 08:23:23', '2023-06-13 08:23:23', NULL, '15666666666');
INSERT INTO `user` VALUES (4, '谢逊', 30, 'qwer1234', '2023-06-13 08:27:42', '2023-06-13 08:27:42', NULL, '12345678901');
INSERT INTO `user` VALUES (6, '谢逊', 30, 'qwer1234', '2023-06-16 16:34:46', '2023-06-16 16:34:46', NULL, '12345678901');

-- ----------------------------
-- Table structure for user_token
-- ----------------------------
DROP TABLE IF EXISTS `user_token`;
CREATE TABLE `user_token`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int NOT NULL COMMENT '关联到用户表id',
  `token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'token',
  `expire_time` timestamp(6) NULL DEFAULT NULL COMMENT '失效时间',
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `deleted_at` timestamp(6) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户Token表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_token
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;

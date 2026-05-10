
SET NAMES utf8mb4;

-- ----------------------------
-- Table structure for pay_order
-- ----------------------------
DROP TABLE IF EXISTS `pay_order`;
CREATE TABLE `pay_order` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `app_id` bigint NOT NULL DEFAULT '0' COMMENT '应用ID',
  `channel_id` bigint NOT NULL DEFAULT '0' COMMENT '渠道ID',
  `channel_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '渠道编码 例如:alipay_pc/wx_lite',
  `merchant_order_id` bigint NOT NULL DEFAULT '0' COMMENT '商户订单ID 例如：trade_order.id',
  `order_group` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '组别 mall商城',
  `subject` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品标题',
  `body` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '商品描述',
  `notify_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付通知回调地址',
  `pay_type` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '支付类型 wxpay微信支付/alipay支付宝/saobei扫呗',
  `trade_type` varchar(16) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '交易类型 saobei(mini小程序) wxpay(mp公众号/mini小程序/appAPP/scan二维码扫码) alipay(web网页/appAPP/scan二维码扫码)\n',
  `price` int NOT NULL DEFAULT '0' COMMENT '支付金额 单位:分',
  `channel_fee_rate` double NOT NULL DEFAULT '0' COMMENT '渠道手续费率 单位:百分比',
  `channel_fee_price` int NOT NULL DEFAULT '0' COMMENT '渠道手续费金额 单位:分',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '支付状态 1未支付/10支付成功/20已退款/30支付关闭',
  `user_ip` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户IP',
  `expire_time` datetime DEFAULT NULL COMMENT '订单失效时间',
  `success_time` datetime DEFAULT NULL COMMENT '订单支付成功时间',
  `extension_id` bigint NOT NULL DEFAULT '0' COMMENT '订单拓展单编号 支付成功后写入',
  `no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付订单号 extension.no 支付成功后写入',
  `refund_price` int NOT NULL DEFAULT '0' COMMENT '退款总金额 单位：分',
  `channel_user_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '渠道用户编号',
  `channel_order_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '渠道订单号',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='支付订单';

-- ----------------------------
-- Records of pay_order
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pay_order_extension
-- ----------------------------
DROP TABLE IF EXISTS `pay_order_extension`;
CREATE TABLE `pay_order_extension` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '支付订单号',
  `pay_order_id` bigint NOT NULL COMMENT '支付订单编号',
  `channel_id` bigint NOT NULL DEFAULT '0' COMMENT '渠道编号',
  `channel_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '渠道编码',
  `user_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户 IP',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '支付状态 1未支付/10支付成功/20已退款/30支付关闭',
  `channel_extras` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '支付渠道的额外参数',
  `channel_error_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '渠道调用报错时，错误码',
  `channel_error_msg` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '渠道调用报错时，错误信息',
  `channel_notify_data` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '支付渠道异步通知的内容',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='支付订单拓展单\n';

-- ----------------------------
-- Records of pay_order_extension
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for pay_refund
-- ----------------------------
DROP TABLE IF EXISTS `pay_refund`;
CREATE TABLE `pay_refund` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '退款单号',
  `app_id` bigint NOT NULL DEFAULT '0' COMMENT '应用ID',
  `channel_id` bigint NOT NULL DEFAULT '0' COMMENT '渠道ID',
  `channel_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '渠道编码 例如:alipay_pc/wx_lite',
  `pay_order_id` bigint NOT NULL DEFAULT '0' COMMENT '支付订单ID',
  `pay_order_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '支付订单编号',
  `merchant_order_id` bigint NOT NULL DEFAULT '0' COMMENT '商户订单ID 例：商城订单ID',
  `merchant_refund_id` bigint NOT NULL DEFAULT '0' COMMENT '商户退款订单ID ',
  `notify_url` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '异步通知商户地址',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '退款状态 0未退款/10退款成功/20退款失败\n',
  `pay_price` int NOT NULL DEFAULT '0' COMMENT '支付金额 单位分',
  `refund_price` int NOT NULL DEFAULT '0' COMMENT '退款金额 单位分',
  `reason` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '申请退款原因',
  `user_ip` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户IP',
  `channel_order_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '渠道订单号，pay_order 中的 channel_order_no 对应',
  `channel_refund_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '渠道退款单号，渠道返回',
  `success_time` datetime DEFAULT NULL COMMENT '退款成功时间',
  `channel_error_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '渠道调用报错时，错误码',
  `channel_error_msg` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '渠道调用报错时，错误信息',
  `channel_notify_data` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付渠道异步通知的内容',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '退款备注',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='退款订单';

-- ----------------------------
-- Records of pay_refund
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `group` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置分组',
  `label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '中文名称 例如:姓名',
  `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置键名 例如:age',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '配置值 例如:18',
  `default_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认值',
  `type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '类型 string,int,uint,bool,datetime,date\n',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '1000' COMMENT '贴牌Id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_tenant_key` (`key`,`tenant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=341 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统配置表';


-- ----------------------------
-- Table structure for sys_cron
-- ----------------------------
DROP TABLE IF EXISTS `sys_cron`;
CREATE TABLE `sys_cron` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标题',
  `tag` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标签 字典表获取',
  `api_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '接口地址',
  `api_param` json DEFAULT NULL COMMENT '接口请求参数',
  `api_header` json DEFAULT NULL COMMENT '接口请求头',
  `pattern` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'cron表达式',
  `status` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '1' COMMENT '任务状态 active运行中/inactive已结束',
  `policy` int NOT NULL DEFAULT '1' COMMENT '策略 1并行/2单例/3单次/4多次',
  `count` int NOT NULL DEFAULT '0' COMMENT '执行次数 policy=4时有效',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `next_run_at` datetime DEFAULT NULL COMMENT '下次预计运行时间',
  `last_run_at` datetime DEFAULT NULL COMMENT '最后一次运行时间',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='定时任务';

-- ----------------------------
-- Records of sys_cron
-- ----------------------------
BEGIN;
INSERT INTO `sys_cron` (`id`, `title`, `tag`, `api_url`, `api_param`, `api_header`, `pattern`, `status`, `policy`, `count`, `sort`, `next_run_at`, `last_run_at`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (5, '删除超过7天的操作日志', '系统任务', 'http://127.0.0.1:8008/admin/operation/record/delete', '{\"day\": 1}', '{\"X-TenantId\": \"1000\", \"Authorization\": \"Bearer AmASj/kdW+Z1WFG3+zipUIehsJA1P2d7YJ1v1oSebV++eLNMNekW8sH9T8OALFS0\"}', '@every 1m', 'inactive', 2, 1, 1, '2024-09-18 09:54:13', '2024-09-18 09:53:13', '30分钟一次 */30 * * * * ', 1, '2024-05-21 17:31:29', 1, '2024-09-18 09:53:32');
COMMIT;

-- ----------------------------
-- Table structure for sys_cron_record
-- ----------------------------
DROP TABLE IF EXISTS `sys_cron_record`;
CREATE TABLE `sys_cron_record` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `cron_id` bigint NOT NULL DEFAULT '0' COMMENT '任务Id sys_cron.id',
  `spend_time` int NOT NULL DEFAULT '0' COMMENT '消耗时间, 单位毫秒',
  `output` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '执行结果或输出，可能包括错误信息',
  `status` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '执行状态 success成功/failure失败',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='定时任务日志';

-- ----------------------------
-- Records of sys_cron_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '上级Id 0是顶级',
  `dept_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '部门编码',
  `dept_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '部门状态 1正常/2停用',
  `level` int NOT NULL COMMENT '关系树层级',
  `tree` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '关系树',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `tenant_id` (`tenant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='部门管理';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_code`, `dept_name`, `status`, `level`, `tree`, `sort`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (56, 0, '', '公司', 1, 1, '', 1, '', 1, '2024-09-27 19:23:41', 137, '2024-10-11 11:44:44', 0);
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_code`, `dept_name`, `status`, `level`, `tree`, `sort`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (57, 56, '', '客服部', 1, 2, '56', 5, '', 1, '2024-09-27 19:23:53', 1, '2024-09-27 19:30:40', 0);
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_code`, `dept_name`, `status`, `level`, `tree`, `sort`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (58, 56, '', '采购部', 1, 2, '56', 0, '', 1, '2024-09-27 19:28:31', 0, '2024-09-27 19:28:31', 0);
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_code`, `dept_name`, `status`, `level`, `tree`, `sort`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (59, 57, '', '客服一部', 1, 3, '56,57', 0, '', 1, '2024-09-27 19:28:49', 138, '2024-10-08 15:59:44', 0);
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_code`, `dept_name`, `status`, `level`, `tree`, `sort`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (60, 59, '', '客服一部01组', 1, 4, '56,57,59', 0, '', 138, '2024-10-08 16:00:13', 0, '2024-10-08 16:00:13', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典id',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典标识',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典名称',
  `sort` int unsigned NOT NULL DEFAULT '1' COMMENT '排序号',
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典';

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict` (`id`, `code`, `name`, `sort`, `note`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (41, 'organizationType', '机构类型', 1, '', 1, '2026-03-21 18:25:28', 0, '2026-03-21 18:25:28');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '字典项id',
  `dict_id` bigint NOT NULL COMMENT '字典id',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典项标识',
  `code_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'string' COMMENT '类型 string/int',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典项名称',
  `sort` int NOT NULL DEFAULT '1' COMMENT '排序号',
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `dict_id` (`dict_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=142 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典项';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_data` (`id`, `dict_id`, `code`, `code_type`, `name`, `sort`, `note`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (141, 41, '1', 'int', '公司', 5, NULL, 1, '2026-03-21 18:25:58', 0, '2026-03-21 18:25:58');
COMMIT;

-- ----------------------------
-- Table structure for sys_email_record
-- ----------------------------
DROP TABLE IF EXISTS `sys_email_record`;
CREATE TABLE `sys_email_record` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮件标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '邮件内容',
  `receiver` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '收件邮箱',
  `sender` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '发件邮箱',
  `create_user_id` int DEFAULT NULL COMMENT '创建人',
  `note` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `tenant_id` int NOT NULL DEFAULT '1' COMMENT '租户id',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `create_user_id` (`create_user_id`) USING BTREE,
  KEY `tenant_id` (`tenant_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='邮件记录';

-- ----------------------------
-- Records of sys_email_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_file_record
-- ----------------------------
DROP TABLE IF EXISTS `sys_file_record`;
CREATE TABLE `sys_file_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '文件名称',
  `path` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '文件存储路径',
  `length` int NOT NULL DEFAULT '0' COMMENT '文件大小',
  `type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'img' COMMENT '文件类型：img图片/file文件',
  `note` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '创建人',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `create_user_id` (`created_by`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件列表';

-- ----------------------------
-- Records of sys_file_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_login_record
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_record`;
CREATE TABLE `sys_login_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `os` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  `device` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '设备名',
  `browser` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '浏览器类型',
  `ip` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'ip地址',
  `ip_city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'ip归属地',
  `login_type` int unsigned NOT NULL DEFAULT '0' COMMENT '操作类型 1登录成功/2登录失败/3退出登录',
  `remark` varchar(400) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='登录日志';

-- ----------------------------
-- Records of sys_login_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '菜单id',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '上级id  0是顶级',
  `title` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单路由地址',
  `component` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单组件地址, 目录可为空',
  `menu_type` int NOT NULL DEFAULT '0' COMMENT '类型 1菜单/2按钮',
  `sort` int NOT NULL DEFAULT '1' COMMENT '排序号',
  `authority` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `api_path` json DEFAULT NULL COMMENT '接口权限 默认存数组',
  `icon` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `hide` int NOT NULL DEFAULT '0' COMMENT '是否隐藏  2否/1是(仅注册路由不显示在左侧菜单)',
  `menu_meta` varchar(800) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '其它路由元信息',
  `tenant_id` bigint NOT NULL DEFAULT '1' COMMENT '租户id',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `tenant_id` (`tenant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=404 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='菜单';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (301, 0, '系统管理', '/system', '', 1, 1, '', NULL, 'IconProSettingOutlined', 2, '{\"props\": {\"badge\": {\"value\": \"New\", \"type\": \"warning\"}}, \"lang\": {\"zh_TW\": \"系統管理\", \"en\": \"System\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (302, 301, '用户管理', '/system/user', '/system/user', 1, 1, '', NULL, 'IconProUserOutlined', 2, '{\"lang\": {\"zh_TW\": \"用戶管理\", \"en\": \"User\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (303, 302, '查询用户', '', '', 2, 1, 'sys:user:page', '[\"get:/admin-api/user/page\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (304, 302, '添加用户', '', '', 2, 2, 'sys:user:save', '[\"post:/admin-api/user/create\"]', '', 2, NULL, 1, NULL, 0, NULL, 137, '2026-03-30 00:08:41');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (305, 302, '修改用户', '', '', 2, 3, 'sys:user:update', '[\"post:/admin-api/user/update\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (306, 302, '删除用户', '', '', 2, 4, 'sys:user:remove', '[\"post:/admin-api/user/delete/{Id}\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (307, 301, '角色管理', '/system/role', '/system/role', 1, 2, '', NULL, 'IconProIdcardOutlined', 2, '{\"lang\": {\"zh_TW\": \"角色管理\", \"en\": \"Role\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (308, 307, '查询角色', '', '', 2, 1, 'sys:role:list', '[\"get:/admin-api/role/page\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (309, 307, '添加角色', '', '', 2, 2, 'sys:role:save', NULL, '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (310, 307, '修改角色', '', '', 2, 3, 'sys:role:update', NULL, '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (311, 307, '删除角色', '', '', 2, 4, 'sys:role:remove', '[\"post:/admin-api/role/delete/{Id}\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (312, 301, '菜单管理', '/system/menu', '/system/menu', 1, 3, '', NULL, 'IconProAppstoreOutlined', 2, '{\"lang\": {\"zh_TW\": \"選單管理\", \"en\": \"Menu\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (313, 312, '查询菜单', '', '', 2, 1, 'sys:menu:list', NULL, '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (314, 312, '添加菜单', '', '', 2, 2, 'sys:menu:save', '[\"post:/admin-api/menu/create\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (315, 312, '修改菜单', '', '', 2, 3, 'sys:menu:update', '[\"post:/admin-api/menu/update\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (316, 312, '删除菜单', '', '', 2, 4, 'sys:menu:remove', '[\"post:/admin-api/menu/delete/{Id}\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (317, 301, '机构管理', '/system/organization', '/system/organization', 1, 4, '', NULL, 'IconProCityOutlined', 2, '{\"hideFooter\":true, \"lang\": {\"zh_TW\": \"機构管理\", \"en\": \"Organization\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (318, 317, '查询机构', '', '', 2, 1, 'sys:org:list', NULL, '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (319, 317, '添加机构', '', '', 2, 2, 'sys:org:save', '[\"post:/admin-api/organization/create\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (320, 317, '修改机构', '', '', 2, 3, 'sys:org:update', '[\"post:/admin-api/organization/update\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (321, 317, '删除机构', '', '', 2, 4, 'sys:org:remove', '[\"post:/admin-api/organization/delete/{Id}\"]', '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (322, 301, '字典管理', '/system/dictionary', '/system/dictionary', 1, 5, '', NULL, 'IconProBookOutlined', 2, '{\"hideFooter\":true, \"lang\": {\"zh_TW\": \"字典管理\", \"en\": \"Dictionary\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (327, 301, '登录日志', '/system/login-record', '/system/login-record', 1, 7, 'sys:login-record:list', '[\"get:/admin-api/login-record/page\"]', 'IconProCalendarOutlined', 2, '{\"lang\": {\"zh_TW\": \"登入日誌\", \"en\": \"LoginRecord\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (328, 301, '操作日志', '/system/operation-record', '/system/operation-record', 1, 8, 'sys:operation-record:list', '[\"get:/admin-api/operation-record/page\"]', 'IconProLogOutlined', 2, '{\"lang\": {\"zh_TW\": \"操作日誌\", \"en\": \"OperationRecord\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (329, 301, '文件管理', '/system/file', '/system/file', 1, 6, '', NULL, 'IconProFolderOutlined', 2, '{\"lang\": {\"zh_TW\": \"檔案管理\", \"en\": \"File\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (330, 329, '上传文件', '', '', 2, 1, 'sys:file:upload', NULL, '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (331, 329, '删除文件', '', '', 2, 2, 'sys:file:remove', NULL, '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (332, 329, '查看记录', '', '', 2, 3, 'sys:file:list', NULL, '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (333, 302, '用户详情', '/system/user/details', '/system/user/details', 1, 5, '', NULL, 'IconProUserOutlined', 1, '{\"active\": \"/system/user\", \"lang\": {\"zh_TW\": \"用戶詳情\", \"en\": \"UserDetails\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (334, 301, '修改个人密码', '', '', 2, 10, 'sys:auth:password', NULL, '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (335, 301, '修改个人资料', '', '', 2, 11, 'sys:auth:user', NULL, '', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (336, 0, 'Dashboard', '/dashboard', '', 1, 0, '', NULL, 'IconProHomeOutlined', 2, NULL, 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (337, 336, '工作台', '/dashboard/workplace', '/dashboard/workplace', 1, 1, '', NULL, 'IconProDesktopOutlined', 2, '{\"lang\": {\"zh_TW\": \"工作臺\", \"en\": \"Workplace\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (338, 336, '分析页', '/dashboard/analysis', '/dashboard/analysis', 1, 2, '', NULL, 'IconProAnalysisOutlined', 2, '{\"props\": {\"badge\": {\"value\": 1}}, \"lang\": {\"zh_TW\": \"分析頁\", \"en\": \"Analysis\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (339, 336, '监控页', '/dashboard/monitor', '/dashboard/monitor', 1, 3, '', NULL, 'IconProDashboardOutlined', 2, '{\"lang\": {\"zh_TW\": \"監控頁\", \"en\": \"Monitor\"}}', 1, NULL, 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (340, 0, '表单页面', '/form', '', 1, 2, '', NULL, 'IconProFormOutlined', 1, '{\"lang\":{\"en\":\"Form\",\"zh_TW\":\"表單頁面\"}}', 1, NULL, 0, NULL, 1, '2026-03-23 12:04:19');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (341, 340, '基础表单', '/form/basic', '/form/basic', 1, 1, '', NULL, 'IconProLinkOutlined', 1, '{\"lang\": {\"zh_TW\": \"基礎表單\", \"en\": \"Basic Form\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:04:19');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (342, 340, '复杂表单', '/form/advanced', '/form/advanced', 1, 2, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"複雜表單\", \"en\": \"Advanced Form\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:04:19');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (343, 340, '分步表单', '/form/step', '/form/step', 1, 3, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"分步表單\", \"en\": \"Step Form\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:04:19');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (344, 0, '列表页面', '/list', '', 1, 3, '', NULL, 'TableOutlined', 1, '{\"lang\":{\"en\":\"List\",\"zh_TW\":\"清單頁面\"},\"props\":{\"iconStyle\":{\"transform\":\"scale(0.88)\"}}}', 1, NULL, 0, NULL, 1, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (345, 344, '基础列表', '/list/basic', '/list/basic', 1, 1, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"基礎清單\", \"en\": \"Basic List\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (346, 344, '复杂列表', '/list/advanced', '/list/advanced', 1, 2, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"複雜清單\", \"en\": \"Advanced List\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (347, 344, '卡片列表', '/list/card', '/list/card', 1, 3, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"卡片清單\", \"en\": \"Card List\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (348, 347, '项目列表', '/list/card/project', '/list/card/project', 1, 1, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"項目清單\", \"en\": \"Project\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (349, 347, '应用列表', '/list/card/application', '/list/card/application', 1, 2, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"應用清單\", \"en\": \"Application\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (350, 347, '文章列表', '/list/card/article', '/list/card/article', 1, 3, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"文章清單\", \"en\": \"Article\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (351, 345, '添加用户', '/list/basic/add', '/list/basic/add', 1, 4, '', NULL, 'Link', 1, '{\"active\": \"/list/basic\", \"lang\": {\"zh_TW\": \"添加用戶\", \"en\": \"Add User\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (352, 345, '修改用户', '/list/basic/edit', '/list/basic/edit', 1, 4, '', NULL, 'Link', 1, '{\"active\": \"/list/basic\", \"lang\": {\"zh_TW\": \"編輯用戶\", \"en\": \"Edit User\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (353, 345, '用户详情', '/list/basic/details/:id', '/list/basic/details', 1, 4, '', NULL, 'Link', 1, '{\"active\": \"/list/basic\"}', 1, NULL, 0, NULL, 0, '2026-03-23 12:02:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (354, 0, '结果页面', '/result', '', 1, 4, '', NULL, 'CircleCheck', 1, '{\"lang\":{\"en\":\"Result\",\"zh_TW\":\"結果頁面\"}}', 1, NULL, 0, NULL, 1, '2026-03-23 12:03:18');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (355, 354, '成功页', '/result/success', '/result/success', 1, 1, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"成功頁\", \"en\": \"Success\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:03:18');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (356, 354, '失败页', '/result/fail', '/result/fail', 1, 2, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"失敗頁\", \"en\": \"Fail\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:03:18');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (357, 0, '异常页面', '/exception', '', 1, 5, '', NULL, 'WarningOutlined', 1, '{\"lang\":{\"en\":\"Exception\",\"zh_TW\":\"异常頁面\"}}', 1, NULL, 0, NULL, 1, '2026-03-23 12:03:35');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (358, 357, '403', '/exception/403', '/exception/403', 1, 1, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:35');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (359, 357, '404', '/exception/404', '/exception/404', 1, 2, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:35');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (360, 357, '500', '/exception/500', '/exception/500', 1, 3, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:35');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (361, 0, '个人中心', '/user', '', 1, 6, '', NULL, 'SetUp', 1, '{\"lang\":{\"en\":\"User\",\"zh_TW\":\"個人中心\"}}', 1, NULL, 0, NULL, 1, '2026-03-23 12:01:42');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (362, 361, '我的资料', '/user/profile', '/user/profile', 1, 1, '', NULL, 'User', 1, '{\"lang\": {\"zh_TW\": \"個人資料\", \"en\": \"Profile\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:01:42');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (363, 361, '我的消息', '/user/message', '/user/message', 1, 2, '', NULL, 'ChatDotSquare', 1, '{\"lang\": {\"zh_TW\": \"我的消息\", \"en\": \"Message\"}}', 1, NULL, 0, NULL, 0, '2026-03-23 12:01:42');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (364, 0, '扩展组件', '/extension', '', 1, 7, '', NULL, 'AppstoreAddOutlined', 1, '{\"props\":{\"iconStyle\":{\"transform\":\"scale(0.92)\"}}}', 1, NULL, 0, NULL, 1, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (365, 364, '标签输入', '/extension/tag', '/extension/tag', 1, 4, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (366, 364, '高级弹窗', '/extension/modal', '/extension/modal', 1, 5, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (367, 364, '文件列表', '/extension/file', '/extension/file', 1, 6, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (368, 364, '图片上传', '/extension/upload', '/extension/upload', 1, 7, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (369, 364, '拖拽排序', '/extension/dragsort', '/extension/dragsort', 1, 24, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (370, 364, '消息提示', '/extension/message', '/extension/message', 1, 1, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (371, 364, '城市选择', '/extension/regions', '/extension/regions', 1, 26, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (372, 364, '打印组件', '/extension/printer', '/extension/printer', 1, 12, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (373, 364, '导入导出', '/extension/excel', '/extension/excel', 1, 27, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (374, 364, '滚动数字', '/extension/count-up', '/extension/count-up', 1, 23, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (375, 364, '状态圆点', '/extension/dot', '/extension/dot', 1, 3, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (376, 364, '表格下拉', '/extension/table-select', '/extension/table-select', 1, 11, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (377, 364, '分割面板', '/extension/split', '/extension/split', 1, 8, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (378, 364, '视频播放', '/extension/player', '/extension/player', 1, 28, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (379, 364, '地图组件', '/extension/map', '/extension/map', 1, 25, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (380, 364, '二维码', '/extension/qr-code', '/extension/qr-code', 1, 20, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (381, 364, '条形码', '/extension/bar-code', '/extension/bar-code', 1, 19, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (382, 364, '富文本框', '/extension/editor', '/extension/editor', 1, 29, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (383, 364, 'markdown', '/extension/markdown', '/extension/markdown', 1, 30, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (384, 364, '头像组合', '/extension/avatar', '/extension/avatar', 1, 2, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (385, 364, '图标选择', '/extension/icon', '/extension/icon', 1, 9, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (386, 364, '文本省略', '/extension/text', '/extension/text', 1, 13, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (387, 364, '高级表格', '/extension/table', '/extension/table', 1, 10, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (388, 364, '仪表盘', '/extension/circle-wave', '/extension/circle-wave', 1, 21, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (389, 364, '标签页', '/extension/tabs', '/extension/tabs', 1, 17, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (390, 364, '步骤条', '/extension/steps', '/extension/steps', 1, 16, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (391, 364, '导航菜单', '/extension/menu', '/extension/menu', 1, 15, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (392, 364, '水印组件', '/extension/watermark', '/extension/watermark', 1, 22, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (393, 364, '引导组件', '/extension/tour', '/extension/tour', 1, 14, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (394, 364, '分段器', '/extension/segmented', '/extension/segmented', 1, 18, '', NULL, 'Link', 1, NULL, 1, NULL, 0, NULL, 0, '2026-03-23 12:03:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (395, 0, '内嵌页面', '/iframe', '', 1, 8, '', NULL, 'Link', 1, '{\"lang\":{\"en\":\"IFrame\",\"zh_TW\":\"內嵌頁面\"}}', 1, '2026-05-04 23:04:17', 0, NULL, 1, '2026-03-23 12:04:04');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (396, 395, '官网', '/iframe/eleadmin', 'https://www.eleadmin.com', 1, 1, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"官網\", \"en\": \"Website\"}}', 1, '2026-05-04 23:04:12', 0, NULL, 0, '2026-03-23 12:04:04');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (397, 395, '文档', '/iframe/eleadmin-doc', 'https://www.eleadmin.com/doc/eleadminpro/', 1, 2, '', NULL, 'Link', 1, '{\"lang\": {\"zh_TW\": \"檔案\", \"en\": \"Document\"}}', 1, '2026-05-04 23:04:10', 0, NULL, 0, '2026-03-23 12:04:04');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (398, 0, '功能演示', '/example', '/example', 1, 9, '', NULL, 'IconProAnalysisOutlined', 2, '{\"lang\":{\"en\":\"Demo\",\"zh_TW\":\"功能演示\"}}', 1, '2026-03-23 11:33:35', 0, NULL, 1, '2026-03-16 15:24:55');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (399, 0, '获取授权', 'https://eleadmin.com/goods/11', '', 1, 10, '', NULL, 'ProtectOutlined', 2, '{\"props\": {\"iconStyle\": {\"transform\": \"scale(0.88)\"}}, \"lang\": {\"zh_TW\": \"獲取授權\", \"en\": \"Authorization\"}}', 1, '2026-03-13 16:24:48', 0, NULL, 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (400, 0, '一级目录', '/demo', '', 1, 100, '', NULL, 'IconProAppstoreOutlined', 2, NULL, 1, NULL, 1, '2026-03-16 14:33:46', 1, '2026-03-16 15:25:47');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (401, 400, '二级菜单', '/demo/reed', '/demo/reed', 1, 5, '', NULL, 'IconProApplicationOutlined', 2, NULL, 1, NULL, 1, '2026-03-16 15:20:25', 1, '2026-03-16 15:24:36');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (402, 0, '测试', '/test', '', 1, 66, '', '[]', '', 2, NULL, 1, '2026-05-04 18:14:49', 1, '2026-05-04 18:14:44', 0, '2026-05-04 18:14:44');
INSERT INTO `sys_menu` (`id`, `parent_id`, `title`, `path`, `component`, `menu_type`, `sort`, `authority`, `api_path`, `icon`, `hide`, `menu_meta`, `tenant_id`, `deleted_at`, `created_by`, `created_at`, `updated_by`, `updated_at`) VALUES (403, 0, '1', '/11', '', 1, 1, '', '[]', '', 2, NULL, 1, '2026-05-04 23:04:07', 142, '2026-05-04 23:04:02', 0, '2026-05-04 23:04:02');
COMMIT;

-- ----------------------------
-- Table structure for sys_operation_record
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_record`;
CREATE TABLE `sys_operation_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '账号',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '请求地址',
  `method` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '请求方式',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '接口模块 例如：用户管理',
  `summary` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '接口描述 例如：添加用户',
  `param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求参数',
  `json_result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '返回结果',
  `error_msg` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '异常信息',
  `spend_time` bigint NOT NULL DEFAULT '0' COMMENT '消耗时间, 单位毫秒',
  `trace_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'trace_id',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态 1成功/2异常',
  `platform` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '平台 admin/api/open',
  `user_agent` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '请求头User-Agent',
  `ip` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '主机地址',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=738 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='操作日志';

-- ----------------------------
-- Records of sys_operation_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_organization
-- ----------------------------
DROP TABLE IF EXISTS `sys_organization`;
CREATE TABLE `sys_organization` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '上级id, 0是顶级',
  `code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '机构代码',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '机构名称',
  `full_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '机构全称',
  `type` int NOT NULL DEFAULT '0' COMMENT '机构类型',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '部门状态 1正常/2停用',
  `link_id` bigint NOT NULL DEFAULT '0' COMMENT '负责人id sys_user.id',
  `link_man` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `link_phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `level` int NOT NULL COMMENT '关系树层级',
  `tree` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '关系树',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `leader_id` (`link_id`) USING BTREE,
  KEY `tenant_id` (`tenant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='组织机构';

-- ----------------------------
-- Records of sys_organization
-- ----------------------------
BEGIN;
INSERT INTO `sys_organization` (`id`, `parent_id`, `code`, `name`, `full_name`, `type`, `status`, `link_id`, `link_man`, `link_phone`, `level`, `tree`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (55, 0, '260322723797', '总公司', '总公司', 1, 1, 0, '', '', 0, '', '', 1, '2026-03-22 22:27:22', 1, '2026-05-04 18:16:01', 0);
INSERT INTO `sys_organization` (`id`, `parent_id`, `code`, `name`, `full_name`, `type`, `status`, `link_id`, `link_man`, `link_phone`, `level`, `tree`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (56, 55, '260504282544', '一级分公司', '一级分公司', 1, 1, 0, '', '', 1, 'tr_55 ', '', 142, '2026-05-04 23:54:49', 0, '2026-05-04 23:54:49', 0);
INSERT INTO `sys_organization` (`id`, `parent_id`, `code`, `name`, `full_name`, `type`, `status`, `link_id`, `link_man`, `link_phone`, `level`, `tree`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (63, 56, '260504131532', '二级分公司', '二级分公司', 1, 1, 0, '', '', 2, 'tr_55 tr_56 ', '', 142, '2026-05-04 23:55:15', 142, '2026-05-05 00:08:55', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '上级角色ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '角色标识',
  `level` int NOT NULL DEFAULT '0' COMMENT '关系树等级',
  `tree` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '关系树',
  `data_scope` tinyint NOT NULL DEFAULT '1' COMMENT '数据范围 1全部/2当前部门/3当前及以下部门/4自定义部门',
  `custom_dept` json DEFAULT NULL COMMENT '自定义部门权限',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '角色状态 1正常/2停用',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`id`, `parent_id`, `name`, `code`, `level`, `tree`, `data_scope`, `custom_dept`, `status`, `sort`, `remark`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (1, 0, '管理员', '', 0, '', 1, NULL, 1, 0, '普通管理员，超管可以不使用角色，user.is_admin字段来设置', 1, '2026-05-04 18:13:52', 1, '2026-05-04 18:13:57', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `role_id` bigint NOT NULL DEFAULT '0' COMMENT '角色id',
  `menu_id` bigint NOT NULL DEFAULT '0' COMMENT '菜单id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_menu_id` (`role_id`) USING BTREE,
  KEY `menu_id` (`menu_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色权限';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (1, 1, 336, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (2, 1, 337, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (3, 1, 338, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (4, 1, 339, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (5, 1, 301, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (6, 1, 302, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (7, 1, 303, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (8, 1, 304, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (9, 1, 305, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (10, 1, 306, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (11, 1, 307, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (12, 1, 308, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (13, 1, 309, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (14, 1, 310, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (15, 1, 311, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (16, 1, 312, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (17, 1, 313, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (18, 1, 314, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (19, 1, 315, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (20, 1, 316, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (21, 1, 317, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (22, 1, 318, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (23, 1, 319, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (24, 1, 320, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (25, 1, 321, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (26, 1, 322, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (27, 1, 323, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (28, 1, 324, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (29, 1, 325, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (30, 1, 326, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (31, 1, 329, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (32, 1, 330, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (33, 1, 331, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (34, 1, 332, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (35, 1, 327, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (36, 1, 328, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (37, 1, 334, '2026-05-04 18:45:26', 0);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `tenant_id`) VALUES (38, 1, 335, '2026-05-04 18:45:26', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenant`;
CREATE TABLE `sys_tenant` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tenant_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '租户名称',
  `link_id` bigint NOT NULL DEFAULT '0' COMMENT '联系人ID sys_user.id',
  `link_man` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人姓名',
  `link_phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人手机',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '租户状态 1正常/2停用',
  `website` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '绑定域名',
  `package_id` bigint NOT NULL COMMENT '租户套餐ID',
  `expire_at` datetime DEFAULT NULL COMMENT '过期时间',
  `account_count` int NOT NULL DEFAULT '0' COMMENT '账号数量',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='租户表';

-- ----------------------------
-- Records of sys_tenant
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_tenant_package
-- ----------------------------
DROP TABLE IF EXISTS `sys_tenant_package`;
CREATE TABLE `sys_tenant_package` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '套餐编号',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '套餐名',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '套餐状态 1正常/2停用',
  `menu_ids` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '菜单ID 逗号分隔',
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='租户套餐表';

-- ----------------------------
-- Records of sys_tenant_package
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `organization_id` bigint NOT NULL DEFAULT '0' COMMENT '机构Id',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '盐加密',
  `role_id` bigint NOT NULL DEFAULT '0' COMMENT '角色Id',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `email_verified` int unsigned NOT NULL DEFAULT '2' COMMENT '邮箱是否验证:  1是 2否',
  `real_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `id_card` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '身份证号',
  `birthday` date DEFAULT NULL COMMENT '出生日期',
  `introduction` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '个人简介',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1正常/2冻结',
  `last_login_at` datetime DEFAULT NULL COMMENT '最新一次登录时间',
  `parent_id` bigint NOT NULL DEFAULT '0' COMMENT '上级用户Id',
  `level` int NOT NULL DEFAULT '0' COMMENT '关系树等级',
  `tree` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '关系树',
  `is_admin` tinyint NOT NULL DEFAULT '2' COMMENT '是否超级管理员 1是/2否 默认：2',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=144 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`id`, `organization_id`, `nickname`, `username`, `password`, `salt`, `role_id`, `phone`, `avatar`, `email`, `email_verified`, `real_name`, `id_card`, `birthday`, `introduction`, `status`, `last_login_at`, `parent_id`, `level`, `tree`, `is_admin`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (1, 55, '超级管理员', 'admin', '28e3e6eabeb9bd70f7a0891ee50f6947', 'koRegu', 1, '', 'http://localhost:8002/avatar/liyuanfang.jpg', '', 0, '', '', '2021-05-19', '遗其欲，则心静！', 1, '2025-01-31 19:31:55', 0, 0, '', 1, 0, NULL, 1, '2025-01-31 19:31:55', 0);
INSERT INTO `sys_user` (`id`, `organization_id`, `nickname`, `username`, `password`, `salt`, `role_id`, `phone`, `avatar`, `email`, `email_verified`, `real_name`, `id_card`, `birthday`, `introduction`, `status`, `last_login_at`, `parent_id`, `level`, `tree`, `is_admin`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (142, 55, '江渚之上', 'reed01', 'c160479daf7127002fa67754cad48d14', 'qJhiRx', 1, '15256569006', '', '', 0, '', '', NULL, '', 1, NULL, 1, 1, 'tr_1 ', 0, 1, '2026-05-04 18:18:11', 142, '2026-05-04 19:11:49', 0);
INSERT INTO `sys_user` (`id`, `organization_id`, `nickname`, `username`, `password`, `salt`, `role_id`, `phone`, `avatar`, `email`, `email_verified`, `real_name`, `id_card`, `birthday`, `introduction`, `status`, `last_login_at`, `parent_id`, `level`, `tree`, `is_admin`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES (143, 56, 'tony', 'tony01', '36b88031bd323cbdae05a9f321bd6e00', 'ZrGeri', 1, '15256569006', '', '', 0, '', '', NULL, '', 1, NULL, 142, 2, 'tr_1 tr_142 ', 0, 142, '2026-05-04 22:40:04', 0, '2026-05-04 22:40:04', 0);
COMMIT;


SET NAMES utf8mb4;

-- ----------------------------
-- Table structure for sys_config_module
-- ----------------------------
DROP TABLE IF EXISTS `sys_config_module`;
CREATE TABLE `sys_config_module` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块编码，代码中使用的唯一标识',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块名称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块说明',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序值，越小越靠前',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1启用/2禁用',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_tenant_code` (`tenant_id`, `code`) USING BTREE,
  KEY `idx_tenant_status_sort` (`tenant_id`, `status`, `sort`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='配置模块分组';

-- ----------------------------
-- Table structure for sys_config_item
-- ----------------------------
DROP TABLE IF EXISTS `sys_config_item`;
CREATE TABLE `sys_config_item` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `module_id` bigint NOT NULL DEFAULT '0' COMMENT '所属配置模块ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配置项名称，后台展示用',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配置项键名，代码读取用唯一标识',
  `config_value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '配置项当前值',
  `default_value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '默认值，读取失败或未配置时兜底',
  `value_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'string' COMMENT '值类型 string/int/uint/bool/json/datetime/date',
  `input_type` tinyint NOT NULL DEFAULT '1' COMMENT '输入类型 1输入框/2范围/3下拉/4单选/5开关/6多选',
  `input_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '控件参数，选项或范围配置，例如 A-1|B-2|C-3',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '配置项说明',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序值，越小越靠前',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1启用/2禁用',
  `is_system` tinyint NOT NULL DEFAULT '2' COMMENT '是否系统内置 1是/2否，内置配置通常由研发维护',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_tenant_config_key` (`tenant_id`, `config_key`) USING BTREE,
  KEY `idx_tenant_module_status_sort` (`tenant_id`, `module_id`, `status`, `sort`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统配置项';

-- ----------------------------
-- Records of sys_config_module
-- ----------------------------
INSERT INTO `sys_config_module` (`id`, `code`, `name`, `description`, `sort`, `status`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES
(1, 'basic', '基础配置', '系统基础展示配置', 1, 2, 0, NULL, 0, NULL, 0),
(2, 'upload', '上传配置', '文件上传与对象存储配置', 2, 2, 0, NULL, 0, NULL, 0),
(3, 'pay', '支付配置', '支付渠道相关配置', 3, 2, 0, NULL, 0, NULL, 0);

-- ----------------------------
-- Records of sys_config_item
-- ----------------------------
INSERT INTO `sys_config_item` (`id`, `module_id`, `name`, `config_key`, `config_value`, `default_value`, `value_type`, `input_type`, `input_params`, `description`, `sort`, `status`, `is_system`, `created_by`, `created_at`, `updated_by`, `updated_at`, `tenant_id`) VALUES
(1, 1, '租户名称', 'tenant_name', '', '', 'string', 1, '', '前端展示的租户名称', 1, 2, 2, 0, NULL, 0, NULL, 0),
(2, 1, '租户 Logo', 'tenant_logo', '', '', 'string', 1, '', '前端展示的租户 Logo 地址', 2, 2, 2, 0, NULL, 0, NULL, 0),
(3, 1, 'ICP备案号', 'icp_code', '', '', 'string', 1, '', '网站ICP备案号', 3, 2, 2, 0, NULL, 0, NULL, 0),
(4, 1, '版权信息', 'copyright', '', '', 'string', 1, '', '网站底部版权信息', 4, 2, 2, 0, NULL, 0, NULL, 0),
(5, 1, '应用 ID', 'appid', '', '', 'string', 1, '', '业务应用 ID', 5, 2, 2, 0, NULL, 0, NULL, 0),
(6, 1, '首页轮播图一', 'home_cover1', '', '', 'string', 1, '', '首页轮播图地址', 6, 2, 2, 0, NULL, 0, NULL, 0),
(7, 1, '首页轮播图二', 'home_cover2', '', '', 'string', 1, '', '首页轮播图地址', 7, 2, 2, 0, NULL, 0, NULL, 0),
(8, 1, '首页轮播图三', 'home_cover3', '', '', 'string', 1, '', '首页轮播图地址', 8, 2, 2, 0, NULL, 0, NULL, 0),
(9, 2, '上传驱动', 'uploadDrive', 'local', 'local', 'string', 3, '本地-local|OSS-oss', '上传驱动，local 或 oss', 1, 2, 2, 0, NULL, 0, NULL, 0),
(10, 2, 'OSS Endpoint', 'uploadOssEndpoint', '', '', 'string', 1, '', '对象存储 Endpoint', 2, 2, 2, 0, NULL, 0, NULL, 0),
(11, 2, 'OSS AccessKeyId', 'uploadOssAccessKeyId', '', '', 'string', 1, '', '对象存储 AccessKeyId', 3, 2, 2, 0, NULL, 0, NULL, 0),
(12, 2, 'OSS AccessKeySecret', 'uploadOssAccessKeySecret', '', '', 'string', 1, '', '对象存储 AccessKeySecret', 4, 2, 2, 0, NULL, 0, NULL, 0),
(13, 2, 'OSS Bucket', 'uploadOssBucket', '', '', 'string', 1, '', '对象存储 Bucket', 5, 2, 2, 0, NULL, 0, NULL, 0),
(14, 2, 'OSS 访问域名', 'uploadOssBucketUrl', '', '', 'string', 1, '', '对象存储访问域名', 6, 2, 2, 0, NULL, 0, NULL, 0),
(15, 3, '支付方式', 'payMethod', '', '', 'string', 3, '微信支付-wechat|扫呗支付-saobei', '默认支付方式', 1, 2, 2, 0, NULL, 0, NULL, 0),
(16, 3, '微信 AppId', 'payWxPayAppId', '', '', 'string', 1, '', '微信支付 AppId', 2, 2, 2, 0, NULL, 0, NULL, 0),
(17, 3, '微信商户号', 'payWxPayMchId', '', '', 'string', 1, '', '微信支付商户号', 3, 2, 2, 0, NULL, 0, NULL, 0),
(18, 3, '微信证书序列号', 'payWxPaySerialNo', '', '', 'string', 1, '', '微信支付证书序列号', 4, 2, 2, 0, NULL, 0, NULL, 0),
(19, 3, '微信 APIv3 Key', 'payWxPayAPIv3Key', '', '', 'string', 1, '', '微信支付 APIv3 Key', 5, 2, 2, 0, NULL, 0, NULL, 0),
(20, 3, '微信支付私钥', 'payWxPayPrivateKey', '', '', 'string', 1, '', '微信支付私钥', 6, 2, 2, 0, NULL, 0, NULL, 0),
(21, 3, '微信 JSAPI 地址', 'payWxPayJsApiUrl', '', '', 'string', 1, '', '微信支付 JSAPI 地址', 7, 2, 2, 0, NULL, 0, NULL, 0),
(22, 3, '扫呗商户号', 'paySaobeiMerchantNo', '', '', 'string', 1, '', '扫呗支付商户号', 8, 2, 2, 0, NULL, 0, NULL, 0),
(23, 3, '扫呗终端号', 'paySaobeiTerminalId', '', '', 'string', 1, '', '扫呗支付终端号', 9, 2, 2, 0, NULL, 0, NULL, 0),
(24, 3, '扫呗访问令牌', 'paySaobeiAccessToken', '', '', 'string', 1, '', '扫呗支付访问令牌', 10, 2, 2, 0, NULL, 0, NULL, 0);

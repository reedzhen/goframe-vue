SET NAMES utf8mb4;

-- ----------------------------
-- Table structure for sys_config_module
-- ----------------------------
CREATE TABLE IF NOT EXISTS `sys_config_module` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块编码，代码中使用的唯一标识',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块名称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '模块说明',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序值，越小越靠前',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1启用/0禁用',
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
CREATE TABLE IF NOT EXISTS `sys_config_item` (
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
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1启用/0禁用',
  `is_system` tinyint NOT NULL DEFAULT '0' COMMENT '是否系统内置 1是/0否，内置配置通常由研发维护',
  `created_by` bigint NOT NULL DEFAULT '0' COMMENT '添加人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` bigint NOT NULL DEFAULT '0' COMMENT '更新人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `tenant_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_tenant_config_key` (`tenant_id`, `config_key`) USING BTREE,
  KEY `idx_tenant_module_status_sort` (`tenant_id`, `module_id`, `status`, `sort`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统配置项';

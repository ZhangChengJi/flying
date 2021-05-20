
# Create Database
# ------------------------------------------------------------
CREATE DATABASE IF NOT EXISTS flying_config DEFAULT CHARACTER SET = utf8mb4;
Use flying_config;
#----------------flying-config------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
                       `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                       `app_id` varchar(60) DEFAULT NULL COMMENT 'AppId',
                       `name` varchar(255) DEFAULT NULL COMMENT '应用名称',
                       `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                       `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app表';



DROP TABLE IF EXISTS `namespace`;
CREATE TABLE `namespace` (
                             `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                             `a_id` bigint DEFAULT NULL COMMENT '应用ID',
                             `app_id` varchar(60) DEFAULT NULL COMMENT 'appId',
                             `format` varchar(60) DEFAULT NULL COMMENT '格式',
                             `node_id` smallint DEFAULT NULL COMMENT '环境ID',
                             `name` varchar(255) DEFAULT NULL COMMENT '命名空间名称',
                             `value` longtext COMMENT '配置信息',
                             `status` tinyint(1) DEFAULT NULL COMMENT '发布状态',
                             `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                             `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                             `release_time` datetime DEFAULT NULL COMMENT '发布时间',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=109 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='namespace表';


DROP TABLE IF EXISTS `release`;
CREATE TABLE `release` (
                           `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                           `app_id` varchar(500) DEFAULT NULL COMMENT '应用ID',
                           `comment` varchar(255) DEFAULT NULL COMMENT '发布说明',
                           `name` varchar(64) DEFAULT NULL COMMENT '发布名称',
                           `namespace_name` varchar(64) DEFAULT NULL COMMENT '命名空间',
                           `release_key` varchar(64) DEFAULT NULL COMMENT '发布key',
                           `format` varchar(60) DEFAULT NULL COMMENT '配置格式',
                           `value` longtext COMMENT '配置信息',
                           `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='release表';

# Create Database
# ------------------------------------------------------------
CREATE DATABASE IF NOT EXISTS flying_admin DEFAULT CHARACTER SET = utf8mb4;
Use flying_admin;
# ------------------------------flying-admin------------------

DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
                       `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                       `app_id` varchar(60) DEFAULT NULL COMMENT 'AppId',
                       `name` varchar(255) DEFAULT NULL COMMENT '应用名称',
                       `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                       `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app项目表';


DROP TABLE IF EXISTS `app_node`;
CREATE TABLE `app_node` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                            `app_id` bigint DEFAULT NULL COMMENT '应用id',
                            `node_id` bigint DEFAULT NULL COMMENT '环境id',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='appnode关联表';


DROP TABLE IF EXISTS `jwt_blacklist`;
CREATE TABLE `jwt_blacklist` (
                                 `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                                 `jwt` text COMMENT 'jwt',
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='黑名单表';




DROP TABLE IF EXISTS `node`;
CREATE TABLE `node` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                        `name` varchar(30) DEFAULT NULL COMMENT '环境名称',
                        `url` varchar(255) DEFAULT NULL COMMENT '环境地址',
                        `key` varchar(255) DEFAULT NULL COMMENT '环境唯一标示',
                        `status` tinyint(1) DEFAULT NULL COMMENT '环境连接状态',
                        `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        `content` varchar(255) DEFAULT NULL COMMENT '环境说明',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='node表';


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                        `username` varchar(191) DEFAULT NULL COMMENT '用户登录名',
                        `password` varchar(191) DEFAULT NULL COMMENT '用户登录密码',
                        `nick_name` varchar(191) DEFAULT '系统用户' COMMENT '用户昵称',
                        `header_img` varchar(191) DEFAULT 'http://qmplusimg.henrongyi.top/head.png' COMMENT '用户头像',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
BEGIN;
INSERT INTO `user` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '系统用户', 'http://qmplusimg.henrongyi.top/head.png');
COMMIT;
#-------------END-------------------------

-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `email` char(64) COLLATE utf8_unicode_ci NOT NULL COMMENT 'email',
    `username` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '註冊帳號名',
    `encrypted_password` char(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '加密密碼',
    `sign_in_count` int(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登入次數',
    `sign_in_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登入時間',
    `sign_in_ip` varchar(39) COLLATE utf8_unicode_ci NULL DEFAULT '' COMMENT '登入ip',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
    `name` char(32) COLLATE utf8_unicode_ci NULL COMMENT '顯示帳號名',
    `first_name` char(32) COLLATE utf8_unicode_ci NULL COMMENT '名字',
    `last_name` char(32) COLLATE utf8_unicode_ci NULL COMMENT '姓氏',
    `user_type` tinyint(1) UNSIGNED NULL DEFAULT '1' COMMENT '帳戶類型',
    `state` tinyint(1) UNSIGNED COLLATE utf8_unicode_ci NOT NULL DEFAULT '1' COMMENT '1=active, 2=blocked',
    `external` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '1=內部用戶, 2=外部用戶',
    `locked_at` timestamp NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '被鎖時間',
    `organization` char(32) COLLATE utf8_unicode_ci NULL DEFAULT '' COMMENT '組織',
    `created_by_id` int UNSIGNED NOT NULL DEFAULT '1' COMMENT '1=自行註冊,建立者ID',
    `location` char(8) COLLATE utf8_unicode_ci NULL DEFAULT '' COMMENT 'country codes',
    `note` varchar(64) COLLATE utf8_unicode_ci NULL DEFAULT '' COMMENT '備註',
    `preferred_language` char(8) COLLATE utf8_unicode_ci NULL DEFAULT '' COMMENT 'language code',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`email`),
    UNIQUE KEY (`username`),
    KEY `name` (`name`) USING BTREE
)  ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;

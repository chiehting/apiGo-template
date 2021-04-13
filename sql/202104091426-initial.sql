-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users` (
    `id` bigint UNSIGNED NOT NULL COMMENT 'ID',
    `email` char(64) COLLATE utf8_unicode_ci NOT NULL COMMENT 'email',
    `encrypted_password` char(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '加密密碼',
    `reset_password_token`  char(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '重設密碼',
    `reset_password_sent_at` timestamp NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '重設密碼時間',
    `sign_in_count` int(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登入次數',
    `sign_in_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '當前登入時間',
    `sign_in_ip` varchar(39) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '登入ip',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
    `name` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '顯示帳號名',
    `admin` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '0=用戶, 1=管理者',
    `locked_at` timestamp NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '被鎖時間',
    `username` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '註冊帳號名',
    `state` tinyint(1) COLLATE utf8_unicode_ci NOT NULL DEFAULT '1' COMMENT '-1=blocked, 1=active',
    `password_expires_at` timestamp NOT NULL DEFAULT '1970-01-01 00:00:01' COMMENT '密碼過期時間',
    `created_by_id` int(20) NOT NULL DEFAULT '0' COMMENT '建立者ID',
    `location` char(8) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT 'country codes',
    `note` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '備註',
    `external` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '0=內部用戶, 1=外部用戶',
    `organization` char(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '組織',
    `preferred_language` char(8) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT 'language code',
    `first_name` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '名字',
    `last_name` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '姓氏',
    `user_type` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '帳戶類型',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`email`),
    UNIQUE KEY (`username`),
    KEY `name` (`name`) USING BTREE,
    KEY `external` (`external`) USING BTREE
)  ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;

-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `settings` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `class` char(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '設定群組',
    `subject` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '設定項目',
    `value` varchar(255) COLLATE utf8_unicode_ci  NULL DEFAULT '' COMMENT '設定值',
    `description` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '設定描述',
    `data_type` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '資料類型',
    `data_edit` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '1=不可編輯, 2=可編輯',
    `max_value` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '1=無上限',
    PRIMARY KEY (`id`),
    UNIQUE KEY `class-subject` (`class`, `subject`) USING BTREE
)  ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci AUTO_INCREMENT=1;

INSERT INTO `settings`
(`id`, `class`, `subject`, `value`, `description`, `data_type`, `data_edit`, `max_value`)
VALUES
(1, 'application', 'name', 'apiGo-template', 'Application name', 'string', 1, 1),
(2, 'application', 'version', '0.0.1', 'Application version', 'string', 1, 1);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `settings`;

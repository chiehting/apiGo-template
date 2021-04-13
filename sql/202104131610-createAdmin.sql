-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO `users` 
(`id`, `email`, `encrypted_password`, `reset_password_token`, `reset_password_sent_at`, `sign_in_count`, `sign_in_at`, `sign_in_ip`, `created_at`, `updated_at`, `name`, `admin`, `locked_at`, `username`, `state`, `password_expires_at`, `created_by_id`, `location`, `note`, `external`, `organization`, `preferred_language`, `first_name`, `last_name`, `user_type`)
VALUES
(0, '', '', '', '1970-01-01 00:00:01', 0, '2021-04-13 08:09:27', '', '2021-04-13 08:09:27', '2021-04-13 08:09:27', '', 0, '1970-01-01 00:00:01', '', 1, '1970-01-01 00:00:01', 0, '', '', 0, '', '', '', '', 0),
(350230461512954596, 'ting911111@gmail.com', '', '', '1970-01-01 00:00:01', 0, '2021-04-13 08:09:27', '', '2021-04-13 08:09:27', '2021-04-13 08:09:27', 'admin', 1, '1970-01-01 00:00:01', 'admin', 1, '1970-01-01 00:00:01', 0, '', 'admin', 0, '', 'en', 'admin', 'admin', 0);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM users WHERE id in (0, 350230461512954596);

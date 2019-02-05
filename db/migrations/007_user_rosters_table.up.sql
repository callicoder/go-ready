CREATE TABLE `user_rosters` (
    `user_id` BIGINT NOT NULL,
    `roster_id` BIGINT NOT NULL,
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp DEFAULT NULL,
    PRIMARY KEY (user_id, roster_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
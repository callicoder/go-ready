CREATE TABLE `messages` (
    `id` BIGINT PRIMARY KEY,
    `message` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `data` JSON DEFAULT NULL,
    `user_id` BIGINT NOT NULL,
    `group_roster_id` BIGINT NOT NULL,     
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `group_members` (
    `group_id` BIGINT NOT NULL,
    `user_id` BIGINT NOT NULL,
    `role` VARCHAR(20) NOT NULL,
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp DEFAULT NULL,
    PRIMARY KEY (group_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
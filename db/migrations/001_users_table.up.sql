CREATE TABLE `users` (
    `id` BIGINT PRIMARY KEY,
    `email` VARCHAR(255) UNIQUE NOT NULL,
    `first_name` VARCHAR(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `last_name` VARCHAR(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `image_url` VARCHAR(255) DEFAULT NULL, 
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
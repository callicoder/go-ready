CREATE TABLE `groups` (
    `id` BIGINT PRIMARY KEY,
    `name` VARCHAR(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `description` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `image_url` VARCHAR(255) DEFAULT NULL,
    `created_by` BIGINT NOT NULL,
    `updated_by` BIGINT NOT NULL,
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
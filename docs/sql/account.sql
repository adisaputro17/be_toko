CREATE TABLE IF NOT EXISTS `account`
(
    `account_id` varchar(100) NOT NULL,
    `fullname` varchar(100) NOT NULL,
    `username` varchar(100) NOT NULL,
    `password` text NOT NULL,
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `created_by` bigint unsigned NOT NULL,
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_by` bigint unsigned NOT NULL,
    PRIMARY KEY (`account_id`),
    UNIQUE KEY `uq_username` (`username`)
)
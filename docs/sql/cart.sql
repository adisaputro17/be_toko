CREATE TABLE IF NOT EXISTS `cart`
(
    `cart_id` varchar(100) NOT NULL,
    `account_id` varchar(100) NOT NULL,
    `product_id` varchar(100) NOT NULL,
    `qty` int NOT NULL,
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `created_by` bigint unsigned NOT NULL,
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_by` bigint unsigned NOT NULL,
    PRIMARY KEY (`cart_id`)
)
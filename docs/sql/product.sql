CREATE TABLE IF NOT EXISTS `product_category`
(
    `product_category_id` varchar(100) NOT NULL,
    `product_category_name` varchar(100) NOT NULL,
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `created_by` bigint unsigned NOT NULL,
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_by` bigint unsigned NOT NULL,
    PRIMARY KEY (`product_category_id`)
)

CREATE TABLE IF NOT EXISTS `product`
(
    `product_id` varchar(100) NOT NULL,
    `product_name` varchar(100) NOT NULL,
    `product_category_id` varchar(100) NOT NULL,
    `description` text NOT NULL,
    `price` decimal(15,2) NOT NULL,
    `stock` int NOT NULL,
    `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `created_by` bigint unsigned NOT NULL,
    `updated_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_by` bigint unsigned NOT NULL,
    PRIMARY KEY (`product_id`)
)
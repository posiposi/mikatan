-- CreateTable
CREATE TABLE `prices` (
    `price_id` VARCHAR(36) NOT NULL,
    `item_id` VARCHAR(36) NOT NULL,
    `price_with_tax` INTEGER NOT NULL,
    `price_without_tax` INTEGER NOT NULL,
    `tax_rate` DOUBLE NOT NULL DEFAULT 10.0,
    `currency` VARCHAR(3) NOT NULL DEFAULT 'JPY',
    `start_date` DATETIME(3) NOT NULL,
    `end_date` DATETIME(3) NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NULL,

    PRIMARY KEY (`price_id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `prices` ADD CONSTRAINT `prices_item_id_fkey` FOREIGN KEY (`item_id`) REFERENCES `items`(`item_id`) ON DELETE RESTRICT ON UPDATE CASCADE;

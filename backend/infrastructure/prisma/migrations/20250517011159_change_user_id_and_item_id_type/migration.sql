/*
  Warnings:

  - The primary key for the `items` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - The primary key for the `users` table will be changed. If it partially fails, the table could be left without primary key constraint.

*/
-- DropForeignKey
ALTER TABLE `items` DROP FOREIGN KEY `items_user_id_fkey`;

-- DropIndex
DROP INDEX `items_user_id_fkey` ON `items`;

-- AlterTable
ALTER TABLE `items` DROP PRIMARY KEY,
    MODIFY `item_id` VARCHAR(36) NOT NULL,
    MODIFY `user_id` VARCHAR(36) NOT NULL,
    ADD PRIMARY KEY (`item_id`);

-- AlterTable
ALTER TABLE `users` DROP PRIMARY KEY,
    MODIFY `user_id` VARCHAR(36) NOT NULL,
    ADD PRIMARY KEY (`user_id`);

-- AddForeignKey
ALTER TABLE `items` ADD CONSTRAINT `items_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`) ON DELETE RESTRICT ON UPDATE CASCADE;

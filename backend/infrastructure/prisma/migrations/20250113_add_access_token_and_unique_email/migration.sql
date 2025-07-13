-- AlterTable
ALTER TABLE `users` ADD COLUMN `access_token` TEXT;

-- Modify email column to varchar and add unique constraint
ALTER TABLE `users` MODIFY `email` VARCHAR(255) NOT NULL;
ALTER TABLE `users` ADD UNIQUE KEY `users_email_key` (`email`);
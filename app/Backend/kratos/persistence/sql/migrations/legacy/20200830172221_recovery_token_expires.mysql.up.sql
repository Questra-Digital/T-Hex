ALTER TABLE `identity_recovery_tokens` ADD COLUMN `expires_at` DATETIME NOT NULL DEFAULT '2000-01-01 00:00:00';
ALTER TABLE `identity_recovery_tokens` ADD COLUMN `issued_at` DATETIME NOT NULL DEFAULT '2000-01-01 00:00:00';
ALTER TABLE `identity_recovery_tokens` MODIFY `selfservice_recovery_flow_id` char(36);
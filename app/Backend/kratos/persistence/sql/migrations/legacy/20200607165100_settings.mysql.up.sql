ALTER TABLE `selfservice_settings_requests` ADD COLUMN `state` VARCHAR (255) NOT NULL DEFAULT 'show_form';
ALTER TABLE `selfservice_settings_requests` DROP COLUMN `update_successful`;
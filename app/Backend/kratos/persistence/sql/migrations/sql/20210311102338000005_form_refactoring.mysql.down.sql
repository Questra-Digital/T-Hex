CREATE TABLE `selfservice_recovery_flow_methods` (
`id` char(36) NOT NULL,
PRIMARY KEY(`id`),
`method` VARCHAR (32) NOT NULL,
`selfservice_recovery_flow_id` char(36) NOT NULL,
`config` JSON NOT NULL,
`created_at` DATETIME NOT NULL,
`updated_at` DATETIME NOT NULL,
FOREIGN KEY (`selfservice_recovery_flow_id`) REFERENCES `selfservice_recovery_flow_methods` (`id`) ON DELETE cascade
) ENGINE=InnoDB;
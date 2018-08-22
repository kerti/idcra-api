CREATE TABLE IF NOT EXISTS `users`
(
  `id` CHAR(36) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `ip_address` VARCHAR(45),
  `created_at` TIMESTAMP DEFAULT NOW(),
  PRIMARY KEY (`id`),
  UNIQUE INDEX `users_idx_1` (`email`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `roles`
(
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY(`id`),
  UNIQUE INDEX `roles_idx_1` (`name`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `rel_users_roles` (
  `user_id` CHAR(36) NOT NULL,
  `role_id` CHAR(36) NOT NULL,
  PRIMARY KEY (`user_id`, `role_id`),
  CONSTRAINT `fk_users_roles_users` FOREIGN KEY (`user_id`)
    REFERENCES `users`(`id`)
    ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_roles_roles` FOREIGN KEY (`role_id`)
    REFERENCES `roles`(`id`)
    ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;

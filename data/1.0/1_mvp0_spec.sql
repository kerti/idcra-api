-- IDCRA API Migration File MVP0 Spec
-- Contents:
-- - Users
-- - Roles
-- - User <--> Role relationship
-- - Schools
-- - Students
-- - DiagnosisAndActions
-- - Surveys
-- - Cases
-- ----------------------------------------------------------------------------

-- Users Table
CREATE TABLE IF NOT EXISTS `users` (
  `id` CHAR(36) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password` BLOB NOT NULL,
  `ip_address` VARCHAR(45),
  `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY (`id`),
  UNIQUE INDEX `users_idx_1` (`email`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;
-- ----------------------------------------------------------------------------

-- Roles Table
CREATE TABLE IF NOT EXISTS `roles` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY(`id`),
  UNIQUE INDEX `roles_idx_1` (`name`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;

-- Roles Data
INSERT INTO `roles` VALUES
('0780e271-5f62-405e-a700-f93d65b70dd4', 'ADMIN'),
('aa8a0b5f-2c48-4f5d-ac49-1f001320a2e1', 'SURVEYOR');
-- ----------------------------------------------------------------------------

-- User <--> Role relationship
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
-- ----------------------------------------------------------------------------

-- Schools Table
CREATE TABLE IF NOT EXISTS `schools` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `schools_idx_1` (`name`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;
-- ----------------------------------------------------------------------------

-- Students Table
CREATE TABLE IF NOT EXISTS `students` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `date_of_birth` DATE NOT NULL,
  `school_id` CHAR(36) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `students_idx_1` (`name`, `school_id`, `date_of_birth`),
  CONSTRAINT `fk_students_schools` FOREIGN KEY (`school_id`)
    REFERENCES `schools`(`id`)
    ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;
-- ----------------------------------------------------------------------------

-- DiagnosisAndActions Table
CREATE TABLE IF NOT EXISTS `diagnosis_and_actions` (
  `id` CHAR(36) NOT NULL,
  `diagnosis` VARCHAR(255) NOT NULL,
  `action` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `diagnosis_and_actions_idx_1` (`diagnosis`, `action`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;

-- DiagnosisAndActions Data
INSERT INTO `diagnosis_and_actions` (`id`, `diagnosis`, `action`) VALUES
('4485a2cb-688e-4e0b-bdcf-bcc008c37a69', 'Pit Fissure Dalam', 'Fissure Sealant'),
('469d07e4-d71a-4dcb-9932-c805aefff48f', 'Karies Superficial', 'ART');
-- ----------------------------------------------------------------------------

-- Surveys Table
CREATE TABLE IF NOT EXISTS `surveys` (
  `id` CHAR(36) NOT NULL,
  `student_id` CHAR(36) NOT NULL,
  `surveyor_id` CHAR(36) NOT NULL,
  `date` DATE NOT NULL,
  `s1q1` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s1q2` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s1q3` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s1q4` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s1q5` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s1q6` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s1q7` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s2q1` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s2q2` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s2q3` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s2q4` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s2q5` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s2q6` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s2q7` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s2q8` ENUM('Low', 'Medium', 'High') NOT NULL,
  `s2q9` ENUM('Low', 'Medium', 'High') NOT NULL,
  `lower_d` INT,
  `lower_e` INT,
  `lower_f` INT,
  `upper_d` INT,
  `upper_m` INT,
  `upper_f` INT,
  `subjective_score` INT,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `survey_idx_1` (`student_id`, `date`),
  INDEX `survey_idx_2` (`subjective_score`),
  CONSTRAINT `fk_surveys_students` FOREIGN KEY (`student_id`)
    REFERENCES `students`(`id`)
    ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_surveys_users` FOREIGN KEY (`surveyor_id`)
    REFERENCES `users`(`id`)
    ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;
-- ----------------------------------------------------------------------------

-- Cases Table
CREATE TABLE IF NOT EXISTS `cases` (
  `id` CHAR(36) NOT NULL,
  `survey_id` CHAR(36) NOT NULL,
  `tooth_number` INT NOT NULL,
  `diagnosis_and_action_id` CHAR(36) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_cases_surveys` FOREIGN KEY (`survey_id`)
    REFERENCES `surveys`(`id`)
    ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_cases_diagnosis_and_actions` FOREIGN KEY (`diagnosis_and_action_id`)
    REFERENCES `diagnosis_and_actions`(`id`)
    ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8;
-- ----------------------------------------------------------------------------

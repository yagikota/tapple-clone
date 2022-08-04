CREATE SCHEMA IF NOT EXISTS `tapple_c` DEFAULT CHARACTER SET utf8mb4;

USE `tapple_c` ;

SET CHARSET utf8mb4;

SET CHARSET utf8mb4;

CREATE TABLE
    IF NOT EXISTS `tapple_c`.`user`(
        `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "ユーザーID",
        `name` VARCHAR(255) NOT NULL COMMENT "ユーザー名",
        `icon` VARCHAR(255) NOT NULL COMMENT "アイコンURL",
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "作成日時",
        `deleted_at` DATETIME DEFAULT NULL COMMENT "論理削除日時",
        PRIMARY KEY(`id`)
    ) ENGINE = InnoDB COMMENT = 'ユーザ';
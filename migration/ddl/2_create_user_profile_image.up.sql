CREATE TABLE
    IF NOT EXISTS `tapple_c`.`user_profile_image`(
        `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "ユーザーID",
        `user_id` int(11) UNSIGNED NOT NULL COMMENT "ユーザーID",
        `image_path` VARCHAR(255) NOT NULL COMMENT "プロフィール画像の取得先",
        PRIMARY KEY(`id`),
        UNIQUE (`user_id`, `image_path`),
        CONSTRAINT `fk_user_profile_image_user`
            FOREIGN KEY (`user_id`)
            REFERENCES `tapple_c`.`user` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
    ) ENGINE = InnoDB COMMENT = 'ユーザー画像';
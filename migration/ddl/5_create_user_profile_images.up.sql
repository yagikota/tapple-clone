CREATE TABLE
    IF NOT EXISTS user_profile_images (
        id INT(11) AUTO_INCREMENT COMMENT "プロフィール写真ID",
        user_id INT(11) NOT NULL COMMENT "ユーザーID",
        image_path VARCHAR(1024) NOT NULL COMMENT "写真URL",
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "プロフィール写真作成日時",
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "プロフィール写真更新日時",
        deteled_at DATETIME COMMENT "プロフィール写真論理削除日時",
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users(id)
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'プロフィール写真';

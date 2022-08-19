CREATE TABLE
    IF NOT EXISTS hobbies (
        id INT(11) AUTO_INCREMENT COMMENT "趣味タグID",
        user_id INT(11) NOT NULL COMMENT "ユーザーID",
        tag VARCHAR(20) NOT NULL COMMENT "趣味タグ",
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "趣味タグ作成日時",
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "趣味タグ更新日時",
        deteled_at DATETIME COMMENT "趣味タグ論理削除日時",
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users(id)
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '趣味タグ';

CREATE TABLE
    IF NOT EXISTS room_users (
        id INT(11) AUTO_INCREMENT COMMENT "ルームユーザーID",
        user_id INT(11) NOT NULL COMMENT "ユーザーID",
        room_id INT(11) NOT NULL COMMENT "ルームID",
        is_pinned BOOLEAN NOT NULL DEFAULT 0 COMMENT "ピン留め(0: されていない, 1: されている)",
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "ルームユーザー作成日時",
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "ルームユーザー作成日時",
        deteled_at DATETIME COMMENT "ルームユーザー論理削除日時",
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (room_id) REFERENCES rooms(id)
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'ルームユーザー';

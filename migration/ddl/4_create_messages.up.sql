CREATE TABLE
    IF NOT EXISTS messages (
        id BIGINT(11) AUTO_INCREMENT COMMENT "メッセージID",
        user_id INT(11) NOT NULL COMMENT "ユーザーID",
        room_id INT(11) NOT NULL COMMENT "ルームID",
        content TEXT NOT NULL COMMENT "メッセージ内容",
        is_read BOOLEAN NOT NULL DEFAULT 0 COMMENT "0(false): 未読, 1(true): 既読",
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "メッセージ作成日時",
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "メッセージ更新日時",
        deteled_at DATETIME COMMENT "メッセージ論理削除日時",
        PRIMARY KEY (id),
        FOREIGN KEY (user_id) REFERENCES users(id),
        FOREIGN KEY (room_id) REFERENCES rooms(id)
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'メッセージ';

CREATE TABLE
    IF NOT EXISTS users (
        id INT(11) AUTO_INCREMENT COMMENT "ユーザーID",
        name VARCHAR(8) NOT NULL COMMENT "ユーザー名",
        icon VARCHAR(1024) NOT NULL COMMENT "アイコンURL",
        gender INT(11) NOT NULL COMMENT "性別（男: 0, 女: 1）",
        birthday DATE NOT NULL COMMENT "誕生日",
        location INT(11) NOT NULL COMMENT "所在地(0: その他, 北海道:1~沖縄: 47)",
        is_principal BOOLEAN NOT NULL DEFAULT 0 COMMENT "本人認証(0: されていない, 1: されている)",
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "ユーザー作成日時",
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "ユーザー更新日時",
        deteled_at DATETIME COMMENT "ユーザー論理削除日時",
        PRIMARY KEY (id)
    ) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'ユーザー';

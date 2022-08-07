-- Active: 1659632193111@@127.0.0.1@3306@tapple_c

use `tapple_c`;

SET CHARSET utf8mb4;

-- user

INSERT INTO
    `user` (`id`, `name`, `icon`)
VALUES (1, "name1", "/icon/1");

INSERT INTO
    `user` (`id`, `name`, `icon`)
VALUES (2, "name2", "/icon/2");

INSERT INTO
    `user` (`id`, `name`, `icon`)
VALUES (3, "name3", "/icon/3");

-- user profile image

INSERT INTO
    `user_profile_image` (`id`, `user_id`, `image_path`)
VALUES (1, 1, "/profile/1");

INSERT INTO
    `user_profile_image` (`id`, `user_id`, `image_path`)
VALUES (2, 1, "/profile/2");

INSERT INTO
    `user_profile_image` (`id`, `user_id`, `image_path`)
VALUES (3, 1, "/profile/3");

INSERT INTO
    `user_profile_image` (`id`, `user_id`, `image_path`)
VALUES (4, 2, "/profile/4");

INSERT INTO
    `user_profile_image` (`id`, `user_id`, `image_path`)
VALUES (5, 2, "/profile/5");

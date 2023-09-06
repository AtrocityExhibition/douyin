CREATE TABLE user (
                      id bigint AUTO_INCREMENT,
                      name varchar(32) NOT NULL DEFAULT '' COMMENT 'The username',
                      password varchar(32) NOT NULL DEFAULT '' COMMENT 'The password',
                      follow_count bigint NOT NULL DEFAULT 0 COMMENT 'The follow_count',
                      follower_count bigint NOT NULL DEFAULT 0 COMMENT 'The follower_count',
                      is_follow varchar(10) NOT NULL DEFAULT 'false' COMMENT 'The is_follow',
                      avatar varchar(255),
                      background_image varchar(255),
                      signature varchar(255),
                      total_favorited bigint NOT NULL DEFAULT 0 COMMENT 'The total_favorited',
                      work_count bigint NOT NULL DEFAULT 0 COMMENT 'The work_count',
                      favorite_count bigint NOT NULL DEFAULT 0 COMMENT 'The favorite_count',
                      create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      delete_at timestamp,
                      PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'user table';
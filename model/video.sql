CREATE TABLE video (
                      id bigint AUTO_INCREMENT,
                      author bigint COMMENT 'The author',
                      play_url varchar(255) NOT NULL DEFAULT '' COMMENT 'The play_url',
                      cover_url varchar(255) NOT NULL DEFAULT '' COMMENT 'The cover_url',
                      title varchar(32) NOT NULL DEFAULT '' COMMENT 'The title',
                      is_favorite  varchar(10) NOT NULL DEFAULT 'false' COMMENT 'The is_favorite',
                      favorite_count bigint NOT NULL DEFAULT 0 COMMENT 'The favorite_count',
                      comment_count  bigint NOT NULL DEFAULT 0 COMMENT 'The comment_count ',
                      create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      delete_at timestamp,
                      PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'video table';
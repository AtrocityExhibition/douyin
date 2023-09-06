CREATE TABLE friend (
                          id bigint AUTO_INCREMENT,
                          user1_id bigint ,
                          user2_id bigint ,
                          create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          delete_at timestamp,
                          PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'friend table';
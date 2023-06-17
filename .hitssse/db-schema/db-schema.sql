SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS annotation_table_annotation_table, annotation_table_asset_table, annotation_table_issue_table, asset_table_annotation_table, asset_table_issue_table, hibernate_sequence, issue_table_annotation_table, issue_table_asset_table, issue_table, issue_information_table, issue_project_table, issue_resolved_hash_table, issue_approved_hash_table, asset_table, asset_information_table, asset_project_table, annotation_table, annotation_commit_table, annotation_information_table, annotation_project_table, annotation_position_table, asset_issue_table, annotation_asset_table, annotation_issue_table, asset_potential_vulnerability_table, annotation_potential_vulnerability_table, potential_vulnerability_table, potential_vulnerability_information_table, potential_vulnerability_resolved_hash_table, potential_vulnerability_approved_hash_table;
SET FOREIGN_KEY_CHECKS = 1;

# Issue
CREATE TABLE `potential_vulnerability_table` (
                                                 `id` varchar(36) NOT NULL,
                                                 `category` varchar(128) NOT NULL,
                                                 `parent_id` varchar(36),
                                                 `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                                 `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                 PRIMARY KEY (`id`),
                                                 CONSTRAINT `fk_issue_issue` FOREIGN KEY (`parent_id`) REFERENCES `potential_vulnerability_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `potential_vulnerability_information_table` (
                                                             `id` varchar(36) NOT NULL,
                                                             `project` varchar(256) NOT NULL,
                                                             `risk_level` varchar(20) DEFAULT 'warm',
                                                             `risk_description` varchar(1024),
                                                             PRIMARY KEY (`id`),
                                                             CONSTRAINT `fk_issue_inforamtation` FOREIGN KEY (`id`) REFERENCES `potential_vulnerability_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `potential_vulnerability_resolved_hash_table` (
                                                               `id` varchar(36) NOT NULL,
                                                               `annotation_id` varchar(36) NOT NULL,
                                                               `annotation_hash` varchar(64) NOT NULL,
                                                               PRIMARY KEY (`id`,`annotation_id`, `annotation_hash`),
                                                               CONSTRAINT `fk_issue_resolved` FOREIGN KEY (`id`) REFERENCES `potential_vulnerability_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `potential_vulnerability_approved_hash_table` (
                                                               `id` varchar(36) NOT NULL,
                                                               `annotation_id` varchar(36) NOT NULL,
                                                               `annotation_hash` varchar(64) NOT NULL,
                                                               PRIMARY KEY (`id`,`annotation_id`, `annotation_hash`),
                                                               CONSTRAINT `fk_issue_approved` FOREIGN KEY (`id`) REFERENCES `potential_vulnerability_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


# Asset
CREATE TABLE `asset_table` (
                               `id` varchar(36) NOT NULL,
                               `category` varchar(128) NOT NULL,
                               `parent_id` varchar(36),
                               `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`),
                               CONSTRAINT `fk_asset_asset` FOREIGN KEY (`parent_id`) REFERENCES `asset_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `asset_information_table` (
                                           `id` varchar(36) NOT NULL,
                                           `project` varchar(256) NOT NULL,
                                           `description` varchar(1024) NOT NULL,
                                           PRIMARY KEY (`id`),
                                           CONSTRAINT `fk_asset_information` FOREIGN KEY (`id`) REFERENCES `asset_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Annotation
CREATE TABLE `annotation_table` (
                                    `annotation_id` varchar(36) NOT NULL,
                                    `annotation_hash` varchar(64) NOT NULL,
                                    `replaces_id` varchar(36),
                                    `replaces_hash` varchar(64),
                                    `filepath` varchar(256) NOT NULL,
                                    `annotation_startcomment_startposition_offset` smallint NOT NULL,
                                    `annotation_startcomment_startposition_linenumber` smallint NOT NULL,
                                    `annotation_startcomment_startposition_column` smallint NOT NULL,
                                    `annotation_startcomment_endposition_offset` smallint NOT NULL,
                                    `annotation_startcomment_endposition_linenumber` smallint NOT NULL,
                                    `annotation_startcomment_endposition_column` smallint NOT NULL,
                                    `annotation_endcomment_startposition_offset` smallint,
                                    `annotation_endcomment_startposition_linenumber` smallint,
                                    `annotation_endcomment_startposition_column` smallint,
                                    `annotation_endcomment_endposition_offset` smallint,
                                    `annotation_endcomment_endposition_linenumber` smallint,
                                    `annotation_endcomment_endposition_column` smallint,
                                    `annotation_codeblock_startposition_offset` smallint NOT NULL,
                                    `annotation_codeblock_startposition_linenumber` smallint NOT NULL,
                                    `annotation_codeblock_startposition_column` smallint NOT NULL,
                                    `annotation_codeblock_endposition_offset` smallint NOT NULL,
                                    `annotation_codeblock_endposition_linenumber` smallint NOT NULL,
                                    `annotation_codeblock_endposition_column` smallint NOT NULL,
                                    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`annotation_id`,`annotation_hash`),
                                    CONSTRAINT fk_annotation_annotation FOREIGN KEY (`replaces_id`, `replaces_hash`) REFERENCES `annotation_table` (`annotation_id`, `annotation_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `annotation_commit_table` (
                                           `commit_id` varchar(40) NOT NULL,
                                           `annotation_id` varchar(36) NOT NULL,
                                           `annotation_hash` varchar(64) NOT NULL,
                                           `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           PRIMARY KEY (`commit_id`, `annotation_id`, `annotation_hash`),
                                           CONSTRAINT fk_annotation_commit FOREIGN KEY (`annotation_id`, `annotation_hash`) REFERENCES `annotation_table` (`annotation_id`, `annotation_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `annotation_information_table` (
                                                `annotation_id` varchar(36) NOT NULL,
                                                `annotation_hash` varchar(64) NOT NULL,
                                                `project` varchar(256) NOT NULL,
                                                `hierarchy_level` varchar(256) NOT NULL,
                                                `annotated_by` varchar(256),
                                                `approved_by` varchar(256),
                                                `description` varchar(1024),
                                                `alias` varchar(256),
                                                `responsible_security_expert` varchar(256),
                                                PRIMARY KEY (`annotation_id`,`annotation_hash`),
                                                CONSTRAINT fk_annotation_information FOREIGN KEY (`annotation_id`,`annotation_hash`) REFERENCES `annotation_table` (`annotation_id`, `annotation_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# Mapping
CREATE TABLE `asset_potential_vulnerability_table` (
                                                       `asset_id` varchar(36) NOT NULL,
                                                       `potential_vulnerability_id` varchar(36) NOT NULL,
                                                       PRIMARY KEY (`asset_id`,`potential_vulnerability_id`),
                                                       CONSTRAINT `fk_asset_issue` FOREIGN KEY (`asset_id`) REFERENCES `asset_table` (`id`),
                                                       CONSTRAINT `fk_asset_issue2` FOREIGN KEY (`potential_vulnerability_id`) REFERENCES `potential_vulnerability_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `annotation_asset_table` (
                                          `annotation_id` varchar(36) NOT NULL,
                                          `annotation_hash` varchar(64) NOT NULL,
                                          `asset_id` varchar(36) NOT NULL,
                                          PRIMARY KEY (`annotation_id`,`annotation_hash`,`asset_id`),
                                          CONSTRAINT `fk_annotation_asset` FOREIGN KEY (`annotation_id`,`annotation_hash`) REFERENCES `annotation_table` (`annotation_id`,`annotation_hash`),
                                          CONSTRAINT `fk_annotation_asset2` FOREIGN KEY (`asset_id`) REFERENCES `asset_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `annotation_potential_vulnerability_table` (
                                                            `annotation_id` varchar(36) NOT NULL,
                                                            `annotation_hash` varchar(64) NOT NULL,
                                                            `potential_vulnerability_id` varchar(36) NOT NULL,
                                                            PRIMARY KEY (`annotation_id`,`annotation_hash`,`potential_vulnerability_id`),
                                                            CONSTRAINT `fk_annotation_issue` FOREIGN KEY (`annotation_id`,`annotation_hash`) REFERENCES `annotation_table` (`annotation_id`,`annotation_hash`),
                                                            CONSTRAINT `fk_annotation_issue2` FOREIGN KEY (`potential_vulnerability_id`) REFERENCES `potential_vulnerability_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;





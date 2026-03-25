-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS entity_tags (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    entity_id BIGINT NOT NULL,
    entity VARCHAR(255) NOT NULL,
    tag_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    UNIQUE KEY unique_entity_tag (entity_id, entity, tag_id),
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS entity_tags;
-- +goose StatementEnd

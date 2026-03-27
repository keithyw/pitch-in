-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS asset_entities (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    entity_id BIGINT NOT NULL,
    entity VARCHAR(255) NOT NULL,
    asset_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    UNIQUE KEY unique_entity_asset (entity_id, entity, asset_id),
    FOREIGN KEY (asset_id) REFERENCES assets(id) ON DELETE CASCADE    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS asset_entities;
-- +goose StatementEnd

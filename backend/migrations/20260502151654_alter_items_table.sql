-- +goose Up
-- +goose StatementBegin
ALTER TABLE items 
ADD COLUMN user_id BIGINT NULL,
ADD CONSTRAINT fk_items_user_id
  FOREIGN KEY(user_id)
  REFERENCES users(id)
  ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE items 
DROP FOREIGN KEY fk_items_user_id,
DROP COLUMN user_id;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS photos (
    user_id INT REFERENCES profiles(id) ON DELETE CASCADE,
    photo_link TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS photos;
-- +goose StatementEnd

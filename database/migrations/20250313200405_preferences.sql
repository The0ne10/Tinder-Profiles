-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS preferences (
    user_id INT REFERENCES profiles(id) ON DELETE CASCADE,
    sex VARCHAR(10),
    age INT,
    radius INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS preferences;
-- +goose StatementEnd

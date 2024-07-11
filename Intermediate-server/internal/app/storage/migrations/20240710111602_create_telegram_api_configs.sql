-- +goose Up
-- +goose StatementBegin
CREATE TABLE telegram_api_configs(
    id SERIAL PRIMARY KEY,
    task_id INTEGER,
    API_ID TEXT,
    API_HASH TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS telegram_api_configs;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_values (
    key TEXT PRIMARY KEY,
    value TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_values;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE values (
    key TEXT PRIMARY KEY,
    value TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE values;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD firebase_uid VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP firebase_uid;
-- +goose StatementEnd

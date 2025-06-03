-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    whiteboard (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description VARCHAR(255)
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE whiteboard
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE c_routers (
    name VARCHAR(32) PRIMARY KEY NOT NULL,
    value VARCHAR(32) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS c_routers;
-- +goose StatementEnd

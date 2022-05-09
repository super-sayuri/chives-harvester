-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE c_users (
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    username VARCHAR(32) UNIQUE NOT NULL,
    email VARCHAR(128),
    mobile VARCHAR(64),
    salt VARCHAR(128) NOT NULL,
    password VARCHAR(128) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_user_email ON c_users(email);
CREATE INDEX idx_user_mobile ON c_users(mobile);

CREATE TRIGGER tri_user_update AFTER UPDATE ON c_users
BEGIN
    UPDATE c_users SET updated_at=CURRENT_TIMESTAMP WHERE id=NEW.id;
END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS c_users;
-- +goose StatementEnd

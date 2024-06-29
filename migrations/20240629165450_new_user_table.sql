-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id SERIAL
        CONSTRAINT users_pk
            PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    rank VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users OWNER TO test;

CREATE UNIQUE INDEX users_id_uindex
    ON users(id);

CREATE TABLE IF NOT EXISTS cars(
    id SERIAL
        CONSTRAINT cars_pk
            PRIMARY KEY,
    user_id INTEGER
        CONSTRAINT cars_users_id_fk
            REFERENCES users,
    colour VARCHAR(20),
    brand VARCHAR(20),
    licence_plate VARCHAR(10)

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

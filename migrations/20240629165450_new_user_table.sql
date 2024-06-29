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
);

ALTER TABLE cars OWNER TO test;

CREATE UNIQUE INDEX cars_id_uindex
    ON cars(id);
INSERT INTO public.users(id, name, rank, created_at) VALUES
                                                         (1, 'John K.', 'CEO', CURRENT_TIMESTAMP),
                                                         (2, 'Lana V.', 'VP Business', CURRENT_TIMESTAMP),
                                                         (3, 'Peter M.', 'Manager', CURRENT_TIMESTAMP),
                                                         (4, 'Sarah W.', 'Software Engineer', CURRENT_TIMESTAMP),
                                                         (5, 'David L.', 'Analyst', CURRENT_TIMESTAMP),
                                                         (6, 'Alice A.', 'Marketing Manager', CURRENT_TIMESTAMP),
                                                         (7, 'Bob B.', 'Sales Representative', CURRENT_TIMESTAMP),
                                                         (8, 'Emily E.', 'Customer Support', CURRENT_TIMESTAMP),
                                                         (9, 'Charles C.', 'Accountant', CURRENT_TIMESTAMP),
                                                         (10, 'Mary M.', 'HR Specialist', CURRENT_TIMESTAMP);
INSERT INTO public.cars(id, user_id, colour, brand, licence_plate) VALUES
                                                                       (1, 3, 'silver', 'Hyundai', 'MN0123'),
                                                                       (2, 5, 'red', 'Volkswagen', 'OP4567'),
                                                                       (3, 1, 'white', 'Kia', 'QR7890'),
                                                                       (4, 7, 'blue', 'Subaru', 'ST1234'),
                                                                       (5, 4, 'black', 'Audi', 'UV5678'),
                                                                       (6, 2, 'gray', 'BMW', 'WX9012'),
                                                                       (7, 8, 'red', 'Mazda', 'YZ3456'),
                                                                       (8, 9, 'green', 'Jeep', 'AB7890'),
                                                                       (9, 6, 'silver', 'Mini', 'CD1234'),
                                                                       (10, 10, 'white', 'Tesla', 'EF5678'),
                                                                       (11, 3, 'blue', 'Peugeot', 'GH9012'),
                                                                       (12, 8, 'black', 'Volvo', 'IJ3456'),
                                                                       (13, 5, 'red', 'Land Rover', 'KL7890'),
                                                                       (14, 1, 'gray', 'Citroen', 'MN1234'),
                                                                       (15, 4, 'silver', 'Fiat', 'OP5678'),
                                                                       (16, 9, 'white', 'Dodge', 'QR9012'),
                                                                       (17, 7, 'blue', 'Mitsubishi', 'ST3456'),
                                                                       (18, 2, 'black', 'Porsche', 'UV7890'),
                                                                       (19, 6, 'red', 'Ferrari', 'WX1234'),
                                                                       (20, 10, 'silver', 'Lamborghini', 'YZ5678');
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS cars
-- +goose StatementEnd

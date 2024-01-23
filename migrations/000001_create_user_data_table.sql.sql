-- +goose Up
CREATE TABLE IF NOT EXISTS user_data 
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT,
    age INT,
    gender TEXT,
    nationality TEXT
);

-- +goose Down
DROP TABLE IF EXISTS user_data;
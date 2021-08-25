-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE "Users" (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    age INT NOT NULL,
    user_type SMALLINT NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL
);

CREATE TABLE "Items" (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    user_id INT NOT NULL REFERENCES "Users"(id) ON DELETE CASCADE,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE "Users" CASCADE;
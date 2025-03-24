-- +goose up

CREATE TABLE users (
	ID UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP NOT NULL DEFAULT now(),
	username TEXT UNIQUE NOT NULL
);

-- +goose down

DROP TABLE users;
CREATE SCHEMA IF NOT EXISTS todoapp;

CREATE TABLE todoapp.users (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    full_name VARCHAR(100) NOT NULL CHECK (char_length(full_name) BETWEEN 3 AND 100),
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND char_length(phone_number) BETWEEN 10 AND 15
    ),
    email VARCHAR(255) NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

CREATE TABLE todoapp.tasks (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    title VARCHAR(100) NOT NULL CHECK (char_length(title) BETWEEN 1 and 100),
    description TEXT,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    completed_at TIMESTAMPTZ,
    author_id INTEGER NOT NULL REFERENCES todoapp.users(id),
    CHECK (
        (completed = FALSE AND completed_at IS NULL)
        OR
        (completed = TRUE AND completed_at >= created_at)
    )
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS schema_migrations_example (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


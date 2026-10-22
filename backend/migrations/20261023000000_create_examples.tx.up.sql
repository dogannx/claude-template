-- Example migration: delete this pair once your first real migration exists.
CREATE TABLE examples (
    id         bigserial PRIMARY KEY,
    name       text        NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

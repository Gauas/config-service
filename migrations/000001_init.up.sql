CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS configs (
    id          UUID      PRIMARY KEY DEFAULT gen_random_uuid(),
    service     TEXT      NOT NULL,
    environment TEXT      NOT NULL,
    config      JSONB     NOT NULL DEFAULT '{}',
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP,
    UNIQUE (service, environment, deleted_at)
);

CREATE INDEX IF NOT EXISTS idx_configs_service_env ON configs (service, environment);

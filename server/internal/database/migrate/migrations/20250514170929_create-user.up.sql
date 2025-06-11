CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS "users" (
  "id"             VARCHAR(255) PRIMARY KEY,
  "name"           VARCHAR(255) NOT NULL,
  "email"          VARCHAR(255) UNIQUE NOT NULL,
  "role"           VARCHAR(255) NOT NULL,
  "password_hash"  VARCHAR(255) NOT NULL,
  "created_at"     TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at"     TIMESTAMP,
  "deleted_at"     TIMESTAMP
);

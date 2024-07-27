CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "email" VARCHAR(256) UNIQUE NOT NULL,
  "hashed_password" VARCHAR(256) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);
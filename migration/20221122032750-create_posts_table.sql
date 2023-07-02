-- +migrate Up
CREATE TABLE IF NOT EXISTS "posts" (
  "id" bigserial NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp,
  "author_id" bigint NOT NULL,
  "caption" text NOT NULL,
  "is_published" boolean NOT NULL DEFAULT true,

  PRIMARY KEY ("id"),
  FOREIGN KEY ("author_id") REFERENCES "users" ("id")
);

-- +migrate Down
DROP TABLE IF EXISTS "posts";

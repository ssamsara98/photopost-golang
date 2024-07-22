
-- +migrate Up
CREATE TABLE IF NOT EXISTS "posts" (
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp,
  "id" bigserial NOT NULL,
  "author_id" bigint,
  "caption" text NOT NULL,
  "is_published" boolean NOT NULL DEFAULT true,

  PRIMARY KEY ("id"),
  FOREIGN KEY ("author_id") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS "posts";

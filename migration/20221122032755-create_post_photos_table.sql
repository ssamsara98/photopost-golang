-- +migrate Up
CREATE TABLE IF NOT EXISTS "post_photos" (
  "id" uuid DEFAULT uuid_generate_v4() NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp,
  "keypath" text NOT NULL,

  PRIMARY KEY ("id")
);

-- +migrate Down
DROP TABLE IF EXISTS "post_photos";


-- +migrate Up
CREATE TABLE IF NOT EXISTS "post_photos" (
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp,
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "keypath" text NOT NULL,

  PRIMARY KEY ("id")
);

-- +migrate Down
DROP TABLE IF EXISTS "post_photos";

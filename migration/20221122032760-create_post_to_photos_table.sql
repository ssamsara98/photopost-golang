-- +migrate Up
CREATE TABLE IF NOT EXISTS "post_to_photos" (
  "id" uuid DEFAULT uuid_generate_v4() NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "updated_at" timestamp NOT NULL DEFAULT now(),
  "deleted_at" timestamp,
  "position" bigint DEFAULT 0 NOT NULL,
  "post_id" bigserial NOT NULL,
  "photo_id" uuid NOT NULL,

  PRIMARY KEY ("id"),
  FOREIGN KEY ("post_id") REFERENCES posts ("id"),
  FOREIGN KEY ("photo_id") REFERENCES post_photos ("id")
);

-- +migrate Down
DROP TABLE IF EXISTS "post_to_photos";

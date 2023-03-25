
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id bigint NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    deleted_at timestamp,
    email text,
    username text,
    password text,
    name text,
    sex_type text DEFAULT 'Unknown'::text,
    birthdate timestamp,

    PRIMARY KEY ("id"),
    UNIQUE ("email"),
    UNIQUE ("username")
);

CREATE TABLE IF NOT EXISTS post (
  id bigint NOT NULL,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  deleted_at timestamp,
  author_id bigint NOT NULL,
  caption text NOT NULL,
  is_published boolean DEFAULT true NOT NULL,

  PRIMARY KEY ("id"),
  FOREIGN KEY ("author_id") REFERENCES users ("id")
);

CREATE TABLE IF NOT EXISTS post_photos (
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  deleted_at timestamp,
  keypath text NOT NULL,

  PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS post_to_photos (
  id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  deleted_at timestamp,
  "position" bigint DEFAULT 0 NOT NULL,
  post_id bigint NOT NULL,
  photo_id uuid NOT NULL,

  PRIMARY KEY ("id"),
  FOREIGN KEY ("post_id") REFERENCES posts ("id")
  FOREIGN KEY ("photo_id") REFERENCES post_photos ("id")
);

-- +migrate Down

DROP TABLE post_to_photos;

DROP TABLE post_photos;

DROP TABLE posts;

DROP TABLE users;

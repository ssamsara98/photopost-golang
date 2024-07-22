
-- +migrate Up
INSERT INTO
  "posts" ("id", "author_id", "caption", "is_published")
VALUES
  (1, 1, 'Self', true),
  (2, 2, 'Waterfall', true),
  (3, 1, 'Dark Sea', true),
  (4, 2, 'Plant', true),
  (5, 2, 'Landscape', true),
  (6, 1, 'HBD', true),
  (7, 1, 'Flo', true),
  (8, 2, 'Yoru', true),
  (9, 1, 'Rain', true),
  (10, 2, 'Re:Dark - Zero Soul', true),
  (11, 2, 'Hibike', true),
  (12, 1, 'Double 👍', true);

-- +migrate Down
DELETE FROM "posts";

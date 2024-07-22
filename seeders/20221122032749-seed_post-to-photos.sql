
-- +migrate Up
INSERT INTO
  "post_to_photos" ("id", "post_id", "photo_id", "position")
VALUES
  ('01f9a2f4-a50c-45f1-be59-de19fee79bb7', 2, '9da1df71-8667-4b55-afc1-f48a9e16c991', 1),
  ('3c442929-40b0-459c-b8f5-604fa598407e', 5, '0d08c931-a6ca-49f3-b78b-2f242a302a8d', 0),
  ('72a94c86-637a-4e49-b156-e9e766d60a05', 4, 'f497da34-a319-48a6-b293-741517de170f', 0),
  ('7445d0cc-8741-4664-9636-ebd1e64a4ec9', 1, '0b53bea5-2162-4bea-a6d3-2e3a280e8b1e', 0),
  ('7d533e3b-37f4-4904-8fd4-a26cc11d0ae9', 6, '0f2a5228-f832-493d-a50e-d2791daf7e0d', 0),
  ('8b5b2ce7-9b59-40ab-a073-4feae4a74c65', 5, '9a3838b0-a8f6-40cc-97b4-3055967268a0', 1),
  ('a99239d4-2a3d-44dd-94b1-bf1a42a19add', 2, 'de996ceb-6d9d-454b-bc0c-c36211b7eed6', 2),
  ('c2fdbfd6-483c-4baf-ad41-a7f3f9bcd4e4', 3, '633a6360-1422-4e63-a125-3434fb32a267', 0),
  ('d6cd6c87-82c4-4f97-8b25-c2eaf431eea8', 4, '8324caa6-a479-413a-a04f-bf64886fd96d', 1),
  ('e8010f02-90af-446c-9739-d59e9f8759b2', 2, '26b4f604-a960-44db-8a8b-effb15eba7ad', 0),
  ('7b4be2d2-f56d-4446-95e2-5810ecd47230', 7, 'e89a4bd8-08cb-4881-8d34-dbec3421e58f', 0),
  ('d0a6d3a1-1e41-4832-8291-6f59974263f0', 8, '4d5a8013-bba6-4936-a8fd-1f2922513d5c', 0),
  ('2fc2dc23-77a5-4e6f-a385-b7d47fc8e310', 9, '1a969ed1-fa64-4fd3-b928-c98c6cf28d22', 0),
  ('998484f2-3488-4a3a-916f-5886092bb72e', 9, 'd0ea94b3-d11b-4d51-b127-eb547dfd8f9f', 1),
  ('d7c2a346-bf2c-48d2-b40d-9c142fa765db', 10, 'e50ad1bb-44c5-4a10-beee-72e3a8f4807e', 0),
  ('fd639ee0-1ef3-415a-9314-453d7391f42f', 10, '6abe1934-953c-4891-a3d7-c5fcf9e88f13', 1),
  ('8af5d85a-6e0e-4741-8bf5-44375d2b10be', 11, '71f2d239-9963-47e9-9108-e15f6a3a5a94', 0),
  ('c8fb5b25-ebf7-4752-98bd-528475884929', 12, 'd0d38d54-8905-420b-a847-dd94d2db1992', 0);

-- +migrate Down
DELETE FROM "post_to_photos";

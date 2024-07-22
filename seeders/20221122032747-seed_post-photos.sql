
-- +migrate Up
INSERT INTO
  "post_photos" ("id", "keypath")
VALUES
  ('0b53bea5-2162-4bea-a6d3-2e3a280e8b1e', 'photopost/77f0ad89-d9da-42d3-b9bc-d6d0bc572454/8oMe9NPmd2Wk8JtGHgfA9D87ZmQJutHC.jpg'),
  ('0d08c931-a6ca-49f3-b78b-2f242a302a8d', 'photopost/8f44b0ea-08b3-41f1-bf01-ad38fb2146a9/HycM_t9Sx70xjgFq7z0Yneo49AMqWql7.jpg'),
  ('0f2a5228-f832-493d-a50e-d2791daf7e0d', 'photopost/2e5e4973-6cd7-48b4-bd1a-8ac31a8b2a45/VFRay6Pv68n6BwE4JjrfH5mwpXnsL9Ze.jpg'),
  ('26b4f604-a960-44db-8a8b-effb15eba7ad', 'photopost/07973959-14c6-4eb6-b653-e10262f55777/yWFz9jagowcN7xoaXzFIkxHJqCiTu2c-.jpg'),
  ('633a6360-1422-4e63-a125-3434fb32a267', 'photopost/6e193892-40f9-4e07-b100-c05059e826c2/JJ7r3fzeqLMDQ6oV6blrwJtJMPfu1vv0.jpg'),
  ('8324caa6-a479-413a-a04f-bf64886fd96d', 'photopost/66fc1b92-49e2-47bc-8a88-8217d9c02086/BLJZwMEtl0SuFOp9lh62IiPAoITInzQO.jpg'),
  ('9a3838b0-a8f6-40cc-97b4-3055967268a0', 'photopost/ac9609de-611e-4204-b0bd-683e5a3a2419/rkXeCHr4_k_dgUzvJLQv_2l4TWi5kGcO.jpg'),
  ('9da1df71-8667-4b55-afc1-f48a9e16c991', 'photopost/ed3f0a66-37b1-420d-8a02-42f9595f8019/bLM6BhG4-Ab9sc0sifhOLgDzWOfchrau.jpg'),
  ('de996ceb-6d9d-454b-bc0c-c36211b7eed6', 'photopost/5867c461-41f5-4a9b-85ca-d5fea1299cf3/iy2zkv_1exwzU4kd_8rS_i4ioqECWnBt.jpg'),
  ('f497da34-a319-48a6-b293-741517de170f', 'photopost/b6bf1361-794e-4f17-a676-184e937ebc11/FQr-UJuF6xOBTNIGlz3HavI9toIL_z3u.jpg'),
  ('800c7d67-0b0a-492f-a3f0-072aa1dc0f07', 'photopost/a7051460-40e8-4310-b5ab-dc4fb118cdc4/y0S0JjecOhJFcdAoHcv44rAHfYrdT2Fy.jpg'),
  ('d8121df8-3f64-43a8-a3c9-d1d576f248c5', 'photopost/08e2af74-9e9e-4eab-a8b0-57aa9147e707/XmPoVRJGKqGDWt5Kmq4GLf850k_oEjhw.jpg'),
  ('ade641f3-e6a4-46eb-a31c-d97e2bf3bbc5', 'photopost/6cfc637c-9945-4858-890a-183028b23e37/8OGxyPIKBjzGPTDCm5mEpQXh8naJbnAY.jpg'),
  ('e89a4bd8-08cb-4881-8d34-dbec3421e58f', 'photopost/569beb89-1471-4c02-a259-4f498de34e92/ii20dBI_wmEbXmls4qvDz1Rjv6Q1tayJ.jpg'),
  ('4d5a8013-bba6-4936-a8fd-1f2922513d5c', 'photopost/67b6011b-5168-497d-8f29-51e38de5e8c7/-UjEbwuwpVHebhm48k4pOoCiGY83hcQU.jpg'),
  ('d0ea94b3-d11b-4d51-b127-eb547dfd8f9f', 'photopost/a68a6822-a91e-4b4b-9398-531a078571ad/QIr-ZqjzLpEWi8TelYO88e9UcG_amW79.jpg'),
  ('1a969ed1-fa64-4fd3-b928-c98c6cf28d22', 'photopost/490e1529-674c-43e1-af17-af3071c600ec/HYsBei8bwS0Z7zxOdpjPEZVM8emTnNGC.png'),
  ('e50ad1bb-44c5-4a10-beee-72e3a8f4807e', 'photopost/81585c6a-782a-4311-abe9-09b26129961f/-cTPGKtNwwL7k4fNQ5QmODbqJ2WhAg1B.jpg'),
  ('6abe1934-953c-4891-a3d7-c5fcf9e88f13', 'photopost/67a98e45-170e-4c56-9782-c574a3bf9168/uwXSLpIDv6fbHc8uSR3bYNxXYoEwaPQh.jpg'),
  ('71f2d239-9963-47e9-9108-e15f6a3a5a94', 'photopost/530b2a41-688d-4c94-b9a7-991c6d703b52/ars0S_nAtNH9Ggt7O-MOkI8OON1O69bs.png'),
  ('d0d38d54-8905-420b-a847-dd94d2db1992', 'photopost/9ac48dd6-1e0a-4b43-a4f0-14e013585995/p2qxOU6vvOOHdh6uv3SZ3mfitFdf70hW.jpg');

-- +migrate Down
DELETE FROM "post_photos";

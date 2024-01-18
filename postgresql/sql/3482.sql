-- Probrably URI description is wrong, and the following schema is wrong because user_id is actually id
CREATE TABLE users (
    user_id NUMERIC PRIMARY KEY,
    user_name VARCHAR(100),
    posts NUMERIC
);

CREATE TABLE followers (
    follower_id NUMERIC PRIMARY KEY,
    user_id_fk NUMERIC REFERENCES users(user_id),
    following_user_id_fk NUMERIC REFERENCES users(user_id)
);

INSERT INTO users (user_id, user_name, posts) VALUES
(1, 'Francisco', 23),
(2, 'Brenda', 51),
(3, 'Helena', 12),
(4, 'Miguel', 70),
(5, 'Laura', 55),
(6, 'Arthur', 2),
(7, 'Alice', 3);

INSERT INTO followers (follower_id, user_id_fk, following_user_id_fk) VALUES
(1, 1, 5),
(2, 2, 4),
(3, 3, 7),
(4, 3, 6),
(5, 4, 2),
(6, 4, 5),
(7, 5, 1),
(8, 5, 3),
(9, 5, 4),
(10, 5, 2),
(11, 7, 3);

-- Same schema but using id instead user_id
CREATE TABLE users (
    id NUMERIC PRIMARY KEY,
    user_name VARCHAR(100),
    posts NUMERIC
);

CREATE TABLE followers (
    follower_id NUMERIC PRIMARY KEY,
    user_id_fk NUMERIC REFERENCES users(id),
    following_user_id_fk NUMERIC REFERENCES users(id)
);

INSERT INTO users (id, user_name, posts) VALUES
(1, 'Francisco', 23),
(2, 'Brenda', 51),
(3, 'Helena', 12),
(4, 'Miguel', 70),
(5, 'Laura', 55),
(6, 'Arthur', 2),
(7, 'Alice', 3);

INSERT INTO followers (follower_id, user_id_fk, following_user_id_fk) VALUES
(1, 1, 5),
(2, 2, 4),
(3, 3, 7),
(4, 3, 6),
(5, 4, 2),
(6, 4, 5),
(7, 5, 1),
(8, 5, 3),
(9, 5, 4),
(10, 5, 2),
(11, 7, 3);

-- DROP TABLE followers;
-- DROP TABLE users;

-- Query using user_id
SELECT 
  s.u1_name AS u1_name, 
  s.u2_name AS u2_name
FROM (
  SELECT
    CASE WHEN u1.posts <= u2.posts THEN u1.user_id ELSE u2.user_id END AS u1_id,
    CASE WHEN u1.posts <= u2.posts THEN u1.user_name ELSE u2.user_name END AS u1_name,
    CASE WHEN u1.posts > u2.posts THEN u1.user_name ELSE u2.user_name END AS u2_name
  FROM (
    SELECT 
      f1.user_id_fk AS user_id_1, 
      f1.following_user_id_fk AS user_id_2
    FROM followers f1
    JOIN followers f2 ON f1.user_id_fk = f2.following_user_id_fk
                     AND f1.following_user_id_fk = f2.user_id_fk
  ) AS mutual_follow
  JOIN users u1 
    ON mutual_follow.user_id_1 = u1.user_id
  JOIN users u2 
    ON mutual_follow.user_id_2 = u2.user_id
  WHERE u1.user_id < u2.user_id
  ORDER BY u1_id
) s;

-- Same query but using id instead user_id
SELECT 
  s.u1_name AS u1_name, 
  s.u2_name AS u2_name
FROM (
  SELECT
    CASE WHEN u1.posts <= u2.posts THEN u1.id ELSE u2.id END AS u1_id,
    CASE WHEN u1.posts <= u2.posts THEN u1.user_name ELSE u2.user_name END AS u1_name,
    CASE WHEN u1.posts > u2.posts THEN u1.user_name ELSE u2.user_name END AS u2_name
  FROM (
    SELECT 
      f1.user_id_fk AS user_id_1, 
      f1.following_user_id_fk AS user_id_2
    FROM followers f1
    JOIN followers f2 ON f1.user_id_fk = f2.following_user_id_fk
                     AND f1.following_user_id_fk = f2.user_id_fk
  ) AS mutual_follow
  JOIN users u1 
    ON mutual_follow.user_id_1 = u1.id
  JOIN users u2 
    ON mutual_follow.user_id_2 = u2.id
  WHERE u1.id < u2.id
  ORDER BY u1_id
) s;

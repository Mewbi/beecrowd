CREATE TABLE nodes (
    node_id NUMERIC,
    pointer NUMERIC
);

INSERT INTO nodes (node_id, pointer) VALUES
(50, 30),
(50, 75),
(30, 15),
(30, 35),
(15, 1),
(15, 20),
(35, 32),
(35, 40),
(1, NULL),
(20, NULL),
(32, NULL),
(40, NULL),
(75, 60),
(75, 90),
(60, 55),
(60, 70),
(90, 80),
(90, 95),
(55, NULL),
(70, NULL),
(80, NULL),
(95, NULL);

-- DROP TABLE nodes;

SELECT 
  DISTINCT node_id,
  CASE
    WHEN ( pointer IS NULL) THEN 'LEAF'
    WHEN ( EXISTS(
      SELECT 1
      FROM nodes n2
      WHERE n2.pointer = n1.node_id
    ) ) THEN 'INNER'
    ELSE 'ROOT'
  END AS type
  FROM nodes n1
ORDER BY node_id ASC;

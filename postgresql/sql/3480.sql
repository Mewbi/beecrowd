CREATE TABLE chairs (
    id SERIAL PRIMARY KEY,
    queue NUMERIC,
    available BOOLEAN
);

INSERT INTO chairs (queue, available) VALUES
(1, FALSE),
(1, FALSE),
(1, TRUE),
(1, FALSE),
(1, FALSE),
(1, FALSE),
(1, TRUE),
(1, TRUE),
(2, TRUE),
(2, FALSE),
(2, TRUE),
(2, TRUE),
(2, FALSE),
(2, TRUE),
(2, TRUE),
(2, FALSE),
(3, TRUE),
(3, FALSE),
(3, TRUE),
(3, FALSE),
(3, TRUE),
(3, TRUE),
(3, FALSE),
(3, FALSE),
(4, TRUE),
(4, FALSE),
(4, FALSE),
(4, TRUE),
(4, TRUE),
(4, FALSE),
(4, FALSE),
(4, TRUE);

-- DROP TABLE chairs;

SELECT *
FROM (
    SELECT *,
           LEAD(available) OVER (ORDER BY id) AS next_available,
           LEAD(queue) OVER (ORDER BY id) AS next_queue
    FROM chairs
) AS subquery
WHERE available = next_available AND available = TRUE AND queue = next_queue;

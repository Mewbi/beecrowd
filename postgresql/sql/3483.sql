CREATE TABLE cities (
    id NUMERIC,
    city_name VARCHAR(50),
    population NUMERIC
);

INSERT INTO cities (id, city_name, population) VALUES
(1, 'São Paulo', 12396372),
(2, 'Rio de Janeiro', 6775561),
(3, 'Brasília', 3094325),
(4, 'Salvador', 2900319),
(5, 'Fortaleza', 2703391),
(6, 'Belo Horizonte', 2530701),
(7, 'Manaus', 2255903),
(8, 'Curitiba', 1963726),
(9, 'Recife', 1661017),
(10, 'Goiânia', 1555626),
(11, 'Belém', 1506420),
(12, 'Porto Alegre', 1492530);

-- DROP TABLE cities;

WITH without_top AS (
  SELECT 
    population
  FROM 
    cities 
  WHERE 
    population != (SELECT MAX(population) FROM cities) AND
    population != (SELECT MIN(population) FROM cities)
)

SELECT 
  city_name, 
  population 
FROM 
  cities 
WHERE 
  population = (SELECT MAX(population) FROM without_top) OR
  population = (SELECT MIN(population) FROM without_top)
ORDER BY population DESC;

--- URI Online Judge SQL
--- Copyright URI Online Judge
--- www.urionlinejudge.com.br
--- Problem 2988

CREATE TABLE teams (
    id integer PRIMARY KEY,
    name varchar(50)
);

GRANT SELECT ON teams TO sql_user;

CREATE TABLE matches  (
    id integer PRIMARY KEY,
    team_1 integer,
    team_2 integer,
    team_1_goals integer,
    team_2_goals integer,
    FOREIGN KEY (team_1) REFERENCES teams(id),
    FOREIGN KEY (team_2) REFERENCES teams(id)
);

GRANT SELECT ON matches TO sql_user;insert into teams
    (id, name)
values
    (1,'CEARA'),
    (2,'FORTALEZA'),
    (3,'GUARANY DE SOBRAL'),
    (4,'FLORESTA');

insert into  matches
    (id, team_1, team_2, team_1_goals, team_2_goals)
values
    (1,4,1,0,4),
    (2,3,2,0,1),
    (3,1,3,3,0),
    (4,3,4,0,1),
    (5,1,2,0,0),
    (6,2,4,2,1);

/* Execute this query to drop the table */
-- DROP TABLE matches;

SELECT
    t.name AS name,
    COUNT(m.id) AS matches,
    SUM(CASE WHEN m.team_1_goals > m.team_2_goals AND m.team_1 = t.id THEN 1 ELSE 0 END) +
    SUM(CASE WHEN m.team_1_goals < m.team_2_goals AND m.team_2 = t.id THEN 1 ELSE 0 END) AS victories,
    SUM(CASE WHEN m.team_1_goals < m.team_2_goals AND m.team_1 = t.id THEN 1 ELSE 0 END) +
    SUM(CASE WHEN m.team_1_goals > m.team_2_goals AND m.team_2 = t.id THEN 1 ELSE 0 END) AS defeats,
    SUM(CASE WHEN m.team_1_goals = m.team_2_goals AND (m.team_1 = t.id OR m.team_2 = t.id) THEN 1 ELSE 0 END) AS draws,
    SUM((CASE WHEN m.team_1_goals > m.team_2_goals AND m.team_1 = t.id THEN 1 ELSE 0 END +
         CASE WHEN m.team_1_goals < m.team_2_goals AND m.team_2 = t.id THEN 1 ELSE 0 END) * 3 +
        CASE WHEN m.team_1_goals = m.team_2_goals AND (m.team_1 = t.id OR m.team_2 = t.id) THEN 1 ELSE 0 END) AS score
FROM
    teams t
LEFT JOIN
    matches m ON t.id = m.team_1 OR t.id = m.team_2
GROUP BY
    t.name
ORDER BY
    score DESC;

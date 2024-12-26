# Write your MySQL query statement below
WITH AvgFeb AS (
    SELECT
        movie_id,
        AVG(rating) AS average_rating
    FROM MovieRating
    WHERE 
        '2020-02-1' <= created_at AND
        created_at < '2020-03-01'
    GROUP BY movie_id
),
HiAvgMov AS (
    SELECT m.title
    FROM AvgFeb af
    INNER JOIN Movies m
    ON af.movie_id = m.movie_id
    WHERE average_rating IN (
        SELECT MAX(average_rating)
        FROM AvgFeb
    )
    ORDER BY m.title
    LIMIT 1
),
UserTotalRated AS (
    SELECT
        u.name,
        COUNT(mr.user_id) AS total_rated
    FROM MovieRating mr
    INNER JOIN Users u
    ON mr.user_id = u.user_id
    GROUP BY mr.user_id
),
MostTotalRatedUser AS (
    SELECT name
    FROM UserTotalRated
    WHERE total_rated IN (
        SELECT MAX(total_rated)
        FROM UserTotalRated
    )
    ORDER BY name
    LIMIT 1
)

SELECT name AS results
FROM MostTotalRatedUser 
UNION ALL
SELECT title
FROM HiAvgMov
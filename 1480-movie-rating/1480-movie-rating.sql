# Write your MySQL query statement below
# less 'WHERE'
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
RankedMovie AS (
    SELECT
        af.movie_id,
        m.title,
        af.average_rating,
        RANK() OVER (ORDER BY af.average_rating DESC, m.title) AS rank_movie
    FROM AvgFeb af
    INNER JOIN Movies m
    ON af.movie_id = m.movie_id
),
RankedUserTotalRated AS (
    SELECT
        u.name,
        COUNT(mr.movie_id) AS total_rated,
        RANK() OVER (ORDER BY COUNT(mr.movie_id) DESC, u.name) AS rank_user
    FROM MovieRating mr
    INNER JOIN Users u
    ON mr.user_id = u.user_id
    GROUP BY mr.user_id
)

SELECT name AS results 
FROM RankedUserTotalRated 
WHERE rank_user = 1
UNION ALL
SELECT title 
FROM RankedMovie 
WHERE rank_movie = 1
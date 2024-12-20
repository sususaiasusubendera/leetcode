# Write your MySQL query statement below
# authentic idea
WITH CumulativeWeight AS (
    SELECT
        person_id,
        person_name,
        weight,
        SUM(weight) OVER (ORDER BY turn) AS cumulative_weight,
        ROW_NUMBER() OVER (ORDER BY turn) AS row_num
    FROM Queue
)

SELECT person_name
FROM CumulativeWeight
WHERE cumulative_weight <= 1000
ORDER BY row_num DESC
LIMIT 1
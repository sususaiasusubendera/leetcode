# Write your MySQL query statement below
# MySQL variable
SELECT person_name
FROM (
    SELECT
        person_name,
        @cumulative_weight := @cumulative_weight + weight AS cumulative_weight
    FROM 
        Queue, 
        (SELECT @cumulative_weight := 0) init
    ORDER BY turn
) cumulative
WHERE cumulative_weight <= 1000
ORDER BY cumulative_weight DESC
LIMIT 1
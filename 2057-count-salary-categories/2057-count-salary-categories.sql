# Write your MySQL query statement below
# CTE, join, union
WITH CategoryCount AS (
    SELECT
        CASE
            WHEN income < 20000 THEN "Low Salary"
            WHEN 20000 <= income AND income <= 50000 THEN "Average Salary"
            ELSE "High Salary"
        END AS category,
        COUNT(*) as accounts_count
    FROM Accounts
    GROUP BY category
)

SELECT
    c.category,
    IFNULL(cc.accounts_count, 0) AS accounts_count
FROM (
    SELECT "Low Salary" AS category
    UNION ALL
    SELECT "Average Salary"
    UNION ALL
    SELECT "High Salary"
) c
LEFT JOIN CategoryCount cc
ON c.category = cc.category


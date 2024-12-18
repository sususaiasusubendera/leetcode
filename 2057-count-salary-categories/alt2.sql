# Write your MySQL query statement below
# 3 select 2 union all
SELECT 
    "Low Salary" AS category,
    COUNT(IF(income < 20000, 1, NULL)) AS accounts_count
FROM Accounts
UNION ALL
SELECT
    "Average Salary",
    COUNT(IF(20000 <= income AND income <= 50000, 1, NULL))
FROM Accounts
UNION ALL
SELECT
    "High Salary",
    COUNT(IF(income > 50000, 1, NULL))
FROM Accounts

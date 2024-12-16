# Write your MySQL query statement below
# 'lead' solution
SELECT DISTINCT num AS ConsecutiveNums
FROM (
    SELECT
        num,
        LEAD(num) OVER (ORDER BY id) AS next_num1,
        LEAD(num, 2) OVER (ORDER BY id) AS next_num2
    FROM Logs
) temp
WHERE num = next_num1 AND num = next_num2



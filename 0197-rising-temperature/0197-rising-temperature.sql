# Write your MySQL query statement below
SELECT n.id FROM Weather m INNER JOIN Weather n ON m.recordDate = DATE_SUB(n.recordDate, INTERVAL 1 DAY)
WHERE m.temperature < n.temperature;
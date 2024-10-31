# Write your MySQL query statement below
SELECT t.id FROM Weather y INNER JOIN Weather t ON y.recordDate = DATE_SUB(t.recordDate, INTERVAL 1 DAY)
WHERE y.temperature < t.temperature;
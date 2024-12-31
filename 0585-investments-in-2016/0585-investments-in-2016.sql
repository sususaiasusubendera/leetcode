# Write your MySQL query statement below
WITH OnePolicyholderOnecity AS (
    SELECT 
        pid,
        CONCAT(lat, lon) AS ConcatLatLon,
        COUNT(CONCAT(lat, lon)) AS CountPolicyholder
    FROM Insurance
    GROUP BY ConcatLatLon
    HAVING CountPolicyholder = 1
),
SameTIV2015PID AS (
    SELECT pid
    FROM Insurance
    WHERE tiv_2015 IN (
        SELECT
            tiv_2015
        FROM Insurance
        GROUP BY tiv_2015
        HAVING COUNT(tiv_2015) > 1
    )
)

SELECT ROUND(SUM(tiv_2016), 2) AS tiv_2016
FROM Insurance
WHERE
    pid IN (SELECT pid FROM OnePolicyholderOnecity) AND
    pid IN (SELECT pid FROM SameTIV2015PID)
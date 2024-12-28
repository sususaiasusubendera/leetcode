# Write your MySQL query statement below
WITH RequesterAccepterAll AS (
    SELECT requester_id AS id
    FROM RequestAccepted
    UNION ALL
    SELECT accepter_id
    FROM RequestAccepted
),
CountID AS (
    SELECT
        id,
        COUNT(ID) AS num
    FROM RequesterAccepterAll
    GROUP BY id
)   

SELECT
    id,
    num
FROM CountID
WHERE num IN (
    SELECT MAX(num)
    FROM CountID
)

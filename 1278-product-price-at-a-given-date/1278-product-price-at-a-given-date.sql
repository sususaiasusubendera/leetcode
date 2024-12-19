# Write your MySQL query statement below
# union before and after 2019-08-16
SELECT DISTINCT
    product_id,
    10 AS price
FROM Products
WHERE product_id NOT IN (
    SELECT DISTINCT product_id
    FROM Products
    WHERE change_date <= "2019-08-16"
)
UNION ALL
SELECT
    product_id,
    new_price AS price
FROM Products
WHERE (product_id, change_date) IN (
    SELECT
        product_id,
        MAX(change_date) AS latest
    FROM Products
    WHERE change_date <= "2019-08-16"
    GROUP BY product_id
)
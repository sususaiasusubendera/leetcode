# Write your MySQL query statement below
# window function
WITH Latest AS (
    SELECT
        product_id,
        new_price,
        change_date
    FROM Products
    WHERE change_date <= "2019-08-16"
),
RankedPrice AS (
    SELECT
        product_id,
        new_price,
        change_date,
        ROW_NUMBER() OVER (PARTITION BY product_id ORDER BY change_date DESC) AS row_num
    FROM Latest
)

SELECT
    p.product_id,
    IFNULL(rp.new_price, 10) AS price
FROM (SELECT DISTINCT product_id FROM Products) p
LEFT JOIN RankedPrice rp
ON p.product_id = rp.product_id AND rp.row_num = 1
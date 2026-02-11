-- Last updated: 2/11/2026, 8:30:08 PM
-- Write your PostgreSQL query statement below
select product_id
from Products
where low_fats = 'Y' and recyclable = 'Y';
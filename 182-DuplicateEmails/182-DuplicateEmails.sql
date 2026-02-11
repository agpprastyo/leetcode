-- Last updated: 2/11/2026, 8:30:23 PM
# Write your MySQL query statement below
SELECT email FROM person GROUP BY email HAVING COUNT(email) > 1;

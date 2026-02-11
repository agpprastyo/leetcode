-- Last updated: 2/11/2026, 8:30:25 PM
-- Write your PostgreSQL query statement below
SELECT 
    P.firstName,
    P.lastName,
    A.city,
    A.state
FROM 
    Person P
LEFT JOIN 
    Address A
ON 
    P.personId = A.personId;

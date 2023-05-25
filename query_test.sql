CREATE TABLE logic_test (
    id INT PRIMARY KEY,
    name VARCHAR(25),
    parent_id INT
);

select * from logic_test;
INSERT INTO logic_test (id, name, parent_id)
VALUES (1, 'Zaki', 2), 
       (2, 'Ilham', null),
       (3, 'Irwan', 2),
       (4, 'Arka', 3);
 
SELECT t1.id, t1.name, t2.name AS parent_name
FROM logic_test t1
LEFT JOIN logic_test t2 ON t1.parent_id = t2.id;
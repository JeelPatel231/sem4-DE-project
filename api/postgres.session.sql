--@block
-- create table which holds uuid and name of restaurant
-- primary key being uuid because same names could occur
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE UUIDidentify(
    id UUID DEFAULT uuid_generate_v4 (),
    name VARCHAR(50),
    PRIMARY KEY (id)
);

-- this statement makes another table which
-- references uuid table primary key and holds
-- layout of every restaurant
CREATE TABLE restaurant (
    id UUID UNIQUE NOT NULL,
    layout TEXT,
    FOREIGN KEY(id) REFERENCES UUIDidentify(id)
);

--@block
DELETE FROM uuididentify
where id='9bf30c27-d822-4d2f-b5c2-6fd43e8abe40';
DELETE FROM restaurant
where id='9bf30c27-d822-4d2f-b5c2-6fd43e8abe40';

--@block
INSERT INTO uuididentify (name)
VALUES ('NAME') --NAME to be replace with restaurant's name
RETURNING id;


--@block
-- insert layout of restaurant using uuid
INSERT into restaurant (id, layout)
VALUES (
    'baea7561-6fe6-4bd1-8b99-aa9d28f274e5', -- UUID OF RESTAURANT
    '[[1,3,1],[2,2,1]]' -- LAYOUT OF RESTAURANT
);

--@block
-- join tables to get name and layout
select * from restaurant
inner join uuididentify using (id);

--@block
-- join tables to get name and layout with null values
select * from restaurant
full join uuididentify using (id);

-- DROP STATEMENTS // dangerous xD
--@block
DROP TABLE uuididentify;
--@block
DROP TABLE restaurant;

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
INSERT INTO uuididentify (name)
VALUES ('NAME') --NAME to be replace with restaurant's name
RETURNING id;


--@block
-- insert layout of restaurant using uuid
INSERT into restaurant (id, layout)
VALUES (
    '90f9fddb-4480-4ab9-9c37-f0b816b66770',
    '[[1,3,1],[2,2,1]]'
);

--@block
-- join tables to get name and layout
select * from restaurant
inner join uuididentify using (id);



-- DROP STATEMENTS // dangerous xD
--@block
DROP TABLE uuididentify;
--@block
DROP TABLE restaurant;

CREATE TABLE persons (
                       document_id varchar(12) unique not null,
                       name varchar(100) not null,
                       phone varchar(11) unique not null
);
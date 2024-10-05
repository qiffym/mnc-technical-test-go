create table users (
    id serial not null primary key,
    name varchar(100) not null,
    email varchar(100) not null unique,
    password varchar(255) not null,
    role varchar(100)
);


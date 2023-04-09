CREATE TABLE users
(
    id serial not null unique,
    "email" varchar(255) not null unique,
    "name" varchar(255) not null,
    lastname varchar(255) not null ,
    password_hash text not null
)
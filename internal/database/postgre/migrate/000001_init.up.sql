CREATE TABLE users
(
    id SERIAL PRIMARY KEY not null unique,
    "email" varchar(255) not null unique,
    "name" varchar(255) not null,
    lastname varchar(255) not null,
    password_hash text not null
);

CREATE TABLE books
(
    id SERIAL PRIMARY KEY not null unique,
    "name" varchar(255) not null,
    price int    
);

CREATE TABLE rent_list
(
    id        BIGSERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users,
    book_id   INTEGER NOT NULL REFERENCES rent_list,
    UNIQUE (user_id, book_id)   
)


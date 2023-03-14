create table users (
    user_id bigserial not null primary key,
    username varchar(30) not null unique,
    email varchar(60) not null unique,
    encrypted_password varchar not null
)
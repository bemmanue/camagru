create table comments (
                          id              bigserial not null primary key,
                          author_id       bigserial not null,
                          post_id         bigserial not null,
                          comment_text    varchar not null,
                          creation_time   timestamp with time zone not null
);

create table images (
                        id              bigserial not null primary key,
                        name            varchar not null,
                        extension       varchar not null,
                        path            varchar not null,
                        upload_time     timestamp with time zone not null,
                        user_id         bigserial not null
);

create table likes (
                       id              bigserial not null primary key,
                       post_id         bigserial not null,
                       user_id         bigserial not null
);

create table posts (
                       id               bigserial not null primary key,
                       image_id         bigserial not null,
                       author_id        bigserial not null,
                       creation_time    timestamp with time zone not null
);

create table users (
                       id                  bigserial not null primary key,
                       username            varchar not null,
                       email               varchar not null,
                       encrypted_password  varchar not null,
                       email_verified      bool not null default false,
                       like_notify         bool not null default true,
                       comment_notify      bool not null default true

);

create table verify_codes (
                              id              bigserial not null primary key,
                              email           varchar not null,
                              code            int not null,
                              user_id         bigserial not null
);

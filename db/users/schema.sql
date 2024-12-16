create table users
(
    id         uuid          not null primary key,
    email      varchar(256)  not null,
    password   varchar(2048) not null,
    salt       varchar(256)  not null,
    created_at timestamp     not null default now(),
    updated_at timestamp     not null default now()
)
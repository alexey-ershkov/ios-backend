CREATE EXTENSION IF NOT EXISTS CITEXT;
CREATE TABLE If Not Exists users
(
    UserID   SERIAL PRIMARY KEY,
    Nickname text,
    Email    text UNIQUE,
    Password text,
    Photo    text
);

create table if not exists currency_info
(
    id          bigserial primary key,
    name        citext collate "C" not null,
    symbol  citext collate "C" not null,
    cmc_id int not null ,
    rank int not null ,
    logo citext collate "C" not null,
    date_added citext collate "C" not null,
    category citext collate "C" not null,
    description citext collate "C" not null,
    platform_cmc_id int,
    platform_name citext collate "C",
    platform_symbol citext collate "C",
    platform_token_address citext collate "C",
    website citext collate "C",
    doc citext collate "C",
    twitter citext collate "C",
    reddit citext collate "C",
    source_code citext collate "C"
);

create table if not exists crypto_cost_in_usd
(
    id           bigserial primary key,
    curr_info_fk bigint                    not null
        constraint currency_info_fk
            references currency_info
            on update cascade on delete cascade,
    cost         double precision          not null,
    date         date default CURRENT_DATE not null
);


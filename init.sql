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
    stock_name  citext collate "C" not null,
    description citext collate "C" not null
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

insert into currency_info (name, stock_name, description) values ('Bitcoin', 'BTC', 'Bitcoin desc');
insert into currency_info (name, stock_name, description) VALUES ('Ethereum', 'ETH', 'Ethereum desc');
insert into currency_info (name, stock_name, description) VALUES ('Waves', 'WVS', 'Waves desc');


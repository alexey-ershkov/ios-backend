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
    cmc_id int primary key ,
    name        citext collate "C" not null,
    symbol  citext collate "C" not null,
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

create table fiat_info
(
    cmc_fiat_id integer not null
        constraint fiat_info_pk
            primary key,
    name        citext  not null,
    sign        citext  not null,
    symbol      citext  not null
);

create table curr_crypto_info
(
    cmc_id             integer          not null
        constraint current_crypto_info_pk
            primary key,
    max                double precision,
    in_market          double precision,
    mined              double precision,
    last_updated       citext           not null,
    percent_change_1h  double precision not null,
    percent_change_24h double precision not null,
    percent_change_7d  double precision not null
);

create table curr_crypto_info_in_fiat
(
    fiat_id      integer          not null
        constraint curr_crypto_info_in_fiat_fiat_info_cmc_fiat_id_fk
            references fiat_info
            on update cascade on delete cascade,
    cmc_id       integer          not null
        constraint curr_crypto_info_in_fiat_curr_crypto_info_cmc_id_fk
            references curr_crypto_info
            on update cascade on delete cascade
        constraint curr_crypto_info_in_fiat_currency_info_cmc_id_fk
            references currency_info
            on update cascade on delete cascade,
    price        double precision not null,
    volume       double precision,
    last_updated citext           not null,
    constraint curr_crypto_info_in_fiat_pk
        primary key (fiat_id, cmc_id)
);









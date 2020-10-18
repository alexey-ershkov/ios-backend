CREATE EXTENSION IF NOT EXISTS CITEXT;
CREATE TABLE If Not Exists users
(
    UserID   SERIAL PRIMARY KEY,
    Nickname text,
    Email    text UNIQUE,
    Password text,
    Photo    text
);



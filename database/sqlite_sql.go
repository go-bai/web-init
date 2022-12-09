package database

var createTableSQL = `
create table if not exists user (
    id integer primary key,
    username text not null unique,
    password text not null,
    created text,
    updated text,
    deleted text
);
`

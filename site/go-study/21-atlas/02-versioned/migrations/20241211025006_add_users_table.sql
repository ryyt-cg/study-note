create table users (
    id integer primary key,
    name text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null
);
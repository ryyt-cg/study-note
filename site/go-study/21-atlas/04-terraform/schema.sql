create table users(
    id integer primary key,
    name varchar(255),
    bio text,
    description text
);

create table blog_posts (
    id integer primary key,
    user_id integer not null,
    title varchar(255),
    body text,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null,
    foreign key (user_id) references users(id)
);

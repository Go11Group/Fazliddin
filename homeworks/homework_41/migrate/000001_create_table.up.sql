create table users(
    id uuid default gen_random_uuid() primary key,
    first_name varchar,
    laset_name varchar,
    age int,
    email varchar
);
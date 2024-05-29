create table users(id int serial primary key, name varchar, age int)

create table orders(id uuid default gen_rand_uuid(), name varchar, user_id references users(id))

insert into users(name, age) values('BittaBola', 17);

insert into orders(name, user_id) values('bitta narse', 2);

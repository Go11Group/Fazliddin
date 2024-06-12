drop table users;

drop table cars;

create table cars(
    id uuid default gen_random_uuid() primary key, 
    name varchar(50), 
    year int, 
    brend varchar(50)
);

create table owners(
    id_owner uuid default gen_random_uuid() primary key, 
    name varchar(50), 
    age int
);

alter table cars rename column id to car_id;

alter table owners rename column id_owner to owner_id;

create table car_owner(
    car_id uuid references cars(car_id),
    owner_id uuid references owners(owner_id),
    constraint car_owner_id primary key (car_id, owner_id)
);

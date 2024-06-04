create table persons(
    id uuid default gen_random_uuid() primary key,
    first_name varchar,
    last_name varchar,
    email varchar,
    age int,
    is_maried bool
);

explain(enelyse)
select * from persons where id = 'fffb58e2-4e0d-4698-b1b5-411355b44aa2';


create index persons_indx on persons(id);

explain(enelyse)
select * from persons where id = 'fffb58e2-4e0d-4698-b1b5-411355b44aa2';

drop index persons_indx;


create index persons_indx on persons using hash (id);

explain(enelyse)
select * from persons where id = 'fffb58e2-4e0d-4698-b1b5-411355b44aa2'
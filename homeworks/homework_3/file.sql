create database dars_db

\c dars_db;

create table students(id serial primary key, name varchar, age int, phone_number varchar(13), email varchar);

\d students

insert into students(name, age, phone_number, email) values('Ahmat', 18, '+998999635522', 'ahmat@gmail.com');
insert into students(name, age, phone_number, email) values('Mehmet', 22, '+998999634455', 'mehmet@gmail.com'), ('Abdulloh', 15, '+998886355272', 'abdulloh@gmail.com');
insert into students(name, age, phone_number, email) values('Bittabola', 20, '+998999634455', 'bittabole@gmail.com'), ('Nmadur', 31, '+998858355272', 'nmadur@gmail.com')('kimdur', 52, '+998955635522', 'kimdur@gmail.com');
insert into students(name, age, phone_number, email) values('Bittabola', 20, '+998999634455', 'bittabole@gmail.com'), ('Nmadur', 31, '+998858355272', 'nmadur@gmail.com'),('kimdur', 52, '+998955635522', 'kimdur@gmail.com');
insert into students(name, age, phone_number, email) values('Mehmet', 22, '+998999634455', 'mehmet@gmail.com');
insert into students(name, age, phone_number, email) values('Avaz', 17, '+998999159455', 'avaz@gmail.com');

update students set age = 25 where age > 30;

select * from students;

delete from students where name = 'Nmadur';

select * from students;

select age, count(*) same_age from students group by age;

select * from students order by age;

select * from students order by age desc;

select * from students inner join students on students.id = students.age;

create table marks(id serial primary key, mark int, name_student varchar, id_students int);

insert into marks(mark, id_students) values(2, 1), (2, 2), (5, 3), (4, 4), (3, 8), (5, 7);

select * from students inner join marks on students.id = marks.id_students;

select * from students left join marks on students.id = marks.id_students;

select * from students right join marks on students.id = marks.id_students;

select * from students full outer join marks on students.id = marks.id_students;

select * from students right join marks on students.id = marks.id_students;

select * from students cross join marks;
--cross ishlatsa bomasakan galatikan
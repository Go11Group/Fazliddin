create table Student(
    id uuid default gen_random_uuid() primary key,
    name varchar,
    age int
);

create table Course(
    id uuid default gen_random_uuid() primary key,
    name varchar
);

create table student_coure(
    id uuid default gen_random_uuid() primary key,
    studentId uuid references Student(id),
    courseId uuid references Course(id)
);

create table grade(
    srudent_courseId uuid references student_course(id),
    grade int
)
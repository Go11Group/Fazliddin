--1--
    -- create qivotganimda shunnaqa qib ketganman

--2--
select c.name, best_student.name from course c
left join(
    select s.id, s.name, sc.courseid, max(g.grade) from student s
    left join student_course sc on sc.studentid = s.id
    left join grade g on g.student_courseid = sc.id
    group by s.id, sc.courseid
) best_student on best_student.courseid = c.id
group by c.id, best_student.name


--3--
select c.name, avg(g.grade) from course c
left join student_course sc on sc.courseid = c.id
join grade g on sc.id = g.student_courseid
group by c.name;


--4--
select c.name, min(s.age) from course c
left join student_course sc on sc.courseid = c.id
left join student s on s.id = sc.studentid
group by c.name;

--5--
select c.name, max(avg_grade) from course c
left join (
    select sc.courseid, avg(g.grade) as avg_grade from student_course sc
    join grade g on sc.id = g.student_courseid
    group by sc.id
) avgg on c.id = avgg.courseid
group by c.name;
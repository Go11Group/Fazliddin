CREATE TABLE "branch" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" VARCHAR(30) NOT NULL
);

CREATE TABLE "teacher" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "branch_id" UUID REFERENCES "branch" ("id"),
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "salary" NUMERIC
);

CREATE TABLE "asisstant_teacher" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "branch_id" UUID REFERENCES "branch" ("id"),
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "salary" NUMERIC
);


CREATE TABLE "course" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "name" VARCHAR(30) NOT NULL,
    "teacher_id" UUID REFERENCES "teacher" ("id"),
    "asisstant_teacher_id" UUID REFERENCES "asisstant_teacher" ("id"),
    "price" NUMERIC
);

CREATE TABLE "student" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "course_id" UUID REFERENCES "course" ("id"),
    "payed_sum" NUMERIC
);





INSERT INTO course (name, teacher_id, asisstant_teacher_id, price) VALUES
('Course 1', '895c2279-1931-4f2f-a8b1-ab185082659c', '0c963693-a26d-4df8-9e7e-c8b5185de22b', 100),
('Course 2', 'fb62e0b5-0f76-4297-ac54-801a44428cbc', '4ed6ed97-729f-41b3-998a-0b724a574d7e', 120),
('Course 3', '00203194-9380-42d3-9249-f0443d1c16f2', 'f51fcc8a-901c-437b-8024-1ad59b9058bc', 150);
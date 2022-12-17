CREATE TABLE public.users
(
    id serial,
    first_name text,
    last_name text,
    email text NOT NULL,
	preferred_time integer,
    PRIMARY KEY (id),
    CONSTRAINT email UNIQUE (email)
);

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;



CREATE TABLE public.courses
(
    id serial,
    name text NOT NULL,
    area text,
    course_id text NOT NULL,
    CONSTRAINT course_id UNIQUE (course_id),
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.courses
    OWNER to postgres;


CREATE TABLE public.requests
(
    id serial,
    course_id integer,
    objective text,
    description text,
    method text,
    meeting_day_of_week integer,
    meeting_time integer,
    contact_method integer,
	created_by text NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT course_id FOREIGN KEY (course_id)
        REFERENCES public.courses (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER TABLE IF EXISTS public.requests
    OWNER to postgres;


CREATE TABLE public.user_to_course
(
    id serial,
    user_id serial,
    course_id serial,
    grade numeric(5, 2) NOT NULL,
    completed_on date NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT user_id FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT course_id FOREIGN KEY (course_id)
        REFERENCES public.courses (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER TABLE IF EXISTS public.user_to_course
    OWNER to postgres;
	
INSERT INTO courses
(name,area,course_id) VALUES 
('Calculus I', 'Math', 'CAL110'), 
('Calculus II', 'Math', 'CAL111'), 
('Geology I', 'Science', 'GEO101'), 
('Geology II', 'Science', 'GEO110'),
('Anthropology', 'Humanities', 'ANT101');

INSERT INTO users
(first_name, last_name, email, preferred_time) VALUES
('Jerry', 'Seinfeld', 'seinfeld@tamu.edu', 2),
('Elaine', 'Benes', 'elaine@tamu.edu', 2),
('George', 'Costanza', 'george@tamu.edu', 1),
('Cosmo', 'Kramer', 'kramer@tamu.edu', 0);

INSERT INTO user_to_course
(user_id, course_id, grade, completed_on) VALUES
(1, 2, 89, '2007-12-13'),
(2, 1, 99, '2007-12-13'),
(2, 2, 100, '2007-12-13'),
(2, 3, 100, '2007-12-13'),
(2, 4, 100, '2007-12-13');

INSERT INTO requests
(course_id, objective, description, method, meeting_day_of_week, meeting_time, contact_method, created_by) VALUES
(2, 'Help with differentials', 'I dont understand derivatives', 'homework help', 1, 2, 4, 'kramer@tamu.edu'),
(3, 'Help with water cycle', 'I dont understand why we need water', 'powerpoint', 1, 2, 4, 'kramer@tamu.edu'),
(4, 'Help with tectonic plate theory', 'I need to understand how the earth works', 'fieldtrip', 1, 2, 4, 'kramer@tamu.edu'),
(2, 'Help solving tridiagonal matrix problems', 'I need to code this in R', 'code tutorial', 1, 2, 4, 'george@tamu.edu');

select * from courses
select * from requests
select * from user_to_course
select * from users

select distinct u.id, u.first_name, u.last_name, u.email, uc.grade, uc.completed_on from users u
join user_to_course uc on u.id = uc.user_id
join requests r on r.course_id = uc.course_id
join courses co on r.course_id = co.id
where co.id = 2 and
u.preferred_time = 2

drop schema public cascade;
create schema public;
	

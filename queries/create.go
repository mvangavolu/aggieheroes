package queries

const UsersCreationQuery = `
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
`

const CoursesCreationQuery = `
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
`

const RequestsCreationQuery = `
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
`

const UserToCourseCreationQuery = `
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
`

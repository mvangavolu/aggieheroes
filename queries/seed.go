package queries

const SeedCoursesQuery = `
INSERT INTO courses
(name,area,course_id) VALUES 
('Calculus I', 'Math', 'CAL110'), 
('Calculus II', 'Math', 'CAL111'), 
('Geology I', 'Science', 'GEO101'), 
('Geology II', 'Science', 'GEO110'),
('Anthropology', 'Humanities', 'ANT101');
`

const SeedUsersQuery = `
INSERT INTO users
(first_name, last_name, email, preferred_time) VALUES
('Jerry', 'Seinfeld', 'seinfeld@tamu.edu', 2),
('Elaine', 'Benes', 'elaine@tamu.edu', 2),
('George', 'Costanza', 'george@tamu.edu', 1),
('Cosmo', 'Kramer', 'kramer@tamu.edu', 0);
`

const SeeduserToCourseQuery = `
INSERT INTO user_to_course
(user_id, course_id, grade, completed_on) VALUES
(1, 2, 89, '2007-12-13'),
(2, 1, 99, '2007-12-13'),
(2, 2, 100, '2007-12-13'),
(2, 3, 100, '2007-12-13'),
(2, 4, 100, '2007-12-13');
`

const SeedRequestsQuery = `
INSERT INTO requests
(course_id, objective, description, method, meeting_day_of_week, meeting_time, contact_method, created_by) VALUES
(2, 'Help with differentials', 'I dont understand derivatives', 'homework help', 1, 2, 4, 'kramer@tamu.edu'),
(3, 'Help with water cycle', 'I dont understand why we need water', 'powerpoint', 1, 2, 4, 'kramer@tamu.edu'),
(4, 'Help with tectonic plate theory', 'I need to understand how the earth works', 'fieldtrip', 1, 2, 4, 'kramer@tamu.edu'),
(2, 'Help solving tridiagonal matrix problems', 'I need to code this in R', 'code tutorial', 1, 2, 4, 'george@tamu.edu');
`

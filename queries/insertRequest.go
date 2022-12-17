package queries

const InsertRequestQuery = `
INSERT INTO requests
(course_id, objective, description, method, meeting_day_of_week, meeting_time, contact_method, created_by) VALUES
($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, course_id, objective, description, method, meeting_day_of_week, meeting_time, contact_method, created_by
`

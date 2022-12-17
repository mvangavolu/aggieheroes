package queries

// GetUserRequestsByEmailWuery gets all requests created under a tami email
const GetUserRequestsByEmailQuery = `
	select id, course_id, objective, description, method, meeting_day_of_week, meeting_time, contact_method, created_by from requests where created_by = $1
`

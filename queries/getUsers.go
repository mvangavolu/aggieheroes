package queries

// GetUsersQuery gets all users in the catalog that have taken a given course and
// match a preferred meeting time with the requestd meeting time
const GetUsersQuery = `
select distinct u.id, u.first_name, u.last_name, u.email, uc.grade, uc.completed_on from users u
join user_to_course uc on u.id = uc.user_id
join requests r on r.course_id = uc.course_id
join courses co on r.course_id = co.id
where co.id = $1 and
r.meeting_time = $2
`

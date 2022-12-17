package queries

// GetCoursesQuery gets all courses in the catalog
const GetCoursesQuery = `
	select id, course_id, name, area from courses 
`

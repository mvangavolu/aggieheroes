package services

import (
	"database/sql"
	"heroes/models"
	"heroes/queries"
)

func GetCourses(db *sql.DB) ([]models.Course, error) {
	rows, err := db.Query(
		queries.GetCoursesQuery,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	courses := []models.Course{}
	for rows.Next() {
		var c models.Course
		if err := rows.Scan(&c.Id, &c.CourseId, &c.Name, &c.Area); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

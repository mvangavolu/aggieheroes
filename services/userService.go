package services

import (
	"database/sql"
	"heroes/models"
	"heroes/queries"
)

func GetUsers(courseId int64, preferredTime int64, db *sql.DB) ([]models.UserCourse, error) {
	rows, err := db.Query(
		queries.GetUsersQuery,
		courseId,
		preferredTime,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	users := []models.UserCourse{}
	for rows.Next() {
		var uc models.UserCourse
		if err := rows.Scan(&uc.User.Id, &uc.User.FirstName, &uc.User.LastName, &uc.User.Email, &uc.Grade, &uc.CompletedOn); err != nil {
			return nil, err
		}
		users = append(users, uc)
	}
	return users, nil
}

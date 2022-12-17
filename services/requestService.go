package services

import (
	"database/sql"
	"heroes/models"
	"heroes/queries"
)

func CreateRequest(r models.Request, db *sql.DB) (models.Request, error) {
	err := db.QueryRow(
		queries.InsertRequestQuery,
		r.CourseId,
		r.Objective,
		r.Description,
		r.Method,
		r.MeetingDayOfWeek,
		r.MeetingTime,
		r.ContactMethod,
		r.CreatedBy,
	).Scan(
		&r.Id,
		&r.CourseId,
		&r.Objective,
		&r.Description,
		&r.Method,
		&r.MeetingDayOfWeek,
		&r.MeetingTime,
		&r.ContactMethod,
		&r.CreatedBy,
	)
	if err != nil {
		return r, err
	}

	return r, nil
}

func GetRequestsByEmail(email string, db *sql.DB) ([]models.Request, error) {
	rows, err := db.Query(
		queries.GetUserRequestsByEmailQuery,
		email,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	requests := []models.Request{}
	for rows.Next() {
		var r models.Request
		if err := rows.Scan(
			&r.Id,
			&r.CourseId,
			&r.Objective,
			&r.Description,
			&r.Method,
			&r.MeetingDayOfWeek,
			&r.MeetingTime,
			&r.ContactMethod,
			&r.CreatedBy,
		); err != nil {
			return nil, err
		}
		requests = append(requests, r)
	}
	return requests, nil
}

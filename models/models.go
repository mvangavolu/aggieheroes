package models

import "time"

type User struct {
	Id            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	PreferredTime int    `json:"preferredTime"`
}

type UserCourse struct {
	User        User      `json:"user"`
	Grade       float64   `json:"grade"`
	CompletedOn time.Time `json:"completedOn"`
}

type Course struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Area     string `json:"area"`
	CourseId string `json:"courseId"`
}

type Request struct {
	Id               int    `json:"id"`
	CourseId         int    `json:"courseId"`
	Objective        string `json:"objective"`
	Description      string `json:"description"`
	Method           string `json:"method"`
	MeetingDayOfWeek int    `json:"meetingDayOfWeek"`
	MeetingTime      int    `json:"meetingTime"`
	ContactMethod    int    `json:"contactMethod"`
	CreatedBy        string `json:"createdBy"`
}

package constants

type DayOfWeek int
type TimeOfDay int
type ContactMethod int

const (
	Sunday DayOfWeek = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	Morning TimeOfDay = iota
	Lunch
	Afternoon
	Evening
	Night
)

const (
	Phone    ContactMethod = 0
	Email    ContactMethod = 1
	InPerson ContactMethod = 2
	Zoom     ContactMethod = 4
)

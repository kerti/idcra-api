package model

// Student is the student entity
type Student struct {
	ID          string
	Name        string
	DateOfBirth string `db:"date_of_birth"`
	SchoolID    string `db:"school_id"`
	// School      *School
	CreatedAt string `db:"created_at"`
}

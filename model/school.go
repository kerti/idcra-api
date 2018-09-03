package model

// School is the school entity
type School struct {
	ID        string
	Name      string
	CreatedAt string `db:"created_at"`
	Students  []*Student
}

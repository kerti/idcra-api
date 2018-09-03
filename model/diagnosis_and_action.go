package model

// DiagnosisAndAction is the diagnosis and action entity
type DiagnosisAndAction struct {
	ID        string
	Diagnosis string
	Action    string
	CreatedAt string `db:"created_at"`
}

package model

// DiagnosisAndAction is the diagnosis and action entity
type DiagnosisAndAction struct {
	ID        string
	Diagnosis string
	Action    string
	UnitCost  float64 `db:"unit_cost"`
	CreatedAt string  `db:"created_at"`
}

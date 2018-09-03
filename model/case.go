package model

// Case is the case entity
type Case struct {
	ID                   string
	SurveyID             string `db:"survey_id"`
	DiagnosisAndActionID string `db:"diagnosis_and_action_id"`
	ToothNumber          int32  `db:"tooth_number"`
	CreatedAt            string `db:"created_at"`
}

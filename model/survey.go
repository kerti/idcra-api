package model

// Survey is the survey entity
type Survey struct {
	ID              string
	StudentID       string `db:"student_id"`
	SurveyorID      string `db:"surveyor_id"`
	Date            string `db:"date"`
	S1Q1            string `db:"s1q1"`
	S1Q2            string `db:"s1q2"`
	S1Q3            string `db:"s1q3"`
	S1Q4            string `db:"s1q4"`
	S1Q5            string `db:"s1q5"`
	S1Q6            string `db:"s1q6"`
	S1Q7            string `db:"s1q7"`
	S2Q1            string `db:"s2q1"`
	S2Q2            string `db:"s2q2"`
	S2Q3            string `db:"s2q3"`
	S2Q4            string `db:"s2q4"`
	S2Q5            string `db:"s2q5"`
	S2Q6            string `db:"s2q6"`
	S2Q7            string `db:"s2q7"`
	S2Q8            string `db:"s2q8"`
	S2Q9            string `db:"s2q9"`
	LowerD          int32  `db:"lower_d"`
	LowerE          int32  `db:"lower_e"`
	LowerF          int32  `db:"lower_f"`
	UpperD          int32  `db:"upper_d"`
	UpperM          int32  `db:"upper_m"`
	UpperF          int32  `db:"upper_f"`
	SubjectiveScore int32  `db:"subjective_score"`
	CreatedAt       string `db:"created_at"`
	Cases           []*Case
}

func (survey *Survey) CalculateScore() {
	// TODO: do this
	survey.SubjectiveScore = 0
}

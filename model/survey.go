package model

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

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

func (s *Survey) CalculateScore() {
	// TODO: do this
	s.SubjectiveScore = 0
}

// SurveyInput is the input for section entity
type SurveyInput struct {
	StudentID  *string
	SurveyorID *string
	Date       *string
	S1Q1       *string
	S1Q2       *string
	S1Q3       *string
	S1Q4       *string
	S1Q5       *string
	S1Q6       *string
	S1Q7       *string
	S2Q1       *string
	S2Q2       *string
	S2Q3       *string
	S2Q4       *string
	S2Q5       *string
	S2Q6       *string
	S2Q7       *string
	S2Q8       *string
	S2Q9       *string
	LowerD     *int32
	LowerE     *int32
	LowerF     *int32
	UpperD     *int32
	UpperM     *int32
	UpperF     *int32
	Cases      *[]*CaseInput
}

func (si *SurveyInput) Validate() error {
	if si.StudentID == nil {
		return fmt.Errorf("student ID is required")
	}

	if si.SurveyorID == nil {
		return fmt.Errorf("surveyor ID is required")
	}

	if si.Date == nil {
		return fmt.Errorf("date is required")
	}

	_, err := time.Parse("2006-01-02", *si.Date)
	if err != nil {
		return fmt.Errorf("invalid date format, expecting yyyy-mm-dd")
	}

	if si.S1Q1 == nil {
		return fmt.Errorf("section 1 question 1 is required")
	}

	if si.S1Q2 == nil {
		return fmt.Errorf("section 1 question 2 is required")
	}

	if si.S1Q3 == nil {
		return fmt.Errorf("section 1 question 3 is required")
	}

	if si.S1Q4 == nil {
		return fmt.Errorf("section 1 question 4 is required")
	}

	if si.S1Q5 == nil {
		return fmt.Errorf("section 1 question 5 is required")
	}

	if si.S1Q6 == nil {
		return fmt.Errorf("section 1 question 6 is required")
	}

	if si.S1Q7 == nil {
		return fmt.Errorf("section 1 question 7 is required")
	}

	if si.S2Q1 == nil {
		return fmt.Errorf("section 2 question 1 is required")
	}

	if si.S2Q2 == nil {
		return fmt.Errorf("section 2 question 2 is required")
	}

	if si.S2Q3 == nil {
		return fmt.Errorf("section 2 question 3 is required")
	}

	if si.S2Q4 == nil {
		return fmt.Errorf("section 2 question 4 is required")
	}

	if si.S2Q5 == nil {
		return fmt.Errorf("section 2 question 5 is required")
	}

	if si.S2Q6 == nil {
		return fmt.Errorf("section 2 question 6 is required")
	}

	if si.S2Q7 == nil {
		return fmt.Errorf("section 2 question 7 is required")
	}

	if si.S2Q8 == nil {
		return fmt.Errorf("section 2 question 8 is required")
	}

	if si.S2Q9 == nil {
		return fmt.Errorf("section 2 question 9 is required")
	}

	if si.LowerD == nil {
		return fmt.Errorf("lower d is required")
	}

	if si.LowerE == nil {
		return fmt.Errorf("lower e is required")
	}

	if si.LowerF == nil {
		return fmt.Errorf("lower f is required")
	}

	if si.UpperD == nil {
		return fmt.Errorf("upper D is required")
	}

	if si.UpperM == nil {
		return fmt.Errorf("upper M is required")
	}

	if si.UpperF == nil {
		return fmt.Errorf("upper F is required")
	}

	return nil
}

func NewSurveyFromInput(input SurveyInput) (s Survey, err error) {
	if err = input.Validate(); err != nil {
		return Survey{}, err
	}

	s = Survey{
		ID:         uuid.NewV4().String(),
		StudentID:  *input.StudentID,
		SurveyorID: *input.SurveyorID,
		Date:       *input.Date,
		S1Q1:       *input.S1Q1,
		S1Q2:       *input.S1Q2,
		S1Q3:       *input.S1Q3,
		S1Q4:       *input.S1Q4,
		S1Q5:       *input.S1Q5,
		S1Q6:       *input.S1Q6,
		S1Q7:       *input.S1Q7,
		S2Q1:       *input.S2Q1,
		S2Q2:       *input.S2Q2,
		S2Q3:       *input.S2Q3,
		S2Q4:       *input.S2Q4,
		S2Q5:       *input.S2Q5,
		S2Q6:       *input.S2Q6,
		S2Q7:       *input.S2Q7,
		S2Q8:       *input.S2Q8,
		S2Q9:       *input.S2Q9,
		LowerD:     *input.LowerD,
		LowerE:     *input.LowerE,
		LowerF:     *input.LowerF,
		UpperD:     *input.UpperD,
		UpperM:     *input.UpperM,
		UpperF:     *input.UpperF,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		Cases:      []*Case{},
	}

	for _, ci := range *input.Cases {
		c, err := NewCaseFromInput(*ci, s.ID)
		if err != nil {
			return Survey{}, err
		}

		s.Cases = append(s.Cases, &c)
	}

	return
}

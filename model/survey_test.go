package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	studentID     = "fake_student_id"
	surveyorID    = "fake_surveyor_id"
	date          = "2018-01-01"
	s1q1          = "Low"
	s1q2          = "Low"
	s1q3          = "Low"
	s1q4          = "Low"
	s1q5          = "Low"
	s1q6          = "Low"
	s1q7          = "Low"
	s2q1          = "Low"
	s2q2          = "Low"
	s2q3          = "Low"
	s2q4          = "Low"
	s2q5          = "Low"
	s2q6          = "Low"
	s2q7          = "Low"
	s2q8          = "Low"
	s2q9          = "Low"
	lowerD        = int32(1)
	lowerE        = int32(2)
	lowerF        = int32(3)
	upperD        = int32(4)
	upperM        = int32(5)
	upperF        = int32(6)
	malformedDate = "fake_date"
)

func getValidSurveyInput() *SurveyInput {
	cases := []*CaseInput{getValidCaseInput()}
	return &SurveyInput{
		StudentID:  &studentID,
		SurveyorID: &surveyorID,
		Date:       &date,
		S1Q1:       &s1q1,
		S1Q2:       &s1q2,
		S1Q3:       &s1q3,
		S1Q4:       &s1q4,
		S1Q5:       &s1q5,
		S1Q6:       &s1q6,
		S1Q7:       &s1q7,
		S2Q1:       &s2q1,
		S2Q2:       &s2q2,
		S2Q3:       &s2q3,
		S2Q4:       &s2q4,
		S2Q5:       &s2q5,
		S2Q6:       &s2q6,
		S2Q7:       &s2q7,
		S2Q8:       &s2q8,
		S2Q9:       &s2q9,
		LowerD:     &lowerD,
		LowerE:     &lowerE,
		LowerF:     &lowerF,
		UpperD:     &upperD,
		UpperM:     &upperM,
		UpperF:     &upperF,
		Cases:      &cases,
	}
}

func TestSurveyInput(t *testing.T) {

	t.Run("calculateScore", func(t *testing.T) {
		input := getValidSurveyInput()
		survey, err := NewSurveyFromInput(*input)

		assert.Nil(t, err)

		survey.CalculateScore()

		assert.Equal(t, int32(0), survey.SubjectiveScore)
	})

	t.Run("validation", func(t *testing.T) {

		t.Run("NoErrors", func(t *testing.T) {
			sut := getValidSurveyInput()
			err := sut.Validate()

			assert.Nil(t, err)
		})

		t.Run("NilStudentID", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.StudentID = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "student ID is required", err.Error())
		})

		t.Run("NilSurveyorID", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.SurveyorID = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "surveyor ID is required", err.Error())
		})

		t.Run("NilDate", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.Date = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "date is required", err.Error())
		})

		t.Run("MalformedDate", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.Date = &malformedDate
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "invalid date format, expecting yyyy-mm-dd", err.Error())
		})

		t.Run("NilSection1Question1", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S1Q1 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 1 question 1 is required", err.Error())
		})

		t.Run("NilSection1Question2", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S1Q2 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 1 question 2 is required", err.Error())
		})

		t.Run("NilSection1Question3", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S1Q3 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 1 question 3 is required", err.Error())
		})

		t.Run("NilSection1Question4", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S1Q4 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 1 question 4 is required", err.Error())
		})

		t.Run("NilSection1Question5", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S1Q5 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 1 question 5 is required", err.Error())
		})

		t.Run("NilSection1Question6", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S1Q6 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 1 question 6 is required", err.Error())
		})

		t.Run("NilSection1Question7", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S1Q7 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 1 question 7 is required", err.Error())
		})

		t.Run("NilSection2Question1", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S2Q1 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 2 question 1 is required", err.Error())
		})

		t.Run("NilSection2Question2", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S2Q2 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 2 question 2 is required", err.Error())
		})

		t.Run("NilSection2Question3", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S2Q3 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 2 question 3 is required", err.Error())
		})

		t.Run("NilSection2Question4", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S2Q4 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 2 question 4 is required", err.Error())
		})

		t.Run("NilSection2Question5", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S2Q5 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 2 question 5 is required", err.Error())
		})

		t.Run("NilSection2Question6", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S2Q6 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 2 question 6 is required", err.Error())
		})

		t.Run("NilSection2Question7", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S2Q7 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 2 question 7 is required", err.Error())
		})

		t.Run("NilSection2Question8", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S2Q8 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 2 question 8 is required", err.Error())
		})

		t.Run("NilSection2Question9", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.S2Q9 = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "section 2 question 9 is required", err.Error())
		})

		t.Run("NilLowerD", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.LowerD = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "lower d is required", err.Error())
		})

		t.Run("NilLowerE", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.LowerE = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "lower e is required", err.Error())
		})

		t.Run("NilLowerF", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.LowerF = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "lower f is required", err.Error())
		})

		t.Run("NilUpperD", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.UpperD = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "upper D is required", err.Error())
		})

		t.Run("NilUpperM", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.UpperM = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "upper M is required", err.Error())
		})

		t.Run("NilUpperF", func(t *testing.T) {
			sut := getValidSurveyInput()
			sut.UpperF = nil
			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "upper F is required", err.Error())
		})

	})
}

func TestNewSurveyFromInput(t *testing.T) {

	t.Run("NoErrors", func(t *testing.T) {
		input := getValidSurveyInput()

		s, err := NewSurveyFromInput(*input)

		assert.Nil(t, err)
		assert.Equal(t, 36, len(s.ID))
		assert.Equal(t, surveyorID, s.SurveyorID)
		assert.Equal(t, studentID, s.StudentID)
		assert.Equal(t, date, s.Date)
		assert.Equal(t, s1q1, s.S1Q1)
		assert.Equal(t, s1q2, s.S1Q2)
		assert.Equal(t, s1q3, s.S1Q3)
		assert.Equal(t, s1q4, s.S1Q4)
		assert.Equal(t, s1q5, s.S1Q5)
		assert.Equal(t, s1q6, s.S1Q6)
		assert.Equal(t, s1q7, s.S1Q7)
		assert.Equal(t, s2q1, s.S2Q1)
		assert.Equal(t, s2q2, s.S2Q2)
		assert.Equal(t, s2q3, s.S2Q3)
		assert.Equal(t, s2q4, s.S2Q4)
		assert.Equal(t, s2q5, s.S2Q5)
		assert.Equal(t, s2q6, s.S2Q6)
		assert.Equal(t, s2q7, s.S2Q7)
		assert.Equal(t, s2q8, s.S2Q8)
		assert.Equal(t, s2q9, s.S2Q9)
		assert.Equal(t, lowerD, s.LowerD)
		assert.Equal(t, lowerE, s.LowerE)
		assert.Equal(t, lowerF, s.LowerF)
		assert.Equal(t, upperD, s.UpperD)
		assert.Equal(t, upperM, s.UpperM)
		assert.Equal(t, upperF, s.UpperF)

		assert.Equal(t, 1, len(s.Cases))
		assert.Equal(t, 36, len(s.Cases[0].ID))
		assert.Equal(t, s.ID, s.Cases[0].SurveyID)
		assert.Equal(t, daID, s.Cases[0].DiagnosisAndActionID)
		assert.Equal(t, toothNumber, s.Cases[0].ToothNumber)
	})

	t.Run("WithErrors", func(t *testing.T) {
		input := getValidSurveyInput()
		input.StudentID = nil

		_, err := NewSurveyFromInput(*input)

		assert.NotNil(t, err)
		assert.Equal(t, "student ID is required", err.Error())
	})

	t.Run("WithErrorsInCases", func(t *testing.T) {
		input := getValidSurveyInput()

		assert.Equal(t, 1, len(*input.Cases))

		cases := *input.Cases
		cases[0].ToothNumber = nil
		input.Cases = &cases

		_, err := NewSurveyFromInput(*input)

		assert.NotNil(t, err)
		assert.Equal(t, "tooth number is required", err.Error())
	})

}

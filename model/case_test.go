package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	daID        = "fakeDiagnosisAndActionID"
	toothNumber = int32(21)
	surveyID    = "fakeSurveyID"
)

func getValidCaseInput() *CaseInput {
	return &CaseInput{
		DiagnosisAndActionID: &daID,
		ToothNumber:          &toothNumber,
	}
}

func TestCaseInput(t *testing.T) {

	t.Run("Validation", func(t *testing.T) {

		t.Run("NoErrors", func(t *testing.T) {
			sut := getValidCaseInput()

			err := sut.Validate()

			assert.Nil(t, err)
		})

		t.Run("NilDiagnosisAndActionID", func(t *testing.T) {
			sut := getValidCaseInput()
			sut.DiagnosisAndActionID = nil

			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "diagnosis and action ID is required", err.Error())
		})

		t.Run("NilToothNumber", func(t *testing.T) {
			sut := getValidCaseInput()
			sut.ToothNumber = nil

			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, "tooth number is required", err.Error())
		})

	})

}

func TestNewCaseFromInput(t *testing.T) {

	t.Run("NoErrors", func(t *testing.T) {
		input := getValidCaseInput()

		c, err := NewCaseFromInput(*input, surveyID)

		assert.Nil(t, err)
		assert.Equal(t, 36, len(c.ID))
		assert.Equal(t, toothNumber, c.ToothNumber)
	})

	t.Run("WithErrors", func(t *testing.T) {
		input := getValidCaseInput()
		input.ToothNumber = nil

		_, err := NewCaseFromInput(*input, surveyID)

		assert.NotNil(t, err)
		assert.Equal(t, "tooth number is required", err.Error())
	})

}

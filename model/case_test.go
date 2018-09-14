package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	daID        *string
	toothNumber *int32
	surveyID    string
)

func init() {
	daIDObj := "fakeDiagnosisAndActionID"
	daID = &daIDObj
	toothNumberObj := int32(21)
	toothNumber = &toothNumberObj
	surveyID = "fakeSurveyID"
}

func TestCaseInput(t *testing.T) {

	t.Run("Validation", func(t *testing.T) {

		t.Run("NoErrors", func(t *testing.T) {
			sut := CaseInput{
				DiagnosisAndActionID: daID,
				ToothNumber:          toothNumber,
			}

			err := sut.Validate()

			assert.Nil(t, err)
		})

		t.Run("NilDiagnosisAndActionID", func(t *testing.T) {
			sut := CaseInput{
				DiagnosisAndActionID: nil,
				ToothNumber:          toothNumber,
			}

			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), "diagnosis and action ID is required")
		})

		t.Run("NilToothNumber", func(t *testing.T) {
			sut := CaseInput{
				DiagnosisAndActionID: daID,
				ToothNumber:          nil,
			}

			err := sut.Validate()

			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), "tooth number is required")
		})

	})

}

func TestNewCaseFromInput(t *testing.T) {

	t.Run("NoErrors", func(t *testing.T) {
		input := CaseInput{
			DiagnosisAndActionID: daID,
			ToothNumber:          toothNumber,
		}

		c, err := NewCaseFromInput(input, surveyID)

		assert.Nil(t, err)
		assert.Equal(t, len(c.ID), 36)
		assert.Equal(t, c.ToothNumber, *toothNumber)
	})

	t.Run("WithErrors", func(t *testing.T) {
		input := CaseInput{
			DiagnosisAndActionID: daID,
			ToothNumber:          nil,
		}

		_, err := NewCaseFromInput(input, surveyID)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "tooth number is required")
	})

}

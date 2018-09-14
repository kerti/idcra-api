package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	daID        *string
	toothNumber *int32
)

func init() {
	daIDObj := "fakeDiagnosisAndActionID"
	daID = &daIDObj
	toothNumberObj := int32(21)
	toothNumber = &toothNumberObj
}

func TestValidateInput(t *testing.T) {
	t.Run("NoError", func(t *testing.T) {
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
}

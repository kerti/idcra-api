package model

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Case is the case entity
type Case struct {
	ID                   string
	SurveyID             string `db:"survey_id"`
	DiagnosisAndActionID string `db:"diagnosis_and_action_id"`
	ToothNumber          int32  `db:"tooth_number"`
	CreatedAt            string `db:"created_at"`
}

// CaseInput is the input for case entity
type CaseInput struct {
	DiagnosisAndActionID *string
	ToothNumber          *int32
}

func (ci *CaseInput) Validate() error {
	if ci.DiagnosisAndActionID == nil {
		return fmt.Errorf("diagnosis and action ID is required")
	}

	if ci.ToothNumber == nil {
		return fmt.Errorf("tooth number is required")
	}

	return nil
}

func NewCaseFromInput(input CaseInput, surveyID string) (c Case, err error) {
	if err = input.Validate(); err != nil {
		return Case{}, err
	}

	c = Case{
		ID:                   uuid.NewV4().String(),
		SurveyID:             surveyID,
		DiagnosisAndActionID: *input.DiagnosisAndActionID,
		ToothNumber:          *input.ToothNumber,
		CreatedAt:            time.Now().Format("2006-01-02 15:04:05"),
	}

	return
}

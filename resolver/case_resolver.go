package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
)

type caseResolver struct {
	c *model.Case
}

func (c *caseResolver) ID() graphql.ID {
	return graphql.ID(c.c.ID)
}

func (c *caseResolver) SurveyID() *string {
	return &c.c.SurveyID
}

func (c *caseResolver) DiagnosisAndActionID() *string {
	return &c.c.DiagnosisAndActionID
}

func (c *caseResolver) ToothNumber() *int32 {
	return &c.c.ToothNumber
}

func (c *caseResolver) CreatedAt() (*graphql.Time, error) {
	if c.c.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, c.c.CreatedAt)
	return &graphql.Time{Time: t}, err
}

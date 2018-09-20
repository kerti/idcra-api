package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
)

type diagnosisAndActionResolver struct {
	d *model.DiagnosisAndAction
}

func (d *diagnosisAndActionResolver) ID() graphql.ID {
	return graphql.ID(d.d.ID)
}

func (d *diagnosisAndActionResolver) Diagnosis() *string {
	return &d.d.Diagnosis
}

func (d *diagnosisAndActionResolver) Action() *string {
	return &d.d.Action
}

func (d *diagnosisAndActionResolver) UnitCost() *float64 {
	return &d.d.UnitCost
}

func (d *diagnosisAndActionResolver) CreatedAt() (*graphql.Time, error) {
	if d.d.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, d.d.CreatedAt)
	return &graphql.Time{Time: t}, err
}

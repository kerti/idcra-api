package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
)

type surveysEdgeResolver struct {
	cursor graphql.ID
	model  *model.Survey
}

func (r *surveysEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *surveysEdgeResolver) Node() *surveyResolver {
	return &surveyResolver{s: r.model}
}

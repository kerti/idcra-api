package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
)

type studentsEdgeResolver struct {
	cursor graphql.ID
	model  *model.Student
}

func (r *studentsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *studentsEdgeResolver) Node() *studentResolver {
	return &studentResolver{s: r.model}
}

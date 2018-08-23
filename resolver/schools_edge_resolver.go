package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
)

type schoolsEdgeResolver struct {
	cursor graphql.ID
	model  *model.School
}

func (r *schoolsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *schoolsEdgeResolver) Node() *schoolResolver {
	return &schoolResolver{s: r.model}
}

package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
)

type schoolsConnectionResolver struct {
	schools    []*model.School
	totalCount int
	from       *string
	to         *string
}

func (r *schoolsConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *schoolsConnectionResolver) Edges() *[]*schoolsEdgeResolver {
	l := make([]*schoolsEdgeResolver, len(r.schools))
	for i := range l {
		l[i] = &schoolsEdgeResolver{
			cursor: service.EncodeCursor(&(r.schools[i].ID)),
			model:  r.schools[i],
		}
	}
	return &l
}

func (r *schoolsConnectionResolver) PageInfo() *pageInfoResolver {
	var startCursor graphql.ID
	var endCursor graphql.ID

	if r.from != nil {
		startCursor = service.EncodeCursor(r.from)
	}

	if r.to != nil {
		endCursor = service.EncodeCursor(r.to)
	}

	return &pageInfoResolver{
		startCursor: &startCursor,
		endCursor:   &endCursor,
		hasNextPage: false,
	}
}

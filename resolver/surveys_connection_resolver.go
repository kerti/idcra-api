package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
)

type surveysConnectionResolver struct {
	surveys    []*model.Survey
	totalCount int
	from       *string
	to         *string
}

func (r *surveysConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *surveysConnectionResolver) Edges() *[]*surveysEdgeResolver {
	l := make([]*surveysEdgeResolver, len(r.surveys))
	for i := range l {
		l[i] = &surveysEdgeResolver{
			cursor: service.EncodeCursor(&(r.surveys[i].ID)),
			model:  r.surveys[i],
		}
	}
	return &l
}

func (r *surveysConnectionResolver) PageInfo() *pageInfoResolver {
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

package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
)

type diagnosisAndActionsConnectionResolver struct {
	diagnosisAndActions []*model.DiagnosisAndAction
	totalCount          int
	from                *string
	to                  *string
}

func (r *diagnosisAndActionsConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *diagnosisAndActionsConnectionResolver) Edges() *[]*diagnosisAndActionsEdgeResolver {
	l := make([]*diagnosisAndActionsEdgeResolver, len(r.diagnosisAndActions))
	for i := range l {
		l[i] = &diagnosisAndActionsEdgeResolver{
			cursor: service.EncodeCursor(&(r.diagnosisAndActions[i].ID)),
			model:  r.diagnosisAndActions[i],
		}
	}
	return &l
}

func (r *diagnosisAndActionsConnectionResolver) PageInfo() *pageInfoResolver {
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

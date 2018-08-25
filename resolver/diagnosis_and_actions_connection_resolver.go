package resolver

import (
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
	return &pageInfoResolver{
		startCursor: service.EncodeCursor(r.from),
		endCursor:   service.EncodeCursor(r.to),
		hasNextPage: false,
	}
}

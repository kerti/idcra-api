package resolver

import (
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
	return &pageInfoResolver{
		startCursor: service.EncodeCursor(r.from),
		endCursor:   service.EncodeCursor(r.to),
		hasNextPage: false,
	}
}

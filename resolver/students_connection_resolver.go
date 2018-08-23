package resolver

import (
	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
)

type studentsConnectionResolver struct {
	students   []*model.Student
	totalCount int
	from       *string
	to         *string
}

func (r *studentsConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *studentsConnectionResolver) Edges() *[]*studentsEdgeResolver {
	l := make([]*studentsEdgeResolver, len(r.students))
	for i := range l {
		l[i] = &studentsEdgeResolver{
			cursor: service.EncodeCursor(&(r.students[i].ID)),
			model:  r.students[i],
		}
	}
	return &l
}

func (r *studentsConnectionResolver) PageInfo() *pageInfoResolver {
	return &pageInfoResolver{
		startCursor: service.EncodeCursor(r.from),
		endCursor:   service.EncodeCursor(r.to),
		hasNextPage: false,
	}
}

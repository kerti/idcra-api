package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
)

type usersConnectionResolver struct {
	users      []*model.User
	totalCount int
	from       *string
	to         *string
}

func (r *usersConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *usersConnectionResolver) Edges() *[]*usersEdgeResolver {
	l := make([]*usersEdgeResolver, len(r.users))
	for i := range l {
		l[i] = &usersEdgeResolver{
			cursor: service.EncodeCursor(&(r.users[i].ID)),
			model:  r.users[i],
		}
	}
	return &l
}

func (r *usersConnectionResolver) PageInfo() *pageInfoResolver {
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

package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
)

type schoolResolver struct {
	s *model.School
}

func (s *schoolResolver) ID() graphql.ID {
	return graphql.ID(s.s.ID)
}

func (s *schoolResolver) Name() *string {
	return &s.s.Name
}

func (s *schoolResolver) CreatedAt() (*graphql.Time, error) {
	if s.s.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, s.s.CreatedAt)
	return &graphql.Time{Time: t}, err
}

func (s *schoolResolver) Students() *[]*studentResolver {
	l := make([]*studentResolver, len(s.s.Students))
	for i := range l {
		l[i] = &studentResolver{
			s: s.s.Students[i],
		}
	}
	return &l
}

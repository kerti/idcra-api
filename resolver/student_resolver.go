package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
)

type studentResolver struct {
	s *model.Student
}

func (s *studentResolver) ID() graphql.ID {
	return graphql.ID(s.s.ID)
}

func (s *studentResolver) Name() *string {
	return &s.s.Name
}

func (s *studentResolver) DateOfBirth() (*graphql.Time, error) {
	if s.s.DateOfBirth == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, s.s.DateOfBirth)
	return &graphql.Time{Time: t}, err
}

func (s *studentResolver) SchoolID() graphql.ID {
	return graphql.ID(s.s.SchoolID)
}

// func (s *studentResolver) School() *schoolResolver {
// 	return &schoolResolver{s: s.s.School}
// }

func (s *studentResolver) CreatedAt() (*graphql.Time, error) {
	if s.s.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, s.s.CreatedAt)
	return &graphql.Time{Time: t}, err
}

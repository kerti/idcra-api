package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
)

type surveyResolver struct {
	s *model.Survey
}

func (s *surveyResolver) ID() graphql.ID {
	return graphql.ID(s.s.ID)
}

func (s *surveyResolver) StudentID() *string {
	return &s.s.StudentID
}

func (s *surveyResolver) SurveyorID() *string {
	return &s.s.SurveyorID
}

func (s *surveyResolver) Date() (*graphql.Time, error) {
	if s.s.Date == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, s.s.Date)
	return &graphql.Time{Time: t}, err
}

func (s *surveyResolver) S1Q1() *string {
	return &s.s.S1Q1
}

func (s *surveyResolver) S1Q2() *string {
	return &s.s.S1Q2
}

func (s *surveyResolver) S1Q3() *string {
	return &s.s.S1Q3
}

func (s *surveyResolver) S1Q4() *string {
	return &s.s.S1Q4
}

func (s *surveyResolver) S1Q5() *string {
	return &s.s.S1Q5
}

func (s *surveyResolver) S1Q6() *string {
	return &s.s.S1Q6
}

func (s *surveyResolver) S1Q7() *string {
	return &s.s.S1Q7
}

func (s *surveyResolver) S2Q1() *string {
	return &s.s.S2Q1
}

func (s *surveyResolver) S2Q2() *string {
	return &s.s.S2Q2
}

func (s *surveyResolver) S2Q3() *string {
	return &s.s.S2Q3
}

func (s *surveyResolver) S2Q4() *string {
	return &s.s.S2Q4
}

func (s *surveyResolver) S2Q5() *string {
	return &s.s.S2Q5
}

func (s *surveyResolver) S2Q6() *string {
	return &s.s.S2Q6
}

func (s *surveyResolver) S2Q7() *string {
	return &s.s.S2Q7
}

func (s *surveyResolver) S2Q8() *string {
	return &s.s.S2Q8
}

func (s *surveyResolver) S2Q9() *string {
	return &s.s.S2Q9
}

func (s *surveyResolver) LowerD() *int32 {
	return &s.s.LowerD
}

func (s *surveyResolver) LowerE() *int32 {
	return &s.s.LowerE
}

func (s *surveyResolver) LowerF() *int32 {
	return &s.s.LowerF
}

func (s *surveyResolver) UpperD() *int32 {
	return &s.s.UpperD
}

func (s *surveyResolver) UpperM() *int32 {
	return &s.s.UpperM
}

func (s *surveyResolver) UpperF() *int32 {
	return &s.s.UpperF
}

func (s *surveyResolver) SubjectiveScore() *int32 {
	return &s.s.SubjectiveScore
}

func (s *surveyResolver) CreatedAt() (*graphql.Time, error) {
	if s.s.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, s.s.CreatedAt)
	return &graphql.Time{Time: t}, err
}

func (s *surveyResolver) Cases() *[]*caseResolver {
	l := make([]*caseResolver, len(s.s.Cases))
	for i := range l {
		l[i] = &caseResolver{
			c: s.s.Cases[i],
		}
	}
	return &l
}

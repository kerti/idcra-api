package resolver

import (
	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) CreateSurvey(ctx context.Context, args *struct {
	Survey *model.SurveyInput
}) (*surveyResolver, error) {
	survey, err := model.NewSurveyFromInput(*args.Survey)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	createdSurvey, err := ctx.Value("surveyService").(*service.SurveyService).TransactionalCreateSurvey(&survey)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created survey : %v", createdSurvey)

	return &surveyResolver{createdSurvey}, nil
}

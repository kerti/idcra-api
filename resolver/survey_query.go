package resolver

import (
	"errors"

	gcontext "github.com/kerti/idcra-api/context"
	"github.com/kerti/idcra-api/loader"
	"github.com/kerti/idcra-api/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) Survey(ctx context.Context, args struct {
	ID string
}) (*surveyResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	survey, err := loader.LoadSurveyByID(ctx, args.ID)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved survey by user_id[%s] : %v", *userID, *survey)

	return &surveyResolver{survey}, nil
}

func (r *Resolver) Surveys(ctx context.Context, args struct {
	First     *int32
	After     *string
	StudentID *string
}) (*surveysConnectionResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	surveys, err := ctx.Value("surveyService").(*service.SurveyService).List(args.First, args.After, args.StudentID)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	count, err := ctx.Value("surveyService").(*service.SurveyService).Count(args.StudentID)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved surveys by user_id[%s] :", *userID)

	config := ctx.Value("config").(*gcontext.Config)
	if config.DebugMode {
		for _, survey := range surveys {
			ctx.Value("log").(*logging.Logger).Debugf("%v", *survey)
		}
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved total surveys count by user_id[%s] : %v", *userID, count)

	if len(surveys) > 0 {
		return &surveysConnectionResolver{surveys: surveys, totalCount: count, from: &(surveys[0].ID), to: &(surveys[len(surveys)-1].ID)}, nil
	}
	return &surveysConnectionResolver{surveys: surveys, totalCount: count, from: nil, to: nil}, nil
}

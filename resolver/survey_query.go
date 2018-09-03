package resolver

import (
	"errors"

	gcontext "github.com/kerti/idcra-api/context"
	"github.com/kerti/idcra-api/loader"
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

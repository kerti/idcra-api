package resolver

import (
	"errors"

	gcontext "github.com/kerti/idcra-api/context"
	"github.com/kerti/idcra-api/loader"
	"github.com/kerti/idcra-api/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) DiagnosisAndAction(ctx context.Context, args struct {
	ID string
}) (*diagnosisAndActionResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	diagnosisAndAction, err := loader.LoadDiagnosisAndActionByID(ctx, args.ID)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved diagnosis and action by user_id[%s] : %v", *userID, *diagnosisAndAction)

	return &diagnosisAndActionResolver{diagnosisAndAction}, nil
}

func (r *Resolver) DiagnosisAndActions(ctx context.Context, args struct {
	First *int32
	After *string
}) (*diagnosisAndActionsConnectionResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	diagnosisAndActions, err := ctx.Value("diagnosisAndActionService").(*service.DiagnosisAndActionService).List(args.First, args.After)
	count, err := ctx.Value("diagnosisAndActionService").(*service.DiagnosisAndActionService).Count()
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved diagnosisAndActions by user_id[%s] :", *userID)
	config := ctx.Value("config").(*gcontext.Config)
	if config.DebugMode {
		for _, diagnosisAndAction := range diagnosisAndActions {
			ctx.Value("log").(*logging.Logger).Debugf("%v", *diagnosisAndAction)
		}
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved total diagnosisAndActions count by user_id[%s] : %v", *userID, count)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &diagnosisAndActionsConnectionResolver{diagnosisAndActions: diagnosisAndActions, totalCount: count, from: &(diagnosisAndActions[0].ID), to: &(diagnosisAndActions[len(diagnosisAndActions)-1].ID)}, nil
}

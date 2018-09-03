package resolver

import (
	"errors"

	gcontext "github.com/kerti/idcra-api/context"
	"github.com/kerti/idcra-api/loader"
	"github.com/kerti/idcra-api/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) School(ctx context.Context, args struct {
	ID string
}) (*schoolResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	school, err := loader.LoadSchoolByID(ctx, args.ID)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved school by user_id[%s] : %v", *userID, *school)

	return &schoolResolver{school}, nil
}

func (r *Resolver) Schools(ctx context.Context, args struct {
	First *int32
	After *string
}) (*schoolsConnectionResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	schools, err := ctx.Value("schoolService").(*service.SchoolService).List(args.First, args.After)
	if err != nil {
		return nil, err
	}

	count, err := ctx.Value("schoolService").(*service.SchoolService).Count()
	if err != nil {
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved schools by user_id[%s] :", *userID)
	config := ctx.Value("config").(*gcontext.Config)

	if config.DebugMode {
		for _, school := range schools {
			ctx.Value("log").(*logging.Logger).Debugf("%v", *school)
		}
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved total schools count by user_id[%s] : %v", *userID, count)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &schoolsConnectionResolver{schools: schools, totalCount: count, from: &(schools[0].ID), to: &(schools[len(schools)-1].ID)}, nil
}

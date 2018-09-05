package resolver

import (
	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) CreateSchool(ctx context.Context, args *struct {
	Name string
}) (*schoolResolver, error) {
	school := &model.School{
		Name: args.Name,
	}

	school, err := ctx.Value("schoolService").(*service.SchoolService).CreateSchool(school)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created school : %v", *school)
	return &schoolResolver{school}, nil
}

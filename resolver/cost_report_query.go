package resolver

import (
	"errors"

	gcontext "github.com/kerti/idcra-api/context"
	"github.com/kerti/idcra-api/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) CostBreakdownBySchoolAndDateRange(ctx context.Context, args struct {
	SchoolID  string
	StartDate string
	EndDate   string
}) (*[]*costReportResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	reports, err := ctx.Value("reportService").(*service.ReportService).CostBreakdownBySchoolAndDateRange(args.SchoolID, args.StartDate, args.EndDate)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved cost report by user_id[%s] : %v", *userID)

	config := ctx.Value("config").(*gcontext.Config)
	if config.DebugMode {
		for _, r := range reports {
			ctx.Value("log").(*logging.Logger).Debugf("%v", *r)
		}
	}

	result := make([]*costReportResolver, 0)
	for _, r := range reports {
		rslv := costReportResolver{r}
		result = append(result, &rslv)
	}

	return &result, nil
}

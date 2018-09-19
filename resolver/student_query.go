package resolver

import (
	"errors"

	gcontext "github.com/kerti/idcra-api/context"
	"github.com/kerti/idcra-api/loader"
	"github.com/kerti/idcra-api/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) Student(ctx context.Context, args struct {
	ID string
}) (*studentResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	student, err := loader.LoadStudentByID(ctx, args.ID)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved student by user_id[%s] : %v", *userID, *student)

	return &studentResolver{student}, nil
}

func (r *Resolver) Students(ctx context.Context, args struct {
	First    *int32
	After    *string
	SchoolID *string
	Keyword  *string
}) (*studentsConnectionResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(gcontext.CredentialsError)
	}
	userID := ctx.Value("user_id").(*string)

	students, err := ctx.Value("studentService").(*service.StudentService).List(args.First, args.After, args.SchoolID, args.Keyword)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	count, err := ctx.Value("studentService").(*service.StudentService).Count(args.SchoolID, args.Keyword)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved students by user_id[%s] :", *userID)
	config := ctx.Value("config").(*gcontext.Config)

	if config.DebugMode {
		for _, student := range students {
			ctx.Value("log").(*logging.Logger).Debugf("%v", *student)
		}
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved total students count by user_id[%s] : %v", *userID, count)

	if len(students) > 0 {
		return &studentsConnectionResolver{students: students, totalCount: count, from: &(students[0].ID), to: &(students[len(students)-1].ID)}, nil
	}
	return &studentsConnectionResolver{students: students, totalCount: count, from: nil, to: nil}, nil
}

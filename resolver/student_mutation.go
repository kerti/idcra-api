package resolver

import (
	"time"

	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) CreateStudent(ctx context.Context, args *struct {
	Name        string
	DateOfBirth string
	SchoolID    string
}) (*studentResolver, error) {
	_, err := time.Parse("2006-01-02", args.DateOfBirth)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	student := &model.Student{
		Name:        args.Name,
		DateOfBirth: args.DateOfBirth,
		SchoolID:    args.SchoolID,
	}

	student, err = ctx.Value("studentService").(*service.StudentService).CreateStudent(student)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Created student : %v", *student)
	return &studentResolver{student}, nil
}

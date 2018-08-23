package loader

import (
	"fmt"
	"sync"

	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
	"golang.org/x/net/context"
	"gopkg.in/nicksrandall/dataloader.v5"
)

type studentLoaderByID struct{}

func newStudentLoaderByID() dataloader.BatchFunc {
	return studentLoaderByID{}.loadBatch
}

func (ldr studentLoaderByID) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			student, err := ctx.Value("studentService").(*service.StudentService).FindByID(key.String())
			results[i] = &dataloader.Result{Data: student, Error: err}
		}(i, key)
	}

	wg.Wait()

	return results
}

func LoadStudentByID(ctx context.Context, key string) (*model.Student, error) {
	var student *model.Student

	ldr, err := extract(ctx, studentLoaderByIDKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	student, ok := data.(*model.Student)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", student, data)
	}

	return student, nil
}

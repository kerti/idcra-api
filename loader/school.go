package loader

import (
	"fmt"
	"sync"

	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
	"golang.org/x/net/context"
	"gopkg.in/nicksrandall/dataloader.v5"
)

type schoolLoaderByID struct{}

func newSchoolLoaderByID() dataloader.BatchFunc {
	return schoolLoaderByID{}.loadBatch
}

func (ldr schoolLoaderByID) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			school, err := ctx.Value("schoolService").(*service.SchoolService).FindByID(key.String())
			results[i] = &dataloader.Result{Data: school, Error: err}
		}(i, key)
	}

	wg.Wait()

	return results
}

func LoadSchoolByID(ctx context.Context, key string) (*model.School, error) {
	var school *model.School

	ldr, err := extract(ctx, schoolLoaderByIDKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	school, ok := data.(*model.School)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", school, data)
	}

	return school, nil
}

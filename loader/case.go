package loader

import (
	"fmt"
	"sync"

	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
	"golang.org/x/net/context"
	"gopkg.in/nicksrandall/dataloader.v5"
)

type caseLoaderByID struct{}

func newCaseLoaderByID() dataloader.BatchFunc {
	return caseLoaderByID{}.loadBatch
}

func (ldr caseLoaderByID) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			caseObj, err := ctx.Value("caseService").(*service.CaseService).FindByID(key.String())
			results[i] = &dataloader.Result{Data: caseObj, Error: err}
		}(i, key)
	}

	wg.Wait()

	return results
}

func LoadCaseByID(ctx context.Context, key string) (*model.Case, error) {
	var caseObj *model.Case

	ldr, err := extract(ctx, caseLoaderByIDKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}
	caseObj, ok := data.(*model.Case)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", caseObj, data)
	}

	return caseObj, nil
}

package loader

import (
	"fmt"
	"sync"

	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
	"golang.org/x/net/context"
	"gopkg.in/nicksrandall/dataloader.v5"
)

type diagnosisAndActionLoaderByID struct{}

func newDiagnosisAndActionLoaderByID() dataloader.BatchFunc {
	return diagnosisAndActionLoaderByID{}.loadBatch
}

func (ldr diagnosisAndActionLoaderByID) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			diagnosisAndAction, err := ctx.Value("diagnosisAndActionService").(*service.DiagnosisAndActionService).FindByID(key.String())
			results[i] = &dataloader.Result{Data: diagnosisAndAction, Error: err}
		}(i, key)
	}

	wg.Wait()

	return results
}

func LoadDiagnosisAndActionByID(ctx context.Context, key string) (*model.DiagnosisAndAction, error) {
	var diagnosisAndAction *model.DiagnosisAndAction

	ldr, err := extract(ctx, diagnosisAndActionLoaderByIDKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}

	diagnosisAndAction, ok := data.(*model.DiagnosisAndAction)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", diagnosisAndAction, data)
	}

	return diagnosisAndAction, nil
}

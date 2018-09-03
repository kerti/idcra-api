package loader

import (
	"fmt"
	"sync"

	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
	"golang.org/x/net/context"
	"gopkg.in/nicksrandall/dataloader.v5"
)

type surveyLoaderByID struct{}

func newSurveyLoaderByID() dataloader.BatchFunc {
	return surveyLoaderByID{}.loadBatch
}

func (ldr surveyLoaderByID) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(keys)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()
			survey, err := ctx.Value("surveyService").(*service.SurveyService).FindByID(key.String())
			results[i] = &dataloader.Result{Data: survey, Error: err}
		}(i, key)
	}

	wg.Wait()

	return results
}

func LoadSurveyByID(ctx context.Context, key string) (*model.Survey, error) {
	var survey *model.Survey

	ldr, err := extract(ctx, surveyLoaderByIDKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(key))()
	if err != nil {
		return nil, err
	}

	survey, ok := data.(*model.Survey)
	if !ok {
		return nil, fmt.Errorf("wrong type: the expected type is %T but got %T", survey, data)
	}

	return survey, nil
}

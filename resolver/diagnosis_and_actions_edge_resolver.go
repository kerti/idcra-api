package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/model"
)

type diagnosisAndActionsEdgeResolver struct {
	cursor graphql.ID
	model  *model.DiagnosisAndAction
}

func (r *diagnosisAndActionsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *diagnosisAndActionsEdgeResolver) Node() *diagnosisAndActionResolver {
	return &diagnosisAndActionResolver{d: r.model}
}

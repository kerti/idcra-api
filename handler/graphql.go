package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/kerti/idcra-api/loader"
)

type GraphQL struct {
	Schema  *graphql.Schema
	Loaders loader.LoaderCollection
}

func (h *GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	var responseJSON []byte

	// set all response params here
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST OPTIONS")

	// check for empty body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if len(body) > 0 {
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := h.Loaders.Attach(r.Context())

		response := h.Schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
		responseJSON, err = json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if _, err := w.Write(responseJSON); err != nil {
		log.Println(err)
	}
}

package handler

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/kerti/idcra-api/model"
	"github.com/kerti/idcra-api/service"
	uuid "github.com/satori/go.uuid"
)

func SurveyReport() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-type", "application/octet-stream")
		w.Header().Set("Content-type", "application/pdf")

		ctx := r.Context()
		filename := strings.TrimPrefix(r.URL.Path, "/reports/surveys/")
		idString := strings.Replace(filename, ".pdf", "", -1)
		id, err := uuid.FromString(idString)
		if err != nil {
			response := &model.Response{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			}
			writeResponse(w, response, response.Code)
			return
		}

		reportData, err := ctx.Value("reportService").(*service.ReportService).GenerateSurveyPDF(id)
		if err != nil {
			response := &model.Response{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			}
			writeResponse(w, response, response.Code)
			return
		}

		reportBytes := bytes.NewReader(reportData.Bytes())
		io.Copy(w, reportBytes)
	})
}

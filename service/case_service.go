package service

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kerti/idcra-api/model"
	"github.com/op/go-logging"
)

type CaseService struct {
	db  *sqlx.DB
	log *logging.Logger
}

func NewCaseService(db *sqlx.DB, log *logging.Logger) *CaseService {
	return &CaseService{db: db, log: log}
}

func (c *CaseService) FindByID(id string) (*model.Case, error) {
	caseObj := &model.Case{}

	caseSQL := `SELECT * FROM cases WHERE id = ?`
	udb := c.db.Unsafe()
	row := udb.QueryRowx(caseSQL, id)
	err := row.StructScan(caseObj)
	if err == sql.ErrNoRows {
		return caseObj, nil
	}
	if err != nil {
		c.log.Errorf("Error in retrieving case : %v", err)
		return nil, err
	}

	return caseObj, nil
}

func (c *CaseService) FindBySurveyID(surveyID *string) ([]*model.Case, error) {
	cases := make([]*model.Case, 0)
	caseSQL := `SELECT * FROM cases WHERE survey_id = ? ORDER BY created_at DESC;`

	err := c.db.Select(&cases, caseSQL, surveyID)
	if err != nil {
		return nil, err
	}

	return cases, nil
}

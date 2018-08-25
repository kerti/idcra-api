package service

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kerti/idcra-api/model"
	"github.com/op/go-logging"
)

type DiagnosisAndActionService struct {
	db  *sqlx.DB
	log *logging.Logger
}

func NewDiagnosisAndActionService(db *sqlx.DB, log *logging.Logger) *DiagnosisAndActionService {
	return &DiagnosisAndActionService{db: db, log: log}
}

func (d *DiagnosisAndActionService) FindByID(id string) (*model.DiagnosisAndAction, error) {
	diagnosisAndAction := &model.DiagnosisAndAction{}

	dnaSQL := `SELECT * FROM diagnosis_and_actions WHERE id = ?`
	udb := d.db.Unsafe()
	row := udb.QueryRowx(dnaSQL, id)
	err := row.StructScan(diagnosisAndAction)
	if err == sql.ErrNoRows {
		return diagnosisAndAction, nil
	}
	if err != nil {
		d.log.Errorf("Error in retrieving diagnosis and action : %v", err)
		return nil, err
	}

	return diagnosisAndAction, nil
}

func (d *DiagnosisAndActionService) List(first *int32, after *string) ([]*model.DiagnosisAndAction, error) {
	diagnosisAndActions := make([]*model.DiagnosisAndAction, 0)
	var fetchSize int32
	if first == nil {
		fetchSize = defaultListFetchSize
	} else {
		fetchSize = *first
	}

	if after != nil {
		dnaSQL := `SELECT * FROM diagnosis_and_actions WHERE created_at < (SELECT created_at FROM diagnosis_and_actions WHERE id = ?) ORDER BY created_at DESC LIMIT ?;`
		decodedIndex, _ := DecodeCursor(after)
		err := d.db.Select(&diagnosisAndActions, dnaSQL, decodedIndex, fetchSize)
		if err != nil {
			return nil, err
		}

		return diagnosisAndActions, nil
	}

	dnaSQL := `SELECT * FROM diagnosis_and_actions ORDER BY created_at DESC LIMIT ?;`
	err := d.db.Select(&diagnosisAndActions, dnaSQL, fetchSize)
	if err != nil {
		return nil, err
	}

	return diagnosisAndActions, nil
}

func (d *DiagnosisAndActionService) Count() (int, error) {
	var count int
	dnaSQL := `SELECT COUNT(*) FROM diagnosis_and_actions`
	err := d.db.Get(&count, dnaSQL)
	if err != nil {
		return 0, err
	}
	return count, nil
}

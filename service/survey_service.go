package service

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kerti/idcra-api/model"
	"github.com/op/go-logging"
)

type SurveyService struct {
	db          *sqlx.DB
	caseService *CaseService
	log         *logging.Logger
}

func NewSurveyService(db *sqlx.DB, caseService *CaseService, log *logging.Logger) *SurveyService {
	return &SurveyService{db: db, caseService: caseService, log: log}
}

func (s *SurveyService) FindByID(id string) (*model.Survey, error) {
	survey := &model.Survey{}

	surveySQL := `SELECT * FROM surveys WHERE id = ?`
	udb := s.db.Unsafe()
	row := udb.QueryRowx(surveySQL, id)
	err := row.StructScan(survey)
	if err == sql.ErrNoRows {
		return survey, nil
	}
	if err != nil {
		s.log.Errorf("Error in retrieving survey : %v", err)
		return nil, err
	}

	cases, err := s.caseService.FindBySurveyID(&survey.ID)
	if err != nil {
		s.log.Errorf("Error in retrieving cases : %v", err)
		return nil, err
	}
	survey.Cases = cases

	return survey, nil
}

func (s *SurveyService) TransactionalCreateSurvey(survey *model.Survey) (*model.Survey, error) {
	surveySQL := `
		INSERT INTO surveys
		(
			id, student_id, surveyor_id, date,
			s1q1, s1q2, s1q3, s1q4, s1q5, s1q6, s1q7,
			s2q1, s2q2, s2q3, s2q4, s2q5, s2q6, s2q7, s2q8, s2q9,
			lower_d, lower_e, lower_f, upper_d, upper_m, upper_f,
			subjective_score, created_at
		) VALUES (
			:id, :student_id, :surveyor_id, :date,
			:s1q1, :s1q2, :s1q3, :s1q4, :s1q5, :s1q6, :s1q7, :s2q1,
			:s2q2, :s2q3, :s2q4, :s2q5, :s2q6, :s2q7, :s2q8, :s2q9,
			:lower_d, :lower_e, :lower_f, :upper_d, :upper_m, :upper_f,
			:subjective_score, :created_at
		)`
	caseFoundSQL := `
		INSERT INTO cases
		(id, survey_id, tooth_number, diagnosis_and_action_id, created_at)
		VALUES
		(:id, :survey_id, :tooth_number, :diagnosis_and_action_id, :created_at)`

	err := Transact(s.db, func(tx *sqlx.Tx) error {
		// store survey
		if _, err := tx.NamedExec(surveySQL, survey); err != nil {
			return err
		}

		// store cases
		for _, c := range survey.Cases {
			if _, err := tx.NamedExec(caseFoundSQL, c); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.FindByID(survey.ID)
}

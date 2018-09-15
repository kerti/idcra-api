package service

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kerti/idcra-api/model"
	"github.com/op/go-logging"
	uuid "github.com/satori/go.uuid"
)

type SchoolService struct {
	db  *sqlx.DB
	log *logging.Logger
}

func NewSchoolService(db *sqlx.DB, log *logging.Logger) *SchoolService {
	return &SchoolService{db: db, log: log}
}

func (s *SchoolService) FindByName(name string) (*model.School, error) {
	school := &model.School{}

	schoolSQL := `SELECT * FROM schools WHERE name = ?`
	udb := s.db.Unsafe()
	row := udb.QueryRowx(schoolSQL, name)
	err := row.StructScan(school)
	if err == sql.ErrNoRows {
		return school, nil
	}
	if err != nil {
		s.log.Errorf("Error in retrieving school : %v", err)
		return nil, err
	}

	return school, nil
}

func (s *SchoolService) FindByID(id string) (*model.School, error) {
	school := &model.School{}

	schoolSQL := `SELECT * FROM schools WHERE id = ?`
	udb := s.db.Unsafe()
	row := udb.QueryRowx(schoolSQL, id)
	err := row.StructScan(school)
	if err == sql.ErrNoRows {
		return school, nil
	}
	if err != nil {
		s.log.Errorf("Error in retrieving school : %v", err)
		return nil, err
	}

	return school, nil
}

func (s *SchoolService) CreateSchool(school *model.School) (*model.School, error) {
	schoolID := uuid.NewV4()
	school.ID = schoolID.String()
	schoolSQL := `INSERT INTO schools (id, name) VALUES (:id, :name)`
	_, err := s.db.NamedExec(schoolSQL, school)
	if err != nil {
		return nil, err
	}
	return s.FindByID(school.ID)
}

func (s *SchoolService) List(first *int32, after *string) ([]*model.School, error) {
	schools := make([]*model.School, 0)
	var fetchSize int32
	if first == nil {
		fetchSize = defaultListFetchSize
	} else {
		fetchSize = *first
	}

	if after != nil {
		schoolSQL := `SELECT * FROM users WHERE created_at < (SELECT created_at FROM schools WHERE id = ?) ORDER BY created_at DESC LIMIT ?;`
		decodedIndex, _ := DecodeCursor(after)
		err := s.db.Select(&schools, schoolSQL, decodedIndex, fetchSize)
		if err != nil {
			return nil, err
		}
		return schools, nil
	}
	schoolSQL := `SELECT * FROM schools ORDER BY created_at DESC LIMIT ?;`
	err := s.db.Select(&schools, schoolSQL, fetchSize)
	if err != nil {
		return nil, err
	}

	return schools, nil
}

func (s *SchoolService) Count() (int, error) {
	var count int
	schoolSQL := `SELECT COUNT(*) FROM schools`
	err := s.db.Get(&count, schoolSQL)
	if err != nil {
		return 0, err
	}
	return count, nil
}

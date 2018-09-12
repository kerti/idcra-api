package service

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kerti/idcra-api/model"
	"github.com/op/go-logging"
	uuid "github.com/satori/go.uuid"
)

type StudentService struct {
	db  *sqlx.DB
	log *logging.Logger
}

func NewStudentService(db *sqlx.DB, log *logging.Logger) *StudentService {
	return &StudentService{db: db, log: log}
}

func (s *StudentService) FindByID(id string) (*model.Student, error) {
	student := &model.Student{}

	studentSQL := `SELECT * FROM students WHERE id = ?`
	udb := s.db.Unsafe()
	row := udb.QueryRowx(studentSQL, id)
	err := row.StructScan(student)
	if err == sql.ErrNoRows {
		return student, nil
	}
	if err != nil {
		s.log.Errorf("Error in retrieving student : %v", err)
		return nil, err
	}

	return student, nil
}

func (s *StudentService) FindBySchoolID(schoolID *string, keyword *string) (students []*model.Student, err error) {
	if keyword != nil {
		strKeyword := fmt.Sprintf("%s%s%s", "%", *keyword, "%")
		studentSQL := `SELECT * FROM students WHERE school_id = ? AND name LIKE ? ORDER BY name ASC;`
		err = s.db.Select(&students, studentSQL, schoolID, strKeyword)
	} else {
		studentSQL := `SELECT * FROM students WHERE school_id = ? ORDER BY name ASC;`
		err = s.db.Select(&students, studentSQL, schoolID)
	}
	return students, err
}

func (s *StudentService) CreateStudent(student *model.Student) (*model.Student, error) {
	studentID := uuid.NewV4()
	student.ID = studentID.String()
	studentSQL := `INSERT INTO students (id, name, date_of_birth, school_id, created_at) VALUES (:id, :name, :date_of_birth, :school_id, NOW())`
	_, err := s.db.NamedExec(studentSQL, student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentService) List(first *int32, after *string) ([]*model.Student, error) {
	students := make([]*model.Student, 0)
	var fetchSize int32
	if first == nil {
		fetchSize = defaultListFetchSize
	} else {
		fetchSize = *first
	}

	if after != nil {
		studentSQL := `SELECT * FROM students WHERE name > (SELECT name FROM students WHERE id = ?) ORDER BY name ASC LIMIT ?;`
		decodedIndex, _ := DecodeCursor(after)
		err := s.db.Select(&students, studentSQL, decodedIndex, fetchSize)
		if err != nil {
			return nil, err
		}
		return students, nil
	}
	studentSQL := `SELECT * FROM students ORDER BY name ASC LIMIT ?;`
	err := s.db.Select(&students, studentSQL, fetchSize)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentService) Count() (int, error) {
	var count int
	studentSQL := `SELECT COUNT(*) FROM students`
	err := s.db.Get(&count, studentSQL)
	if err != nil {
		return 0, err
	}
	return count, nil
}

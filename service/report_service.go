package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/kerti/idcra-api/model"
	"github.com/op/go-logging"
)

type ReportService struct {
	db  *sqlx.DB
	log *logging.Logger
}

func NewReportService(db *sqlx.DB, log *logging.Logger) *ReportService {
	return &ReportService{db: db, log: log}
}

func (s *ReportService) CostBreakdownBySchoolAndDateRange(schoolID string, startDate string, endDate string) ([]*model.CostReport, error) {
	results := make([]*model.CostReport, 0)

	reportSQL := `
	select
		d.action description,
		sum(d.unit_cost) cost
	from
		cases c
		left join surveys s on c.survey_id = s.id
		left join students st on s.student_id = st.id
		left join diagnosis_and_actions d on c.diagnosis_and_action_id = d.id
	where
		st.school_id = ?
		and s.date >= ?
		and s.date < ?
	group by
		d.action;`

	err := s.db.Select(&results, reportSQL, schoolID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	var totalCost float64
	for _, r := range results {
		totalCost += r.Cost
	}

	summary := &model.CostReport{
		Description: "Total",
		Cost:        totalCost,
	}

	results = append(results, summary)

	return results, nil
}

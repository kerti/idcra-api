package model

// CostReport is the cost report entity
type CostReport struct {
	Description string  `db:"description"`
	Cost        float64 `db:"cost"`
}

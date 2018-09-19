package resolver

import (
	"github.com/kerti/idcra-api/model"
)

type costReportResolver struct {
	c *model.CostReport
}

func (c *costReportResolver) Description() string {
	return c.c.Description
}

func (c *costReportResolver) Cost() float64 {
	return c.c.Cost
}

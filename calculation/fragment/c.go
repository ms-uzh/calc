package fragment

import (
	"github.com/fforootd/calc/calculation"
	"github.com/fforootd/calc/models"
)

func calculateC(previous float64, polyamine models.Polyamine, isFirst bool) float64 {
	c := previous + polyamine.Mass - float64(polyamine.Quaternary)*calculation.H
	if isFirst {
		c -= calculation.NH3
	}
	return c
}

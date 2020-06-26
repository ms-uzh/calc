package fragment

import (
	"github.com/ms-uzh/calc/calculation"
	"github.com/ms-uzh/calc/models"
)

func calculateC(previous float64, polyamine models.Polyamine, isFirst bool) float64 {
	c := previous + polyamine.Mass - float64(polyamine.Quaternary)*calculation.H - massElectron
	if isFirst {
		c -= calculation.NH3
	}
	return c
}

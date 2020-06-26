package fragment

import (
	"github.com/ms-uzh/calc/calculation"
	"github.com/ms-uzh/calc/models"
)

func calculateA(previous float64, polyamine models.Polyamine, isFirst bool) float64 {
	a := previous + polyamine.Mass - (float64(polyamine.Quaternary) * calculation.H)
	if isFirst {
		a -= calculation.MassElectron
	}
	return a
}

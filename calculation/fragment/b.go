package fragment

import (
	"github.com/ms-uzh/calc/calculation"
	"github.com/ms-uzh/calc/models"
)

func calculateB(previous float64, polyamine models.Polyamine, isFirst bool) float64 {
	b := previous + polyamine.Mass - (float64(polyamine.Quaternary) * calculation.H)
	if isFirst {
		b -= calculation.H2O
		b -= calculation.MassElectron
	}
	return b
}

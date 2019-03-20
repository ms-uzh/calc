package fragment

import (
	"github.com/fforootd/calc/calculation"
	"github.com/fforootd/calc/models"
)

func calculateB(previous float64, polyamine models.Polyamine, isFirst bool) float64 {
	b := previous + polyamine.Mass - (float64(polyamine.Quaternary) * calculation.H)
	if isFirst {
		b -= calculation.H2O
	}
	return b
}

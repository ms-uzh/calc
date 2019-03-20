package fragment

import (
	"github.com/fforootd/calc/calculation"
	"github.com/fforootd/calc/models"
)

func calculateTz(previous float64, currentPoly models.Polyamine, isFirst bool) float64 {
	if isFirst {
		return previous + currentPoly.Mass - float64(currentPoly.Quaternary)*calculation.H - calculation.NH + calculation.NH3
	}
	return previous + currentPoly.Mass - float64(currentPoly.Quaternary)*calculation.H
}

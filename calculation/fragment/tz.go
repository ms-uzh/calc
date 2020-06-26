package fragment

import (
	"github.com/ms-uzh/calc/calculation"
	"github.com/ms-uzh/calc/models"
)

func calculateTz(previous float64, currentPoly models.Polyamine, isFirst bool) float64 {
	if isFirst {
		return previous + currentPoly.Mass - float64(currentPoly.Quaternary)*calculation.H - calculation.NH + calculation.NH3 - massElectron
	}
	return previous + currentPoly.Mass - float64(currentPoly.Quaternary)*calculation.H - massElectron
}

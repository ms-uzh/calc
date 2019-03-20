package fragment

import (
	"github.com/fforootd/calc/calculation"
	"github.com/fforootd/calc/models"
)

func calculateY(previous float64, tail models.Tail, currentPoly, previousPoly models.Polyamine, isFirst bool) float64 {
	if isFirst {
		return previous + currentPoly.Mass - currentPoly.Sub.Mass - float64(tail.Quaternary)*calculation.H - calculation.NH - calculation.NH3
	}
	return previous + currentPoly.Mass + previousPoly.Sub.Mass - currentPoly.Sub.Mass - float64(tail.Quaternary)*calculation.H
}

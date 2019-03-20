package fragment

import (
	"github.com/fforootd/calc/calculation"
	"github.com/fforootd/calc/models"
)

func calculateA(previous float64, polyamine models.Polyamine) float64 {
	return previous + polyamine.Mass - (float64(polyamine.Quaternary) * calculation.H)
}

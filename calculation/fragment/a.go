package fragment

import (
	"github.com/ms-uzh/calc/calculation"
	"github.com/ms-uzh/calc/models"
)

func calculateA(previous float64, polyamine models.Polyamine) float64 {
	return previous + polyamine.Mass - (float64(polyamine.Quaternary) * calculation.H) - massElectron
}

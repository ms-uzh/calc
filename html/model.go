package html

import (
	"github.com/fforootd/calc/models"
)

type PageData struct {
	*Choose
	*Calculation
}

type Calculation struct {
	GenericName     string
	ChemicalFormula []string
	MolecularMass   float64
}

type Choose struct {
	Heads      models.Heads
	Polyamines []models.Polyamines
	Tails      models.Tails
}

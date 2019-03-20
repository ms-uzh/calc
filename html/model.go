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
	Fragments       models.Fragments
}

type Choose struct {
	Heads      models.Heads
	Polyamines []models.Polyamines
	Tails      models.Tails
}

type Chosen struct {
	Head           string   `schema:"head"`
	PolyamineUnits []string `schema:"polyamineUnits"`
	Tail           string   `schema:"tail"`
}

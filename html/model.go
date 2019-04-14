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
	Precursor1      float64
	Precursor2      float64
	PrecursorHDX1   float64
	PrecursorHDX2   float64
	Fragments       models.Fragments
	HDX             uint
	Quaternary      int
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

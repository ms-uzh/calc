package models

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
	Fragments       Fragments
	HDX             uint
	Quaternary      int
	Hugo            string
}

type Choose struct {
	Heads      Heads
	Polyamines []Polyamines
	Tails      Tails
	Spiders    []Spiders
}

type Chosen struct {
	Head           string   `schema:"head"`
	PolyamineUnits []string `schema:"polyamineUnits"`
	Tail           string   `schema:"tail"`
	Spiders        []string `schema:"spiders"`
}

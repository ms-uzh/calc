package html

import (
	"html/template"
	"net/http"

	calc "github.com/fforootd/calc/calculation"
	"github.com/fforootd/calc/calculation/fragment"

	"github.com/fforootd/calc/models"

	"github.com/fforootd/calc/config"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type Server struct {
	handler http.Handler
	tmpl    *template.Template
	conf    *config.Config
}

func NewServer(conf *config.Config) *Server {
	polyamineSelectors := []models.Polyamines{}
	for i := uint(0); i < conf.App.MaxPolyamineSelectors; i++ {
		polyamineSelectors = append(polyamineSelectors, conf.Polynamines)
	}
	server := &Server{conf: conf}
	server.tmpl = template.Must(template.New("index.html").Funcs(server.getFuncs()).ParseFiles("./html/index.html", "./html/calculation.html", "./html/choose.html"))
	return server
}

func (s *Server) processPost(head models.Head, tail models.Tail, polyamines ...models.Polyamine) *PageData {
	data := &PageData{Choose: &Choose{}, Calculation: &Calculation{}}
	data.Choose.Polyamines = make([]models.Polyamines, s.conf.App.MaxPolyamineSelectors)

	data.Choose.Heads = s.conf.Heads.SetSelected(head.Name)
	for i := uint(0); i < s.conf.App.MaxPolyamineSelectors; i++ {
		polyName := "-"
		if int(i) < len(polyamines) {
			polyName = polyamines[i].Name
		}
		data.Choose.Polyamines[i] = s.conf.Polynamines.SetSelected(polyName)
	}
	data.Choose.Tails = s.conf.Tails.SetSelected(tail.Name)

	data.Calculation.GenericName = calc.CalculateName(head, tail, polyamines...)
	data.Calculation.ChemicalFormula, _ = calc.CalculateFormula(head, tail, polyamines...)
	data.Calculation.MolecularMass = calc.CalculateMass(head, tail, polyamines...)
	data.Calculation.Fragments = fragment.CalculateFragments(head, tail, polyamines...)
	data.Calculation.Precursor1 = calc.CalculatePrecursor1(head, tail, polyamines...)
	data.Calculation.Precursor2 = calc.CalculatePrecursor2(head, tail, polyamines...)
	data.Calculation.PrecursorHDX1 = calc.CalculatePrecursorHDX1(head, tail, polyamines...)
	data.Calculation.PrecursorHDX2 = calc.CalculatePrecursorHDX2(head, tail, polyamines...)
	data.Calculation.HDX = calc.CalculateHDX(head, tail, polyamines...)
	data.Calculation.Quaternary = calc.CalculateQuaternary(head, tail, polyamines...)
	return data
}

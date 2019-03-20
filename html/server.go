package html

import (
	"html/template"
	"log"
	"net/http"

	calc "github.com/fforootd/calc/calculation"
	"github.com/fforootd/calc/calculation/fragment"

	"github.com/fforootd/calc/models"

	"github.com/yabslabs/utils/logging"

	"github.com/fforootd/calc/config"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func (s *Server) ListenAndServe() {
	routeHandler := mux.NewRouter()
	routeHandler.PathPrefix("/resources").Handler(s.handleResources())
	routeHandler.Handle("/calc", s.handle())
	s.handler = routeHandler
	// http.Handle("/", s.handle())

	if err := http.ListenAndServe(":"+s.conf.App.Port, s.handler); err != nil {
		log.Fatal("oups: ", err)
	}
}

type Server struct {
	data    *PageData
	handler http.Handler
	tmpl    *template.Template
	conf    *config.Config
}

func NewServer(conf *config.Config) *Server {
	polyamineSelectors := []models.Polyamines{}
	for i := uint(0); i < conf.App.MaxPolyamineSelectors; i++ {
		polyamineSelectors = append(polyamineSelectors, conf.Polynamines)
	}
	server := &Server{
		data: &PageData{
			Choose: &Choose{
				Heads:      conf.Heads,
				Polyamines: polyamineSelectors,
				Tails:      conf.Tails,
			},
		},
		conf: conf,
	}
	server.tmpl = template.Must(template.New("index.html").Funcs(server.getFuncs()).ParseFiles("./html/index.html", "./html/calculation.html", "./html/choose.html"))
	return server
}

func (s *Server) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data *PageData
		switch r.Method {
		case http.MethodPost:
			data = s.handlePost(r)
		case http.MethodGet:
			data = s.handleGet(r)
		}
		if err := s.tmpl.Execute(w, data); err != nil {
			log.Println("CALC-sAMak: ", err)
		}
	}
}

func (s *Server) handleGet(r *http.Request) *PageData {
	logging.Log("HTML-Uko25").Info("get")
	return s.data
}

func (s *Server) handlePost(r *http.Request) *PageData {
	if err := r.ParseForm(); err != nil {
		log.Fatal("HTML-TzPhd: ", err)
	}
	var chosenInput Chosen

	err := decoder.Decode(&chosenInput, r.PostForm)
	logging.Log("HTML-nYlSm").OnError(err).Warn("decode failed")

	head := s.data.Heads.GetByName(chosenInput.Head)
	tail := s.data.Tails.GetByName(chosenInput.Tail)
	polyamines := s.getPolyamines(chosenInput.PolyamineUnits...)

	data := new(PageData)
	*data = *s.data
	data.Calculation = &Calculation{}
	data.Calculation.GenericName = calc.CalculateName(head, tail, polyamines...)
	data.Calculation.ChemicalFormula, _ = calc.CalculateFormula(head, tail, polyamines...)
	data.Calculation.MolecularMass = calc.CalculateMass(head, tail, polyamines...)
	data.Calculation.Fragments = fragment.CalculateFragments(head, tail, polyamines...)
	return data
}

func (s *Server) getPolyamines(names ...string) models.Polyamines {
	polyamines := models.Polyamines{}
	for i, name := range names {
		if name == "-" {
			return polyamines
		}
		polyamines = append(polyamines, s.data.Polyamines[i].GetByName(name))
	}
	return polyamines
}

func (s *Server) handleResources() http.Handler {
	return http.StripPrefix("/resources", http.FileServer(http.Dir("./html/resources/")))
}

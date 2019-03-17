package html

import (
	"html/template"
	"log"
	"net/http"

	calc "github.com/fforootd/calc/calculation"

	"github.com/fforootd/calc/models"

	"github.com/yabslabs/utils/logging"

	"github.com/fforootd/calc/config"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func (s *Server) ListenAndServce() {
	http.Handle("/calc", s.handle())
	http.Handle("/", s.handle())
	if err := http.ListenAndServe(":"+s.conf.App.Port, nil); err != nil {
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

type Post struct {
	Head           string   `schema:"head"`
	PolyamineUnits []string `schema:"polyamineUnits"`
	Tail           string   `schema:"tail"`
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
	var formPost Post

	err := decoder.Decode(&formPost, r.PostForm)
	logging.Log("HTML-nYlSm").OnError(err).Warn("decode failed")

	head := s.data.Heads.GetByName(formPost.Head)
	tail := s.data.Tails.GetByName(formPost.Tail)
	polyamines := s.getPolyamines(formPost.PolyamineUnits...)

	data := new(PageData)
	*data = *s.data
	data.Calculation = &Calculation{}
	data.Calculation.GenericName = calc.CalculateName(head, tail, polyamines...)
	data.Calculation.ChemicalFormula, _ = calc.CalculateFormula(head, tail, polyamines...)
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

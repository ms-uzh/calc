package html

import (
	"log"
	"net/http"

	"github.com/fforootd/calc/models"
	"github.com/gorilla/mux"

	"github.com/yabslabs/utils/logging"
)

func (s *Server) ListenAndServe() {
	routeHandler := mux.NewRouter()
	routeHandler.PathPrefix("/resources").Handler(s.handleResources())
	routeHandler.Handle("/calc", s.handle())
	routeHandler.Handle("/", s.handle())
	s.handler = routeHandler

	if err := http.ListenAndServe(":"+s.conf.App.Port, s.handler); err != nil {
		log.Fatal("oups: ", err)
	}
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

func (s *Server) handleResources() http.Handler {
	return http.StripPrefix("/resources", http.FileServer(http.Dir("./html/resources/")))
}

func (s *Server) handleGet(r *http.Request) *PageData {
	logging.Log("HTML-Uko25").Info("get")
	data := &PageData{Choose: &Choose{
		Heads:      s.conf.Heads,
		Tails:      s.conf.Tails,
		Polyamines: make([]models.Polyamines, s.conf.App.MaxPolyamineSelectors),
	}}
	for i := uint(0); i < s.conf.App.MaxPolyamineSelectors; i++ {
		data.Choose.Polyamines[i] = s.conf.Polynamines
	}
	return data
}

func (s *Server) handlePost(r *http.Request) *PageData {
	if err := r.ParseForm(); err != nil {
		log.Fatal("HTML-TzPhd: ", err)
	}
	var chosenInput Chosen

	err := decoder.Decode(&chosenInput, r.PostForm)
	logging.Log("HTML-nYlSm").OnError(err).Warn("decode failed")

	head := s.conf.Heads.GetByName(chosenInput.Head)
	tail := s.conf.Tails.GetByName(chosenInput.Tail)
	polyamines := s.getPolyamines(chosenInput.PolyamineUnits...)

	return s.processPost(head, tail, polyamines...)
}

func (s *Server) getPolyamines(names ...string) models.Polyamines {
	polyamines := models.Polyamines{}
	for _, name := range names {
		if name == "-" {
			return polyamines
		}
		polyamines = append(polyamines, s.conf.Polynamines.GetByName(name))
	}
	return polyamines
}

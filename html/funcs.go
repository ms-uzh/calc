package html

import (
	"fmt"

	"github.com/fforootd/calc/templates"
)

func (s *Server) getFuncs() map[string]interface{} {
	return map[string]interface{}{
		"joinChemicalFormula":  templates.JoinChemicalFormula,
		"getCalcAddress":       s.generateGetCalcAddress(),
		"generateChemicalName": templates.GenerateChemicalName,
		"getCSSPath":           s.generateCSSPath(),
		"round":                templates.Round,
		"index":                templates.Index,
	}
}

func (s *Server) generateGetCalcAddress() func() string {
	return func() string {
		if s.conf.App.RedirectPort == "" {
			return fmt.Sprintf("%s/calc", s.conf.App.URL)
		}
		return fmt.Sprintf("%s:%s/calc", s.conf.App.URL, s.conf.App.RedirectPort)
	}
}

func (s *Server) generateCSSPath() func() string {
	return func() string {
		return "/resources/css/styles.css"
	}
}

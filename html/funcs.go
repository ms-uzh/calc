package html

import (
	"encoding/base64"
	"fmt"

	"github.com/fforootd/calc/templates"
)

func (s *Server) getFuncs() map[string]interface{} {
	return map[string]interface{}{
		"joinChemicalFormula":  templates.JoinChemicalFormula,
		"getCalcAddress":       s.generateGetCalcAddress(),
		"generateChemicalName": templates.GenerateChemicalName,
		"getCSSPath":           s.generateCSSPath(),
		"getJavascriptPath":    s.generateJavascriptPath(),
		"round":                templates.Round,
		"index":                templates.Index,
		"toBase64":             toBase64,
	}
}

func toBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
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

func (s *Server) generateJavascriptPath() func() string {
	return func() string {
		return "/resources/js"
	}
}

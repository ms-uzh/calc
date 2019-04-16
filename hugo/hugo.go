package hugo

import (
	"bytes"
	"html/template"

	"github.com/yabslabs/utils/logging"

	"github.com/fforootd/calc/models"

	"github.com/fforootd/calc/templates"
)

type Hugo struct {
	tmpl *template.Template
}

func NewHugo() (hugo Hugo) {
	hugo.tmpl = template.Must(template.New("hugo.md.tmpl").Funcs(hugo.getFuncs()).ParseFiles("./templates/hugo/hugo.md.tmpl"))
	return hugo
}

func (h Hugo) Exec(calc *models.Calculation) string {
	var writer bytes.Buffer
	if err := h.tmpl.Execute(&writer, calc); err != nil {
		logging.Log("HUGO-hqbK8").WithError(err).Warn("unable to execute hugo tempalte")
		return ""
	}
	return writer.String()
}

func (s *Hugo) getFuncs() map[string]interface{} {
	return map[string]interface{}{
		"joinChemicalFormula":  templates.JoinChemicalFormula,
		"generateChemicalName": templates.GenerateChemicalName,
		"round":                templates.Round,
		"index":                templates.Index,
	}
}

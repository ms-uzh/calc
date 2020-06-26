package hugo

import (
	"bytes"
	"html/template"

	"github.com/caos/logging"

	"github.com/ms-uzh/calc/models"

	"github.com/ms-uzh/calc/templates"
)

type Hugo struct {
	tmpl *template.Template
}

type hugoData struct {
	*models.Calculation
	models.Spiders
}

func NewHugo() (hugo Hugo) {
	hugo.tmpl = template.Must(template.New("hugo.md.tmpl").Funcs(hugo.getFuncs()).ParseFiles("./templates/hugo/hugo.md.tmpl"))
	return hugo
}

func (h Hugo) Exec(calc *models.Calculation, spiders models.Spiders) string {
	var writer bytes.Buffer
	data := hugoData{
		calc,
		spiders,
	}
	if err := h.tmpl.Execute(&writer, &data); err != nil {
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

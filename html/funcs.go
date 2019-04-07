package html

import (
	"fmt"
	"regexp"
)

var (
	superRegexp      = regexp.MustCompile(`([0-9]|)\+`)
	numberRegexp     = regexp.MustCompile("[0-9]+")
	inBracketsRegexp = regexp.MustCompile(`\(.*[0-9]+.*\)`)
)

func (s *Server) getFuncs() map[string]interface{} {
	return map[string]interface{}{
		"joinChemicalFormula":  joinChemicalFormula,
		"getCalcAddress":       s.generateGetCalcAddress(),
		"generateChemicalName": generateChemicalName,
		"getCSSPath":           s.generateCSSPath(),
		"round":                round,
	}
}

func (s *Server) generateGetCalcAddress() func() string {
	return func() string {
		return fmt.Sprintf("%s:%s/calc", s.conf.App.URL, s.conf.App.Port)
	}
}

func (s *Server) generateCSSPath() func() string {
	return func() string {
		return "/resources/css/styles.css"
	}
}

func round(number float64) string {
	return fmt.Sprintf("%.5f", number)
}

func joinChemicalFormula(texts []string) (joined string) {
	for _, text := range texts {
		if text != "" {
			joined += text
		}
	}
	joined = superRegexp.ReplaceAllStringFunc(joined, super)
	joined = numberRegexp.ReplaceAllStringFunc(joined, subscript)
	return joined
}

func generateChemicalName(text string) string {
	text = superRegexp.ReplaceAllStringFunc(text, super)
	// text = inBracketsRegexp.ReplaceAllStringFunc(text, subscript)
	return text
}

func super(text string) (replaced string) {
	for _, numb := range text {
		switch string(numb) {
		case "1":
			replaced = replaced + "¹"
		case "2":
			replaced = replaced + "²"
		case "3":
			replaced = replaced + "³"
		case "4":
			replaced = replaced + "⁴"
		case "5":
			replaced = replaced + "⁵"
		case "6":
			replaced = replaced + "⁶"
		case "7":
			replaced = replaced + "⁷"
		case "8":
			replaced = replaced + "⁸"
		case "9":
			replaced = replaced + "⁹"
		case "0":
			replaced = replaced + "⁰"
		case "+":
			replaced = replaced + "⁺"
		}
	}
	return replaced
}

func subscript(text string) (replaced string) {
	for _, numb := range text {
		switch string(numb) {
		case "1":
			replaced = replaced + "₁"
		case "2":
			replaced = replaced + "₂"
		case "3":
			replaced = replaced + "₃"
		case "4":
			replaced = replaced + "₄"
		case "5":
			replaced = replaced + "₅"
		case "6":
			replaced = replaced + "₆"
		case "7":
			replaced = replaced + "₇"
		case "8":
			replaced = replaced + "₈"
		case "9":
			replaced = replaced + "₉"
		case "0":
			replaced = replaced + "₀"
		}
	}
	return replaced
}

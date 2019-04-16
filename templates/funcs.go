package templates

import (
	"fmt"
	"regexp"
)

var (
	superRegexp  = regexp.MustCompile(`([0-9]|)\+`)
	numberRegexp = regexp.MustCompile("[0-9]+")
)

func Round(number float64) string {
	return fmt.Sprintf("%.5f", number)
}

func JoinChemicalFormula(texts []string) (joined string) {
	for _, text := range texts {
		text = superRegexp.ReplaceAllStringFunc(text, super)
		text = numberRegexp.ReplaceAllStringFunc(text, subscript)
		if text != "" {
			joined += text
		}
	}
	return joined
}

func GenerateChemicalName(text string) string {
	text = superRegexp.ReplaceAllStringFunc(text, super)
	return text
}

func Index(numb int) int {
	return numb + 1
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

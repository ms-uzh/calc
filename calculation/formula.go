package calculation

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/pkg/errors"

	"github.com/ms-uzh/calc/models"
)

var numbRegexp = regexp.MustCompile("[0-9]")

type Formulas []*Formula

func (f Formulas) Len() int           { return len(f) }
func (f Formulas) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f Formulas) Less(i, j int) bool { return f[i].name < f[j].name }

type Formula struct {
	name  string
	count int
}

func (f Formulas) set(formula *Formula) Formulas {
	for _, existingFormula := range f {
		if existingFormula.name == formula.name {
			existingFormula.count += formula.count
			return f
		}
	}
	f = append(f, formula)
	return f
}

func (f Formulas) join() []string {
	formulas := []string{}
	sort.Sort(f)
	for _, formula := range f {
		formulas = append(formulas, fmt.Sprintf("%s%v", formula.name, formula.count))
	}
	return formulas
}

func CalculateFormula(head models.Head, tail models.Tail, polyamines ...models.Polyamine) (formula []string, err error) {
	formulas := Formulas{}
	formulas, err = caclulateFormulas(formulas, head.Formula...)
	if err != nil {
		return nil, err
	}
	for _, polyamine := range polyamines {
		formulas, err = caclulateFormulas(formulas, polyamine.Formula...)
		if err != nil {
			return nil, err
		}
	}
	formulas, err = caclulateFormulas(formulas, tail.Formula...)
	if err != nil {
		return nil, err
	}

	quat := CalculateQuaternary(head, tail, polyamines...)
	quatName := CalculateQuaternaryName(quat)

	formulasJoined := formulas.join()
	if quatName != "" {
		formulasJoined = append(formulas.join(), quatName)
	}

	return formulasJoined, nil
}

func caclulateFormulas(existingFormulas Formulas, formulas ...string) (Formulas, error) {
	for _, formulaText := range formulas {
		formula, err := createFormula(formulaText)
		if err != nil {
			return nil, err
		}
		existingFormulas = existingFormulas.set(formula)
	}
	return existingFormulas, nil
}

func createFormula(formula string) (_ *Formula, err error) {
	indexes := numbRegexp.FindStringIndex(formula)
	count := 1
	name := formula
	if indexes != nil {
		countText := formula[indexes[0]:]
		count, err = strconv.Atoi(countText)
		if err != nil {
			return nil, errors.WithMessage(err, "wrong formula. should be characters numbers. eg C10")
		}
		name = formula[:indexes[0]]
	}

	return &Formula{name: name, count: count}, nil
}

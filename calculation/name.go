package calculation

import (
	"fmt"

	"github.com/ms-uzh/calc/models"
)

func CalculateName(head models.Head, tail models.Tail, polyamines ...models.Polyamine) (name string) {
	quaternary := CalculateQuaternary(head, tail, polyamines...)

	name = head.Name
	for _, polyamine := range polyamines {
		name += polyamine.Name
	}
	name += tail.Name
	name += CalculateQuaternaryName(quaternary)

	return name
}

func CalculateQuaternaryName(quaternary int) string {
	if quaternary == 0 {
		return ""
	}
	if quaternary == 1 {
		return "+"
	}
	return fmt.Sprint(quaternary, "+")
}

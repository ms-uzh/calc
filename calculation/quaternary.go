package calculation

import "github.com/ms-uzh/calc/models"

func CalculateQuaternary(head models.Head, tail models.Tail, polyamines ...models.Polyamine) (quaternary int) {
	quaternary = head.Quaternary
	quaternary += tail.Quaternary
	for _, polyamine := range polyamines {
		quaternary += polyamine.Quaternary
	}
	return quaternary
}

package calculation

import "github.com/ms-uzh/calc/models"

func CalculateMass(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	mass := head.Mass
	for _, polyamine := range polyamines {
		mass += polyamine.Mass
	}
	mass += tail.Mass

	quaternary := CalculateQuaternary(head, tail, polyamines...)
	mass -= MassElectron * float64(quaternary)

	return mass
}

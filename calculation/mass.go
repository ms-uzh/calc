package calculation

import "github.com/fforootd/calc/models"

func calculateMass(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	mass := head.Mass
	for _, polyamine := range polyamines {
		mass += polyamine.Mass
	}
	mass += tail.Mass
	return mass
}

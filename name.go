package main

import (
	"fmt"

	"github.com/fforootd/calc/models"
)

func generateName(head models.Head, tail models.Tail, polyamines ...models.Polyamine) (name string) {
	quaternary := calculateQuaternary(head, tail, polyamines...)

	name = head.Name
	for _, polyamine := range polyamines {
		name += polyamine.Name
	}
	name += tail.Name
	name += generateQuaternaryName(quaternary)

	return name
}

func generateQuaternaryName(quaternary int) string {
	if quaternary == 0 {
		return ""
	}
	if quaternary == 1 {
		return "+"
	}
	return fmt.Sprint(quaternary, "+")
}

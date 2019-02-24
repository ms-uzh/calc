package utils

import (
	"github.com/fforootd/calc/models"
)

func CalcTotal(input ...float64) float64 {
	total := 0.0
	for _, value := range input {
		total += value
	}
	return total
}

func CallCulateFragements(head, poly, tail float64) models.Fragment {
	a := CallculateA(head, poly)
	b := CallculateB(head, poly)
	c := CallculateC(head, poly)
	ta := CallculateTa(head, poly)
	z := CallculateZ(tail, poly)
	y := CallculateY(tail, poly)
	tz := CallculateTz(tail, poly)

	x := models.Fragment{A: a, B: b, C: c, Ta: ta, Z: z, Y: y, Tz: tz}

	return x
}

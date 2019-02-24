package utils

import (
	"github.com/fforootd/calc/models"
)

const c1 = 1.0078250

//CalcTotal calculates the total mass
func CalcTotal(input ...float64) float64 {
	total := 0.0
	for _, value := range input {
		total += value
	}
	return total
}

//CallCulateFragements offloads the generation of the fragments
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

//CalcPreCursorOne
//If no Object is quanterary add c1 to the mass
//If one Object is "" do nothing, return total mass
//If more than one "" return null
func CalcPreCursorOne(input ...float64) float64 {

	return x
}

//if no object is quartenary add twice c1
//if one object is "" add once c1
//if two "" do nothing, return total mass
//if quanternry >2 return null
//
//always half the return value
func CalcPreCursorTwo(input ...float64) float64 {
	y := 12 // TODO implement
	x := y / 2
	return x
}

//add up HDX values
func CalcHDX(input ...float64) float64 {

	return x
}

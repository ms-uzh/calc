package fragment

import (
	"github.com/fforootd/calc/models"
)

func CalculateFragments(head models.Head, tail models.Tail, polyamines ...models.Polyamine) models.Fragments {
	fragments := make([]models.Fragment, len(polyamines))
	fragments = calculateFromHead(fragments, head, tail, polyamines...)
	fragments = calculateFromTail(fragments, tail, polyamines...)

	return fragments
}

func calculateFromTail(fragments models.Fragments, tail models.Tail, polyamines ...models.Polyamine) models.Fragments {
	insertIndex := 0
	previous := models.Fragment{Z: tail.Mass, Y: tail.Mass, Tz: tail.Mass}

	for iterateIndex := len(polyamines) - 1; iterateIndex >= 0; iterateIndex-- {
		polyamine := polyamines[iterateIndex]
		previousPolyamine := models.Polyamine{}
		if iterateIndex != len(polyamines)-1 {
			previousPolyamine = polyamines[iterateIndex+1]
		}
		isFirst := iterateIndex == len(polyamines)-1

		z := calculateZ(previous.Z, tail, polyamine, previousPolyamine, isFirst)
		previous.Z = z
		fragments[insertIndex].Z = z

		y := calculateY(previous.Y, tail, polyamine, previousPolyamine, isFirst)
		previous.Y = y
		fragments[insertIndex].Y = y

		tz := calculateTz(previous.Tz, polyamine, isFirst)
		previous.Tz = tz
		fragments[insertIndex].Tz = tz

		insertIndex++
	}
	return fragments
}

func calculateFromHead(fragments models.Fragments, head models.Head, tail models.Tail, polyamines ...models.Polyamine) models.Fragments {
	previous := models.Fragment{A: head.Mass, B: head.Mass, C: head.Mass, Ta: head.Mass}
	for i, polyamine := range polyamines {
		a := calculateA(previous.A, polyamine)
		previous.A = a

		b := calculateB(previous.B, polyamine, i == 0)
		previous.B = b

		c := calculateC(previous.C, polyamine, i == 0)
		previous.C = c

		isLast := i == len(polyamines)-1
		followingPolyamine := models.Polyamine{}
		if !isLast {
			followingPolyamine = polyamines[i+1]
		}
		ta := calculateTA(previous.Ta, polyamine, followingPolyamine, tail, i == 0, isLast)
		previous.Ta = ta

		fragment := models.Fragment{
			A:  a,
			B:  b,
			C:  c,
			Ta: ta,
		}

		fragments[i] = fragment
	}
	return fragments
}

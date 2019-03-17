package calculation

import (
	"github.com/fforootd/calc/models"
)

func calculateFragments(head models.Head, polyamines ...models.Polyamine) []models.Fragment {
	fragments := make([]models.Fragment, len(polyamines))
	previous := models.Fragment{A: head.Mass, B: head.Mass, C: head.Mass, Ta: head.Mass}
	for i, polyamine := range polyamines {
		a := calculateA(previous.A, polyamine)
		previous.A = a
		b := calculateB(previous.B, polyamine, i == 0)
		previous.B = b
		c := calculateC(previous.C, polyamine, i == 0)
		previous.C = c
		// ta := calculateTA(previous.Ta, polyamine)
		// previous.Ta = ta
		fragment := models.Fragment{
			A: a,
			B: b,
			C: c,
			// Ta: ta,
		}

		fragments[i] = fragment
	}

	return fragments
}

func calculateA(previous float64, polyamine models.Polyamine) float64 {
	return previous + polyamine.Mass - (float64(polyamine.Quaternary) * h)
}

func calculateB(previous float64, polyamine models.Polyamine, isFirst bool) float64 {
	b := previous + polyamine.Mass - (float64(polyamine.Quaternary) * h)
	if isFirst {
		b -= h2o
	}
	return b
}

func calculateC(previous float64, polyamine models.Polyamine, isFirst bool) float64 {
	c := previous + polyamine.Mass - float64(polyamine.Quaternary)*h
	if isFirst {
		c -= nh3
	}
	return c
}

func calculateTA(previous float64, polyamine, followingPolyamine models.Polyamine, isFirst bool) float64 {
	ta := previous + polyamine.Mass
	if isFirst {
		ta += followingPolyamine.Sub.Mass
		ta -= float64(followingPolyamine.Quaternary) * h
	}
	ta += nh3
	return ta
}

func calculateFirstTA(head float64, polyamine, following *models.Polyamine) float64 {
	ta1 := head + polyamine.Mass - float64(following.Quaternary)*h + nh3
	ta1 += following.Sub.Mass
	return ta1
}

func calculateMiddleTAs(previous float64, polyamine, following *models.Polyamine) float64 {
	ta2 := previous + polyamine.Mass - float64(following.Quaternary)*h
	ta2 += polyamine.Sub.Mass
	ta2 += following.Sub.Mass
	return ta2
}

func calculateLastTA(previous float64, polyamine, previousPolyamine models.Polyamine, tail models.Tail) float64 {
	ta3 := previous + polyamine.Mass - float64(tail.Quaternary)*h
	ta3 -= previousPolyamine.Sub.Mass
	ta3 += tail.Sub.Mass
	return ta3
}

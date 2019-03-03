package main

import (
	"github.com/fforootd/calc/models"
)

func calculatePrecursor1(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	precursor := generateMass(head, tail, polyamines...)

	quaternary := calculateQuaternary(head, tail, polyamines...)
	correcture := precursor1Correcture(quaternary)
	if correcture == -1 {
		return -1
	}

	precursor += correcture

	return precursor
}

func calculatePrecursor2(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	precursor := generateMass(head, tail, polyamines...)

	quaternary := calculateQuaternary(head, tail, polyamines...)
	correcture := precursor2Correcture(quaternary)
	if correcture == -1 {
		return -1
	}

	precursor += correcture
	precursor *= 0.5

	return precursor
}

func calculatePrecursorHDX1(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	precursor := generateMass(head, tail, polyamines...)

	quaternary := calculateQuaternary(head, tail, polyamines...)
	hdx := calculateHDX(head, tail, polyamines...)
	correcture := precursorHDX1Correcture(quaternary, float64(hdx))
	if correcture == -1 {
		return -1
	}

	precursor += correcture
	precursor *= 0.5

	return precursor
}

func calculatePrecursorHDX2(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	precursor := generateMass(head, tail, polyamines...)

	quaternary := calculateQuaternary(head, tail, polyamines...)
	hdx := calculateHDX(head, tail, polyamines...)
	correcture := precursorHDX2Correcture(quaternary, float64(hdx))
	if correcture == -1 {
		return -1
	}

	precursor += correcture
	precursor *= 0.5

	return precursor
}

func precursor1Correcture(quaternary int) float64 {
	if quaternary == 0 {
		return h
	}
	if quaternary == 1 {
		return 0
	}
	return -1
}

func precursor2Correcture(quaternary int) float64 {
	if quaternary == 0 {
		return 2 * h
	}
	if quaternary == 1 {
		return h
	}
	if quaternary == 2 {
		return 0
	}
	return -1
}

func precursorHDX1Correcture(quaternary int, hdx float64) float64 {
	if quaternary == 0 {
		return hdx*h + (hdx+1)*d
	}
	if quaternary == 1 {
		return hdx*h + hdx*d
	}
	return -1
}

func precursorHDX2Correcture(quaternary int, hdx float64) float64 {
	if quaternary == 0 {
		return hdx*h + (hdx+2)*d
	}
	if quaternary == 1 {
		return hdx*h + (hdx+1)*d
	}
	if quaternary == 2 {
		return hdx*h + hdx*d
	}
	return -1
}

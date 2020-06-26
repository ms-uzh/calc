package calculation

import (
	"github.com/ms-uzh/calc/models"
)

func CalculatePrecursor1(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	precursor := CalculateMass(head, tail, polyamines...)

	quaternary := CalculateQuaternary(head, tail, polyamines...)
	correcture := precursor1Correcture(quaternary)
	if correcture == -1 {
		return -1
	}

	if quaternary == 0 {
		precursor -= MassElectron
	}
	precursor += correcture

	return precursor
}

func CalculatePrecursor2(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	precursor := CalculateMass(head, tail, polyamines...)

	quaternary := CalculateQuaternary(head, tail, polyamines...)
	correcture := precursor2Correcture(quaternary)
	if correcture == -1 {
		return -1
	}

	if quaternary == 0 {
		precursor -= 2 * MassElectron
	} else if quaternary == 1 {
		precursor -= MassElectron
	}

	precursor += correcture
	precursor *= 0.5

	return precursor
}

func CalculatePrecursorHDX1(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	precursor := CalculateMass(head, tail, polyamines...)
	quaternary := CalculateQuaternary(head, tail, polyamines...)
	hdx := CalculateHDX(head, tail, polyamines...)

	correcture := precursorHDX1Correcture(quaternary, float64(hdx))
	if correcture == -1 {
		return -1
	}

	if quaternary == 0 {
		precursor -= MassElectron
	}

	precursor += correcture

	return precursor
}

func CalculatePrecursorHDX2(head models.Head, tail models.Tail, polyamines ...models.Polyamine) float64 {
	precursor := CalculateMass(head, tail, polyamines...)
	quaternary := CalculateQuaternary(head, tail, polyamines...)
	hdx := CalculateHDX(head, tail, polyamines...)

	correcture := precursorHDX2Correcture(quaternary, float64(hdx))
	if correcture == -1 {
		return -1
	}

	if quaternary == 0 {
		precursor -= 2 * MassElectron
	} else if quaternary == 1 {
		precursor -= MassElectron
	}

	precursor += correcture
	precursor *= 0.5

	return precursor
}

func precursor1Correcture(quaternary int) float64 {
	if quaternary == 0 {
		return H
	}
	if quaternary == 1 {
		return 0
	}
	return -1
}

func precursor2Correcture(quaternary int) float64 {
	if quaternary == 0 {
		return 2 * H
	}
	if quaternary == 1 {
		return H
	}
	if quaternary == 2 {
		return 0
	}
	return -1
}

func precursorHDX1Correcture(quaternary int, hdx float64) float64 {
	if quaternary == 0 {
		return -1*hdx*H + (hdx+1)*D
	}
	if quaternary == 1 {
		return -1*hdx*H + hdx*D
	}
	return -1
}

func precursorHDX2Correcture(quaternary int, hdx float64) float64 {
	if quaternary == 0 {
		return -1*hdx*H + (hdx+2)*D
	}
	if quaternary == 1 {
		return -1*hdx*H + (hdx+1)*D
	}
	if quaternary == 2 {
		return -1*hdx*H + hdx*D
	}
	return -1
}

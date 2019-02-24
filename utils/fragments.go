package utils

const c2 = 18.01056
const c3 = 17.02655
const c4 = 17.02655 //TODO insert correct float here

func CallculateA(head, poly float64) float64 {
	a := head + poly
	return a
}

func CallculateB(head, poly float64) float64 {
	b := head + poly - c2
	return b
}

func CallculateC(head, poly float64) float64 {
	c := head + poly - c3
	return c
}

func CallculateTa(head, poly float64) float64 {
	ta := head + poly + c3
	return ta
}

func CallculateZ(tail, poly float64) float64 {
	z := tail + poly - c4
	return z
}

func CallculateY(tail, poly float64) float64 {
	y := tail + poly - c4 - c3
	return y
}

func CallculateTz(tail, poly float64) float64 {
	tz := tail + poly - c4 + c3
	return tz
}

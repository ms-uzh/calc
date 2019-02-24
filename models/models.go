package models

type Head struct {
	Mass float64
}

type PolyamineUnit struct {
	Mass            float64
	Quaternaryamine bool
}

type Tail struct {
	Mass            float64
	Quaternaryamine bool
}

type Result struct {
	Mass      float64
	Fragments []Fragment
}

type Fragment struct {
	A  float64
	B  float64
	C  float64
	Ta float64
	Z  float64
	Y  float64
	Tz float64
}

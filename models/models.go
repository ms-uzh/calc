package models

type Molecule struct {
	name string
}

type Head struct {
	Name string
	Mass float64
	HDX  int
}

type PolyamineUnit struct {
	Name            string
	Mass            float64
	Quaternaryamine bool
	HDX             int
}

type Tail struct {
	Name            string
	Mass            float64
	Quaternaryamine bool
	HDX             int
}

type Result struct {
	Mass         float64
	PreCursorOne float64
	PreCursorTwo float64
	Fragments    []Fragment
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

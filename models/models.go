package models

type Head struct {
	Name       string
	Formula    []string
	Mass       float64
	HDX        uint
	Quaternary int
}

type Polyamine struct {
	Name       string
	Formula    []string
	Mass       float64
	HDX        uint
	Sub        Sub
	Quaternary int
}

type Tail struct {
	Name       string
	Formula    []string
	Mass       float64
	HDX        uint
	Sub        Sub
	Quaternary int
}

type Sub struct {
	Name string
	Mass float64
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

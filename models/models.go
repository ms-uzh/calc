package models

type Head struct {
	Name       string
	Formula    []string
	Mass       float64
	HDX        uint
	Quaternary int
}

type Heads []Head

func (heads Heads) GetByName(name string) Head {
	for _, head := range heads {
		if head.Name == name {
			return head
		}
	}
	return Head{}
}

type Polyamines []Polyamine

type Polyamine struct {
	Name       string
	Formula    []string
	Mass       float64
	HDX        uint
	Sub        Sub
	Quaternary int
}

func (polys Polyamines) GetByName(name string) Polyamine {
	for _, poly := range polys {
		if poly.Name == name {
			return poly
		}
	}
	return Polyamine{}
}

type Tails []Tail

type Tail struct {
	Name       string
	Formula    []string
	Mass       float64
	HDX        uint
	Sub        Sub
	Quaternary int
}

func (tails Tails) GetByName(name string) Tail {
	for _, tail := range tails {
		if tail.Name == name {
			return tail
		}
	}
	return Tail{}
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

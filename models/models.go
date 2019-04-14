package models

type Head struct {
	Name       string
	Formula    []string
	Mass       float64
	HDX        uint
	Quaternary int
	IsSelected bool
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

func (heads Heads) SetSelected(name string) Heads {
	for i, head := range heads {
		heads[i].IsSelected = head.Name == name
	}
	return heads
}

type Polyamines []Polyamine

type Polyamine struct {
	Name       string
	Formula    []string
	Mass       float64
	HDX        uint
	Sub        Sub
	Quaternary int
	IsSelected bool
}

func (polys Polyamines) GetByName(name string) Polyamine {
	for _, poly := range polys {
		if poly.Name == name {
			return poly
		}
	}
	return Polyamine{}
}

func (polys Polyamines) SetSelected(name string) Polyamines {
	newPolys := make(Polyamines, len(polys))
	copy(newPolys, polys)
	for i, poly := range newPolys {
		newPolys[i].IsSelected = poly.Name == name
	}
	return newPolys
}

func containsName(name string, names ...string) bool {
	for _, checkName := range names {
		if name == checkName {
			return true
		}
	}
	return false
}

type Tails []Tail

type Tail struct {
	Name       string
	Formula    []string
	Mass       float64
	HDX        uint
	Sub        Sub
	Quaternary int
	IsSelected bool
}

func (tails Tails) GetByName(name string) Tail {
	for _, tail := range tails {
		if tail.Name == name {
			return tail
		}
	}
	return Tail{}
}

func (tails Tails) SetSelected(name string) Tails {
	for i, tail := range tails {
		tails[i].IsSelected = tail.Name == name
	}
	return tails
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

type Fragments []Fragment
type Fragment struct {
	ID int
	A  float64
	B  float64
	C  float64
	Ta float64
	Z  float64
	Y  float64
	Tz float64
}

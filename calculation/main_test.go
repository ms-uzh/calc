package calculation

import (
	"os"
	"testing"

	"github.com/ms-uzh/calc/models"
)

var (
	example1 exampleData
	example2 exampleData
	example3 exampleData
)

type exampleData struct {
	Head  models.Head
	Tail  models.Tail
	Polys models.Polyamines
}

func TestMain(t *testing.M) {
	example1 = example1Data()
	example2 = example2Data()
	example3 = example3Data()
	code := t.Run()
	os.Exit(code)
}

func example1Data() exampleData {
	head := models.Head{
		Name:    "4-OH-IndAc",
		Formula: []string{"C10", "H8", "N", "O2"},
		Mass:    174.05550,
		HDX:     2,
	}
	tail := models.Tail{
		Name:    "",
		Formula: []string{"H2", "N"},
		Mass:    16.01872,
		HDX:     2,
	}
	polyamnies := models.Polyamines{
		models.Polyamine{
			Name:    "3",
			Formula: []string{"C3", "H7", "N"},
			Mass:    57.05785,
			HDX:     1,
		},
		models.Polyamine{
			Name:    "(OH)3",
			Formula: []string{"C3", "H7", "N", "O"},
			Mass:    73.05276,
			HDX:     1,
			Sub: models.Sub{
				Name: "O",
				Mass: 15.99491,
			},
		},
		models.Polyamine{
			Name:    "4",
			Formula: []string{"C4", "H9", "N"},
			Mass:    71.07350,
			HDX:     1,
		},
	}
	return exampleData{head, tail, polyamnies}
}

func example2Data() exampleData {
	head := models.Head{
		Name:    "4-OH-IndAc",
		Formula: []string{"C10", "H8", "N", "O2"},
		Mass:    174.05550,
		HDX:     2,
	}
	tail := models.Tail{
		Name:    "(NMe₃)",
		Formula: []string{"C3", "H9", "N"},
		Mass:    59.07350,
		HDX:     0,
		Sub: models.Sub{
			Name: "C3H7",
			Mass: 43.05478,
		},
		Quaternary: 1,
	}
	polyamnies := models.Polyamines{
		models.Polyamine{
			Name:    "3",
			Formula: []string{"C3", "H7", "N"},
			Mass:    57.05785,
			HDX:     1,
		},
		models.Polyamine{
			Name:    "(OH)3",
			Formula: []string{"C3", "H7", "N", "O"},
			Mass:    73.05276,
			HDX:     1,
			Sub: models.Sub{
				Name: "O",
				Mass: 15.99491,
			},
		},
		models.Polyamine{
			Name:    "5",
			Formula: []string{"C5", "H11", "N"},
			Mass:    85.08915,
			HDX:     1,
		},
	}
	return exampleData{head, tail, polyamnies}
}

func example3Data() exampleData {
	head := models.Head{
		Name:    "IndLac",
		Formula: []string{"C11", "H10", "N1", "O2"},
		Mass:    188.07115,
		HDX:     2,
	}
	tail := models.Tail{
		Name:    "",
		Formula: []string{"H2", "N"},
		Mass:    16.01872,
		HDX:     2,
	}
	polyamnies := models.Polyamines{
		models.Polyamine{
			Name:    "4",
			Formula: []string{"C4", "H9", "N"},
			Mass:    71.07350,
			HDX:     1,
		},
		models.Polyamine{
			Name:    "(Me₂)3",
			Formula: []string{"C5", "H12", "N"},
			Mass:    86.09697,
			HDX:     0,
			Sub: models.Sub{
				Name: "C2H5",
				Mass: 29.03913,
			},
			Quaternary: 1,
		},
		models.Polyamine{
			Name:    "(Me₂)3",
			Formula: []string{"C5", "H12", "N"},
			Mass:    86.09697,
			HDX:     0,
			Sub: models.Sub{
				Name: "C2H5",
				Mass: 29.03913,
			},
			Quaternary: 1,
		},
	}
	return exampleData{head, tail, polyamnies}
}

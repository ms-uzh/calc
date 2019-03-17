package calculation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fforootd/calc/models"
)

func TestExampleName1(t *testing.T) {
	head := models.Head{
		Name:       "4-OH-IndAc",
		Quaternary: 0,
	}

	tail := models.Tail{
		Name:       "",
		Quaternary: 0,
	}

	pa1 := models.Polyamine{
		Name:       "3",
		Quaternary: 0,
	}

	pa2 := models.Polyamine{
		Name:       "(OH)3",
		Quaternary: 0,
	}

	pa3 := models.Polyamine{
		Name:       "4",
		Quaternary: 0,
	}

	assert.Equal(t, "4-OH-IndAc3(OH)34", CalculateName(head, tail, pa1, pa2, pa3))
}

func TestQuaternaryZero(t *testing.T) {
	assert.Equal(t, "", CalculateQuaternaryName(0))
}

func TestQuaternaryOne(t *testing.T) {
	assert.Equal(t, "+", CalculateQuaternaryName(1))
}

func TestQuaternaryMulti(t *testing.T) {
	quaternary := 4
	assert.Equal(t, fmt.Sprint(quaternary, "+"), CalculateQuaternaryName(quaternary))
}

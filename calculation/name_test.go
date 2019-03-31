package calculation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateNameExample1(t *testing.T) {
	name := CalculateName(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, "4-OH-IndAc3(OH)34", name)
}

func TestCalculateNameExample2(t *testing.T) {
	name := CalculateName(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, "4-OH-IndAc3(OH)35(NMe₃)⁺", name)
}

func TestCalculateNameExample3(t *testing.T) {
	name := CalculateName(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, "IndLac4(Me₂)3(Me₂)3²⁺", name)
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

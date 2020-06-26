package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculcateMassExample1(t *testing.T) {
	mass := CalculateMass(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 391.25832999999994, mass)
}

func TestCalculcateMassExample2(t *testing.T) {
	mass := CalculateMass(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 448.32821142010005, mass)
}

func TestCalculcateMassExample3(t *testing.T) {
	mass := CalculateMass(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, 447.35621284019993, mass)
}

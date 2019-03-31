package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePrecursor1Example1(t *testing.T) {
	precursor1 := calculatePrecursor1(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 392.26616, precursor1)
}

func TestCalculatePrecursor1Example2(t *testing.T) {
	precursor1 := calculatePrecursor1(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 448.32876, precursor1)
}

func TestCalculatePrecursor1Example3(t *testing.T) {
	precursor1 := calculatePrecursor1(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, float64(-1), precursor1)
}

func TestCalculatePrecursor2Example1(t *testing.T) {
	precursor1 := calculatePrecursor2(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 196.63700, precursor1)
}

func TestCalculatePrecursor2Example2(t *testing.T) {
	precursor1 := calculatePrecursor2(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 224.66829, precursor1)
}

func TestCalculatePrecursor2Example3(t *testing.T) {
	precursor1 := calculatePrecursor2(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, 223.67866, precursor1)
}

func TestCalculatePrecursorHDX1Example1(t *testing.T) {
	precursor1 := calculatePrecursorHDX1(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 400.31637, precursor1)
}

func TestCalculatePrecursorHDX1Example2(t *testing.T) {
	precursor1 := calculatePrecursorHDX1(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 453.36015, precursor1)
}

func TestCalculatePrecursorHDX1Example3(t *testing.T) {
	precursor1 := calculatePrecursorHDX1(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, float64(-1), precursor1)
}

func TestCalculatePrecursorHDX2Example1(t *testing.T) {
	precursor1 := calculatePrecursorHDX2(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 201.16523, precursor1)
}

func TestCalculatePrecursorHDX2Example2(t *testing.T) {
	precursor1 := calculatePrecursorHDX2(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 227.68712, precursor1)
}

func TestCalculatePrecursorHDX2Example3(t *testing.T) {
	precursor1 := calculatePrecursorHDX2(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, 226.19436, precursor1)
}

package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePrecursor1Example1(t *testing.T) {
	precursor1 := calculatePrecursor1(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 392.26615503222996, precursor1)
}

func TestCalculatePrecursor1Example2(t *testing.T) {
	precursor1 := calculatePrecursor1(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 448.32876000000005, precursor1)
}

func TestCalculatePrecursor1Example3(t *testing.T) {
	precursor1 := calculatePrecursor1(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, float64(-1), precursor1)
}

func TestCalculatePrecursor2Example1(t *testing.T) {
	precursor1 := calculatePrecursor2(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 196.63699003222996, precursor1)
}

func TestCalculatePrecursor2Example2(t *testing.T) {
	precursor1 := calculatePrecursor2(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 224.66829251611503, precursor1)
}

func TestCalculatePrecursor2Example3(t *testing.T) {
	precursor1 := calculatePrecursor2(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, 223.67865499999996, precursor1)
}

func TestCalculatePrecursorHDX1Example1(t *testing.T) {
	precursor1 := calculatePrecursorHDX1(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 400.31637077438995, precursor1)
}

func TestCalculatePrecursorHDX1Example2(t *testing.T) {
	precursor1 := calculatePrecursorHDX1(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 453.36014483885003, precursor1)
}

func TestCalculatePrecursorHDX1Example3(t *testing.T) {
	precursor1 := calculatePrecursorHDX1(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, float64(-1), precursor1)
}

func TestCalculatePrecursorHDX2Example1(t *testing.T) {
	precursor1 := calculatePrecursorHDX2(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 201.16523638719497, precursor1)
}

func TestCalculatePrecursorHDX2Example2(t *testing.T) {
	precursor1 := calculatePrecursorHDX2(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 227.68712341942503, precursor1)
}

func TestCalculatePrecursorHDX2Example3(t *testing.T) {
	precursor1 := calculatePrecursorHDX2(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, 226.19434741942496, precursor1)
}

package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePrecursor1Example1(t *testing.T) {
	precursor1 := CalculatePrecursor1(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 392.26560645232996, precursor1)
}

func TestCalculatePrecursor1Example2(t *testing.T) {
	precursor1 := CalculatePrecursor1(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 448.32766284020005, precursor1)
}

func TestCalculatePrecursor1Example3(t *testing.T) {
	precursor1 := CalculatePrecursor1(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, float64(-1), precursor1)
}

func TestCalculatePrecursor2Example1(t *testing.T) {
	precursor1 := CalculatePrecursor2(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 196.63644145232996, precursor1)
}

func TestCalculatePrecursor2Example2(t *testing.T) {
	precursor1 := CalculatePrecursor2(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 224.66746964626503, precursor1)
}

func TestCalculatePrecursor2Example3(t *testing.T) {
	precursor1 := CalculatePrecursor2(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, 223.67755784019997, precursor1)
}

func TestCalculatePrecursorHDX1Example1(t *testing.T) {
	precursor1 := CalculatePrecursorHDX1(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 400.31582219448995, precursor1)
}

func TestCalculatePrecursorHDX1Example2(t *testing.T) {
	precursor1 := CalculatePrecursorHDX1(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 453.35904767905004, precursor1)
}

func TestCalculatePrecursorHDX1Example3(t *testing.T) {
	precursor1 := CalculatePrecursorHDX1(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, float64(-1), precursor1)
}

func TestCalculatePrecursorHDX2Example1(t *testing.T) {
	precursor1 := CalculatePrecursorHDX2(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 201.16468780729497, precursor1)
}

func TestCalculatePrecursorHDX2Example2(t *testing.T) {
	precursor1 := CalculatePrecursorHDX2(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 227.68630054957504, precursor1)
}

func TestCalculatePrecursorHDX2Example3(t *testing.T) {
	precursor1 := CalculatePrecursorHDX2(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, 226.19325025962496, precursor1)
}

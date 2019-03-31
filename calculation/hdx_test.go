package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHDXExample1(t *testing.T) {
	hdx := calculateHDX(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, uint(7), hdx)
}
func TestCalculateHDXExample2(t *testing.T) {
	hdx := calculateHDX(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, uint(5), hdx)
}
func TestCalculateHDXExample3(t *testing.T) {
	hdx := calculateHDX(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, uint(5), hdx)
}

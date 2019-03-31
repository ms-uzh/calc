package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateQuaternaryExample1(t *testing.T) {
	quat := calculateQuaternary(example1.Head, example1.Tail, example1.Polys...)
	assert.Equal(t, 0, quat)
}

func TestCalculateQuaternaryExample2(t *testing.T) {
	quat := calculateQuaternary(example2.Head, example2.Tail, example2.Polys...)
	assert.Equal(t, 1, quat)
}

func TestCalculateQuaternaryExample3(t *testing.T) {
	quat := calculateQuaternary(example3.Head, example3.Tail, example3.Polys...)
	assert.Equal(t, 2, quat)
}

package templates

import (
	"testing"

	"gotest.tools/assert"
)

func TestRound(t *testing.T) {
	rounded := Round(391.25832999999994)
	assert.Equal(t, "391.25833", rounded)

	rounded = Round(391.25832399999994)
	assert.Equal(t, "391.25832", rounded)
}

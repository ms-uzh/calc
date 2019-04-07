package html

import (
	"testing"

	"gotest.tools/assert"
)

func TestRound(t *testing.T) {
	rounded := round(391.25832999999994)
	assert.Equal(t, "391.25833", rounded)

	rounded = round(391.25832399999994)
	assert.Equal(t, "391.25832", rounded)
}

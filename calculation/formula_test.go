package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateFormula(t *testing.T) {
	formulaTexts := []string{"C10", "H8", "N", "O2"}
	expectedFormulas := []*Formula{
		&Formula{count: 10, name: "C"},
		&Formula{count: 8, name: "H"},
		&Formula{count: 1, name: "N"},
		&Formula{count: 2, name: "O"},
	}
	for i, formulaText := range formulaTexts {
		formula, err := createFormula(formulaText)
		require.NoError(t, err)
		assert.Equal(t, expectedFormulas[i], formula)
	}
}

func TestGenerateFormulaInvalidText(t *testing.T) {
	formula, err := createFormula("C10N")
	require.Nil(t, formula)
	assert.Error(t, err)
}

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

func TestGenerateFormulaExample1(t *testing.T) {
	formulas, err := CalculateFormula(example1.Head, example1.Tail, example1.Polys...)
	require.NoError(t, err)
	assert.Len(t, formulas, 4)
	expected := []string{"C20", "H33", "N5", "O3"}
	assert.Equal(t, expected, formulas)
}
func TestGenerateFormulaExample2(t *testing.T) {
	formulas, err := CalculateFormula(example2.Head, example2.Tail, example2.Polys...)
	require.NoError(t, err)
	assert.Len(t, formulas, 4)

	expected := []string{"C24", "H42", "N5", "O3"}
	assert.Equal(t, expected, formulas)
}

func TestGenerateFormulaExample3(t *testing.T) {
	formulas, err := CalculateFormula(example3.Head, example3.Tail, example3.Polys...)
	require.NoError(t, err)
	assert.Len(t, formulas, 4)
	expected := []string{"C25", "H45", "N5", "O2"}
	assert.Equal(t, expected, formulas)
}

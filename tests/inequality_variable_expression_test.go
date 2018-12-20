package tests

import (
	"testing"

	. "github.com/delaneyj/cassgowary"
	"github.com/stretchr/testify/assert"
)

func TestVariableExpressionLessThanEqualTo(t *testing.T) {
	x := NewVariable("x")
	solver := NewSolver()
	solver.AddConstraint(x.LessThanOrEqualToExpression(NewExpression(100)))
	solver.UpdateVariables()
	assert.True(t, x.Value <= 100)
	solver.AddConstraint(x.EqualsFloat(90))
	solver.UpdateVariables()
	assert.InDelta(t, x.Value.Raw(), 90, FloatEpsilon)
}

func TestVariableExpressionLessThanEqualToUnsatisfiable(t *testing.T) {
	x := NewVariable("x")
	solver := NewSolver()
	solver.AddConstraint(x.LessThanOrEqualToExpression(NewExpression(100)))
	solver.UpdateVariables()
	assert.True(t, x.Value <= 100)
	solver.AddConstraint(x.EqualsFloat(110))
	solver.UpdateVariables()
}

func TestVariableExpressionGreaterThanEqualTo(t *testing.T) {
	x := NewVariable("x")
	solver := NewSolver()
	solver.AddConstraint(x.GreaterThanOrEqualToExpression(NewExpression(100)))
	solver.UpdateVariables()
	assert.True(t, x.Value >= 100)
	solver.AddConstraint(x.EqualsFloat(110))
	solver.UpdateVariables()
	assert.InDelta(t, x.Value.Raw(), 110, FloatEpsilon)
}

func TestVariableExpressionGreaterThanEqualToUnsatisfiable(t *testing.T) {
	x := NewVariable("x")
	solver := NewSolver()
	solver.AddConstraint(x.GreaterThanOrEqualToFloat(100))
	solver.UpdateVariables()
	assert.True(t, x.Value >= 100)
	solver.AddConstraint(x.EqualsFloat(90))
	solver.UpdateVariables()
}

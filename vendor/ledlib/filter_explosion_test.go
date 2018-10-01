package ledlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestMatrixMul(t *testing.T) {
	p1 := mat.NewDense(3, 1, []float64{
		2, 4, 6})
	p2 := mat.NewDense(1, 1, []float64{
		2})
	var c mat.Dense
	c.Mul(p1, p2)

	assert.Equal(t, float64(4), c.At(0, 0))
	assert.Equal(t, float64(8), c.At(1, 0))
	assert.Equal(t, float64(12), c.At(2, 0))

}

package formas

import (
	"testing"
	"math"
	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		assert.Equal(t, areaEsperada, areaRecebida)
	})

	t.Run("Círculo", func(t *testing.T) {
		circulo := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		assert.Equal(t, areaEsperada, areaRecebida)
	})
}
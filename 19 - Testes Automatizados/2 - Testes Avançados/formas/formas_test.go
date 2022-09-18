package formas

import (
	"math"
	"testing"
)


func testeArea(t *testing.T){
	t.Run("Retangulo", func(t *testing.T){
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.area()

		if areaEsperada != areaRecebida{
			t.Errorf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo",func(t *testing.T){
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.area()

		if areaEsperada != areaRecebida{
			t.Errorf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
package formas

import (
	"fmt"
	"math"
)

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f", f.area())
}

func (c Circulo) area() float64{
	return math.Pi * math.Pow(c.raio, 2)
}

func (r Retangulo) area()float64{
	return r.altura * r.largura
}

func main() {

	r := Retangulo{10, 15}
	escreverArea(r)

	c := Circulo{10}
	escreverArea(c)

}
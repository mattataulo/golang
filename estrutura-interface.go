package main

import "math"

type Formato interface {
	Area() float64
	Perimetro() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
	return (r.Largura + r.Altura) * 2
}

type Circulo struct {
	Radius float64
}

func (c Circulo) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangulo struct {
	Base   float64
	Altura float64
	Sides  []float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

func (t Triangulo) Perimetro() float64 {
	return AddFloat(t.Sides)
}

func AddFloat(integers []float64) (sum float64) {
	for _, v := range integers {
		sum += v
	}
	return sum
}

package main

import (
	"reflect"
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		name      string
		Formato   Formato
		area      float64
		Perimetro float64
	}{
		{
			name:      "Retangulo",
			Formato:   Retangulo{Largura: 12, Altura: 6},
			area:      72.0,
			Perimetro: 36,
		},
		{
			name:      "Circulo",
			Formato:   Circulo{Radius: 10},
			area:      314.1592653589793,
			Perimetro: 62.83185307179586,
		},
		{
			name:      "Triangulo",
			Formato:   Triangulo{Base: 12, Altura: 6, Sides: []float64{3.0, 3.0, 3.0}},
			area:      36.0,
			Perimetro: 9,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name+" calculating area", func(t *testing.T) {
			got := tt.Formato.Area()
			if !reflect.DeepEqual(got, tt.area) {
				t.Errorf("%#v got %g want %g", tt.Formato, got, tt.area)
			}
		})

		t.Run(tt.name+" calculating Perimetro", func(t *testing.T) {
			got := tt.Formato.Perimetro()
			if !reflect.DeepEqual(got, tt.Perimetro) {
				t.Errorf("%#v got %g want %g", tt.Formato, got, tt.Perimetro)
			}
		})

	}

}

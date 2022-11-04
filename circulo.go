package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (cir *Circulo) area() float64 {
	pi := math.Pi
	return pi * (cir.Radio * cir.Radio)
}
func (cir *Circulo) perimetro() float64 {
	pi2 := math.Pi
	return 2 * pi2 * cir.Radio

}

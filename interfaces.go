package figuras

import "fmt"

type Figura interface {
	area() float64
	perimetro() float64
}

func Calcular(figura Figura) {
	fmt.Println("Medidas: ", figura)
	fmt.Println("Area: ", figura.area())
	fmt.Println("Perímetro: ", figura.perimetro())
}

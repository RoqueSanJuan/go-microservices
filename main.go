package main

import (
	"fmt"
)

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func main() {
	x := 10
	name := "DevOps"
	isWorking := true

	var diasSemana = map[string]int{"April": 7}

	fmt.Println("Hola Mundo")
	fmt.Println(x, name, isWorking)
	fmt.Printf("Estoy haciendo un ejemplo de %v, %v, %v", x, name, isWorking)
	fmt.Println()
	a, p := rectProps(10, 2)
	fmt.Printf("El area es %f y el perimetro %f ", a, p)
	fmt.Println(diasSemana)
}

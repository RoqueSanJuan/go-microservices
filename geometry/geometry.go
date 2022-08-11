package geometry

import "math"

func Area(length, width float64) float64 {
	return length * width
}

//Diagonal Con mayus para que pueda ser importada
func Diagonal(length, width float64) float64 {
	return math.Sqrt((length * length) + (width * width))
}

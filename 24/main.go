package main

import (
	"fmt"
	"math"
)

type Точка struct {
	x, y float64
}

func НоваяТочка(x float64, y float64) *Точка {
	return &Точка{x: x, y: y}
}

// Расчет расстояния между двумя точками
func (т1 *Точка) расстояниеДо(т2 *Точка) float64 {
	return math.Sqrt(math.Pow(т2.x-т1.x, 2.0) + math.Pow(т2.y-т1.y, 2.0))
}

func main() {
	//Создаём 2 точки
	т1 := НоваяТочка(0, 0)
	т2 := НоваяТочка(1, 1)
	// Выводим расстояние между точками
	fmt.Printf("Расстояние между точками = %f", т1.расстояниеДо(т2))
}

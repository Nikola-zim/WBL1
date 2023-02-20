package main

import (
	"fmt"
	"math"
)

// Есть структура Circle с необходимым методом нахождения площади круга
type Circle struct {
	radius float32
}

func (c *Circle) CircleArea() float32 {
	return c.radius * c.radius * math.Pi
}

// Но существует страрая структура в которой нет нужного метода
type LegacyCircle struct {
	radius float32
}

func (c *LegacyCircle) GetRadius() float32 {
	return c.radius
}

// Создадим структуру с методом Адаптера
type LegacyCircleAdapter struct {
	legacyCircle *LegacyCircle
}

func (a *LegacyCircleAdapter) CircleArea() float32 {
	return a.legacyCircle.GetRadius() * a.legacyCircle.GetRadius() * math.Pi
}

// Интерфейс, определяющий метод CircleArea()
type CircleAreaCalculator interface {
	CircleArea() float32
}

func PrintCircleArea(c CircleAreaCalculator) {
	fmt.Printf("Circle area: %f\n", c.CircleArea())
}

func main() {
	//Простой случай, когда адаптер не нужен
	circle := &Circle{radius: 2.5}
	PrintCircleArea(circle)
	//использование адаптера для старой структуры
	legacyCircle := &LegacyCircle{radius: 2.5}
	legacyCircleAdapter := &LegacyCircleAdapter{legacyCircle: legacyCircle}
	PrintCircleArea(legacyCircleAdapter)
}

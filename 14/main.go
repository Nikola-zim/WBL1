package main

import (
	"fmt"
	"reflect"
)

func testingStrategy(typeDeterminer func(interface{})) {
	typeDeterminer(1)
	typeDeterminer(uint64(2))
	typeDeterminer("Wildberries")
	typeDeterminer(true)
	typeDeterminer(new(chan int))
	typeDeterminer(typeDeterminer)
}

func typeDeterminer1(input interface{}) {
	fmt.Printf("Тип переменной со значением %v: %v\n", input, reflect.TypeOf(input))
}

func typeDeterminer2(input interface{}) {
	fmt.Printf("Тип переменной со значением %v: %T\n", input, input)
}

func typeDeterminer3(input interface{}) {
	fmt.Printf("Тип переменной со значением %v: %v\n", input, reflect.ValueOf(input).Kind())
}

func main() {
	//1. С помощью reflect.TypeOf
	fmt.Println("\nОпределение типа с помощью reflect.TypeOf")
	testingStrategy(typeDeterminer1)
	//2. С помощью Printf
	fmt.Println("\nОпределение типа с помощью Printf")
	testingStrategy(typeDeterminer2)
	//2. С помощью ValueOf(input).Kind()
	fmt.Println("\nОпределение типа с помощью ValueOf(input).Kind()")
	testingStrategy(typeDeterminer3)
}

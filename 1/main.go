package main

import (
	"errors"
	"fmt"
)

// Встраивание методов
// Имеем родительскую структуру Human
type Human struct {
	age  uint
	name string
}

func (h *Human) sayHi() {
	fmt.Printf("Привет, я - human и меня зовут %s \n", h.name)
}

func NewHuman(age uint, name string) *Human {
	return &Human{age: age, name: name}
}

// Встраиваем родительскую структуру Human в структуру Action
type Action struct {
	isHungry bool
	Human
}

func (act *Action) toEat() error {
	if act.isHungry {
		fmt.Println("Human кушает")
		return nil
	} else {
		return errors.New("Human уже не голоден")
	}
}

func NewAction(human Human) *Action {
	return &Action{isHungry: true, Human: human}
}

func main() {
	human := NewHuman(24, "Nick")
	action := NewAction(*human)
	action.sayHi()
	err := action.toEat()
	if err != nil {
		fmt.Println(err)
	}
	action.isHungry = false
	err = action.toEat()
	if err != nil {
		fmt.Println(err)
	}
}

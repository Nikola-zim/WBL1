package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

var justString string

func createHugeString(num int) string {
	char := "А"                      // задаем символ
	str := strings.Repeat(char, num) // создаем строку
	return str
}

func someWrongFunc() {
	v := createHugeString(1 << 10)
	fmt.Println(unsafe.Pointer((*(*reflect.StringHeader)(unsafe.Pointer(&v))).Data))
	justString = v[:100]
	fmt.Println(unsafe.Pointer((*(*reflect.StringHeader)(unsafe.Pointer(&justString))).Data))
}

/*
Проблема функции someWrongFunc заключается в том, что происходит утечка памяти,
и выделенная память под переменную v не очиститься до того как justString выйдет из области видимости.
Решением этого является использование метода Clone
*/
func someRightFunc() {
	v := createHugeString(1 << 10)
	fmt.Println(unsafe.Pointer((*(*reflect.StringHeader)(unsafe.Pointer(&v))).Data))
	justString = strings.Clone(v[:100])
	fmt.Println(unsafe.Pointer((*(*reflect.StringHeader)(unsafe.Pointer(&justString))).Data))
}

func main() {
	someWrongFunc()
	someRightFunc()
}

package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	//Создадим множество, т.е. map, где ключ - значение из массива, а значение - пустой struct{}{}
	var res = make(map[string]struct{})
	for _, val := range arr {
		res[val] = struct{}{}
	}
	for key := range res {
		fmt.Println("Key:", key)
	}
}

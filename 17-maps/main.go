package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 100)
	for i := range slice1 {
		slice1[i] = rand.Intn(50)
	}

	map1 := make(map[int]int, 80)

	for _, v1 := range slice1 {
		v, ok := map1[v1]
		if !ok {
			map1[v1] = 1
		} else {
			map1[v1] = v + 1
		}
	}

	fmt.Println("Slice")
	fmt.Println(slice1)
	fmt.Println(map1)

}

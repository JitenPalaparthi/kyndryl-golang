package main

import "fmt"

func main() {
	greet()
	greetM("Hello WOrld!")
	arr := [5]int{10, 12, 13, 14, 53}
	s0 := sumOf(arr)
	fmt.Println("Sum:", s0)
	s1, m1 := sumAndMultipyOf(arr)
	fmt.Println("Sum:", s1, "Muliply:", m1)
	s2, _ := sumAndMultipyOf(arr)
	fmt.Println("Sum:", s2)
	_, m3 := sumAndMultipyOf(arr)
	fmt.Println("Muliply:", m3)

	IncrBy(arr, 10)
	fmt.Println(arr)
	arr1 := [4]int{10, 11, 12, 14}

	// slice, map, channel, and any pointer type is used to pass to a function is a referece

	slice1 := []int{10, 12, 13, 14, 53}
	slice2 := []int{10, 12, 13, 14, 53, 43, 34, 34, 343, 34, 343, 234, 4534, 5345346, 54746, 7, 75667}
	s4, m4 := sumAndMultiplySlice(slice1)
	fmt.Println("Sum:", s4, "Muliply:", m4)
	s5, m5 := sumAndMultiplySlice(arr1[:])
	fmt.Println("Sum:", s5, "Muliply:", m5)
	s6, m6 := sumAndMultiplySlice(arr[:])
	fmt.Println("Sum:", s6, "Muliply:", m6)
	s7, m7 := sumAndMultiplySlice(slice2)
	fmt.Println("Sum:", s7, "Muliply:", m7)

	IncrSliceBy(slice1, 10)
	fmt.Println(slice1)

}

func greet() {
	fmt.Println("hello Kyndryl")
}

func greetM(msg string) {
	fmt.Println(msg)
}

func sumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func sumAndMultipyOf(arr [5]int) (sum int, mul int) {
	mul = 1
	for _, v := range arr {
		sum += v
		if v != 0 {
			mul *= v
		}
	}
	return sum, mul
}

func IncrBy(arr [5]int, inc int) { // generally arrays and normal variables are pass by value
	for i, v := range arr {
		arr[i] = v + inc
	}
	fmt.Println("Inside func")
	fmt.Println(arr)
	fmt.Println("-----------")
}

func sumAndMultiplySlice(arr []int) (sum int, mul int) {
	mul = 1
	for _, v := range arr {
		sum += v
		if v != 0 {
			mul *= v
		}
	}
	return sum, mul
}

func IncrSliceBy(arr []int, inc int) { // generally arrays and normal variables are pass by value
	for i, v := range arr {
		arr[i] = v + inc
	}
	fmt.Println("Inside func")
	fmt.Println(arr)
	fmt.Println("-----------")
}

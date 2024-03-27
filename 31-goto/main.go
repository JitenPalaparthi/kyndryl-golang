package main

func main() {

	arr := [5]int{10, 12, 14, 16, 18}
	index := 0
loop:
	if index <= 4 {
		println(index, "-->", arr[index], " ")
		index++
		goto loop
	}
}

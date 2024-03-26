package main

import (
	"fmt"
)

func main() {

	var num1 Integer = 12

	str1 := num1.ToString()
	sq1 := num1.Square()

	fmt.Printf("str1 value:%s type:%T\n", str1, str1)
	fmt.Println("Sq1:", sq1)

	var num2 int = 12

	sq2 := Integer(num2).Square()
	fmt.Println("Sq2:", sq2)

	var flaot1 float32 = 12.34

	sq3 := Integer(flaot1).Square()
	fmt.Println("Sq3:", sq3)

	var num3 Integer = 10
	var num4 Integer3 = 10
	var num5 int = 10

	cube1 := num4.Cube()
	fmt.Println("Cube1:", cube1)
	sq4 := Integer(num4).Square()
	fmt.Println("Sq4:", sq4)
	cube2 := Integer3(num3).Cube()
	fmt.Println("Cube2:", cube2)

	sq5 := Integer(num5).Square()
	cube3 := Integer3(num5).Cube()
	fmt.Println("Sq5:", sq5)
	fmt.Println("Cube3:", cube3)

}

type Integer int // new type. extending the existing type int
type Integer3 Integer
type Integer2 = int // alias for the same type

func (i Integer) ToString() string {
	return fmt.Sprintf("%d", i)
	//return strconv.Itoa(int(i))
}

func (i Integer) Square() Integer {
	return i * i
}

func (i Integer3) Cube() Integer3 {
	return i * i * i
}

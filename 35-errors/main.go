package main

import "fmt"

func main() {

	// str1 := "5000"
	// num1, err := strconv.Atoi(str1)

	// num2 := 5000
	// str2 := strconv.Itoa(num2)

	// fmt.Println(str1, str2, num1, num2, err)

	slice1 := make([]*int, 5)
	i1 := new(int)
	*i1 = 10
	slice1[0] = i1
	*i1 = 11
	slice1[1] = i1
	*i1 = 12
	slice1[2] = i1
	*i1 = 13
	slice1[3] = i1
	*i1 = 14
	slice1[4] = i1

	sum, err := SumOf(slice1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
	sum, err = SumOf(nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}

	var num2 []*int
	sum, err = SumOf(num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
}

func SumOf(nums []*int) (int, error) {
	if nums == nil {
		return 0, fmt.Errorf("nil input is given")
	}
	sum := 0
	for _, v := range nums {
		sum += *v
	}
	return sum, nil
}

func SumOfS(nums []*int) (int, string) {
	if nums == nil {
		return 0, fmt.Sprint("nil input is given")
	}
	sum := 0
	for _, v := range nums {
		sum += *v
	}
	return sum, ""
}

// slice, map , chan, pointer, interface can be nil

// What is an error
// What is exception
// What is excepting handling

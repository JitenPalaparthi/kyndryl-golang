package main

func main() {

	day := 5

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	num := 50

	switch { // empty switch
	case num >= 0 && num <= 50:
		println(num, " is between 0-50")
	case num > 50 && num <= 150:
		println(num, " is between 50 and 150")
	case num > 150 && num <= 250:
		println(num, " is between 150 and 250")
	default:
		println(num)
	}

	switch char := 'B'; char {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is vovel")
	default:
		println(string(char), "is either consonent or special char")
	}

	num = 48
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
	case num%4 == 0:
		println(num, "is divisible by 4")
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

	println("fall through logic")
	num = 44
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	default:
		println(num, "is not divisible by 2,4 or 8")
	}

}

package main

func main() {

	var age uint8 = 18

	if age >= 18 {
		println("eligible for vote")
	} else {
		println("not eligible for vote")
	}

	gender := 'M'
	if age >= 18 && (gender == 'F' || gender == 'f') {
		println("she is eligible for marraige")
	} else if age >= 21 && (gender == 'M' || gender == 'm') {
		println("he is eligible for marraige")
	} else {
		println("not eligible for marriage")
	}

	//num>100?println("number is greater han 100"):println("num is less than or equal to 100")

	if char := '池'; char <= 255 {
		println(string(char), char, " is ASCII char")
	} else {
		println(string(char), char, " is not an ASCII char")
	}
}

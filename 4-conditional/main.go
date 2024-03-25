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

}

package main

import (
	"fmt"
	"regexp"
)

func main() {

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	pattern1 := `^[a-zA-Z]`

	// str := "Hello World"

	Validemail := "Jitenp@outlook.com"
	Invalidemail1 := "Jitenp@outlook"
	//	InvalidEmail2 := "jiten123.com"

	ok, err := regexp.MatchString(pattern, Validemail)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok {
		fmt.Println("valid email")
	} else {
		fmt.Println("invalid email")
	}

	ok, err = regexp.Match(pattern, []byte(Invalidemail1))
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok {
		fmt.Println("valid email")
	} else {
		fmt.Println("invalid email")
	}

	rex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return
	}

	ok1 := rex.MatchString(Validemail)

	if ok1 {
		fmt.Println("valid email")
	} else {
		fmt.Println("Invalid email")
	}

	rex1, err := regexp.Compile(pattern1)
	if err != nil {
		fmt.Println(err)
		return
	}

	ok1 = rex1.MatchString("hello world")
	if ok1 {
		fmt.Println("valid string")
	} else {
		fmt.Println("Invalid string")
	}
	//regexp.MatchReader()

}

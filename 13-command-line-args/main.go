package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	str := ""
	if len(args) >= 2 {
		switch args[1] {
		case "en", "EN", "english", "ENGLISH":
			str = "Hello World!"
		case "ch", "CH", "chinees", "CHINEES":
			str = "世界您好!"
		case "hindi", "Hindi", "HINDI":
			str = "हैलो वर्ल्ड"
		default:
			fmt.Println("unsupported langauge")
			return
		}
		//fmt.Println(str)
	} else {
		fmt.Println("not enough arguments given")
		return
	}
	if len(args) >= 3 {
		rstr := ""
		switch args[2] {

		case "west":
			for _, v := range str {
				rstr = string(v) + rstr
			}
			str = rstr
		}
	}
	fmt.Println("\u2713" + " " + str)
}

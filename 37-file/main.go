package main

import (
	"demo/person"
	"log"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	if len(args) == 4 {
		name := args[1]
		age, err := strconv.Atoi(args[2])
		if err != nil {
			log.Println("Error:", err)
			return
		}
		email := args[3]

		p := person.Person{Name: name, Age: uint8(age), Email: email}
		ok, err := p.Add("employees.txt")
		if err != nil {
			log.Println("Error:", err)
			return
		}
		if ok {
			log.Println("Successfully added")
		}

	} else {
		log.Println("Invalid commond line parameters")
		return
	}
}

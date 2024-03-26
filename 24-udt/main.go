package main

import "fmt"

func main() {

	p1 := Person{Name: "JP", Email: "JitenP@Outlook.Com", Mobile: "9618558500", Address: struct { // inline structure
		City, Country, Pincode string
	}{City: "Hyd", Country: "India", Pincode: "544000"}}
	fmt.Println(p1)

	p2 := new(Person)
	p2.Address.City = "bangalore"
	p2.Name = "Jiten"
	p2.Address.Pincode = "560086"
	fmt.Println(p2)
}

type Person struct {
	Name, Email, Mobile string
	Address             struct { // inline structure
		City, Country, Pincode string
	}
}

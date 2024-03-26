package main

import "fmt"

func main() {

	p1 := Person{Name: "JP", Email: "JitenP@Outlook.Com", Mobile: "9618558500", Address: Address{City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560086"}}
	fmt.Println(p1)
	p1.Address.City = "Guntur"
	p1.Address.State = "Andhra Pradesh"
	p1.Address.PinCode = "522000"
	fmt.Println(p1)

	s1 := Student{Name: "JP", Grade: "6th", Address: Address{City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560086"}}

	s1.City = "Guntur"
	s1.State = "AP"
	s1.PinCode = "522000"
	s1.Address.State = "Active"
	s1.State = "Active"

	fmt.Println(s1)
}

type Person struct {
	Name, Email, Mobile string
	Address             Address
}

type Student struct {
	Name, Grade, Status string
	Address             // Promoted field
}

type Address struct {
	City, State, Country, PinCode, Status string
}

// class Student extends Address{
// 	string Name
// 	string Grade
// }

// Student s1 = new Student()
// s1.Name
// s1.City
// s1.State

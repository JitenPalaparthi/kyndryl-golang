package main

import "fmt"

func main() {
	var p1 Person
	p1.Name = "Jiten"
	p1.Email = "JitenP@Outlook.Com"
	p1.Age = 38
	fmt.Println(p1) // this is formatted and developers/ debug kind of a print
	println("Name:", p1.Name)
	println("Email:", p1.Email)
	println("Age:", p1.Age)

	p2 := Person{} // shorthand declaration of type
	println("Name:", p2.Name)
	println("Email:", p2.Email)
	println("Age:", p2.Age)

	p3 := Person{Name: "Jiten", Age: 38, Email: "Jiten.Palaparthi@Gmail.Com"}
	println("Name:", p3.Name)
	println("Email:", p3.Email)
	println("Age:", p3.Age)
	println("Mobile:", p3.Mobile)

	p4 := new(Person)

	p4.Name = "Jiten"
	p4.Email = "JitenP@Outlook.Com"
	p4.Age = 38
	//fmt.Println(p4) // this is formatted and developers/ debug kind of a print
	println("Name:", (*p4).Name)
	println("Email:", p4.Email)
	println("Age:", p4.Age)

	p5 := &Person{Name: "Jiten", Age: 38, Email: "Jiten.Palaparthi@Gmail.Com"}
	var mobile string = "9618558500"
	//*p5.Mobile = "9618558500" // pointer
	p5.Mobile = &mobile
	println("Name:", p5.Name)
	println("Email:", p5.Email)
	println("Age:", p5.Age)
	println("Mobile:", *p5.Mobile)

	p6 := &Person{} // shorthand declaration of type
	println("Name:", p6.Name)
	println("Email:", p6.Email)
	println("Age:", p6.Age)

}

type Person struct {
	Name   string
	Email  string
	Age    uint8
	Mobile *string // the field is a pointer
}

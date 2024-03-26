package main

import "fmt"

func main() {
	p1 := struct { // Anonymous struct
		Name    string
		Email   string
		Address struct {
			City string
		}
	}{
		Name:  "jiten",
		Email: "jitenP@outlook.com",
		Address: struct {
			City string
		}{
			City: "Hyderabad"},
	}
	fmt.Println(p1)
}

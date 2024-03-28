package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("abcd.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 400)
	defer f.Close() // 1. You dont forget to close. 2. defer code is reachable at any cost even during panic
	if err != nil {
		fmt.Println(err)
		return
	}

	//defer f.Close()
	_, err = f.WriteString("Hello World\n")
	if err != nil {
		fmt.Println("2-", err)
		return
	}
	//f.Close()
	// Here derfe f.Close is called automatically
}

// Issues here are 1. may forget to close the file 2. close call might have not be reached
// To mitigate these two issues can use defer

// File IO operations
// Network IO Operations
// Open Sockets, Close Sockets, Open File Descripters, Close File descripters
// You have to close sockets , file descripters , db connections

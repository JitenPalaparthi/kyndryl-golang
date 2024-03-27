package main

import (
	"log"
	"os"
)

func main() {
	//f, err := os.Open("abcd.txt")
	f, err := os.OpenFile("abcd.txt", os.O_WRONLY|os.O_APPEND, os.ModePerm)
	log.Println("open a file named abcd.txt")
	if err != nil {
		//fmt.Println(err)
		log.Println("Error:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("hello World\n")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully written to the file")
}

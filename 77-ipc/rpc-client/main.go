package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8089")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()
	args := new(Args)
	args.A = 101
	args.B = 105
	var result int

	err = client.Call("Calc.Multiply", args, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Multiply :", *&result)
}

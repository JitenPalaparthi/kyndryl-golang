package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Calc int

func (c *Calc) Multiply(a *Args, result *int) error {
	if a == nil {
		return errors.New("invalid arguments")
	}
	if result == nil {
		return errors.New("invalid arguments.Nil result")
	}
	*result = a.A * a.B
	return nil
}

func (c *Calc) Add(a *Args, result *int) error {
	if a == nil {
		return errors.New("invalid arguments")
	}
	if result == nil {
		return errors.New("invalid arguments.Nil result")
	}
	*result = a.A + a.B
	return nil
}

var (
	PORT string
)

func init() {
	flag.StringVar(&PORT, "port", "8089", "--port=8089 or -port 8089")

	// PORT *string
	//PORT = flag.String("port", "8089", "--port=8089 or -port 8089")
}

func main() {
	flag.Parse() // imp
	c := new(Calc)
	err := rpc.Register(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on port:" + PORT)
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()
	rpc.Accept(lis)
}

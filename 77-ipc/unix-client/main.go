package main

import (
	"flag"
	"fmt"
	"net"
	"runtime"
)

var (
	msg string
)

func init() {
	flag.StringVar(&msg, "message", "Hello World", `--message="Hello World"`)
}

func main() {
	flag.Parse()
	switch runtime.GOOS {
	case "linux", "darwin":
		conn, err := net.Dial("unix", "/tmp/socket")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		//msg := "Hello Socket, How are you doing"
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Read:", string(buf[:n]))
	case "windows":

		conn, err := net.Dial("unix", `\\.pipe\somepipe`)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		//msg := "Hello Socket, How are you doing"
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Read:", string(buf[:n]))

	}
}

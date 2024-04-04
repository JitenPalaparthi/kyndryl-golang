package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	os.Remove("/tmp/socket")
	lis, err := net.Listen("unix", "/tmp/socket")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("unix socket is up and running, listening on /tmp/socket")
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		go func(c net.Conn) {
			buf := make([]byte, 1024)
			for {
				n, err := c.Read(buf)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Read data:", string(buf[:n]))
				_, err = c.Write(buf[:n])
				if err != nil {
					fmt.Println(err)
					return
				}
			}

		}(conn)
	}
}

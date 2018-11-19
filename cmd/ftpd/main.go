package main

import (
	"fmt"
	"net"
	"os"
)

const Address  = "localhost:8000"

func main() {
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error starting listener: %v \n", err)
		os.Exit(1)
	}
	fmt.Println("server listening " +  Address)
	for {
		_, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error accepting connection: %v \n", err)
		}
	}
}

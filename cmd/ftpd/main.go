package main

import (
	"fmt"
	"github.com/4hel/ftpd/logger"
	"github.com/4hel/ftpd/server"
	"net"
)

const Address  = "localhost:80"

func main() {
	// start listening socket
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		//logger.Error.Fatal(fmt.Sprintf("error starting listener: %v \n", err))
		logger.Error.Fatal("error starting listener: ", err)
	}
	fmt.Println("server listening " +  Address)

	// acceppt connections and start server go routine
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error.Println("error accepting connection: ", err)
		} else {
			go server.ReadLoop(conn)
		}
	}
}

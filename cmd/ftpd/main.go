package main

import (
	"github.com/4hel/ftpd/logger"
	"github.com/4hel/ftpd/server"
	"net"
)

const Address = "localhost:8000"

func main() {
	// start listening socket
	listener, err := net.Listen("tcp", Address)
	if err != nil {
		logger.Error.Fatal("error starting listener: ", err)
	}
	logger.Info.Println("server listening:", Address)

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

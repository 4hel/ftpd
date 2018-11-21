package server

import (
	"bufio"
	"context"
	"fmt"
	"github.com/4hel/ftpd/command"
	"github.com/4hel/ftpd/logger"
	"net"
	"strings"
)

func ReadLoop(conn net.Conn) {

	defer conn.Close()
	ctx := context.WithValue(context.Background(), command.ContextKeyConnection, conn)

	fmt.Fprintln(conn, "220 FTP Server ready.")
	scanner := bufio.NewScanner(conn)
	var parsedCommand command.Command
	var parsingError error

	// parse command from the client
	for scanner.Scan() {
		msg := scanner.Text()
		cmd := strings.Fields(msg)[0]
		switch cmd {
		case "CLOSE":
			parsedCommand = command.NewCommandClose(ctx)
			parsedCommand.Execute()
		case "USER":
			parsedCommand, parsingError = command.NewCommandUser(ctx, msg)
			if parsingError == nil {
				parsedCommand.Execute()
			} else {
				logger.Error.Println(parsingError)
				break
			}
		case "SYST":
			parsedCommand = command.NewCommandSyst(ctx)
			parsedCommand.Execute()
		case "PASV":
			c := make(chan net.Conn)
			port, err := openDataConnection(c)
			parsedCommand = command.NewCommandPasv(ctx, port)
			parsedCommand.Execute()
			if err == nil {
				conn := <- c
				ctx = context.WithValue(ctx, command.ContextKeyDataConnection, conn)
			}
		default:
			fmt.Println(msg)
			continue
		}


	}

	if scanner.Err() != nil {
		if parsedCommand.Type() != command.Close {
			logger.Error.Println("error reading from connection: ", scanner.Err())
		}
	}
}

func openDataConnection(c chan net.Conn) (int, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		//fmt.Fprintln(c.Connection(), "425 Can't open data connection.")
		return 0, err
	}
	go acceptData(listener, c)
	return listener.Addr().(*net.TCPAddr).Port, nil
}

func acceptData(listener net.Listener, c chan net.Conn) {
	conn, err := listener.Accept()
	if err != nil {
		logger.Error.Println("error accepting connection: ", err)
	} else {
		logger.Info.Println("data connection accepted")
		c <- conn
	}
}

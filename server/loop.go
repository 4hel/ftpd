package server

import (
	"bufio"
	"fmt"
	"github.com/4hel/ftpd/command"
	"github.com/4hel/ftpd/logger"
	"net"
	"strings"
)

func ReadLoop(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintln(conn, "220 FTP Server ready.")
	scanner := bufio.NewScanner(conn)
	//
	// parse command from the client
	//
	for scanner.Scan() {
		msg := scanner.Text()
		cmd := strings.Fields(msg)[0]
		var parsedCommand command.Command
		switch cmd {
		case "CLOSE":
			parsedCommand = parseClose(msg)
		case "USER":
			parsedCommand = parseUser(msg)
		}
		fmt.Println(parsedCommand)
	}

	if scanner.Err() != nil {
		logger.Error.Println("error reading from connection: ", scanner.Err())
	}
}

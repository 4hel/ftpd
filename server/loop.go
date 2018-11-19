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

	// parse command from the client
	for scanner.Scan() {
		msg := scanner.Text()
		cmd := strings.Fields(msg)[0]
		switch cmd {
		case "CLOSE":
			parsedCommand = command.NewCommandClose(ctx)
		case "USER":
			//parsedCommand = parseUser(msg)
		}
		parsedCommand.Execute()
	}

	if scanner.Err() != nil {
		if parsedCommand.Type() != command.Close {
			logger.Error.Println("error reading from connection: ", scanner.Err())
		}
	}
}

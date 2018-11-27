package server

import (
	"bufio"
	"context"
	"fmt"
	"github.com/4hel/ftpd/command"
	"github.com/4hel/ftpd/logger"
	"net"
	"os"
	"strings"
)

func ReadLoop(conn net.Conn) {

	defer conn.Close()
	dir, err := os.Getwd()
	if err != nil {
		logger.Error.Print(err)
		return
	}

	ctx := context.WithValue(context.Background(), command.ContextKeyConnection, conn)
	ctx = context.WithValue(ctx, command.ContextKeyDir, dir)

	fmt.Fprintln(conn, "220 FTP Server ready.")
	scanner := bufio.NewScanner(conn)
	var parsedCommand command.Command

	// parse command from the client
	for scanner.Scan() {
		msg := scanner.Text()
		cmd := strings.Fields(msg)[0]
		switch cmd {
		case "QUIT":
			parsedCommand = command.NewCommandClose(ctx)
		case "CLOSE":
			parsedCommand = command.NewCommandClose(ctx)
		case "USER":
			parsedCommand = command.NewCommandUser(ctx, msg)
		case "SYST":
			parsedCommand = command.NewCommandSyst(ctx)
		case "PASV":
			parsedCommand = command.NewCommandPasv(ctx)
		case "LIST":
			parsedCommand = command.NewCommandList(ctx)
		case "PWD":
			parsedCommand = command.NewCommandPwd(ctx)

		default:
			fmt.Println(msg)
			continue
		}
		ctx = parsedCommand.Execute()

	}

	if scanner.Err() != nil {
		if parsedCommand.Type() != command.Close {
			logger.Error.Println("error reading from connection: ", scanner.Err())
		}
	}
}

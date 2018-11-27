package command

import (
	"context"
	"fmt"
	"github.com/4hel/ftpd/logger"
	"net"
)

type CommandPasv struct {
	CommandBase
}

func NewCommandPasv(ctx context.Context) Command {
	retval := CommandPasv{
		CommandBase{Kind: Pasv, Ctx: ctx},
	}
	return retval
}

func (c CommandPasv) Execute() context.Context {

	ch := make(chan net.Conn)
	port, err := openDataConnection(ch)

	conn := c.Connection()
	lowByte := port & 0x00FF
	highByte := port >> 8
	msg := fmt.Sprintf("227 Entering Passive Mode (%d,%d,%d,%d,%d,%d)\n", 127, 0, 0, 1, highByte, lowByte)
	fmt.Fprint(conn, msg)
	logger.Info.Print(msg, "decimal port", port)

	if err == nil {
		conn := <-ch
		c.Ctx = context.WithValue(c.Ctx, ContextKeyDataConnection, conn)
	}

	return c.Ctx
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

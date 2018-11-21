package command

import (
	"context"
	"fmt"
	"github.com/4hel/ftpd/logger"
)

type CommandPasv struct {
	CommandBase
	port int
}

func NewCommandPasv(ctx context.Context, port int) (Command) {
	retval := CommandPasv{
		CommandBase{Kind: Pasv, Ctx: ctx},
		port,
	}
	return retval
}

func (c CommandPasv) Execute() {
	conn := c.Connection()
	port := c.port
	lowByte := port & 0x00FF
	highByte := port >> 8
	msg := fmt.Sprintf( "227 Entering Passive Mode (%d,%d,%d,%d,%d,%d)\n", 127,0,0,1,highByte, lowByte)
	fmt.Fprint(conn, msg)
	logger.Info.Print(msg, "decimal port", port)
}

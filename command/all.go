package command

import (
	"context"
	"github.com/4hel/ftpd/logger"
	"net"
)

type CommandType int

const (
	ContextKeyConnection = "connection"

	Close CommandType = 1 << iota
	User
)

type Command interface {
	Execute()
	Type() CommandType
}

type CommandBase struct {
	Kind CommandType
	Ctx  context.Context
}

func (c CommandBase) Connection() net.Conn {
	conn, ok := c.Ctx.Value(ContextKeyConnection).(net.Conn)
	if !ok {
		logger.Error.Fatalln("connection not found in context")
	}
	return conn
}

func (c CommandBase) Type() CommandType {
	return c.Kind
}

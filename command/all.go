package command

import (
	"context"
	"github.com/4hel/ftpd/logger"
	"net"
)

type CommandType int

const (
	ContextKeyConnection     = "connection"
	ContextKeyDataConnection = "data"
	ContextKeyDir            = "dir"

	Close CommandType = 1 << iota
	User
	Syst
	Pasv
	List
)

type Command interface {
	Execute() context.Context
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

func (c CommandBase) DataConnection() net.Conn {
	conn, ok := c.Ctx.Value(ContextKeyDataConnection).(net.Conn)
	if !ok {
		logger.Error.Fatalln("data connection not found in context")
	}
	return conn
}

func (c CommandBase) Type() CommandType {
	return c.Kind
}

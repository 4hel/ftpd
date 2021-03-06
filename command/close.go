package command

import "context"

type CommandClose struct {
	CommandBase
}

func NewCommandClose(ctx context.Context) Command {
	retval := CommandClose{
		CommandBase{Kind: Close, Ctx: ctx},
	}
	return retval
}

func (c CommandClose) Execute() context.Context {
	conn := c.Connection()
	conn.Close()
	return c.Ctx
}

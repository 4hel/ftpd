package command

import (
	"context"
	"fmt"
)

type CommandPwd struct {
	CommandBase
}

func NewCommandPwd(ctx context.Context) Command {
	retval := CommandPwd{
		CommandBase{Kind: Pwd, Ctx: ctx},
	}
	return retval
}

func (c CommandPwd) Execute() context.Context {
	conn := c.Connection()

	dir := c.Ctx.Value(ContextKeyDir).(string)

	fmt.Fprint(conn, "257 \"", dir, "\"\n")

	return c.Ctx
}

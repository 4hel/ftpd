package command

import (
	"context"
	"fmt"
)

type CommandSyst struct {
	CommandBase
}

func NewCommandSyst(ctx context.Context) (Command) {
		retval := CommandSyst{
			CommandBase{Kind: Syst, Ctx: ctx},
		}
		return retval
}

func (c CommandSyst) Execute() context.Context {
	conn := c.Connection()
	fmt.Fprintln(conn, "215 LINUX")

	return c.Ctx
}


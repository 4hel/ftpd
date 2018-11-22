package command

import (
	"context"
	"fmt"
)


type CommandList struct {
	CommandBase
}

func NewCommandList(ctx context.Context) (Command) {
	retval := CommandList{
		CommandBase{Kind: List, Ctx: ctx},
	}
	return retval
}

func (c CommandList) Execute() {
	conn := c.DataConnection()
	fmt.Fprintln(conn, "https://cr.yp.to/ftp/list/binls.html")
}

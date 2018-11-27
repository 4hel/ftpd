package command

import (
	"context"
	"fmt"
	"strings"
)

type CommandUser struct {
	CommandBase
	User string
}

func NewCommandUser(ctx context.Context, msg string) Command {
	parts := strings.Fields(msg)
	user := ""
	if len(parts) > 1 {
		user = parts[1]
	}
		retval := CommandUser{
			CommandBase: CommandBase{Kind: User, Ctx: ctx},
			User:        user,
		}
		return retval

}

func (c CommandUser) Execute() context.Context {
	conn := c.Connection()
	fmt.Fprintf(conn, "230 User is %s.\n", c.User)

	return c.Ctx
}

package command

import (
	"context"
	"fmt"
	"github.com/go-errors/errors"
	"strings"
)

type CommandUser struct {
	CommandBase
	User string
}

func NewCommandUser(ctx context.Context, msg string) (Command, error) {
	parts := strings.Fields(msg)
	if len(parts) != 2 {
		return nil, errors.New("user command must have two words")
	} else {
		retval := CommandUser{
			CommandBase: CommandBase{Kind: User, Ctx: ctx},
			User: parts[1],
		}
		return retval, nil
	}
}

func (c CommandUser) Execute() {
	conn := c.Connection()
	fmt.Fprintf(conn, "230 User is %s.\n", c.User)
}

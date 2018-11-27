package command

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CommandCwd struct {
	CommandBase
	Msg string
}

func NewCommandCwd(ctx context.Context, msg string) Command {
	retval := CommandCwd{
		CommandBase{Kind: Cwd, Ctx: ctx},
		msg,
	}
	return retval
}

func (c CommandCwd) Execute() context.Context {
	conn := c.Connection()

	oldDirAbs := c.Ctx.Value(ContextKeyDir).(string)

	fields := strings.Fields(c.Msg)
	if len(fields) < 2 {
		fmt.Fprint(conn, "500 No directory given\n")
		return c.Ctx
	}

	newDir := fields[1]
	newDirAbs := "/"
	if newDir[0] == "/"[0] {
		newDirAbs = newDir
	} else {
		newDirAbs = oldDirAbs + "/" + newDir
	}

	_, err := os.Open(newDirAbs)
	path, err := filepath.Abs(newDirAbs)
	if err != nil {
		fmt.Fprintln(conn, "550 "+newDir+": No such file or directory.")
	} else {
		c.Ctx = context.WithValue(c.Ctx, ContextKeyDir, path)
		fmt.Fprintln(conn, "250 Okay.")
	}

	return c.Ctx
}

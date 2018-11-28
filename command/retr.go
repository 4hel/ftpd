package command

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

type CommandRetr struct {
	CommandBase
	File string
}

func NewCommandRetr(ctx context.Context, msg string) Command {
	parts := strings.Fields(msg)
	file := ""
	if len(parts) > 1 {
		file = parts[1]
	}
	retval := CommandRetr{
		CommandBase: CommandBase{Kind: User, Ctx: ctx},
		File:        file,
	}
	return retval

}

func (c CommandRetr) Execute() context.Context {
	conn := c.Connection()
	dataConn := c.DataConnection()
	defer dataConn.Close()

	dirAbs := c.Ctx.Value(ContextKeyDir).(string)

	f, err := os.Open(dirAbs + "/" + c.File)
	defer f.Close()
	if err != nil {
		fmt.Fprintln(conn, "451 Could not open file")
		return c.Ctx
	}

	fileInfo, err := os.Stat(dirAbs + "/" + c.File)
	if !fileInfo.Mode().IsRegular() {
		fmt.Fprintln(conn, "451 Not a regular file")
		return c.Ctx
	}

	fmt.Fprintln(conn, "150 Transfer starting")

	//-, err := io.Copy(dataConn, f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Fprint(dataConn, scanner.Text()+"\r\n")
	}

	if scanner.Err() != nil {
		fmt.Fprintln(conn, "451 Error transfering file")
		return c.Ctx
	}

	fmt.Fprintln(conn, "226 Transfer OK")

	return c.Ctx
}

package command

import (
	"context"
	"fmt"
	"github.com/4hel/ftpd/logger"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"time"
)

type CommandList struct {
	CommandBase
}

func NewCommandList(ctx context.Context) Command {
	retval := CommandList{
		CommandBase{Kind: List, Ctx: ctx},
	}
	return retval
}

func (c CommandList) Execute() context.Context {
	conn := c.Connection()
	dataConn := c.DataConnection()
	defer dataConn.Close()

	fmt.Fprintln(conn, "150 Opening ASCII mode data connection for file list")

	f, err := os.Open(c.Ctx.Value(ContextKeyDir).(string))
	defer f.Close()

	if err != nil {
		logger.Error.Fatal(err)
	}
	infos, err := f.Readdir(0)
	if err != nil {
		logger.Error.Fatal(err)
	}
	for _, info := range infos {
		//
		// -rw-r--r-- 1 owner group           213 Aug 26 16:31 README
		//
		fmt.Fprint(dataConn, info.Mode().String())
		uid := info.Sys().(*syscall.Stat_t).Uid
		gid := info.Sys().(*syscall.Stat_t).Gid
		fileUser, _ := user.LookupId(strconv.Itoa(int(uid)))
		fileGroup, _ := user.LookupGroupId(strconv.Itoa(int(gid)))
		fmt.Fprint(dataConn, " "+fileUser.Name)
		fmt.Fprint(dataConn, " "+fileGroup.Name+" ")
		fmt.Fprintf(dataConn, "%13d ", info.Size())

		t := time.Unix(info.Sys().(*syscall.Stat_t).Mtim.Unix())
		fmt.Fprint(dataConn, t.Month().String()[:3]+" ")
		fmt.Fprintf(dataConn, "%3d ", t.Day())
		fmt.Fprintf(dataConn, "%2d:%02d ", t.Hour(), t.Minute())
		fmt.Fprint(dataConn, info.Name()+"\r\n")
	}

	fmt.Fprintln(conn, "226 Transfer complete")

	return c.Ctx
}

package server

import "github.com/4hel/ftpd/command"

func parseClose (msg string) command.Command {
	return command.Command{Kind:command.Close}
}

func parseUser(msg string) command.Command  {
	return command.Command{Kind:command.User}
}

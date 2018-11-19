package server

import (
	"bufio"
	"fmt"
	"github.com/4hel/ftpd/logger"
	"net"
)

func ReadLoop(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintln(conn, "220 FTP Server ready.")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Println(cmd)
	}
	if scanner.Err() != nil {
		logger.Error.Println("error reading from connection: ", scanner.Err())
	}
}

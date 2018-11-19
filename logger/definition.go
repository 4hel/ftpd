package logger

import (
	"log"
	"os"
)

var Info = log.New(os.Stdout, "", 0)

var Error = log.New(os.Stderr, "", log.Lshortfile)
package logger

import (
	"log"
	"os"
)

var Info = log.New(os.Stdout, "", log.Lshortfile)

var Error = log.New(os.Stderr, "", log.Lshortfile)
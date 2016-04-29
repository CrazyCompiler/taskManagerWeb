package config

import (
	"os"
)

type Context struct {
	ErrorLogFile *os.File
	ServerAddress string
}

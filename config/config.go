package config

import (
	"os"
)

type ContextObject struct {
	ErrorLogFile *os.File
	ServerAddress string
}

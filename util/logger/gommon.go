package logger

import (
	gommonLog "github.com/labstack/gommon/log"
	"github.com/labstack/gommon/random"
)

func New(prefix string) Logger {
	if prefix == "-" {
		prefix = random.New().String(16, random.Alphanumeric)
	}
	logger := gommonLog.New(prefix)
	return logger
}

package util

import (
	"context"
	"io"
)

type Context interface {
	context.Context
	Logger() io.Writer
	Print(string)
	Println(string)
}

type baseContext struct {
	context.Context
	logger io.Writer
}

func (bf baseContext) Logger() io.Writer {
	return bf.logger
}

func (bf baseContext) Print(msg string) {
	_, _ = bf.logger.Write([]byte(msg))
}

func (bf baseContext) Println(msg string) {
	_, _ = bf.logger.Write([]byte(msg + "\n"))
}

func followerBaseContext() baseContext {
	return baseContext{}
}

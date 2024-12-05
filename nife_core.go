package nife

import (
	"context"
	"errors"
)

var (
	ErrorCommandNotFound = errors.New("Command not found")
)

type Cmd struct {
	Name       string
	Title      string
	Usage      string
	Short      string
	Long       string
	CmdHandler func(ctx context.Context, cmd string, args []string) error
}

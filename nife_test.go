package nife_test

import (
	"errors"
	"testing"

	"github.com/hi100e/nife"
)

func TestRegisterCmd(t *testing.T) {
	cmd := &nife.Cmd{
		Name:  "new-test-cmd",
		Title: "New Test Command",
		Usage: "new-test-cmd",
		Short: "This is a new test command",
		Long:  "This is a new test command that does nothing",
	}

	nocmd, err := nife.GetCmd(cmd.Name)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if !errors.Is(err, nife.ErrorCommandNotFound) {
		t.Fatalf("Expected ErrorCommandNotFound, got %v", err)
	}

	if nocmd != nil {
		t.Fatalf("Expected nil, got %v", cmd)
	}

	nife.RegisterCmd(cmd)

	cmd2, err := nife.GetCmd(cmd.Name)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if cmd2 == nil {
		t.Fatalf("Expected cmd, got nil")
	}

	if cmd2.Name != "new-test-cmd" {
		t.Fatalf("Expected new-test-cmd, got %v", cmd.Name)
	}

	cmds := nife.GetCmds()
	if len(cmds) != 1 {
		t.Fatalf("Expected 1, got %v", len(cmds))
	}
	listCmd := cmds[0]
	if listCmd.Name != "new-test-cmd" {
		t.Fatalf("Expected new-test-cmd, got %v", listCmd.Name)
	}

}

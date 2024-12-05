package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/hi100e/nife"
	_ "github.com/hi100e/nife/init"
)

func main() {
	//get the command line arguments
	args := os.Args[1:]
	//use the first argument as the command
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	command := args[0]

	if command == "" {
		printUsage()
		os.Exit(1)
	}
	//use the rest of the arguments as the command arguments
	commandArgs := args[1:]

	//get the command from the nife package
	cmd, err := nife.GetCmd(command)
	if err != nil {
		//if the command is not found, print an error message
		if errors.Is(err, nife.ErrorCommandNotFound) {
			fmt.Printf("Command not found: %s\n", command)
			os.Exit(1)
		}
		//if there is another error, print the error message
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = RunCmd(cmd, commandArgs)
}

func printUsage() {
	fmt.Println("Usage: nife <command> [args]")
	fmt.Printf("Available Commands:\n")
	for _, cmd := range nife.GetCmds() {
		fmt.Printf("  %s: %s\n", cmd.Name, cmd.Title)
	}
}

func RunCmd(cmd *nife.Cmd, args []string) error {
	ctx := context.WithValue(context.Background(), "cmd", cmd)
	err := cmd.CmdHandler(ctx, cmd.Name, args)
	return err
}

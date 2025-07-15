package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: fmt.Sprintf("Exit the %s", replName),
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func commandExit() error {
	defer os.Exit(0)
	if replName == "" {
		return errors.New("unnamed repl. exiting.")
	}
	fmt.Println(exitMessage)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the %s!\nUsage:\n\n", replName)
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

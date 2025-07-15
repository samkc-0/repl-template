package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	help, ok := commands["help"]
	if !ok {
		log.Fatal("no 'help' command defined in 'commands.go'")
	}
	for {
		fmt.Print(prompt)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatalf("program terminated with error: %v", err)
		}
		input := cleanInput(scanner.Text())
		command, ok := commands[input[0]]
		if !ok {
			err := help.callback()
			if err != nil {
				log.Fatalf("error in 'help' command: %v", err)
			}
		} else {
			err := command.callback()
			if err != nil {
				log.Fatalf("something went wrong calling '%s' command: %v", command.name, err)
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

package main

import (
	"bufio"
	"fmt"
	"github.com/samkc/repl-template/internal/pokeapi"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	repl(cfg)
}

func repl(cfg *config) {
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
			err := help.callback(cfg)
			if err != nil {
				log.Fatalf("error in 'help' command: %v", err)
			}
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

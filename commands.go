package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/samkc-0/repl-template/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsUrl *string
	prevLocationsUrl *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"map": {
			name:        "map",
			description: "Displays the next page of location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the pokemon in a location area.\nusage: explore <location name>",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *config, args ...string) error {
	if replName == "" {
		return errors.New("unnamed repl. exiting.")
	}
	fmt.Println(exitMessage)
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("Welcome to the %s!\nUsage:\n\n", replName)
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	response, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = response.Next
	cfg.prevLocationsUrl = response.Previous
	for _, loc := range response.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsUrl == nil {
		return errors.New("You're on the first page")
	}

	response, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = response.Next
	cfg.prevLocationsUrl = response.Previous
	for _, loc := range response.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: explore <location name>")
	}
	response, err := cfg.pokeapiClient.GetLocationDetails(args[0])
	if err != nil {
		return err
	}

	for _, pokemonEncounter := range response.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}
	return nil
}

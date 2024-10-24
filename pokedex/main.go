package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Victor-AvestaAI/pokedex/pokeapi"
	"github.com/Victor-AvestaAI/pokedex/pokecache"
)

var cache pokecache.Cache

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	prev          *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
	}
}

func commandHelp(config *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage: ")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *config) error {
	os.Exit(0)
	return nil
}

func commandMap(config *config) error {

	locs, err := config.pokeapiClient.GetLocations(config.next)
	if err != nil {
		return err
	}
	config.next = locs.Next
	config.prev = locs.Previous

	fmt.Println()
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	return nil
}

func commandMapb(config *config) error {
	if config.prev == nil {
		return fmt.Errorf("no previous page")
	}
	locs, err := config.pokeapiClient.GetLocations(config.prev)
	if err != nil {
		return err
	}
	config.next = locs.Next
	config.prev = locs.Previous

	fmt.Println()
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println()

	return nil
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func main() {

	fmt.Println(cache)

	scanner := bufio.NewScanner(os.Stdin)

	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeapiClient: pokeClient,
	}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

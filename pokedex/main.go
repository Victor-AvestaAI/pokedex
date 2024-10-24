package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Victor-AvestaAI/pokedex/pokeapi"
	"github.com/Victor-AvestaAI/pokedex/pokecache"
)

type config struct {
	pokeapiClient pokeapi.Client
	cache         pokecache.Cache
	next          *string
	prev          *string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	cache := pokecache.NewCache(5 * time.Minute)

	config := &config{
		pokeapiClient: pokeClient,
		cache:         cache,
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

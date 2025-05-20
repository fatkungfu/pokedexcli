package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatkungfu/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}

// 1. Create support for a simple REPL
// Start an infinite for loop. This loop will execute once for every command the user types in (we don't want to exit the program after just one command)
func startRepl(cfg *config) {
	// Wait for user input using bufio.NewScanner
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ") // Print the prompt (no newline)
		reader.Scan()           // Wait for user input

		words := cleanInput(reader.Text()) // Clean the input using the cleanInput function
		if len(words) == 0 {               // Check if the user entered anything
			continue // If not, continue to the next iteration of the loop
		}

		commandName := words[0] // Get the first word of the input (the command name)
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName] // Check if the command exists in the map
		if exists {                                   // If the command doesn't exist
			err := command.callback(cfg, args...) // Call the command's callback function
			if err != nil {                       // If there was an error
				fmt.Println(err) // Print the error message
			}
			continue
		} else {
			fmt.Println("Unknown command") // Print an error message
			continue                       // Continue to the next iteration of the loop
		}
	}
	// fmt.Printf("Your command was: %s\n", commandName) // Print the first word
}

func cleanInput(text string) []string {
	/* The purpose of this function will be to split the users input into "words" based on whitespace.
	It should also lowercase the input and trim any leading or trailing whitespace. */
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
	}
}

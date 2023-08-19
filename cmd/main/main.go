package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mughilan16/pokedex-go/internal/api"
)

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

var locationOffset = -20
var locationCount = 850

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
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
			name:        "map ",
			description: "Display location areas in pokemon",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb ",
			description: "Display previous location page areas in pokemon",
			callback:    commandMapB,
		},
	}
}

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\npokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command, ok := commands[input]
		if ok {
			err := command.callback()
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			fmt.Println("Invalid commands!! use \"help\" to see the available commands")
		}
	}
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := getCommands()
	fmt.Print(`
Welcome to the Pokedex!
Usage:
`)
	for _, cmd := range commands {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap() error {
	if locationOffset >= locationCount {
		fmt.Println("There is no next page")
	}
	locationOffset += api.LocationLimit
	list, err := api.FetchLocation(locationOffset)
	if err != nil {
		return err
	}
	locationCount = list.Count
	for i, place := range list.Results {
		fmt.Printf("%d. %s\n", locationOffset+i+1, place.Name)
	}
	return nil
}

func commandMapB() error {
	if locationOffset-api.LocationLimit < 0 {
		fmt.Println("There is no previous page")
		return nil
	}
	if locationOffset > 20 {
		locationOffset -= api.LocationLimit
	} else {
		locationOffset = 0
	}
	list, err := api.FetchLocation(locationOffset)
	if err != nil {
		return err
	}
	locationCount = list.Count
	for i, place := range list.Results {
		fmt.Printf("%d. %s\n", locationOffset+i+1, place.Name)
	}
	return nil
}

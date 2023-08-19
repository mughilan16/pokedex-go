package command

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mughilan16/pokedex-go/internal/command/location"
)

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

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
			callback:    location.CommandMap,
		},
	}
}

func Start() {
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

package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := map[string]cliCommand{
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
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		commands[input].callback()
	}
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print(`

Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex

`)
	return nil
}

package location

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type locationList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type mapCommand struct {
	name        string
	description string
	callback    func() error
}

var offset = 0
var count = 850

const (
	limit   = 20
	baseUrl = "https://pokeapi.co/api/v2/location?offset=%d&limit=%d"
)

func getCommands() map[string]mapCommand {
	return map[string]mapCommand{
		"next": {
			name:        "next",
			description: "show next page of the places list",
			callback:    commandNext,
		},
		"prev": {
			name:        "prev",
			description: "show previous page of the places list",
			callback:    commandPrev,
		},
		"print": {
			name:        "print",
			description: "print the current page of the place list",
			callback:    commandCurrent,
		},
		"help": {
			name:        "help",
			description: "provides help for map command",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit from map command",
			callback:    nil,
		},
	}
}

func CommandMap() error {
	mapcmds := getCommands()
	offset = 0
	scanner := bufio.NewScanner(os.Stdin)
	list, err := fetch(offset)
	if err != nil {
		return err
	}
	count = list.Count
	display(offset, list)
	fmt.Printf("use help to see all the commands available\n")
	for {
		fmt.Printf("\nmap > ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}
		cmd, ok := mapcmds[input]
		if ok {
			cmd.callback()
		} else {
			fmt.Println("Invalid comamnds! use \"help\" to see the available commands")
		}
	}
	return nil
}

func display(offset int, list locationList) {
	for i, place := range list.Results {
		fmt.Printf("%d. %s\n", offset+i+1, place.Name)
	}
}

func fetch(offset int) (locationList, error) {
	url := fmt.Sprintf(baseUrl, offset, limit)
	res, err := http.Get(url)
	if err != nil {
		return locationList{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return locationList{}, err
	}

	list := locationList{}
	err = json.Unmarshal(body, &list)
	if err != nil {
		return locationList{}, err
	}
	return list, nil
}

func commandHelp() error {
	commands := getCommands()
	fmt.Print(`
Map commands shows location in pokemon games
Usage:
`)
	for _, cmd := range commands {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandNext() error {
	if offset+limit >= 850 {
		fmt.Println("There is no next page")
	}
	offset += limit
	list, err := fetch(offset)
	if err != nil {
		return err
	}
	display(offset, list)
	return nil
}

func commandPrev() error {
	if offset-limit < 0 {
		fmt.Println("There is no previous page")
	}
	offset -= limit
	list, err := fetch(offset)
	if err != nil {
		return err
	}
	display(offset, list)
	return nil
}

func commandCurrent() error {
	list, err := fetch(offset)
	if err != nil {
		return err
	}
	display(offset, list)
	return nil
}

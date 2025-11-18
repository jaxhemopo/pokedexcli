package main

 

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/jaxhemopo/pokedexcli/internal/pokeapi"
)


type cliCommand struct {
	name string 
	description string 
	callback func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return  map[string]cliCommand{
		"exit": {
			name:		"exit",
			description: "Exit the Pokedex!",
			callback: 	commandExit,
		},
		"help": {
			name: 		"help",
			description: "Displays a help message",
			callback:	 commandHelp,
		},
		"map": {
			name:			"map",
			description:	"Displays the names of 20 location areas in Pokemon world",
			callback: 		commandMap,
		},
			"mapb": {
			name:			"mapb",
			description:	"Displays the previous 20 location areas in Pokemon world",
			callback: 		commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	fields := strings.Fields(lowered)
	return fields
}

func commandMap(cfg *config, args []string) error {
	res, err := pokeapi.GetLocationAreas(cfg.Next)
	if err != nil {
		fmt.Printf("error getting location areas")
	}
	for _,la := range res.Results {
		fmt.Printf("%s\n", la.Name)
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	return nil
}

func commandMapb(cfg *config, args []string) error {
	res, err := pokeapi.GetLocationAreas(cfg.Previous)
	if err != nil {
		fmt.Printf("error getting location areas")
	}
	for _,la := range res.Results {
		fmt.Printf("%s\n", la.Name)
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	return nil
}

func commandExit(cfg *config, args []string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	commands := getCommands()
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n")
	for _, cmd := range commands {
		cString := cmd.name + ":" + cmd.description
		fmt.Printf("%s\n", cString)
	}

	return nil
}


func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Printf("Pokedex >")
		scanner.Scan()
		line := scanner.Text()
		cleaned := cleanInput(line)
				if len(cleaned) == 0 {
   		 continue
			}
		cmd := cleaned[0]
		args := cleaned[1:]

		c,ok := commands[cmd]
		if !ok {
			fmt.Printf("Unknown command\n")
			continue
		}

		err := c.callback(cfg, args)
		if err != nil {
			fmt.Printf("Error running command: %v\n", err)
		}

	}
}





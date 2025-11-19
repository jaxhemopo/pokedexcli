package main

 

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/jaxhemopo/pokedexcli/internal/pokeapi"
	"math/rand"
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
		"explore": {
			name:			"explore",
			description:	"See a list of all Pokemon located in an area",
			callback: 		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Catch a Pokemon to add to your Pokedex!",
			callback: 		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description:	"inspect a Pokemon in your Pokedex!",
			callback: 		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description:	"list all Pokemon in your Pokedex!",
			callback: 		commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	fields := strings.Fields(lowered)
	return fields
}

func commandPokedex(cfg *config, args []string) error {
	for key,_ := range cfg.Pokedex {
		fmt.Printf("%s\n", key)
	}
	return nil
}

func commandInspect(cfg *config, args []string) error {
	name := args[0]
	_,ok := cfg.Pokedex[name]
	if !ok {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}

	fmt.Printf("Name: %s\n", cfg.Pokedex[name].Name)
	fmt.Printf("Height: %d\n", cfg.Pokedex[name].Height)
	fmt.Printf("Weightt: %d\n", cfg.Pokedex[name].Weight)
	fmt.Printf("Stats:\n")
	for _,s := range cfg.Pokedex[name].Stats {
		fmt.Printf(". -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _,t := range cfg.Pokedex[name].Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}

func commandCatch(cfg *config, args []string) error {
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	res, err := pokeapi.CatchPokemon(name)
	if err != nil{
		fmt.Printf("error catching pokemon: %v\n", err)
		return err
	}
	// max exp == 400 so divide base by 4
	chance := (res.BaseExp / 4)
	chance = 100 - chance
	r := rand.Intn(100)
	if r < chance {
		fmt.Printf("%s was caught!\n", name)
		cfg.Pokedex[name] = res
		return nil
	} else {
		fmt.Printf("%s escaped\n", name)
		return nil
	}
}



func commandMap(cfg *config, args []string) error {
	res, err := pokeapi.GetLocationAreas(cfg.Next)
	if err != nil {
		fmt.Printf("error getting location areas: %v\n", err)
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
		fmt.Printf("error getting location areas: %v", err)
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

func commandExplore(cfg *config, args []string) error {
	area, err := pokeapi.GetPokemonList(args[0])
	if err != nil {
		fmt.Printf("area not found")
		return err
	} 
	for _,encounter := range area.PokemonEncounters {
		fmt.Printf("%v\n", encounter.Pokemon.Name)
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





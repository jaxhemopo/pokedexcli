package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	fields := strings.Fields(lowered)
	return fields
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex >")
		scanner.Scan()
		line := scanner.Text()
		cleaned := cleanInput(line)
		fmt.Printf("Your command was: %s\n", cleaned[0])
	}
}





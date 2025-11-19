package main

import (
	"github.com/jaxhemopo/pokedexcli/internal/pokecache"
	"github.com/jaxhemopo/pokedexcli/internal/pokeapi"
	"time"
)
type config struct {
	Cache *pokecache.Cache 
	Next string 
	Previous string 
	Pokedex map[string]pokeapi.Pokemon
}


func main() {
	cache := pokecache.NewCache(5 * time.Second)
	baseurl := "https://pokeapi.co/api/v2/location-area/"
	cfg := &config{
	Cache: cache,
	Next: baseurl,
	Pokedex: make(map[string]pokeapi.Pokemon),
}
	startREPL(cfg)
}


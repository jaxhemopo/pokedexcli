package main

import "github.com/jaxhemopo/pokedexcli/internal/pokecache"

type config struct {
	Cache Cache 
	Next string 
	Previous string 
}


func main() {
	cache := pokecache.NewCache(5 * time.Second)
	baseurl := "https://pokeapi.co/api/v2/location-area/"
	cfg := &config{
	Cache: cache,
	Next: baseurl,
}
	startREPL(cfg)
}


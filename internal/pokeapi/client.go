package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"github.com/jaxhemopo/pokedexcli/internal/pokecache"
	"time"
)

var cache = pokecache.NewCache(5 * time.Second)

type LocationArea struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type LAResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []LocationArea `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	BaseExp int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []PokemonStat `json:"stats"`
	Types []PokemonType `json:"types"`
}

type PokemonType struct {
    Slot int              `json:"slot"`
    Type TypeNameWrapper  `json:"type"`
}

type TypeNameWrapper struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type PokemonStat struct {
    BaseStat int           `json:"base_stat"`
    Effort   int           `json:"effort"`
    Stat     StatInfo      `json:"stat"`
}

type StatInfo struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type AreaResponse struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

func GetLocationAreas(url string) (LAResponse, error) {
	data, found := cache.Get(url)
	if found {
		var resp LAResponse
		if err := json.Unmarshal(data, &resp); err != nil {
			return LAResponse{}, err
		}
		return resp, nil
	} 
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data from URL")
		return LAResponse{}, err
	}
	
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading JSON body")
		return LAResponse{}, err
	}

	cache.Add(url, data)

	var resp LAResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return LAResponse{}, err
	}
	return resp, nil
}

func GetPokemonList(loc string) (AreaResponse, error){
	fullUrl := "https://pokeapi.co/api/v2/location-area/" + loc + "/"
	data, found := cache.Get(fullUrl)
	if found {
		var resp AreaResponse
		if err := json.Unmarshal(data, &resp); err != nil {
			return AreaResponse{}, err
		}
		return resp, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return AreaResponse{}, err
	}

	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return AreaResponse{}, err
	}

	cache.Add(fullUrl, data)

	var resp AreaResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return AreaResponse{}, err
	} 
	return resp, nil

}

func CatchPokemon(name string) (Pokemon, error) {
	fullUrl := "https://pokeapi.co/api/v2/pokemon/" + name + "/"
	data, found := cache.Get(fullUrl)
	var resp Pokemon
	if found {
		if err := json.Unmarshal(data, &resp); err != nil {
			return Pokemon{}, err
		}
		return resp, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()
	
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	cache.Add(fullUrl, data)

	if err := json.Unmarshal(data, &resp); err != nil{
		return Pokemon{}, err
	}
	return resp, nil


}
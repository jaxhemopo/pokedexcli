package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

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

func GetLocationAreas(url string) (LAResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data from URL")
		return LAResponse{}, err
	}
	
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading JSON body")
		return LAResponse{}, err
	}

	var resp LAResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Printf("error decoding body")
		return LAResponse{}, err
	}
	return resp, nil
}
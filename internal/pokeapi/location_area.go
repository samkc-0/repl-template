package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaPokemonEncounters struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	}
}

func (client *Client) ListLocations(pageUrl *string) (LocationAreaResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	data, err := fetchPokeapi(client, url)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var locationAreaResponse LocationAreaResponse
	if err := json.Unmarshal(data, &locationAreaResponse); err != nil {
		return LocationAreaResponse{}, err
	}
	return locationAreaResponse, nil
}

func (client *Client) GetLocationDetails(locationName string) (LocationAreaPokemonEncounters, error) {
	url := baseUrl + "/location-area/" + locationName

	data, err := fetchPokeapi(client, url)
	if err != nil {
		return LocationAreaPokemonEncounters{}, err
	}

	var pokemonEncounters LocationAreaPokemonEncounters
	if err := json.Unmarshal(data, &pokemonEncounters); err != nil {
		return LocationAreaPokemonEncounters{}, err
	}
	return pokemonEncounters, nil
}

func fetchPokeapi(client *Client, url string) ([]byte, error) {
	if data, ok := client.pokecache.Get(url); ok {
		return data, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	client.pokecache.Add(url, data)
	return data, nil
}
